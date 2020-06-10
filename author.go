package embeds

import (
	"github.com/bwmarrin/discordgo"
)

// Author wraps the discordgo.MessageEmbedAuthor type and adds features
type Author struct {
	*discordgo.MessageEmbedAuthor
}

// Finalize strips away the extra functions and returns the wrapped type
func (a *Author) Finalize() *discordgo.MessageEmbedAuthor {
	return a.MessageEmbedAuthor
}

// NewAuthor creates and returns a blank author struct
func NewAuthor() *Author {
	res := &Author{
		MessageEmbedAuthor: &discordgo.MessageEmbedAuthor{},
	}
	return res
}

func (a *Author) SetURL(url string) *Author {
	a.URL = url
	return a
}

func (a *Author) SetIconURL(iconUrl string) *Author {
	a.IconURL = iconUrl
	return a
}

func (a *Author) SetName(name string) *Author {
	a.Name = name
	return a
}

func (a *Author) SetProxyIconURL(proxyIconUrl string) *Author {
	a.ProxyIconURL = proxyIconUrl
	return a
}
