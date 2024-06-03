package config

type Config interface {
	LoadConfig ()
	GetEnv (key string) string
}