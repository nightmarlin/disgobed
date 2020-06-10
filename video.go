package discordgoembedwrapper

import (
	"github.com/bwmarrin/discordgo"
)

// Video wraps the discordgo.MessageEmbedVideo type and adds features
type Video struct {
	*discordgo.MessageEmbedVideo
}

// Finalize strips away the extra functions and returns the wrapped type
func (v *Video) Finalize() *discordgo.MessageEmbedVideo {
	return v.MessageEmbedVideo
}
