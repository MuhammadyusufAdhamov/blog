package main

import (
	"blog/storage"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var (
	PostgresUser     = "postgres"
	PostgresPassword = "7"
	PostgresHost     = "localhost"
	PostgresPort     = 5432
	PostgresDatabase = "golang"
)

func main() {
	connStr := fmt.Sprintf(
		"host = %s port=%d user=%s password=%s dbname=%s sslmode=disable",
		PostgresHost,
		PostgresPort,
		PostgresUser,
		PostgresPassword,
		PostgresDatabase,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to open connection: %v", err)
	}

	dbManager := storage.NewDBManager(db)

	blogs, err := dbManager.GetAll(&storage.GetBlogsQueryParam{
		Title: "Consequater",
		Limit: 20,
		Page:  2,
	})
	if err != nil {
		log.Fatalf("failed to get all blogs: %v", err)
	}
	PrintBlogs(blogs)

	//err = InsertFakeDate(dbManager)
	//if err != nil {
	//	log.Fatalf("failed to insert fake date: %v", err)
	//}

}

func PrintBlogs(blogs []*storage.Blog) {
	for _, blog := range blogs {
		PrintBlog(blog)
	}
}

func PrintBlog(blog *storage.Blog) {
	fmt.Println("----------- Blog ----------")
	fmt.Println("ID:", blog.ID)
	fmt.Println("Title:", blog.Title)
	fmt.Println("Description:", blog.Description)
	fmt.Println("Author:", blog.Author)
	fmt.Println("Created at:", blog.CreatedAt.Format("2006-01-02 15:04"))
	fmt.Println("--------------------------")
}
