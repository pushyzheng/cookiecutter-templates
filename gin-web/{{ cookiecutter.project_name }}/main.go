package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	initRouter(r)

	log.Println("listen and serve on http://localhost:8080 (for windows 'localhost:8080')")
	err := r.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
