package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

type PokemonResult struct {
	EntryNo int    `json:"entry_number"`
	Name    string `json:"name"`
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

	r.GET("/", func(c *gin.Context) {

		var flag int = 0

		for i := 0; i <= 20; i++ {
			flag = 0
			for j := 2; j < responseObject.Pokemon[i].EntryNo/2; j++ {
				if responseObject.Pokemon[i].EntryNo%j == 0 {
					flag = 1
					break
				}
			}

			if flag == 0 && responseObject.Pokemon[i].EntryNo > 1 {
				fmt.Printf("%d ", responseObject.Pokemon[i].EntryNo)
				result := PokemonResult{

					EntryNo: responseObject.Pokemon[i].EntryNo,
					Name:    responseObject.Pokemon[i].Species.Name,
				}

				c.JSON(http.StatusOK, map[string]interface{}{
					"status":  true,
					"code":    200,
					"message": "Success",
					"data":    result,
				})
			}

		}

	})
	r.Run(":5000") // listen and server on 0.0.0.0:8080
}
