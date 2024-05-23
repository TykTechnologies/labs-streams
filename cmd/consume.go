package cmd

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"github.com/spf13/cobra"
)

// consumerCmd represents the consumer command
var consumerCmd = &cobra.Command{
	Use:   "consume",
	Short: "Consumes from a stream to STDIO",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("consumer called")

		consumer, err := sarama.NewConsumer([]string{"localhost:9093"}, nil)
		if err != nil {
			log.Fatal("error creating consumer", err.Error())
		}

		topic := "instrument.AMZN"
		partition := 0

		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			log.Fatal("error creating partition consumer", err.Error())
		}

		msgs := make(chan *sarama.ConsumerMessage)
		go func() {
			for {
				select {
				case msg := <-msgs:
					fmt.Printf("Consumed message: %s\n", string(msg.Value))
				}
			}
		}()

		for {
			select {
			case msg := <-pc.Messages():
				msgs <- msg
			}
		}
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
