package cache

import "github.com/Param-Scalent/Go_Practice/Go-Crash-Course/REST-API-with-MUX/entity"

type PostCache interface {
	Set(key string, value *entity.Post)
	Get(key string) *entity.Post
}
