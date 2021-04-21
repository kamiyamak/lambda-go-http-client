package main

import (
  "fmt"
  "os"
  "net/http"
  "io/ioutil"
  _ "github.com/aws/aws-lambda-go/lambda"
)


func main() {
  url := os.Getenv("REQUEST_URL") // lambda環境変数から取得

  resp, _ := http.Get(url)
  defer resp.Body.Close()

  byteArray, _ := ioutil.ReadAll(resp.Body)
  fmt.Println(string(byteArray)) // htmlをstringで取得
}