package disgobed

import (
	"fmt"

	"github.com/Nightmarlin/disgobed/validation"
	"github.com/andersfylling/disgord"
)

/*
VideoBuilder wraps the disgord.EmbedVideo type and adds features. This wrapper ignores MessageEmbedVideo.ProxyURL as
the API would ignore that field if present
*/
type VideoBuilder struct {
	*disgord.EmbedVideo
	Errors *[]error
}

/*
Finalize strips away the extra functions and returns the wrapped type. It should always be called before an thumbnail is
attached. Finalize will also purge the error cache!
*/
func (v *VideoBuilder) Finalize() (*disgord.EmbedVideo, *[]error) {
	defer func(v *VideoBuilder) { v.Errors = nil }(v)
	return v.EmbedVideo, v.Errors
}

/*
addError takes a message string and adds it to the error slice stored in AuthorBuilder. If the pointer is nil a new error slice
is created. This function takes the same inputs as fmt.Sprintf
*/
func (v *VideoBuilder) addError(format string, values ...interface{}) {
	if v.Errors == nil {
		v.Errors = &[]error{}
	}
	*v.Errors = append(*v.Errors, fmt.Errorf(format, values...))
}

/*
SetURL sets the video source url to the value given to it then returns a pointer to the VideoBuilder structure
*/
func (v *VideoBuilder) SetURL(url string) *VideoBuilder {
	v.URL = url
	return v
}

/*
SetHW sets the video embed height and width to the values given then returns a pointer to the VideoBuilder structure. If either
h <= 0 or w <= 0, this operation does nothing
(This function fails silently)
*/
func (v *VideoBuilder) SetHW(h int, w int) *VideoBuilder {
	if h > 0 && w > 0 {
		v.Height = h
		v.Width = w
	} else {
		v.addError(validation.InvalidHWErrTemplateString, `video`, h, `video`, w)
	}
	return v
}

/*
SetHeight sets the video embed height to the value given then returns a pointer to the VideoBuilder structure. If h <= 0, this
operation does nothing
(This function fails silently)
*/
func (v *VideoBuilder) SetHeight(h int) *VideoBuilder {
	if h > 0 {
		v.Height = h
	} else {
		v.addError(validation.ValueNotBetweenErrTemplateString, `video height`, h, 0, `infinity`)
	}
	return v
}

/*
SetWidth sets the video embed width to the value given then returns a pointer to the VideoBuilder structure. If w <= 0, this
operation does nothing
(This function fails silently)
*/
func (v *VideoBuilder) SetWidth(w int) *VideoBuilder {
	if w > 0 {
		v.Width = w
	} else {
		v.addError(validation.ValueNotBetweenErrTemplateString, `video width`, w, 0, `infinity`)
	}
	return v
}
