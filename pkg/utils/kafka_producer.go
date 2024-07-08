package utils

import (
	"hex/mjolnir-core/pkg/interfaces"
	"hex/mjolnir-core/pkg/utils/logging"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var producer *kafka.Producer

type Email struct {
	ReceiverAddress string
	SenderAddress   string
	Subject         string
	TemplateName    string
	Locale          string
	TemplateData    map[string]string
}

type KafkaProducerImpl struct{}

func NewKafkaProducer() interfaces.KafkaProducerInterface {
	return &KafkaProducerImpl{}
}

func (kp *KafkaProducerImpl) InitKafkaProducer(brokers []string) error {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": strings.Join(brokers, ","),
		"client.id":         "mjolnirCore",
		"acks":              "all",
	})

	if err != nil {
		return err
	}

	producer = p
	return nil
}

func (kp *KafkaProducerImpl) SendMessageToKafka(topic string, message []byte) error {
	logging.Info(producer.String())

	deliveryChan := make(chan kafka.Event, 10000)

	err := producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: message,
	},
		deliveryChan,
	)

	if err != nil {
		return err
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		return m.TopicPartition.Error
	}

	return nil
}

func (kp *KafkaProducerImpl) CloseKafkaProducer() {
	producer.Close()
}
