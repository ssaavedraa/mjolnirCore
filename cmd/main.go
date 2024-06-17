package main

import (
	"log"

	"hex/mjolnir-core/pkg/config"
	"hex/mjolnir-core/pkg/routes"
	utils "hex/mjolnir-core/pkg/utils/wrappers"
)

func main() {
	config := config.NewConfig()

	config.LoadConfig()

	bcrypt := &utils.BcryptWrapper{}
	jwt := &utils.JwtWrapper{}

	r := routes.SetupRouter(bcrypt, jwt, config)

	port := config.GetEnv("PORT")

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
