package discordgoembedwrapper

import (
	"github.com/bwmarrin/discordgo"
)

/*
Author wraps the discordgo.MessageEmbedAuthor type and adds features
*/
type Author struct {
	*discordgo.MessageEmbedAuthor
}

/*
Finalize strips away the extra functions and returns the wrapped type
*/
func (a *Author) Finalize() *discordgo.MessageEmbedAuthor {
	return a.MessageEmbedAuthor
}

/*
NewAuthor creates and returns a blank author struct
*/
func NewAuthor() *Author {
	res := &Author{
		MessageEmbedAuthor: &discordgo.MessageEmbedAuthor{},
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
*/
func (a *Author) SetIconURL(iconUrl string) *Author {
	if checkValidIconURL(iconUrl) {
		a.IconURL = iconUrl
	}
	return a
}

/*
SetName takes a string and sets the Author's name to that value. It then returns the pointer to the Author
*/
func (a *Author) SetName(name string) *Author {
	a.Name = name
	return a
}

/*
SetProxyIconURL takes an image address string prefixed with https:// / http:// / attachment:// and adds it to the Author
(if the string does not start with one of these, no URL will be added). It then returns the pointer to the Author
structure
*/
func (a *Author) SetProxyIconURL(proxyIconUrl string) *Author {
	if checkValidIconURL(proxyIconUrl) {
		a.ProxyIconURL = proxyIconUrl
	}
	return a
}
