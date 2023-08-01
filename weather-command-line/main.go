package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func helloWorld(c *gin.Context) {
	fmt.Fprintf(c.Writer, "Hello World")
}
func handleRequests() {
	router := gin.Default()

	router.GET("/", helloWorld)

	http.ListenAndServe(":8082", router)
}


func main() {

	handleRequests()
}