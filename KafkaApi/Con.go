package KafkaApi

import (
	"vamshi/Models"
    "fmt"
    "os"
    "os/signal"
    "github.com/confluentinc/confluent-kafka-go/kafka"
)


func KafkaProducerSetup(bootstrap_server string) (producer *kafka.Producer){
	// Create Kafka producer configuration
	producerConfig := &kafka.ConfigMap{
		// "bootstrap.servers": "localhost:9092", // Kafka broker address
		"bootstrap.servers": bootstrap_server,
	}
	
	
	// Create Kafka producer
	producer, err := kafka.NewProducer(producerConfig)
	if err != nil {
			fmt.Printf("Failed to create Kafka producer: %v\n", err)
			os.Exit(1)
	}
	return producer

}



func KafkaPublisher(message Models.EventJson,topic string,producer *kafka.Producer) {

    // Produce message
    deliveryChan := make(chan kafka.Event)
    err := producer.Produce(&kafka.Message{
        TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
        Value:          []byte(message),
    }, deliveryChan)

	if err != nil{
		fmt.Println("unable to publish message to kafka due to this error ",err)
	}

    // Wait for delivery report
    ev := <-deliveryChan
    m := ev.(*kafka.Message)
    if m.TopicPartition.Error != nil {
        fmt.Printf("Failed to deliver message: %v\n", m.TopicPartition.Error)
    } else {
        fmt.Printf("Message delivered to topic %s [%d] at offset %v\n", *m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
    }

    // // Close the producer
    // producer.Close()

    // // Handle graceful shutdown
    // sigchan := make(chan os.Signal, 1)
    // signal.Notify(sigchan, os.Interrupt)
    // <-sigchan
    // fmt.Println("Shutting down...")
}
