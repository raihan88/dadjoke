/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get a random Dad Joke",
	Long: "This command fetches a random joke from the icanhazdadjoke api",
	Run: func(cmd *cobra.Command, args []string) {
		getRandomJoke()
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// randomCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// randomCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


type Joke struct{
	ID string `json:"id"`
	Joke string `json:"joke"`
	Status int `json:"status"`
}


func getRandomJoke(){
	url := "https://icanhazdadjoke.com/"
	responseBytes := getJokeData(url)
	joke := Joke{}
	if err := json.Unmarshal(responseBytes, &joke); err != nil{
		log.Printf("Could not unmarshal response -%v", err)
	}

	fmt.Println(string(joke.Joke))

}


func getJokeData(baseAPI string) []byte{
	request, err := http.NewRequest(http.MethodGet,baseAPI,nil)

	if err != nil{
		log.Printf("Could not request a dadjoke -%v",err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "Dadjoke CLI(github.com/raihan88/dadjoke)")
	
	response, err := http.DefaultClient.Do(request)
	if err != nil{
		log.Printf("Could not make a request -%v", err)
	}

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil{
		log.Printf("Could not make a response -%v", err)
	}

	return responseBytes
}