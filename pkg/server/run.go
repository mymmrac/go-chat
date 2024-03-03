package server

import (
	"context"
	"fmt"
	"net"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/moby/moby/pkg/namesgenerator"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/mymmrac/go-chat/pkg/api/chat"
)

func Run() error {
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", viper.GetInt("port")))
	if err != nil {
		return fmt.Errorf("listen: %w", err)
	}

	grpcServer := grpc.NewServer()
	chat.RegisterChatServer(grpcServer, newServer())

	log.Info("Listening...")
	if err = grpcServer.Serve(listener); err != nil {
		return fmt.Errorf("serve: %w", err)
	}

	return nil
}

type server struct {
	chat.UnimplementedChatServer

	registerClient   chan communication
	unregisterClient chan communication

	clients   map[string]communication
	responses chan *chat.ChatResponse
}

func newServer() *server {
	srv := &server{
		UnimplementedChatServer: chat.UnimplementedChatServer{},
		registerClient:          make(chan communication, 1),
		unregisterClient:        make(chan communication, 1),
		clients:                 make(map[string]communication),
		responses:               make(chan *chat.ChatResponse, 64),
	}
	go srv.handleClients()
	return srv
}

type communication struct {
	user      string
	responses chan<- *chat.ChatResponse
}

func (s *server) handleClients() {
	for {
		select {
		case client, ok := <-s.registerClient:
			if !ok {
				return
			}

			s.clients[client.user] = client
		case client, ok := <-s.unregisterClient:
			if !ok {
				return
			}

			delete(s.clients, client.user)
		case response, ok := <-s.responses:
			if !ok {
				return
			}

			for _, client := range s.clients {
				select {
				case client.responses <- response:
					// Sent
				default:
					log.Warn("Lost event", "user", client.user)
				}
			}
		}
	}
}

func (s *server) Communication(stream chat.Chat_CommunicationServer) error {
	ctx, cancel := context.WithCancel(stream.Context())
	defer cancel()

	responses := make(chan *chat.ChatResponse, 64)

	client := communication{
		user:      namesgenerator.GetRandomName(0),
		responses: responses,
	}

	log.Info("Register client", "user", client.user)
	s.registerClient <- client
	defer func() {
		log.Info("Unregister client", "user", client.user)
		s.unregisterClient <- client
	}()

	state := &clientState{
		Context:   ctx,
		user:      client.user,
		responses: responses,
	}

	state.sendResponse(&chat.ChatResponse{
		Type: &chat.ChatResponse_Me{
			Me: client.user,
		},
	})

	s.responses <- &chat.ChatResponse{
		Type: &chat.ChatResponse_Connected{
			Connected: client.user,
		},
	}
	defer func() {
		s.responses <- &chat.ChatResponse{
			Type: &chat.ChatResponse_Disconnected{
				Disconnected: client.user,
			},
		}
	}()

	go func() {
		defer cancel()

		for {
			select {
			case <-ctx.Done():
				return
			case response, ok := <-responses:
				if !ok {
					return
				}

				if err := stream.Send(response); err != nil {
					log.Error("Send response", "error", err)
					return
				}
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			// Continue
		}

		request, err := stream.Recv()
		if err != nil {
			return fmt.Errorf("recive: %w", err)
		}

		if err = s.processRequest(state, request); err != nil {
			return fmt.Errorf("process request: %w", err)
		}
	}
}

type clientState struct {
	context.Context
	user      string
	responses chan<- *chat.ChatResponse
}

func (s *clientState) sendResponse(response *chat.ChatResponse) {
	select {
	case <-s.Done():
		// Disconnected
	case s.responses <- response:
		// Sent
	}
}

func (s *clientState) sendError(err *chat.ChatResponse_Error) {
	s.sendResponse(&chat.ChatResponse{
		Type: &chat.ChatResponse_Error_{
			Error: err,
		},
	})
}

func (s *server) processRequest(state *clientState, request *chat.ChatRequest) error {
	switch req := request.Type.(type) {
	case *chat.ChatRequest_Message:
		text := strings.TrimSpace(req.Message)
		if text == "" {
			state.sendError(&chat.ChatResponse_Error{
				Message: "Empty messages aren't allowed",
			})
			return nil
		}

		log.Info("Message", "user", state.user, "text", text)
		s.responses <- &chat.ChatResponse{
			Type: &chat.ChatResponse_Message_{
				Message: &chat.ChatResponse_Message{
					User: state.user,
					Text: text,
				},
			},
		}
	default:
		return fmt.Errorf("unsupported request: %T", req)
	}
	return nil
}
