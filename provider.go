package discordgoembedwrapper

import (
	"github.com/bwmarrin/discordgo"
)

/*
Provider wraps the discordgo.MessageEmbedProvider type and adds features.
Provider is an esoteric part of the discord API, and is likely to be deprecated in a future version. It is recommended
you don't use it... As such it does not have any of the standard wrappings... Use at your own risk
*/
type Provider struct {
	*discordgo.MessageEmbedProvider
}

/*
Finalize strips away the extra functions and returns the wrapped type
*/
func (p *Provider) Finalize() *discordgo.MessageEmbedProvider {
	return p.MessageEmbedProvider
}
