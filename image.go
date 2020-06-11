package discordgoembedwrapper

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Image wraps the discordgo.MessageEmbedImage type and adds features
type Image struct {
	*discordgo.MessageEmbedImage
	Errors *[]error
}

/*
Finalize strips away the extra functions and returns the wrapped type. It should always be called before an image is
sent. Finalize will also purge the error cache!
*/
func (i *Image) Finalize() (*discordgo.MessageEmbedImage, *[]error) {
	defer func(i *Image) { i.Errors = nil }(i)
	return i.MessageEmbedImage, i.Errors
}

/*
NewImage creates and returns an empty image structure
*/
func NewImage() *Image {
	return &Image{
		MessageEmbedImage: &discordgo.MessageEmbedImage{},
		Errors:            nil,
	}
}

/*
addError takes a message string and adds it to the error slice stored in Author. If the pointer is nil a new error slice
is created. this function takes the same inputs as fmt.Sprintf
*/
func (i *Image) addError(format string, values ...interface{}) {
	if i.Errors == nil {
		i.Errors = &[]error{}
	}
	*i.Errors = append(*i.Errors, fmt.Errorf(format, values...))
}

func (i *Image) SetURL(url string) *Image {
	i.URL = url
	return i
}

/*
SetHW sets the image embed height and width to the values given then returns a pointer to the Image structure. If either
h <= 0 or w <= 0, this operation does nothing
(This function fails silently)
*/
func (i *Image) SetHW(h int, w int) *Image {
	if h > 0 && w > 0 {
		i.Height = h
		i.Width = w
	} else {
		i.addError(`image height '%i' or video width '%i' is less than or equal to 0`, h, w)
	}
	return i
}

/*
SetHeight sets the image embed height to the value given then returns a pointer to the Image structure. If h <= 0, this
operation does nothing
(This function fails silently)
*/
func (i *Image) SetHeight(h int) *Image {
	if h > 0 {
		i.Height = h
	} else {
		i.addError(`image height '%i' is less than or equal to 0`, h)
	}
	return i
}

/*
SetWidth sets the image embed width to the value given then returns a pointer to the Image structure. If w <= 0, this
operation does nothing
(This function fails silently)
*/
func (i *Image) SetWidth(w int) *Image {
	if w > 0 {
		i.Width = w
	} else {
		i.addError(`image width '%i' is less than or equal to 0`, w)
	}
	return i
}
