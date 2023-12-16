package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/suhail34/sorting_server/internal/server"
)

func main() {
  srv := server.NewServer()
  if err := srv.Listen(":8000"); err != nil {
    logrus.Error("Error Starting the server : ", err)
    os.Exit(1)
  }
}
