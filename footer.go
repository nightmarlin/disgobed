package embeds

import (
	"github.com/bwmarrin/discordgo"
)

// Footer wraps the discordgo.MessageEmbedFooter type and adds features
type Footer struct {
	*discordgo.MessageEmbedFooter
}

// Finalize strips away the extra functions and returns the wrapped type
func (f *Footer) Finalize() *discordgo.MessageEmbedFooter {
	return f.MessageEmbedFooter
}
