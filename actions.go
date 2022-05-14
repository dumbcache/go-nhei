package nhentai

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func fetch(url string, raw interface{}) error {

	log.Println(url)

	res, err := http.Get(url)
	if err != nil {
		return fetchErr
	}
	if err := unmarshal(raw, res); err != nil {
		return err
	}
	return nil
}

// used to get the individual doujin by doujin id
func FetchDoujin(id int) (*Doujin, error) {

	fetchURL := fmt.Sprintf("%s%d", GalleryURL, id)
	raw := new(RawDoujin)
	d := new(Doujin)
	err := fetch(fetchURL, raw)
	d.transform(raw)
	return d, err
}

func RelatedDoujin(id int) ([]Doujin, error) {
	fetchURL := fmt.Sprintf("%s%d/related", GalleryURL, id)
	raw := new(RawDoujinList)
	err := fetch(fetchURL, raw)
	dlist := raw.transform()
	return dlist, err
}

// fetch the homepage, nothing but the recent 25 doujins. you can pass page number to get respective page.
//
// page number should be 1,2,......so on
func HomePage(page int) ([]Doujin, error) {

	fetchURL := fmt.Sprintf("%s%d", AllGalleryURL, page)
	raw := new(RawDoujinList)
	err := fetch(fetchURL, raw)
	dlist := raw.transform()
	return dlist, err

}
func FetchPopular(page int) ([]Doujin, error) {

	fetchURL := fmt.Sprintf("%s*&page=%d&sort=%s", SearchGalleryURL, page, PopularAllTime)
	raw := new(RawDoujinList)
	err := fetch(fetchURL, raw)
	dlist := raw.transform()
	return dlist, err

}



func FetchRandom() (*Doujin, error) {

	client := new(http.Client)
	req, err := http.NewRequest(http.MethodHead, RandomURL, nil)
	if err != nil {
		return nil, ErrFormat("inside fetchRandom", err)
	}
	res, err := client.Do(req)

	if err != nil {
		return nil, ErrFormat("inside fetchRandom", err)
	}
	newUrl := strings.Trim(res.Request.URL.Path, "g/")
	id, _ := strconv.Atoi(newUrl)
	d, err := FetchDoujin(id)
	return d, err
}

// get the doujin by a string value provided.
// you can pass page number to get respective page.
//
// page number should be 1,2,......so on
//
// you can pass below sort options to sort accordingly
//Recent          = "" (default)
//PopularAllTime  = "popular"
//PopularThisWeek = "popular-week"
//PopularToday    = "popular-today"
func Search(query string, page int, sort string) ([]Doujin, error) {

	fetchURL := fmt.Sprintf("%s%s&page=%d&%s", SearchGalleryURL, query, page, sort)
	raw := new(RawDoujinList)
	err := fetch(fetchURL, raw)
	dlist := raw.transform()
	return dlist, err

}

func SearchByTagID(id int, page int, sort string) ([]Doujin, error) {

	fetchURL := fmt.Sprintf("%s%d&page=%d&%s", TaggedGalleryURL, id, page, sort)
	raw := new(RawDoujinList)
	err := fetch(fetchURL, raw)
	dlist := raw.transform()
	return dlist, err

}

// converting raw doujin format to Doujin format
func (d *Doujin) transform(raw *RawDoujin) {

	d.transformImages(raw.Images, raw.MediaID)
	d.transformTags(raw.Tags)

	switch raw.ID.(type) {
	case float64:
		d.ID = int(raw.ID.(float64))
	case string:
		d.ID, _ = strconv.Atoi(raw.ID.(string))
	}
	d.MediaID = raw.MediaID
	d.Titles = DoujinTitle(raw.Title)
	d.URL = fmt.Sprintf("%s%d", HostURL, d.ID)
	d.UploadDate = time.Unix(int64(raw.UploadDate), 0)
	d.UploadTimeStamp = raw.UploadDate
	d.PageCount = raw.PageCount
	d.Favourites = raw.Favourites
}

func (raw *RawDoujinList) transform() []Doujin {
	dlist := []Doujin{}
	for _, v := range raw.Result {
		d := new(Doujin)
		d.transform(&v)
		dlist = append(dlist, *d)
	}
	return dlist
}

func unmarshal(i interface{}, res *http.Response) error {

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return ErrFormat("inside unmarshal", err)
	}
	defer res.Body.Close()
	file, _ := os.Create("data.json")
	file.Write(data)
	json.Unmarshal(data, i)
	return nil
}

func (d *Doujin) transformImages(images rawDoujinImage, mediaID string) {

	for i, v := range images.Pages {
		page := new(DoujinPage)
		page.Extension = imgExtension(v.T)
		page.Height = v.H
		page.Width = v.W
		page.Number = i + 1
		page.URL = fmt.Sprintf("%s%s/%d.%s", ImageURL, mediaID, i+1, page.Extension)
		d.Pages = append(d.Pages, *page)
	}
	d.Cover = fmt.Sprintf("%s%s/%s%s", ThumbsURL, mediaID, "cover.", imgExtension(images.Cover.T))
	d.Thumb = fmt.Sprintf("%s%s/%s%s", ThumbsURL, mediaID, "thumb.", imgExtension(images.Thumb.T))
}

func (d *Doujin) transformTags(tags []RawDoujinTag) {

	for _, v := range tags {
		tag := new(RawDoujinTag)
		tag.ID = v.ID
		tag.Type = v.Type
		tag.Name = v.Name
		tag.Count = v.Count
		tag.URL = Host + v.URL
		d.Tags = append(d.Tags, *tag)
	}
}

func imgExtension(s string) string {

	switch s {
	case "j":
		return "jpg"
	case "p":
		return "png"
	case "g":
		return "gif"
	}
	return ""
}

func ErrFormat(head string, body error) error {
	return fmt.Errorf("%s\n---->\t%s%s%s: %w...%s\n--------------------------------", Green, Red, head, Yellow, body, Reset)
}

func (d *Doujin) FilterTags(name string) []RawDoujinTag {
	filtered := []RawDoujinTag{}
	for _, v := range d.Tags {
		if v.Type == name {
			filtered = append(filtered, v)
		}
	}
	return filtered
}
