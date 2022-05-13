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

* SearchGalleryURL = Get doujins by tagid
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
	Host      = "https://nhentai.net"
	HostURL   = "https://nhentai.net/g/"
	HostAPI   = "https://nhentai.net/api/"
	ImageURL  = "https://i.nhentai.net/galleries/"
	ThumbsURL = "https://t.nhentai.net/galleries/"

	GalleryURL   = HostAPI + "gallery/"
	GalleriesURL = HostAPI + "galleries/"
)

// Search galleries
const (
	AllGalleryUrl    = GalleriesURL + "all?page="
	TaggedGalleryURL = GalleriesURL + "tagged?tag_id="
	SearchGalleryURL = GalleriesURL + "search?query="
)

// Sorting methods
const (
	Recent          = ""
	PopularAllTime  = "popular"
	PopularThisWeek = "popular-week"
	PopularToday    = "popular-today"
)

var fetchErr = errors.New("error while retrieving doujin")