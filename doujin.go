package nhentai

type RawDoujin struct {
	Id         int             `json:"id"`
	MediaId    string          `json:"media_id"`
	Title      rawDoujinTitle  `json:"title"`
	Images     rawDoujinImages `json:"images"`
	Scanlator  string          `json:"scanlator"`
	UploadDate int             `json:"upload_date"`
	Tags       []rawDoujinTags `json:"tags"`
	TotalPages int             `json:"num_pages"`
	Favourites int             `json:"num_favourites"`
}

type rawDoujinTitle struct {
	English  string `json:"english"`
	Japanese string `json:"japanese"`
	Pretty   string `json:"pretty"`
}

type rawDoujinImages struct {
	Pages []rawDoujinPages `json:"pages"`
}

type rawDoujinPages struct {
	T string `json:"t"`
	W int `json:"w"`
	H int `json:"h"`
}

type rawDoujinTags struct {
	Id    int    `json:"id"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	Url   string `json:"url"`
	Count int    `json:"count"`
}
