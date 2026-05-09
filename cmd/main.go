package main

import (
	"context"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/kirillivontyev/finance_tracker/internal/db"
	"github.com/kirillivontyev/finance_tracker/internal/repository"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
		return
	}

	ctx := context.Background()

	conn, err := db.CreateConnectionDB(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("База подрублена!")

	//lim := 1293.23

	/*categ := models.Category{
		ID:           1,
		CategoryName: "sdfjdksjfbsl",
		MonthlyLimit: &lim,
	}*/

	rep := repository.NewCategoryRepositoty(conn)

	if err := rep.DeleteCategory(ctx, 3); err != nil {
		fmt.Println(err.Error())
		return
	}

	cat, err := rep.GetCategory(ctx)
	fmt.Println(cat, err)

}
