package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/TykTechnologies/labs-streams/model"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cobra"
)

var (
	to         string
	instrument string
)

var publisherCmd = &cobra.Command{
	Use:   "publish",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("publisher called")

		feedBytes, err := os.ReadFile("./feeds/feed.json")
		if err != nil {
			log.Fatal("error reading file", err.Error())
			return
		}
		var feedSlice []*model.PM
		json.Unmarshal(feedBytes, &feedSlice)

		switch to {
		case "redis":
			fmt.Println("publishing to Redis stream")
			err := toRedis(feedSlice)
			if err != nil {
				log.Fatal("error publishing to Redis", err.Error())
			}
		default:
			log.Fatal("unsupported destination", to)
		}

	},
}

func toRedis(feedSlice []*model.PM) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:             "localhost:6379",
		Password:         "",
		DB:               0,
		DisableIndentity: true, // Disable set-info on connect
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return err
	}

	ctx := context.Background()
	for _, tick := range feedSlice {
		r := rand.Intn(10)
		time.Sleep(time.Duration(r) * time.Millisecond * 100)

		t := model.Tick{
			Price1000: int(tick.V[0] * 1000),
		}

		payload := map[string]interface{}{
			"price_1000": t.Price1000,
			"timestamp":  time.Now().Unix(),
		}

		payloadBytes, _ := json.Marshal(payload)
		fmt.Println(string(payloadBytes))

		err = rdb.XAdd(ctx, &redis.XAddArgs{
			Stream: instrument,
			MaxLen: 0,
			ID:     "",
			Values: payload,
		}).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

func init() {
	rootCmd.AddCommand(publisherCmd)
	rootCmd.PersistentFlags().StringVar(&to, "to", "redis", "where the publisher should publish to")
	rootCmd.PersistentFlags().StringVar(&instrument, "instrument", "AMZN", "the instrument to publish")
}
