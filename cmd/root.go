package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	pb "shorten/protos"
)

var RootCmd = &cobra.Command{
	Use:   "shorten",
	Short: "shorten",
}

var Exit = os.Exit

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		RootCmd.Println(err)
		Exit(-1)
	}
}

func rpc(f func(pb.ShortenServiceClient) error) error {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("error dialing shorten server: %v", err)
	}

	client := pb.NewShortenServiceClient(conn)
	return f(client)
}
