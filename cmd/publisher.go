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

var publisherCmd = &cobra.Command{
	Use:   "publisher",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("publisher called")

		feedBytes, err := os.ReadFile("./feeds/feed.json")
		if err != nil {
			log.Fatal("error reading file", err.Error())
			return
		}
		var feedSlice []model.PM
		json.Unmarshal(feedBytes, &feedSlice)

		rdb := redis.NewClient(&redis.Options{
			Addr:             "localhost:6379",
			Password:         "",
			DB:               0,
			DisableIndentity: true, // Disable set-info on connect
		})
		_, err = rdb.Ping(context.Background()).Result()
		if err != nil {
			log.Fatal("Unbale to connect to Redis", err.Error())
		}

		ctx := context.Background()
		for _, tick := range feedSlice {
			r := rand.Intn(10)
			time.Sleep(time.Duration(r) * time.Millisecond * 100)

			t := model.Tick{
				Instrument: "AMZN",
				Price1000:  int(tick.V[0] * 1000),
			}

			payload := map[string]interface{}{
				"instrument": t.Instrument,
				"price_1000": t.Price1000,
				"timestamp":  time.Now().Unix(),
			}

			payloadBytes, _ := json.Marshal(payload)
			fmt.Println(string(payloadBytes))

			err = rdb.XAdd(ctx, &redis.XAddArgs{
				Stream: "AMZN",
				MaxLen: 0,
				ID:     "",
				Values: payload,
			}).Err()
			if err != nil {
				log.Fatal("error publishing to Redis stream", err.Error())
			}
			//fmt.Printf("tick %s %v\n", time.Now().String(), t)
		}
	},
}

func init() {
	rootCmd.AddCommand(publisherCmd)
}
