package thegamesdb

//go:generate generategamesdb -output get_art.go -method GetArt -type GetArtResponse -endpoint GetArt.php

type GetArtResponse struct {
	Images ImagesEntity `xml:"Images"`
}

type ImagesEntity struct {
	Fanarts     []Fanart `xml:"fanart"`
	Boxarts     []Boxart `xml:"boxart"`
	Banners     []Image  `xml:"banner"`
	Screenshots []Fanart `xml:"screenshot"`
}
