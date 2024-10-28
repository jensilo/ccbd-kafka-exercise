package main

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/jensilo/ccbd-kafka-exercise/go-src"
)

func main() {
	consumer, _ := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	defer consumer.Close()

	partitionConsumer, _ := consumer.ConsumePartition("live-scores", 0, sarama.OffsetNewest)
	defer partitionConsumer.Close()

	fmt.Println("Live Score Updates:")
	fmt.Println("==================")

	for msg := range partitionConsumer.Messages() {
		var score src.Score
		json.Unmarshal(msg.Value, &score)
		fmt.Printf("\n%s vs %s: %d-%d (%s)\n",
			score.Team1, score.Team2,
			score.Score1, score.Score2,
			score.TimeStamp)
	}
}
