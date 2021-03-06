package nhentai

import "errors"

/*

*
* for random doujin
* EndpointRandomDoujin = https://nhentai.net/random


* GalleryURL = Get doujin gallery by id
------------------------------------------------
* GalleryURL + "{id}"
* GalleryURL + "{id}" + "/related"
* GalleryURL + "{id}" + "/comments"

* AllGalleryUrl = Get all doujins
------------------------------------------------
* GAllGalleryUrl
* GalleryURL + "&page=2"

* SearchGalleryURL = Get doujins by query
------------------------------------------------
* SearchGalleryURL + "konosuba"
* SearchGalleryURL + "konosuba&page=3&sort=popular"

* TaggedGalleryURL = Get doujins by tagid
------------------------------------------------
* TaggedGalleryURL + "2"
* TaggedGalleryURL + "2&page=3&sort=popular"

* ThumbsURL = url to get cover and thumb pics of doujins
*-----------------------------------------------
* ThumbsURL + {mediaid} + "cover.jpg"
* ThumbsURL + 1234 + "thumb.jpg"

* ImageURL = url to get pages of doujins
*-----------------------------------------------
* ImageURL + {mediaid} + "/1.jpg"  (mediaid is different from doujinid)
* imageURL + "12324/1.jpg"

* image formats
*------------------------------------------------------
* p = "png"
* j = "jpg"
* g = "gif"
*/

// Host endpoints
const (
	Host      = "https://nhentai.net/"
	HostURL   = "https://nhentai.net/g/"
	HostAPI   = "https://nhentai.net/api/"
	ImageURL  = "https://i.nhentai.net/galleries/"
	ThumbsURL = "https://t.nhentai.net/galleries/"

	GalleryURL   = HostAPI + "gallery/"
	GalleriesURL = HostAPI + "galleries/"
)

// Search galleries
const (
	AllGalleryURL    = GalleriesURL + "all?page="
	TaggedGalleryURL = GalleriesURL + "tagged?tag_id="
	SearchGalleryURL = GalleriesURL + "search?query="
	RandomURL        = Host + "random"
)

const (
	PopularAllTime  = "sort=popular"
	PopularThisWeek = "sort=popular-week"
	PopularToday    = "sort=popular-today"
)

var TagTypes = []string{"artist", "category", "character", "group", "language", "parody", "tag"}

var fetchErr = errors.New("error while retrieving doujin")

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
