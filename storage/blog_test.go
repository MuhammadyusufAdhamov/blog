package storage

import (
	"github.com/bxcodec/faker/v4"
	"github.com/stretchr/testify/require"
	"testing"
)

func createBlog(t *testing.T) *Blog {
	blog, err := dbManager.Create(&Blog{
		Title:       faker.Sentence(),
		Description: faker.Sentence(),
		Author:      faker.Name(),
	})
	require.NoError(t, err)
	require.NotEmpty(t, blog)

	return blog
}

func deleteBlog(id int, t *testing.T) {
	err := dbManager.Delete(id)
	require.NoError(t, err)
}

func TestGetBlog(t *testing.T) {
	b := createBlog(t)

	blog, err := dbManager.Get(b.ID)
	require.NoError(t, err)
	require.NotEmpty(t, blog)

	deleteBlog(blog.ID, t)
}

func TestCreateBlog(t *testing.T) {
	b := createBlog(t)
	deleteBlog(b.ID, t)
}

func TestUpdateBlog(t *testing.T) {
	b := createBlog(t)

	b.Author = faker.Name()
	b.Description = faker.Sentence()
	b.Title = faker.Sentence()

	blog, err := dbManager.Update(b)
	require.NoError(t, err)
	require.NotEmpty(t, blog)

	deleteBlog(blog.ID, t)
}

func TestDeleteBlog(t *testing.T) {
	b := createBlog(t)

	deleteBlog(b.ID, t)
}

func TestGetAll(t *testing.T) {
	b := createBlog(t)

	blogs, err := dbManager.GetAll(&GetBlogsQueryParam{
		Limit: 10,
		Page:  1,
	})

	require.NoError(t, err)
	require.GreaterOrEqual(t, 1, len(blogs))

	deleteBlog(b.ID, t)
}
