package discordgoembedwrapper

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

/*
SetIconURL takes an image address string prefixed with https:// / http:// / attachment:// and adds it to the Footer (if
the string does not start with one of these, no URL will be added). It then returns the pointer to the Footer structure
(This function fails silently)
*/
func (f *Footer) SetIconURL(iconUrl string) *Footer {
	if checkValidIconURL(iconUrl) {
		f.IconURL = iconUrl
	}
	return f
}

/*
SetText takes a string and sets the Footer's text to that value. It then returns the pointer to the Footer. The discord
API limits Footer values to 2048 characters, so this function will do nothing if len(val) > 2048
(This function fails silently)
*/
func (f *Footer) SetText(val string) *Footer {
	if len(val) <= 2048 {
		f.Text = val
	}
	return f
}

/*
SetProxyIconURL takes an image address string prefixed with https:// / http:// / attachment:// and adds it to the Footer
(if the string does not start with one of these, no URL will be added). It then returns the pointer to the Footer
structure
(This function fails silently)
*/
func (f *Footer) SetProxyIconURL(proxyIconUrl string) *Footer {
	if checkValidIconURL(proxyIconUrl) {
		f.ProxyIconURL = proxyIconUrl
	}
	return f
}
