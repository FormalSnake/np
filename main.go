package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type Profile struct {
	Track struct {
		Name  string `json:"name"`
		Url   string `json:"url"`
		Image []struct {
			Url  string `json:"#text"`
			Size string `json:"size"`
		} `json:"image"`
		Artist struct {
			Name string `json:"#text"`
		} `json:"artist"`
	} `json:"track"`
}

const PROFILE_ENDPOINT = "https://lastfm-last-played.biancarosa.com.br/"

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		println("Usage: np <username>")
		os.Exit(1)
		return
	}
	username := args[0]

	profile := getProfile(username)
	if profile == nil {
		println("Profile not found")
		os.Exit(1)
		return
	}

	printInfo(profile)
}

func getProfile(query string) *Profile {

	res, err := http.Get(PROFILE_ENDPOINT + query + "/latest-song")
	if err != nil {
		return nil
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var profile Profile
	json.Unmarshal(body, &profile)
	return &profile

}
