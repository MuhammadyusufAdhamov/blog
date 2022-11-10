package main

import (
	"blog/storage"
	"fmt"
	"github.com/bxcodec/faker/v4"
)

func InsertFakeDate(db *storage.DBManager) error {
	var blog storage.Blog
	for i := 1; i <= 1000; i++ {
		blog.Title = faker.Sentence()
		blog.Description = faker.Sentence()
		blog.Author = faker.FirstName() + " " + faker.LastName()

		_, err := db.Create(&blog)
		if err != nil {
			return err
		}
		if i%100 == 0 {
			fmt.Println(i)
		}
	}
	return nil
}
