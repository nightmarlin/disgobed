package discordgoembedwrapper

import (
	"github.com/bwmarrin/discordgo"
)

// Provider wraps the discordgo.MessageEmbedProvider type and adds features
type Provider struct {
	*discordgo.MessageEmbedProvider
}

// Finalize strips away the extra functions and returns the wrapped type
func (p *Provider) Finalize() *discordgo.MessageEmbedProvider {
	return p.MessageEmbedProvider
}
