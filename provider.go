package discordgoembedwrapper

import (
	"github.com/bwmarrin/discordgo"
)

/*
Provider wraps the discordgo.MessageEmbedProvider type and adds features.
Provider is an esoteric part of the discord API, and is likely to be deprecated in a future version. It is recommended
you don't use it... Use at your own risk. No Provider fields are validated
*/
type Provider struct {
	*discordgo.MessageEmbedProvider
	Errors *[]error
}

/*
Finalize strips away the extra functions and returns the wrapped type. Provider does not perform validation on inputs,
so Finalize() should always return nil for its errors
*/
func (p *Provider) Finalize() (*discordgo.MessageEmbedProvider, *[]error) {
	defer func(p *Provider) { p.Errors = nil }(p)
	return p.MessageEmbedProvider, p.Errors
}

/*
NewProvider creates and returns a pointer to an empty provider struct
*/
func NewProvider() *Provider {
	return &Provider{
		MessageEmbedProvider: &discordgo.MessageEmbedProvider{},
		Errors:               nil,
	}
}

/*
SetURL sets the provider's URL field and returns the pointer to the Provider
*/
func (p *Provider) SetURL(url string) *Provider {
	p.URL = url
	return p
}

/*
SetURL sets the provider's Name field and returns the pointer to the Provider
*/
func (p *Provider) SetName(name string) *Provider {
	p.Name = name
	return p
}
