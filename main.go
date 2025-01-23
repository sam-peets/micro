package main

import (
	"fmt"

	"github.com/sam-peets/micro/auth"
	"github.com/sam-peets/micro/db"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func main() {
	client := db.GetClient()
	defer client.Close()
	context, cancel := client.Context()
	defer cancel()
	client.Client.Ping(context, readpref.Primary())

	user, err := auth.NewUser(&auth.UserPayload{
		Username: "tester2",
		Password: "asdf",
	})
	fmt.Println(user, err)
}
