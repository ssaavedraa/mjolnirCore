package main

import (
	"log"

	"hex/mjolnir-core/pkg/config"
	"hex/mjolnir-core/pkg/routes"
	wrappers "hex/mjolnir-core/pkg/utils/wrappers"
)

func main() {
	cfg := config.NewConfig()
	cfg.LoadConfig()

	bcrypt := &wrappers.BcryptWrapper{}
	jwt := &wrappers.JwtWrapper{}

	r := routes.SetupRouter(
		bcrypt,
		jwt,
		cfg,
	)

	port := cfg.GetEnv("PORT")
	if port == "" {
		port = "8080" // default to port 8080 if not specified
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
