package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

// A Pokemon Struct to map every pokemon to.
type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

// A struct to map our Pokemon's Species which includes it's name
type PokemonSpecies struct {
	Name string `json:"name"`
}

func main() {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	r := gin.Default()

	for i := 3; i < len(responseObject.Pokemon); i++ {
		checkPrimeNumber(responseObject.Pokemon[i].EntryNo, responseObject.Pokemon[i].Species.Name)

	}

	r.GET("/", func(c *gin.Context) {

		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  true,
			"code":    200,
			"message": "Success",
			"data":    responseObject,
		})
	})
	r.Run() // listen and server on 0.0.0.0:8080
}

func checkPrimeNumber(num int, name string) {

	sq_root := int(math.Sqrt(float64(num)))
	for i := 2; i <= sq_root; i++ {
		if num%i == 0 {
			return
		}
	}
	fmt.Printf("%d %s \n", num, name)

}
