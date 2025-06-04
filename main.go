package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MaleneJung/go-mux-frontend-routing-walker/frontend"
)

type Config struct {
	Port uint16 `json:"port"`
}

func main() {

	configBuffer, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
		return
	}

	var config Config
	if err := json.Unmarshal(configBuffer, &config); err != nil {
		log.Fatal(err)
		return
	}

	serveMux := http.NewServeMux()

	if err := frontend.MuxFrontendWalker(serveMux, "/", "lenes-modular-user/", true); err != nil {
		log.Fatal(err)
		return
	}

	http.ListenAndServe(":"+fmt.Sprint(config.Port), serveMux)

}
