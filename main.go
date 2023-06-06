package main

import (
	kvc "key-value-client/key_value_client"
	"log"
)

func main() {
  client := kvc.New("localhost:6379")

  client.Connect()

  var res string
  var err error

  res, err = client.Set("key", "value")

  if err != nil {
    log.Println(err)
  }

  log.Println(res)

  res, err = client.Get("key")

  if err != nil {
    log.Println(err)
  }

  log.Println(res)
}
