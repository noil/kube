package main

import (
	"net/http"
	"github.com/sirupsen/logrus"
	"os"
)

var log = logrus.New()

func main() {
	port := os.Getenv("SERVICE_PORT")
	if len(port) == 0 {
		log.Fatal("required parameter service port is not set")
	}
	//port := flag.String("port", "8080", "a port")
	//flag.Parse()
	kubeService := NewService()
	http.HandleFunc("/state", kubeService.CheckState)
	http.HandleFunc("/health", kubeService.CheckHealth)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Error(err)
		panic(err)
	}
}
