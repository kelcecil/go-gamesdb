package thegamesdb

import (
	"testing"
)

func TestGetPlatformGames(t *testing.T) {
	parameters := map[string]string{
		"platform": "1",
	}
	games, err := GetPlatformGames(parameters)
	if err != nil {
		t.Fatalf("%s", err.Error())
	}
	println(games.Games[0].Id)
}

func TestGetPlatformList(t *testing.T) {
	parameters := map[string]string{}
	platforms, err := GetPlatformsList(parameters)
	if err != nil {
		t.Fatalf("%s", err.Error())
	}
	println(platforms.PlatformTag.Platforms[0].Name)
}

func TestGetGamesList(t *testing.T) {
	parameters := map[string]string{
		"name": "Double Dragon",
	}
	games, err := GetGamesList(parameters)
	if err != nil {
		t.Fatalf("%s", err.Error())
	}
	println(len(games.Games))
	println(games.Games[0].GameTitle)
	println(games.Games[0].Platform)
}

func TestGetGame(t *testing.T) {
	parameters := map[string]string{
		"id": "2",
	}
	game, err := GetGame(parameters)
	if err != nil {
		t.Fatalf("%s", err.Error())
	}
	println(game.Game.GameTitle)
}

func TestGetPlatform(t *testing.T) {
	parameters := map[string]string{
		"id": "15",
	}
	platform, err := GetPlatform(parameters)
	if err != nil {
		t.Fatalf("%s", err.Error())
	}
	println(string(platform.Platform.Console))
}
