package main

import (
	"github.com/sirupsen/logrus"
	"youtube-manager/databases"
	"youtube-manager/models"
)

func main() {
	db, err := databases.Connect()
	defer db.Close()

	if err != nil {
		logrus.Fatal(err)
	}

	db.Debug().AutoMigrate(&models.User{})
	db.Debug().AutoMigrate(&models.Favorite{})
}
