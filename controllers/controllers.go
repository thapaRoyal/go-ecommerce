package controllers

import "github.com/gin-gonic/gin"

func HashPassword(password string) string {}

func VerifyPassword(userPassword string, givenPassword string) (bool, string) {}

func SignUp() gin.HandlerFunc {}

func Login() gin.HandlerFunc {}

func ProductViewAdmin() gin.HandlerFunc {}

func SearchProduct() gin.HandlerFunc {}

func SearchProductByQuery() gin.HandlerFunc {}
