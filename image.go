package discordgoembedwrapper

import (
	"errors"
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
addError takes a message string and adds it to the error slice stored in Author. If the pointer is nil a new error slice
is created. this function takes the same inputs as fmt.Sprintf
*/
func (i *Image) addError(format string, values ...interface{}) {
	if i.Errors == nil {
		i.Errors = &[]error{}
	}
	*i.Errors = append(*i.Errors, errors.New(fmt.Sprintf(format, values...)))
}
