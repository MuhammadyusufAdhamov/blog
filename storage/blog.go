package storage

import (
	"database/sql"
	"fmt"
	"time"
)

type DBManager struct {
	db *sql.DB
}

func NewDBManager(db *sql.DB) *DBManager {
	return &DBManager{db: db}
}

type Blog struct {
	ID          int
	Title       string
	Description string
	Author      string
	CreatedAt   time.Time
}

type GetBlogsQueryParam struct {
	Author string
	Title  string
	Page   int32
	Limit  int32
}

func (b *DBManager) Create(blog *Blog) (*Blog, error) {
	query := `
		insert into blogs(
		                  title,
		                  description,
		                  author
		) VALUES ($1,$2,$3)
		returning id,title,description,author,created_at
`
	row := b.db.QueryRow(
		query,
		blog.Title,
		blog.Description,
		blog.Author,
	)

	var result Blog
	err := row.Scan(
		&result.ID,
		&result.Title,
		&result.Description,
		&result.Author,
		&result.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (b *DBManager) Get(id int) (*Blog, error) {
	var result Blog

	query := `
		select 
			id,
			title,
			description,
			author,
			created_at
		from blogs
		where id=$1
	`

	row := b.db.QueryRow(query, id)
	err := row.Scan(
		&result.ID,
		&result.Title,
		&result.Description,
		&result.Author,
		&result.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (b *DBManager) GetAll(params *GetBlogsQueryParam) ([]*Blog, error) {
	var blogs []*Blog

	offset := (params.Page - 1) + params.Limit

	limit := fmt.Sprintf("Limit %d offset %d ", params.Limit, offset)

	filter := " Where true "
	if params.Author != " " {
		filter += " and author like '%" + params.Author + "%' "
	}

	query := `
		select 
			id,
			title,
			description,
			author,
			created_at
		from blogs
		` + filter + `
		order by created_at desc
` + limit

	rows, err := b.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var blog Blog
		err := rows.Scan(
			&blog.ID,
			&blog.Title,
			&blog.Description,
			&blog.Author,
			&blog.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		blogs = append(blogs, &blog)
	}

	return blogs, nil
}

func (b *DBManager) Update(blog *Blog) (*Blog, error) {
	query := `
		update blogs set
		                 title=$1,
		                 description=$2,
		                 author=$3
		where id=$4
		returning id,title,description,author,created_at
`
	row := b.db.QueryRow(
		query,
		blog.Title,
		blog.Description,
		blog.Author,
		blog.ID,
	)

	var result Blog
	err := row.Scan(
		&result.ID,
		&result.Title,
		&result.Description,
		&result.Author,
		&result.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return blog, nil
}

func (b *DBManager) Delete(id int) error {
	query := "delete from blogs where id=$1"

	result, err := b.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowscount, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowscount == 0 {
		return sql.ErrNoRows
	}
	return nil
}
