package discordgoembedwrapper

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

/*
Author wraps the discordgo.MessageEmbedAuthor type and adds features
*/
type Author struct {
	*discordgo.MessageEmbedAuthor
	Errors *[]error
}

/*
Finalize strips away the extra functions and returns the wrapped type. It should always be called before an author is
sent. Finalize will also purge the error cache!
*/
func (a *Author) Finalize() (*discordgo.MessageEmbedAuthor, *[]error) {
	defer func(a *Author) { a.Errors = nil }(a)
	return a.MessageEmbedAuthor, a.Errors
}

/*
addError takes a message string and adds it to the error slice stored in Author. If the pointer is nil a new error slice
is created. This function takes the same inputs as fmt.Sprintf
*/
func (a *Author) addError(format string, values ...interface{}) {
	if a.Errors == nil {
		a.Errors = &[]error{}
	}
	*a.Errors = append(*a.Errors, fmt.Errorf(format, values...))
}

/*
NewAuthor creates and returns a blank author struct
*/
func NewAuthor() *Author {
	res := &Author{
		MessageEmbedAuthor: &discordgo.MessageEmbedAuthor{},
		Errors:             nil,
	}
	return res
}

/*
SetURL sets the author field link to the value given, then returns the pointer to the Author
*/
func (a *Author) SetURL(url string) *Author {
	a.URL = url
	return a
}

/*
SetIconURL takes an image address string prefixed with https:// / http:// / attachment:// and adds it to the Author (if
the string does not start with one of these, no URL will be added). It then returns the pointer to the Author structure
(This function fails silently)
*/
func (a *Author) SetIconURL(iconUrl string) *Author {
	if checkValidIconURL(iconUrl) {
		a.IconURL = iconUrl
	} else {
		a.addError(`author iconUrl '%v' does not start with "http://" | "https://" | "attachment://"`,
			iconUrl)
	}
	return a
}

/*
SetName takes a string and sets the Author's name to that value. It then returns the pointer to the Author. The discord
API limits Author names to 256 characters, so this function will do nothing if len(name) > 256
(This function fails silently)
*/
func (a *Author) SetName(name string) *Author {
	if len(name) <= 256 {
		a.Name = name
	} else {
		a.addError(`author name exceeds 256 characters: len(name) = %v | '%v'`, len(name), name)
	}
	return a
}

/*
SetProxyIconURL takes an image address string prefixed with https:// / http:// / attachment:// and adds it to the Author
(if the string does not start with one of these, no URL will be added). It then returns the pointer to the Author
structure
(This function fails silently)
*/
func (a *Author) SetProxyIconURL(proxyIconUrl string) *Author {
	if checkValidIconURL(proxyIconUrl) {
		a.ProxyIconURL = proxyIconUrl
	} else {
		a.addError(`author proxyIconUrl '%v' does not start with "http://" | "https://" | "attachment://"`,
			proxyIconUrl)
	}
	return a
}
