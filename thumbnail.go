package discordgoembedwrapper

import (
	"github.com/bwmarrin/discordgo"
)

// Thumbnail wraps the discordgo.MessageEmbedThumbnail type and adds features
type Thumbnail struct {
	*discordgo.MessageEmbedThumbnail
}

// Finalize strips away the extra functions and returns the wrapped type
func (t *Thumbnail) Finalize() *discordgo.MessageEmbedThumbnail {
	return t.MessageEmbedThumbnail
}
