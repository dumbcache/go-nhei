package nhentai

import "time"

type Doujin struct {
	ID              int            `json:"id"`
	MediaID         string         `json:"mediaId"`
	Titles          DoujinTitle    `json:"title"`
	URL             string         `json:"url"`
	Cover           string         `json:"cover"`
	Thumb           string         `json:"thumb"`
	Pages           []DoujinPage   `json:"pages"`
	UploadDate      time.Time      `json:"uploadDate"`
	UploadTimeStamp int            `json:"timeStamp"`
	Tags            []RawDoujinTag `json:"tags"`
	PageCount       int            `json:"pageCount"`
	Favourites      int            `json:"favourites"`
}

type DoujinTitle struct {
	English  string `json:"english"`
	Japanese string `json:"japanese"`
	Pretty   string `json:"pretty"`
}

type DoujinPage struct {
	Extension string `json:"t"`
	Width     int    `json:"w"`
	Height    int    `json:"h"`
	URL       string `json:"u"`
	Number    int    `json:"n"`
}
