package thegamesdb

import (
	"testing"
)

func TestGetGamesList(t *testing.T) {
	parameters := map[string]string{
		"name": "Battletoads",
	}
	games, err := GetGamesList(parameters)
	if err != nil {
		t.Fatalf("%s", err.Error())
	}
	println(len(games.Games))
	println(games.Games[0].GameTitle)
	println(games.Games[0].Platform)
}
