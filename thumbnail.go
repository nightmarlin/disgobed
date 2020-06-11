package disgobed

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

/*
Thumbnail wraps the discordgo.MessageEmbedThumbnail type and adds features
*/
type Thumbnail struct {
	*discordgo.MessageEmbedThumbnail
	Errors *[]error
}

/*
Finalize strips away the extra functions and returns the wrapped type. It should always be called before an thumbnail is
attached. Finalize will also purge the error cache!
*/
func (t *Thumbnail) Finalize() (*discordgo.MessageEmbedThumbnail, *[]error) {
	defer func(t *Thumbnail) { t.Errors = nil }(t)
	return t.MessageEmbedThumbnail, t.Errors
}

/*
NewThumbnail creates and returns a pointer to a new empty thumbnail
*/
func NewThumbnail() *Thumbnail {
	return &Thumbnail{
		MessageEmbedThumbnail: &discordgo.MessageEmbedThumbnail{},
		Errors:                nil,
	}
}

/*
addError takes a message string and adds it to the error slice stored in Author. If the pointer is nil a new error slice
is created. This function takes the same inputs as fmt.Sprintf
*/
func (t *Thumbnail) addError(format string, values ...interface{}) {
	if t.Errors == nil {
		t.Errors = &[]error{}
	}
	*t.Errors = append(*t.Errors, fmt.Errorf(format, values...))
}

/*
SetURL takes an image address string prefixed with https:// / http:// / attachment:// and adds it to the Thumbnail (if
the string does not start with one of these, no URL will be added). It then returns the pointer to the Thumbnail
structure
(This function fails silently)
*/
func (t *Thumbnail) SetURL(url string) *Thumbnail {
	if checkValidIconURL(url) {
		t.URL = url
	} else {
		t.addError(`thumbnail url '%v' does not start with "http://" | "https://" | "attachment://"`, url)
	}
	return t
}

/*
SetProxyURL takes an image address string prefixed with https:// / http:// / attachment:// and adds it to the Thumbnail
(if the string does not start with one of these, no URL will be added). It then returns the pointer to the Thumbnail
structure
(This function fails silently)
*/
func (t *Thumbnail) SetProxyURL(proxyUrl string) *Thumbnail {
	if checkValidIconURL(proxyUrl) {
		t.URL = proxyUrl
	} else {
		t.addError(`thumbnail proxyUrl '%v' does not start with "http://" | "https://" | "attachment://"`, proxyUrl)
	}
	return t
}

/*
SetHW sets the Thumbnail embed height and width to the values given then returns a pointer to the Thumbnail structure.
If either h <= 0 or w <= 0, this operation does nothing
(This function fails silently)
*/
func (t *Thumbnail) SetHW(h int, w int) *Thumbnail {
	if h > 0 && w > 0 {
		t.Height = h
		t.Width = w
	} else {
		t.addError(`thumbnail height '%t' or thumbnail width '%t' is less than or equal to 0`, h, w)
	}
	return t
}

/*
SetHeight sets the Thumbnail embed height to the value given then returns a pointer to the Thumbnail structure. If h <=
0, this operation does nothing
(This function fails silently)
*/
func (t *Thumbnail) SetHeight(h int) *Thumbnail {
	if h > 0 {
		t.Height = h
	} else {
		t.addError(`thumbnail height '%t' is less than or equal to 0`, h)
	}
	return t
}

/*
SetWidth sets the Thumbnail embed width to the value given then returns a pointer to the Thumbnail structure. If w <= 0,
this operation does nothing
(This function fails silently)
*/
func (t *Thumbnail) SetWidth(w int) *Thumbnail {
	if w > 0 {
		t.Width = w
	} else {
		t.addError(`thumbnail width '%t' is less than or equal to 0`, w)
	}
	return t
}
