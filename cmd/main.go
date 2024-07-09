package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"hex/mjolnir-core/pkg/config"
	"hex/mjolnir-core/pkg/routes"
	"hex/mjolnir-core/pkg/utils"
	wrappers "hex/mjolnir-core/pkg/utils/wrappers"
)

func main() {
	kafkaproducer := utils.NewKafkaProducer()
	kafkaBroker := os.Getenv("KAFKA_BROKER")
	log.Printf("KafkaProducer: %+v", kafkaproducer)

	err := kafkaproducer.InitKafkaProducer([]string{kafkaBroker})
	if err != nil {
		log.Fatalf("Failed to initialize Kafka producer: %v", err)
	}
	defer kafkaproducer.CloseKafkaProducer()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	cfg := config.NewConfig()
	cfg.LoadConfig()

	bcrypt := &wrappers.BcryptWrapper{}
	jwt := &wrappers.JwtWrapper{}

	r := routes.SetupRouter(
		kafkaproducer,
		bcrypt,
		jwt,
		cfg,
	)

	port := cfg.GetEnv("PORT")
	if port == "" {
		port = "8080" // default to port 8080 if not specified
	}

	go func() {
		if err := r.Run(":" + port); err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	<-signals
	log.Println("Shutting down...")
}
