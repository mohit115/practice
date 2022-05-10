package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func initializeRoutes() {
	router.GET("/get")
	router.POST("/store")
	router.GET("/status")
}

func getA() gin.HandlerFunc {
	
	return func(c *gin.Context) {
		// If there's no error or if the token is not empty
		// the user is already logged in
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if loggedIn {
			// if token, err := c.Cookie("token"); err == nil || token != "" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
