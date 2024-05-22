package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"

	"github.com/spf13/cobra"
)

// consumerCmd represents the consumer command
var consumerCmd = &cobra.Command{
	Use:   "consumer",
	Short: "Consumes from a stream to STDIO",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("consumer called")

		rdb := redis.NewClient(&redis.Options{
			Addr:             "localhost:1379",
			Password:         "",
			DB:               0,
			DisableIndentity: true, // Disable set-info on connect
		})
		_, err := rdb.Ping(context.Background()).Result()
		if err != nil {
			log.Fatal("Unbale to connect to Redis", err.Error())
		}

		//ctx := context.Background()
		//
		//subject := "AMZN"

	},
}

func init() {
	rootCmd.AddCommand(consumerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// consumerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// consumerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
