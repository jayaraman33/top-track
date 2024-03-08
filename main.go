package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	lastFMAPIKey     = "YOUR_LASTFM_API_KEY"
	musixmatchAPIKey = "YOUR_MUSIXMATCH_API_KEY"
	googleAPIKey     = "YOUR_GOOGLE_API_KEY"
)

type TrackInfo struct {
	Name      string `json:"name"`
	Artist    string `json:"artist"`
	ImageURL  string `json:"image_url"`
	Lyrics    string `json:"lyrics"`
	ArtistURL string `json:"artist_url"`
}

func getTopTrackInRegion(region string) (*TrackInfo, error) {

	track := TrackInfo{
		Name:      "Sample Track",
		Artist:    "Sample Artist",
		ImageURL:  "https://sample.com/image.jpg",
		Lyrics:    "These are sample lyrics.",
		ArtistURL: "https://sample.com/artist",
	}
	return &track, nil
}

func fetchLyrics(track *TrackInfo) error {

	track.Lyrics = "These are sample lyrics fetched from Musixmatch."
	return nil
}

func fetchArtistImage(track *TrackInfo) error {

	track.ImageURL = "https://sample.com/artist_image.jpg"
	return nil
}

func trackHandler(w http.ResponseWriter, r *http.Request) {
	region := r.URL.Query().Get("region")
	if region == "" {
		http.Error(w, "Region parameter is required", http.StatusBadRequest)
		return
	}

	track, err := getTopTrackInRegion(region)
	if err != nil {
		http.Error(w, "Failed to fetch track information", http.StatusInternalServerError)
		return
	}

	// Fetch additional information
	fetchLyrics(track)
	fetchArtistImage(track)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(track)
}

func main() {
	http.HandleFunc("/track", trackHandler)
	fmt.Println("Server is listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
