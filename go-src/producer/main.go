package main

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/jensilo/ccbd-kafka-exercise/go-src"
	"time"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, _ := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	defer producer.Close()

	for {
		var score src.Score
		fmt.Print("Enter Team 1: ")
		fmt.Scanf("%s", &score.Team1)
		fmt.Print("Enter Team 2: ")
		fmt.Scanf("%s", &score.Team2)
		fmt.Print("Enter Score (e.g., 2 1): ")
		fmt.Scanf("%d %d", &score.Score1, &score.Score2)

		score.TimeStamp = time.Now().Format(time.RFC3339)

		scoreJSON, _ := json.Marshal(score)
		msg := &sarama.ProducerMessage{
			Topic: "live-scores",
			Value: sarama.StringEncoder(scoreJSON),
		}
		producer.SendMessage(msg)
		fmt.Println("Score updated!")
	}
}
