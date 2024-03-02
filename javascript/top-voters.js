const axios = require('axios');

async function fetchData() {
    try {
        // Replace "YOUR_SERVER_TOKEN" with your actual server token
        const serverToken = "YOUR_SERVER_TOKEN";

        // The period can be "current" or "lastMonth"
        const period = "lastMonth";

        // Build the URL with the required parameters
        const url = `https://api.top-games.net/v1/servers/${serverToken}/players-ranking?type=${period}`;

        // Make a GET request using Axios
        const response = await axios.get(url);

        const result = response.data;

        // Check if the request was successful (code 200)
        if (result.code === 200 && result.success) {
            // Access the list of players
            const players = result.players;

            // Iterate through the list of players and display information
            players.forEach(player => {
                console.log("Player Name: " + player.playername);
                console.log("Votes: " + player.votes + "\n");
            });
        } else {
            // Display an error message if the request failed
            console.error("Request failed. Error code: " + result.code);
            console.error("Error message: " + result.message);
        }
    } catch (error) {
        // Handle any errors that occurred during the request
        console.error("Error making the request:", error.message);
    }
}

// Call the function to fetch data
fetchData();
