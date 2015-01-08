package thegamesdb

import (
	"strings"
)

var apiEndpoint string = "http://thegamesdb.net/api/"

type GamesList struct {
	Games []GameEntity `xml:"Game"`
}

type GetGameResponse struct {
	Game GameEntity `xml:"Game"`
}

type GameEntity struct {
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

type PlatformList struct {
	PlatformTag Platforms `xml:"Platforms"`
}

type GetPlatformResponse struct {
	Platform PlatformEntity `xml:"Platform"`
}

type Platforms struct {
	Platforms []ConcisePlatformEntity `xml:"Platform"`
}

type ConcisePlatformEntity struct {
	Id    string `xml:"id"`
	Name  string `xml:"name"`
	Alias string `xml:"alias"`
}

type PlatformEntity struct {
	Id             string `xml:"id"`
	Platform       string `xml:"Platform"`
	Console        string `xml:"console"`
	Controller     string `xml:"controller"`
	Overview       string `xml:"overview"`
	Developer      string `xml:"developer"`
	Manufacturer   string `xml:"manufacturer"`
	CPU            string `xml:"cpu"`
	Memory         string `xml:"memory"`
	Graphics       string `xml:"graphics"`
	Sound          string `xml:"sound"`
	Display        string `xml:"display"`
	Media          string `xml:"media"`
	MaxControllers string `xml:"maxcontrollers"`
	Ratings        string `xml:"Rating"`
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
		if value != "" {
			params = append(params, key+"="+value)
		}
	}
	return strings.Join(params, "&")
}

//go:generate generategamesdb -output get_games_list.go -method GetGamesList -type GamesList -endpoint GetGamesList.php
//go:generate generategamesdb -output get_platforms_list.go -method GetPlatformsList -type PlatformList -endpoint GetPlatformsList.php
//go:generate generategamesdb -output get_game.go -method GetGame -type GetGameResponse -endpoint GetGame.php
//go:generate generategamesdb -output get_platform.go -method GetPlatform -type GetPlatformResponse -endpoint GetPlatform.php
