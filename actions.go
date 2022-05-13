package nhentai

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

// used to get the individual doujin by doujin id
func FetchDoujin(id int) (*Doujin, error) {

	fetchURL := fmt.Sprintf("%s%d", GalleryURL, id)
	raw := new(RawDoujin)
	d := new(Doujin)

	res, err := http.Get(fetchURL)
	if err != nil {
		return nil, fetchErr
	}
	if err := unmarshal(raw, res); err != nil {
		return nil, err
	}
	d.transform(raw)

	return d, nil
}

// fetch the homepage, nothing but the recent 25 doujins. you can pass page number to get respective page.
//
// page number should be 1,2,......so on
func HomePage(page int) ([]Doujin,error){
	
	raw := new(RawDoujinList)
	dlist := []Doujin{}
	fetchURL := fmt.Sprintf("%s%d",AllGalleryUrl,page)
	res, err := http.Get(fetchURL)
	if err != nil {
		return nil, fetchErr
	}
	if err := unmarshal(raw, res); err != nil {
		return nil, err
	}
	
	for _, v := range raw.Result {
		d := new(Doujin)
		d.transform(&v)
		dlist =append(dlist,*d)
	}

	return dlist, nil
}

// converting raw doujin format to Doujin format
func (d *Doujin) transform(raw *RawDoujin) {

	d.transformImages(raw.Images, raw.MediaID)
	d.transformTags(raw.Tags)

	switch raw.ID.(type) {
	case float64:
		d.ID = int(raw.ID.(float64))
	case string:
		d.ID,_ = strconv.Atoi(raw.ID.(string))
	}
	d.MediaID = raw.MediaID
	d.Titles = DoujinTitle(raw.Title)
	d.URL = fmt.Sprintf("%s%d", HostURL, raw.ID)
	d.UploadDate = time.Unix(int64(raw.UploadDate), 0)
	d.UploadTimeStamp = raw.UploadDate
	d.PageCount = raw.PageCount
	d.Favourites = raw.Favourites
}

func unmarshal(i interface{}, res *http.Response) error {

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.New("error while unmarshallling")
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
