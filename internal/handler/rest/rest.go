package rest

import "github.com/gin-gonic/gin"

type User interface {
	CreateUser() gin.HandlerFunc
	GetUser() gin.HandlerFunc
	DeleteUser() gin.HandlerFunc
	LoginUser() gin.HandlerFunc
}

type Location interface {
	CreateLocation() gin.HandlerFunc
	GetLocation() gin.HandlerFunc
}

type PriceEstimation interface {
	CreatePriceEstimation() gin.HandlerFunc
	GetPriceEstimation() gin.HandlerFunc
	DeletePriceEstimation() gin.HandlerFunc
	UpdatePriceEstimation() gin.HandlerFunc
	ListPriceEstimation() gin.HandlerFunc
}
