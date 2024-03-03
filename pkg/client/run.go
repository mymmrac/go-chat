package client

import (
	"context"
	"fmt"
	"time"

	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/mymmrac/go-chat/pkg/api/chat"
)

func Run() error {
	ctx := context.Background()

	conn, err := grpc.DialContext(
		ctx,
		viper.GetString("address"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return fmt.Errorf("connect: %w", err)
	}
	defer func() {
		_ = conn.Close()
	}()

	client := chat.NewChatClient(conn)

	stream, err := client.Communication(ctx)
	if err != nil {
		return fmt.Errorf("communication: %w", err)
	}

	go handleInput(stream)

	var response *chat.ChatResponse
	for {
		response, err = stream.Recv()
		if err != nil {
			return fmt.Errorf("recive: %w", err)
		}

		if err = processResponse(response); err != nil {
			return fmt.Errorf("process response: %w", err)
		}
	}
}

func handleInput(stream chat.Chat_CommunicationClient) {
	var text string
	for {
		_, err := fmt.Scanln(&text)
		if err != nil {
			log.Error("Scan", "error", err)
		}

		err = stream.Send(&chat.ChatRequest{
			Type: &chat.ChatRequest_Message{
				Message: text,
			},
		})
		if err != nil {
			log.Error("Send message", "error", err)
		}
	}
}

func processResponse(response *chat.ChatResponse) error {
	switch resp := response.Type.(type) {
	case *chat.ChatResponse_Me:
		fmt.Printf("Me [%s]\n", resp.Me)
	case *chat.ChatResponse_Connected:
		fmt.Printf("Connected [%s]\n", resp.Connected)
	case *chat.ChatResponse_Disconnected:
		fmt.Printf("Disconnected [%s]\n", resp.Disconnected)
	case *chat.ChatResponse_Error_:
		fmt.Println("Error:", resp.Error.Message)
	case *chat.ChatResponse_Message_:
		fmt.Printf("[%s][%s]: %s\n", time.Now().Format(time.TimeOnly), resp.Message.User, resp.Message.Text)
	default:
		return fmt.Errorf("unsupported response: %T", resp)
	}
	return nil
}
