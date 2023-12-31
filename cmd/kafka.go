/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/lucasti79/codepix-go/application/kafka"
	"github.com/lucasti79/codepix-go/infra/db"
	"github.com/spf13/cobra"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

// kafkaCmd represents the kafka command
var kafkaCmd = &cobra.Command{
	Use:   "kafka",
	Short: "Start consuming transactions using apache kafka",
	Run: func(cmd *cobra.Command, args []string) {
		database := db.ConnectDB(os.Getenv("env"))
		deliveryChannel := make(chan ckafka.Event)
		producer := kafka.NewKafkaProducer()

		go kafka.DeliveryReport(deliveryChannel)

		kafkaProcessor := kafka.NewKafkaProcessor(database, producer, deliveryChannel)
		kafkaProcessor.Consume()
	},
}

func init() {
	rootCmd.AddCommand(kafkaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kafkaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kafkaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
