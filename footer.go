package discordgoembedwrapper

import (
	"errors"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Footer wraps the discordgo.MessageEmbedFooter type and adds features
type Footer struct {
	*discordgo.MessageEmbedFooter
	Errors *[]error
}

// Finalize strips away the extra functions and returns the wrapped type
func (f *Footer) Finalize() (*discordgo.MessageEmbedFooter, *[]error) {
	defer func(f *Footer) { f.Errors = nil }(f)
	return f.MessageEmbedFooter, f.Errors
}

/*
addError takes a message string and adds it to the error slice stored in Footer. If the pointer is nil a new error slice
is created. This function takes the same inputs as fmt.Sprintf
*/
func (f *Footer) addError(format string, values ...interface{}) {
	if f.Errors == nil {
		f.Errors = &[]error{}
	}
	*f.Errors = append(*f.Errors, errors.New(fmt.Sprintf(format, values...)))
}

/*
SetIconURL takes an image address string prefixed with https:// / http:// / attachment:// and adds it to the Footer (if
the string does not start with one of these, no URL will be added). It then returns the pointer to the Footer structure
(This function fails silently)
*/
func (f *Footer) SetIconURL(iconUrl string) *Footer {
	if checkValidIconURL(iconUrl) {
		f.IconURL = iconUrl
	} else {
		f.addError(`footer iconUrl '%v' does not start with "http://" | "https://" | "attachment://"`)
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
	} else {
		f.addError(`footer text exceeds 2048 characters: len(val) = %v`, len(val))
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
	} else {
		f.addError(`footer proxyIconUrl '%v' does not start with "http://" | "https://" | "attachment://"`)
	}
	return f
}
