package config

import "os"

var (
	ORDER_CODE string
)

func LoadConfig() {
	ORDER_CODE = os.Getenv("ORDER_CODE")
}