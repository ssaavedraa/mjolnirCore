package interfaces

type KafkaProducerInterface interface {
	InitKafkaProducer(brokers []string) error
	SendMessageToKafka(topic string, message []byte) error
	CloseKafkaProducer()
}
