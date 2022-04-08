package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

func PostMethod(c *gin.Context) {
    message := "Posted"
    c.JSON(http.StatusOK, message)
}

func GetMethod(c *gin.Context) {
  message := "Getted"
  c.JSON(http.StatusOK, message)
}

func main() {
  router := gin.Default()

  router.POST("/", PostMethod)
  router.GET("/", GetMethod)

  listenPort := "4000"

  router.Run(":"+listenPort)
}
