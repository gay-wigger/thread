package main

import (
	"fmt"
	"net/http"
  "os"
  "encoding/json"
  "twitter"
  "terminal"
)

type Config struct {
  APIKey string `json:"apiKey"`
  APISecret string `json:"apiSecret"`
}

func main() {
  if len(os.Args) < 2 {
    fmt.Println("Please provide a twitter username")
    return
  }

  config, err := loadConfig("config.json")
  if err != nil {
    fmt.Println("Error:", err)
    return
  }

	apiKey := config.APIKey
	apiSecret := config.APISecret
  username := os.Args[1]
	client := &http.Client{}
  
  accessToken := twitter.getAccessToken(client, apiKey, apiSecret)
  id := getTwitterUserId(client, accessToken, username)
  getFollowers(client, accessToken, id)
}

func loadConfig(filename string) (Config, error) {
  var config Config
  configFile, err := os.Open(filename)
  defer configFile.Close()
  if err != nil {
    return config, err
  }
  json.NewDecoder(configFile).Decode(&config)
  return config, nil
}
