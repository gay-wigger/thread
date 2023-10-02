package twitter 

import (
  "fmt" 
  "net/http"
  "net/url"
  "encoding/json"
  "strings"
)

func getAccessToken(client *http.Client, apiKey string, apiSecret string) string { 
  data := url.Values{}
  data.Set("grant_type", "client_credentials")

  req, _ := http.NewRequest("POST", "https://api.twitter.com/oauth2/token", strings.NewReader(data.Encode()))
  req.SetBasicAuth(apiKey, apiSecret)
  req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")

  resp, err := client.Do(req)
  if err != nil {
    fmt.Println("Error:", err)
  }
  defer resp.Body.Close()
  
  var result = decodeResults(resp)
  return result["access_token"].(string)
}

func getTwitterUserId(client *http.Client, accessToken string, username string) (string, error) {
  req, _ := http.NewRequest("GET", fmt.Sprintf("https://api.twitter.com/2/users/by/username/%s", username), nil)
  req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
  fmt.Println()
  resp, err := client.Do(req)
  if err != nil {
    fmt.Println("Error:", err)
    return "", err
  }
  defer resp.Body.Close()
  
  var result map[string]interface{}
  json.NewDecoder(resp.Body).Decode(&result)
  return result["data"].(map[string]interface{})["id"].(string), nil
}

func getFollowers(client *http.Client, accessToken string, id string) {
  req, _ := http.NewRequest("GET", fmt.Sprintf("https://api.twitter.com/2/users/%s/followers", id), nil)
  req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
  resp, err := client.Do(req)
  if err != nil {
    fmt.Println("Error:", err)
    return
  }
  defer resp.Body.Close()

  var result = decodeResults(resp)
  followers := result["data"].([]interface{})
  for _, follower := range followers {
    fmt.Println(follower.(map[string]interface{})["username"])
  }
}

func decodeResults(resp *http.Response) map[string]interface{} {
  var result map[string]interface{}
  json.NewDecoder(resp.Body).Decode(&result)
  return result
}
