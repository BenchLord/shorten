package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	pb "shorten/protos"
)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create a shortened url",
	Run: func(cmd *cobra.Command, args []string) {
		url, err := cmd.Flags().GetString("url")
		if err != nil {
			fmt.Println(err)
			Exit(-1)
		}

		keepAlive, err := cmd.Flags().GetInt64("keep_alive")
		if err != nil {
			fmt.Println(err)
			Exit(-1)
		}

		if err := rpc(func(client pb.ShortenServiceClient) error {
			res, err := client.CreateUrl(context.Background(), &pb.CreateUrlReq{
				Url:       url,
				KeepAlive: keepAlive,
			})
			if err != nil {
				return err
			}
			fmt.Printf("Created shortened url with id '%s'\n", res.GetId())
			return nil
		}); err != nil {
			fmt.Println(err)
			Exit(-1)
		}
	},
}

func init() {
	CreateCmd.Flags().String("url", "", "the url that will be shortened")
	CreateCmd.Flags().Int64("keep_alive", 0, "how many seconds the url will be valid. urls will be valid forever by default")
	RootCmd.AddCommand(CreateCmd)
}
