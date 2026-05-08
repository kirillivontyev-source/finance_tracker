package main

import (
	"context"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/kirillivontyev/finance_tracker/internal/db"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
		return
	}

	ctx := context.Background()

	if _, err := db.CreateConnectionDB(ctx); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("База подрублена!")

}
