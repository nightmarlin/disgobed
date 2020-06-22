package disgobed

import (
	"fmt"

	"github.com/Nightmarlin/disgobed/validation"
	"github.com/andersfylling/disgord"
)

/*
AuthorBuilder wraps the discordgo.MessageEmbedAuthor type and adds features
*/
type AuthorBuilder struct {
	*disgord.EmbedAuthor
	Errors *[]error
}

/*
Finalize strips away the extra functions and returns the wrapped type. It should always be called before an author is
sent. Finalize will also purge the error cache!
*/
func (a *AuthorBuilder) Finalize() (*disgord.EmbedAuthor, *[]error) {
	defer func(a *AuthorBuilder) { a.Errors = nil }(a)
	return a.EmbedAuthor, a.Errors
}

/*
addError takes a message string and adds it to the error slice stored in AuthorBuilder. If the pointer is nil a new error slice
is created. This function takes the same inputs as fmt.Sprintf
*/
func (a *AuthorBuilder) addError(format string, values ...interface{}) {
	if a.Errors == nil {
		a.Errors = &[]error{}
	}
	*a.Errors = append(*a.Errors, fmt.Errorf(format, values...))
}

/*
NewAuthor creates and returns a blank author struct
*/
func NewAuthor() *AuthorBuilder {
	res := &AuthorBuilder{
		EmbedAuthor: &disgord.EmbedAuthor{},
		Errors:      nil,
	}
	return res
}

/*
SetURL sets the author field link to the value given, then returns the pointer to the AuthorBuilder
*/
func (a *AuthorBuilder) SetURL(url string) *AuthorBuilder {
	a.URL = url
	return a
}

/*
SetIconURL takes an image address string prefixed with https:// / http:// / attachment:// and adds it to the AuthorBuilder (if
the string does not start with one of these, no URL will be added). It then returns the pointer to the AuthorBuilder structure
(This function fails silently)
*/
func (a *AuthorBuilder) SetIconURL(iconUrl string) *AuthorBuilder {
	if validation.CheckValidIconURL(iconUrl) {
		a.IconURL = iconUrl
	} else {
		a.addError(validation.InvalidUrlErrTemplateString, `author iconUrl`, iconUrl)
	}
	return a
}

/*
SetName takes a string and sets the AuthorBuilder's name to that value. It then returns the pointer to the AuthorBuilder. The discord
API limits AuthorBuilder names to 256 characters, so this function will do nothing if len(name) > 256
(This function fails silently)
*/
func (a *AuthorBuilder) SetName(name string) *AuthorBuilder {
	var limit = 256
	if len(name) <= limit {
		a.Name = name
	} else {
		a.addError(validation.CharacterCountExceedsLimitErrTemplateString, `author name`, limit, len(name), name)
	}
	return a
}

/*
SetProxyIconURL takes an image address string prefixed with https:// / http:// / attachment:// and adds it to the AuthorBuilder
(if the string does not start with one of these, no URL will be added). It then returns the pointer to the AuthorBuilder
structure
(This function fails silently)
*/
func (a *AuthorBuilder) SetProxyIconURL(proxyIconUrl string) *AuthorBuilder {
	if validation.CheckValidIconURL(proxyIconUrl) {
		a.ProxyIconURL = proxyIconUrl
	} else {
		a.addError(validation.InvalidUrlErrTemplateString, `author proxyIconUrl`, proxyIconUrl)
	}
	return a
}
