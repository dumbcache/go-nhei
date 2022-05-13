package nhentai

type RawDoujinList struct{
	Result []RawDoujin	`json:"result"`
}

type RawDoujin struct {
	ID         int            `json:"id"`
	MediaID    string         `json:"media_id"`
	Title      rawDoujinTitle `json:"title"`
	Images     rawDoujinImage `json:"images"`
	Scanlator  string         `json:"scanlator"`
	UploadDate int            `json:"upload_date"`
	Tags       []RawDoujinTag `json:"tags"`
	PageCount  int            `json:"num_pages"`
	Favourites int            `json:"num_favorites"`
}

type rawDoujinTitle struct {
	English  string `json:"english"`
	Japanese string `json:"japanese"`
	Pretty   string `json:"pretty"`
}

type rawDoujinImage struct {
	Pages []rawDoujinPage `json:"pages"`
	Cover rawDoujinPage   `json:"cover"`
	Thumb rawDoujinPage   `json:"thumbnail"`
}

type rawDoujinPage struct {
	T string `json:"t"`
	W int    `json:"w"`
	H int    `json:"h"`
}

type RawDoujinTag struct {
	ID    int    `json:"id"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	URL   string `json:"url"`
	Count int    `json:"count"`
}
