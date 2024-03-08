Overview

This is a simple HTTP API written in Go that provides information about the top music tracks in a specified region. It fetches track details including track name, artist, image URL, lyrics, and artist URL from various sources and returns the information in JSON format.

Setup
1. Clone the repository to your local machine:

* git clone https://github.com/jayaraman33/top-track
  
2. Navigate to the project directory:

* cd top-track
  
3. Install dependencies:

go mod tidy

4. Replace the placeholders with your actual API keys:
Replace YOUR_LASTFM_API_KEY with your Last.fm API key.
Replace YOUR_MUSIXMATCH_API_KEY with your Musixmatch API key.
Replace YOUR_GOOGLE_API_KEY with your Google API key.

5. Build and run the application:

go build
./music-track-info-api

Usage
Endpoint

_  GET /track?region=REGION_CODE
_ Returns information about the top track in the specified region.
_ Required query parameter:
region: The code representing the region for which you want to retrieve the top track.

Example
curl http://localhost:8080/track?region=us



Dependencies
encoding/json: Package json implements encoding and decoding of JSON objects.
net/http: Package http provides HTTP client and server implementations.
log: Package log implements a simple logging package.
Contributing
Contributions are welcome! Please feel free to submit issues or pull requests.

License
This project is licensed under the MIT License - see the LICENSE file for details.
