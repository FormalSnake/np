package main

import (
	"encoding/json"
	"image"
	"io"
	"net/http"
	"os"

	"github.com/dolmen-go/kittyimg"
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
	img := getImage(profile.Track.Image[0].Url)
	if img == nil {
		println("Image not found")
		return
	}
	kittyimg.Fprintln(os.Stdout, img)

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

func getImage(url string) image.Image {
	res, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer res.Body.Close()
	img, _, err := image.Decode(res.Body)
	if err != nil {
		return nil
	}
	return img
}
