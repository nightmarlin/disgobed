package discordgoembedwrapper

import (
	"github.com/bwmarrin/discordgo"
)

// Field wraps the discordgo.MessageEmbedField type and adds features
type Field struct {
	*discordgo.MessageEmbedField
}

// Finalize strips away the extra functions and returns the wrapped type
func (f *Field) Finalize() *discordgo.MessageEmbedField {
	return f.MessageEmbedField
}
