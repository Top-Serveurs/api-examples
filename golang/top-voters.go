package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Player struct to represent player data
type Player struct {
	Votes      int    `json:"votes"`
	PlayerName string `json:"playername"`
}

// APIResponse struct to represent the structure of the API response
type APIResponse struct {
	Code    int      `json:"code"`
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Players []Player `json:"players"`
}

func fetchData() {
	// Replace "YOUR_SERVER_TOKEN" with your actual server token
	serverToken := "YOUR_SERVER_TOKEN"

	// The period can be "current" or "lastMonth"
	period := "lastMonth"

	// Build the URL with the required parameters
	url := fmt.Sprintf("https://api.top-games.net/v1/servers/%s/players-ranking?type=%s", serverToken, period)

	// Make a GET request
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return
	}
	defer response.Body.Close()

	// Check for a successful response (status code 200)
	if response.StatusCode != http.StatusOK {
		fmt.Printf("Request failed. Status code: %d\n", response.StatusCode)
		return
	}

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Check if the response body is empty
	if len(body) == 0 {
		fmt.Println("Empty response body.")
		return
	}

	// Parse the JSON response
	var apiResponse APIResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Check if the request was successful (code 200)
	if apiResponse.Code == 200 && apiResponse.Success {
		// Access the list of players
		players := apiResponse.Players

		// Iterate through the list of players and display information
		for _, player := range players {
			fmt.Printf("Player Name: %s\n", player.PlayerName)
			fmt.Printf("Votes: %d\n\n", player.Votes)
		}
	} else {
		// Display an error message if the request failed
		fmt.Printf("Request failed. Error code: %d\n", apiResponse.Code)
		fmt.Printf("Error message: %s\n", apiResponse.Message)
	}
}

func main() {
	// Call the function to fetch data
	fetchData()
}
