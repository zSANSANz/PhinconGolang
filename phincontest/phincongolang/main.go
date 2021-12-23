package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	response, err := http.Get("https://pokeapi.co/api/v2/pokemon?limit=151")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(responseData))

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  true,
			"code":    200,
			"message": "Success",
			"data":    string(responseData),
		})
	})
	r.Run() // listen and server on 0.0.0.0:8080
}
