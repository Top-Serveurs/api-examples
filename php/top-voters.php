<?php

// Replace "YOUR_SERVER_TOKEN" with your actual server token
$serverToken = "YOUR_SERVER_TOKEN";

// The period can be "current" or "lastMonth"
$period = "lastMonth";

// Build the URL with the required parameters
$url = "https://api.top-games.net/v1/servers/{$serverToken}/players-ranking?type={$period}";

// Initialize cURL session
$ch = curl_init($url);

// Set cURL options
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);

// Execute the cURL request and get the response
$response = curl_exec($ch);

// Check for cURL errors
if (curl_errno($ch)) {
    echo 'cURL Error: ' . curl_error($ch);
}

// Close the cURL session
curl_close($ch);

// Decode the JSON response
$result = json_decode($response, true);

// Check if the request was successful (code 200)
if ($result['code'] == 200 && $result['success']) {
    // Access the list of players
    $players = $result['players'];

    // Iterate through the list of players and display information
    foreach ($players as $player) {
        echo "Player Name: " . $player['playername'] . "<br>";
        echo "Votes: " . $player['votes'] . "<br><br>";
    }
} else {
    // Display an error message if the request failed
    echo "Request failed. Error code: " . $result['code'] . "<br>";
    echo "Error message: " . $result['message'];
}
