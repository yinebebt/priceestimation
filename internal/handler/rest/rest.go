package rest

import "github.com/gin-gonic/gin"

type User interface {
	CreateUser() gin.HandlerFunc
	GetUser() gin.HandlerFunc
	DeleteUser() gin.HandlerFunc
}
