package embeds

import (
	"github.com/bwmarrin/discordgo"
)

// Image wraps the discordgo.MessageEmbedImage type and adds features
type Image struct {
	*discordgo.MessageEmbedImage
}

// Finalize strips away the extra functions and returns the wrapped type
func (i *Image) Finalize() *discordgo.MessageEmbedImage {
	return i.MessageEmbedImage
}
