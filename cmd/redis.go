package cmd

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	pb "shorten/protos"
	"shorten/server"
	"shorten/server/adapters"
)

// This should really be a flag on the server command
// instead of a seperate command but not that important.

var RedisCmd = &cobra.Command{
	Use:   "redis",
	Short: "starts the shorten service with redis storage",
	Run: func(cmd *cobra.Command, args []string) {
		address := cmd.Flag("address").Value.String()
		address, err := cmd.Flags().GetString("address")
		if err != nil {
			fmt.Println(err)
			Exit(-1)
		}
		port := cmd.Flag("port").Value.String()
		storage, err := adapters.NewRedisStorage(address, port)
		if err != nil {
			fmt.Println(err)
			Exit(-1)
		}

		service := server.NewService().WithStorage(storage)
		lis, err := net.Listen("tcp", ":8080")
		if err != nil {
			fmt.Println(err)
			Exit(-1)
		}

		grpcServer := grpc.NewServer()
		pb.RegisterShortenServiceServer(grpcServer, service)
		grpcServer.Serve(lis)
	},
}

func init() {
	RedisCmd.Flags().String("address", "localhost", "the redis server address")
	RedisCmd.Flags().Int64("port", 6379, "the redis server port")
	ServerCmd.AddCommand(RedisCmd)
}
