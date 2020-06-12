package disgobed

import (
	"fmt"

	"github.com/andersfylling/disgord"
)

/*
ThumbnailBuilder wraps the disgord.EmbedThumbnail type and adds features
*/
type ThumbnailBuilder struct {
	*disgord.EmbedThumbnail
	Errors *[]error
}

/*
Finalize strips away the extra functions and returns the wrapped type. It should always be called before an thumbnail is
attached. Finalize will also purge the error cache!
*/
func (t *ThumbnailBuilder) Finalize() (*disgord.EmbedThumbnail, *[]error) {
	defer func(t *ThumbnailBuilder) { t.Errors = nil }(t)
	return t.EmbedThumbnail, t.Errors
}

/*
NewThumbnail creates and returns a pointer to a new empty thumbnail
*/
func NewThumbnail() *ThumbnailBuilder {
	return &ThumbnailBuilder{
		EmbedThumbnail: &disgord.EmbedThumbnail{},
		Errors:         nil,
	}
}

/*
addError takes a message string and adds it to the error slice stored in AuthorBuilder. If the pointer is nil a new error slice
is created. This function takes the same inputs as fmt.Sprintf
*/
func (t *ThumbnailBuilder) addError(format string, values ...interface{}) {
	if t.Errors == nil {
		t.Errors = &[]error{}
	}
	*t.Errors = append(*t.Errors, fmt.Errorf(format, values...))
}

/*
SetURL takes an image address string prefixed with https:// / http:// / attachment:// and adds it to the ThumbnailBuilder (if
the string does not start with one of these, no URL will be added). It then returns the pointer to the ThumbnailBuilder
structure
(This function fails silently)
*/
func (t *ThumbnailBuilder) SetURL(url string) *ThumbnailBuilder {
	if checkValidIconURL(url) {
		t.URL = url
	} else {
		t.addError(invalidUrlErrTemplateString, `thumbnail url`, url)
	}
	return t
}

/*
SetProxyURL takes an image address string prefixed with https:// / http:// / attachment:// and adds it to the ThumbnailBuilder
(if the string does not start with one of these, no URL will be added). It then returns the pointer to the ThumbnailBuilder
structure
(This function fails silently)
*/
func (t *ThumbnailBuilder) SetProxyURL(proxyUrl string) *ThumbnailBuilder {
	if checkValidIconURL(proxyUrl) {
		t.URL = proxyUrl
	} else {
		t.addError(invalidUrlErrTemplateString, `thumbnail proxyUrl`, proxyUrl)
	}
	return t
}

/*
SetHW sets the ThumbnailBuilder embed height and width to the values given then returns a pointer to the ThumbnailBuilder structure.
If either h <= 0 or w <= 0, this operation does nothing
(This function fails silently)
*/
func (t *ThumbnailBuilder) SetHW(h int, w int) *ThumbnailBuilder {
	if h > 0 && w > 0 {
		t.Height = h
		t.Width = w
	} else {
		t.addError(invalidHWErrTemplateString, `thumbnail`, h, `thumbnail`, w)
	}
	return t
}

/*
SetHeight sets the ThumbnailBuilder embed height to the value given then returns a pointer to the ThumbnailBuilder structure. If h <=
0, this operation does nothing
(This function fails silently)
*/
func (t *ThumbnailBuilder) SetHeight(h int) *ThumbnailBuilder {
	if h > 0 {
		t.Height = h
	} else {
		t.addError(valueNotBetweenErrTemplateString, `thumbnail height`, h, 0, `infinity`)
	}
	return t
}

/*
SetWidth sets the ThumbnailBuilder embed width to the value given then returns a pointer to the ThumbnailBuilder structure. If w <= 0,
this operation does nothing
(This function fails silently)
*/
func (t *ThumbnailBuilder) SetWidth(w int) *ThumbnailBuilder {
	if w > 0 {
		t.Width = w
	} else {
		t.addError(valueNotBetweenErrTemplateString, `thumbnail width`, w, 0, `infinity`)
	}
	return t
}
