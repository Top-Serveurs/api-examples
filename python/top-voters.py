import requests

# Replace "YOUR_SERVER_TOKEN" with your actual server token
server_token = "YOUR_SERVER_TOKEN"

# The period can be "current" or "lastMonth"
period = "lastMonth"

# Build the URL with the required parameters
url = f"https://api.top-games.net/v1/servers/{server_token}/players-ranking?type={period}"

try:
    # Make a GET request
    response = requests.get(url)

    # Check if the request was successful (status code 200)
    if response.status_code == 200:
        # Parse the JSON response
        api_response = response.json()

        # Check if the request was successful (code 200)
        if api_response["code"] == 200 and api_response["success"]:
            # Access the list of players
            players = api_response["players"]

            # Iterate through the list of players and display information
            for player in players:
                print(f"Player Name: {player['playername']}")
                print(f"Votes: {player['votes']}\n")
        else:
            # Display an error message if the request failed
            print(f"Request failed. Error code: {api_response['code']}")
            print(f"Error message: {api_response['message']}")
    else:
        # Display an error message if the request failed
        print(f"Request failed. Status code: {response.status_code}")

except requests.RequestException as e:
    # Handle request exceptions
    print(f"Error making the request: {e}")
