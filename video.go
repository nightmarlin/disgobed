package disgobed

import (
	"fmt"

	"github.com/andersfylling/disgord"
)

/*
Video wraps the disgord.EmbedVideo type and adds features. This wrapper ignores MessageEmbedVideo.ProxyURL as
the API would ignore that field if present
*/
type Video struct {
	*disgord.EmbedVideo
	Errors *[]error
}

/*
Finalize strips away the extra functions and returns the wrapped type. It should always be called before an thumbnail is
attached. Finalize will also purge the error cache!
*/
func (v *Video) Finalize() (*disgord.EmbedVideo, *[]error) {
	defer func(v *Video) { v.Errors = nil }(v)
	return v.EmbedVideo, v.Errors
}

/*
addError takes a message string and adds it to the error slice stored in Author. If the pointer is nil a new error slice
is created. This function takes the same inputs as fmt.Sprintf
*/
func (v *Video) addError(format string, values ...interface{}) {
	if v.Errors == nil {
		v.Errors = &[]error{}
	}
	*v.Errors = append(*v.Errors, fmt.Errorf(format, values...))
}

/*
SetURL sets the video source url to the value given to it then returns a pointer to the Video structure
*/
func (v *Video) SetURL(url string) *Video {
	v.URL = url
	return v
}

/*
SetHW sets the video embed height and width to the values given then returns a pointer to the Video structure. If either
h <= 0 or w <= 0, this operation does nothing
(This function fails silently)
*/
func (v *Video) SetHW(h int, w int) *Video {
	if h > 0 && w > 0 {
		v.Height = h
		v.Width = w
	} else {
		v.addError(invalidHWErrTemplateString, `video`, h, `video`, w)
	}
	return v
}

/*
SetHeight sets the video embed height to the value given then returns a pointer to the Video structure. If h <= 0, this
operation does nothing
(This function fails silently)
*/
func (v *Video) SetHeight(h int) *Video {
	if h > 0 {
		v.Height = h
	} else {
		v.addError(valueNotBetweenErrTemplateString, `video height`, h, 0, `infinity`)
	}
	return v
}

/*
SetWidth sets the video embed width to the value given then returns a pointer to the Video structure. If w <= 0, this
operation does nothing
(This function fails silently)
*/
func (v *Video) SetWidth(w int) *Video {
	if w > 0 {
		v.Width = w
	} else {
		v.addError(valueNotBetweenErrTemplateString, `video width`, w, 0, `infinity`)
	}
	return v
}
