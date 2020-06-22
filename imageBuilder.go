package disgobed

import (
	"fmt"

	"github.com/Nightmarlin/disgobed/validation"
	"github.com/andersfylling/disgord"
)

// ImageBuilder wraps the disgord.EmbedImage type and adds features
type ImageBuilder struct {
	*disgord.EmbedImage
	Errors *[]error
}

/*
Finalize strips away the extra functions and returns the wrapped type. It should always be called before an image is
sent. Finalize will also purge the error cache!
*/
func (i *ImageBuilder) Finalize() (*disgord.EmbedImage, *[]error) {
	defer func(i *ImageBuilder) { i.Errors = nil }(i)
	return i.EmbedImage, i.Errors
}

/*
NewImage creates and returns an empty image structure
*/
func NewImage() *ImageBuilder {
	return &ImageBuilder{
		EmbedImage: &disgord.EmbedImage{},
		Errors:     nil,
	}
}

/*
addError takes a message string and adds it to the error slice stored in AuthorBuilder. If the pointer is nil a new error slice
is created. this function takes the same inputs as fmt.Sprintf
*/
func (i *ImageBuilder) addError(format string, values ...interface{}) {
	if i.Errors == nil {
		i.Errors = &[]error{}
	}
	*i.Errors = append(*i.Errors, fmt.Errorf(format, values...))
}

/*
SetURL takes an image address string prefixed with https:// / http:// / attachment:// and adds it to the ImageBuilder (if
the string does not start with one of these, no URL will be added). It then returns the pointer to the ImageBuilder structure
(This function fails silently)
*/
func (i *ImageBuilder) SetURL(url string) *ImageBuilder {
	if validation.CheckValidIconURL(url) {
		i.URL = url
	} else {
		i.addError(validation.InvalidUrlErrTemplateString, `image url`, url)
	}
	return i
}

/*
SetProxyURL takes an image address string prefixed with https:// / http:// / attachment:// and adds it to the ImageBuilder (if
the string does not start with one of these, no URL will be added). It then returns the pointer to the ImageBuilder structure
(This function fails silently)
*/
func (i *ImageBuilder) SetProxyURL(proxyUrl string) *ImageBuilder {
	if validation.CheckValidIconURL(proxyUrl) {
		i.ProxyURL = proxyUrl
	} else {
		i.addError(validation.InvalidUrlErrTemplateString, `image proxyUrl`, proxyUrl)
	}
	return i
}

/*
SetHW sets the image embed height and width to the values given then returns a pointer to the ImageBuilder structure. If either
h <= 0 or w <= 0, this operation does nothing
(This function fails silently)
*/
func (i *ImageBuilder) SetHW(h int, w int) *ImageBuilder {
	if h > 0 && w > 0 {
		i.Height = h
		i.Width = w
	} else {
		i.addError(validation.InvalidHWErrTemplateString, `image`, h, `image`, w)
	}
	return i
}

/*
SetHeight sets the image embed height to the value given then returns a pointer to the ImageBuilder structure. If h <= 0, this
operation does nothing
(This function fails silently)
*/
func (i *ImageBuilder) SetHeight(h int) *ImageBuilder {
	if h > 0 {
		i.Height = h
	} else {
		i.addError(validation.ValueNotBetweenErrTemplateString, `image height`, h, 0, `infinity`)
	}
	return i
}

/*
SetWidth sets the image embed width to the value given then returns a pointer to the ImageBuilder structure. If w <= 0, this
operation does nothing
(This function fails silently)
*/
func (i *ImageBuilder) SetWidth(w int) *ImageBuilder {
	if w > 0 {
		i.Width = w
	} else {
		i.addError(validation.ValueNotBetweenErrTemplateString, `image width`, w, 0, `infinity`)
	}
	return i
}
