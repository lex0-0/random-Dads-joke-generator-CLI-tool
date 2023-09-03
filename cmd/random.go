/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get a ramdon dad joke",
	Long:  `It just givs u random dads joke form an api or some thing like that`,
	Run: func(cmd *cobra.Command, args []string) {
		getRandomjoke()
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)

}

type Joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func getRandomjoke() {
	url := "https://icanhazdadjoke.com/"
	responseBytes := Getjokedata(url)
	joke := Joke{}
	if err := json.Unmarshal(responseBytes, &joke); err != nil {
		log.Printf("could not unmarshal response - %v", err)
	}

	fmt.Println(string(joke.Joke))
}
func Getjokedata(baseAPT string) []byte {
	request, err := http.NewRequest(
		http.MethodGet,
		baseAPT,
		nil,
	)
	if err != nil {
		log.Printf("could not request a dad joke - %v", err)
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "Dad joke CLI")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("could not make a request - %v", err)
	}
	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("could not read response body - %v", err)
	}
	return responseBytes
}
