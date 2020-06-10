package discordgoembedwrapper

import (
	"github.com/bwmarrin/discordgo"
)

/*
Video wraps the discordgo.MessageEmbedVideo type and adds features. This wrapper ignores MessageEmbedVideo.ProxyURL as
the API would ignore that field if present
*/
type Video struct {
	*discordgo.MessageEmbedVideo
}

// Finalize strips away the extra functions and returns the wrapped type
func (v *Video) Finalize() *discordgo.MessageEmbedVideo {
	return v.MessageEmbedVideo
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
	}
	return v
}
