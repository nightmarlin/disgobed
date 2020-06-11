package discordgoembedwrapper

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
addError takes a message string and adds it to the error slice stored in Author. If the pointer is nil a new error slice
is created. This function takes the same inputs as fmt.Sprintf
*/
func (t *Thumbnail) addError(format string, values ...interface{}) {
	if t.Errors == nil {
		t.Errors = &[]error{}
	}
	*t.Errors = append(*t.Errors, fmt.Errorf(format, values...))
}
