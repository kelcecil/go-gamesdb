package thegamesdb

import (
	"strings"
)

var apiEndpoint string = "http://thegamesdb.net/api/"

type GamesList struct {
	Games []Game `xml:"Game"`
}

type Game struct {
	Id          string `xml:"id"`
	GameTitle   string `xml:"GameTitle"`
	ReleaseDate string `xml:"ReleaseDate"`
	Platform    string `xml:"Platform"`
	Overview    string `xml:"Overview"`
	//Add genre
	Players string   `xml:"Players"`
	Coop    string   `xml:"Co-op"`
	ESRB    string   `xml:"ESRB"`
	YouTube string   `xml:"Youtube"`
	Fanarts []Fanart `xml:"fanart"`
	Boxarts []Boxart `xml:"boxart"`
}

type Fanart struct {
	Original Image `xml:"original"`
	Thumb    Image `xml:"thumb"`
}

type Image struct {
	Width  int    `xml:"width,attr"`
	Height int    `xml:"height,attr"`
	Path   string `xml:",chardata"`
}

type Boxart struct {
	Side          string `xml:"side,attr"`
	Width         int    `xml:"width,attr"`
	Height        int    `xml:"height,attr"`
	ThumbnailPath string `xml:thumb,attr"`
	ImagePath     string `xml:",chardata"`
}

func ConvertMapIntoGetParams(parameters map[string]string) string {
	params := make([]string, 0)
	for key, value := range parameters {
		if  value != "" {
			params = append(params, key+"="+value)
		}
	}
	return strings.Join(params, "&")
}

//go:generate generategamesdb -output get_games_list.go -method GetGamesList -type GamesList -endpoint GetGamesList.php
//go:generate generategamesdb -output get_game.go -method GetGame -type Game -endpoint GetGame.php
