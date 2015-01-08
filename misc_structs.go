package thegamesdb

var apiEndpoint string = "http://thegamesdb.net/api/"

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