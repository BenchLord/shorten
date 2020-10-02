package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	pb "shorten/protos"
)

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "retrieve a shortened url",
	Run: func(cmd *cobra.Command, args []string) {
		id := cmd.Flag("id").Value.String()
		if err := rpc(func(client pb.ShortenServiceClient) error {
			res, err := client.GetUrl(context.Background(), &pb.GetUrlReq{
				Id: id,
			})
			if err != nil {
				return err
			}
			fmt.Printf("Url: %s\nTTL: %d\n", res.GetUrl(), res.GetKeepAlive())
			return nil
		}); err != nil {
			fmt.Println(err)
			Exit(-1)
		}
	},
}

func init() {
	GetCmd.Flags().String("id", "", "the id of the shortened url")
	RootCmd.AddCommand(GetCmd)
}
