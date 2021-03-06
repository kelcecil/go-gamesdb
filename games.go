package thegamesdb

//go:generate generategamesdb -output get_game.go -method GetGame -type GetGameResponse -endpoint GetGame.php
//go:generate generategamesdb -output get_games_list.go -method GetGamesList -type GetGamesListResponse -endpoint GetGamesList.php

type GetGamesListResponse struct {
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
