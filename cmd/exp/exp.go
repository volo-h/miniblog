package main

import (
  "fmt"

  "github.com/volo-h/miniblog/models"
)


func main() {
  cfg := models.DefaultPostgresConfig()
  db, err := models.Open(cfg)

  if err != nil {
    panic(err)
  }
  defer db.Close()

  err = db.Ping()
  if err != nil {
    panic(err)
  }

  fmt.Println("Connected!")

  us := models.UserService{
    DB: db,
  }

  user, err := us.Create("em2@google.com", "volo")
  if err != nil {
    panic(err)
  }
  fmt.Println(user)
}
