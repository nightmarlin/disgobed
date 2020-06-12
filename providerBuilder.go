package disgobed

import (
	"github.com/andersfylling/disgord"
)

/*
ProviderBuilder wraps the disgord.EmbedProvider type and adds features.
ProviderBuilder is an esoteric part of the discord API, and is likely to be deprecated in a future version. It is recommended
you don't use it... Use at your own risk. No ProviderBuilder fields are validated
*/
type ProviderBuilder struct {
	*disgord.EmbedProvider
	Errors *[]error
}

/*
Finalize strips away the extra functions and returns the wrapped type. ProviderBuilder does not perform validation on inputs,
so Finalize() should always return nil for its errors
*/
func (p *ProviderBuilder) Finalize() (*disgord.EmbedProvider, *[]error) {
	defer func(p *ProviderBuilder) { p.Errors = nil }(p)
	return p.EmbedProvider, p.Errors
}

/*
NewProvider creates and returns a pointer to an empty provider struct
*/
func NewProvider() *ProviderBuilder {
	return &ProviderBuilder{
		EmbedProvider: &disgord.EmbedProvider{},
		Errors:        nil,
	}
}

/*
SetURL sets the provider's URL field and returns the pointer to the ProviderBuilder
*/
func (p *ProviderBuilder) SetURL(url string) *ProviderBuilder {
	p.URL = url
	return p
}

/*
SetName sets the provider's Name field and returns the pointer to the ProviderBuilder
*/
func (p *ProviderBuilder) SetName(name string) *ProviderBuilder {
	p.Name = name
	return p
}
