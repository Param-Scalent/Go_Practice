package repository

import "github.com/Param-Scalent/Go_Practice/Go-Crash-Course/REST-API-with-MUX/entity"

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}
