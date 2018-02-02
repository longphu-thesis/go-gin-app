package repository

import "github.com/longphu-thesis/go-gin-app/blog/author"

type AuthorRepository interface {
	GetByID(id int64) (*author.Author, error)
}
