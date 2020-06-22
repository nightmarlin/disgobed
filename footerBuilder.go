package disgobed

import (
	"fmt"

	"github.com/Nightmarlin/disgobed/validation"
	"github.com/andersfylling/disgord"
)

/*
FooterBuilder wraps the disgord.EmbedFooter type and adds features
*/
type FooterBuilder struct {
	*disgord.EmbedFooter
	Errors *[]error
}

/*
Finalize strips away the extra functions and returns the wrapped type. It should always be called before a footer is
attached. Finalize will also purge the error cache!
*/
func (f *FooterBuilder) Finalize() (*disgord.EmbedFooter, *[]error) {
	defer func(f *FooterBuilder) { f.Errors = nil }(f)
	return f.EmbedFooter, f.Errors
}

/*
NewFooter creates and returns a new empty footer
*/
func NewFooter() *FooterBuilder {
	return &FooterBuilder{
		EmbedFooter: &disgord.EmbedFooter{},
		Errors:      nil,
	}
}

/*
addError takes a message string and adds it to the error slice stored in FooterBuilder. If the pointer is nil a new error slice
is created. This function takes the same inputs as fmt.Sprintf
*/
func (f *FooterBuilder) addError(format string, values ...interface{}) {
	if f.Errors == nil {
		f.Errors = &[]error{}
	}
	*f.Errors = append(*f.Errors, fmt.Errorf(format, values...))
}

/*
SetIconURL takes an image address string prefixed with https:// / http:// / attachment:// and adds it to the FooterBuilder (if
the string does not start with one of these, no URL will be added). It then returns the pointer to the FooterBuilder structure
(This function fails silently)
*/
func (f *FooterBuilder) SetIconURL(iconUrl string) *FooterBuilder {
	if validation.CheckValidIconURL(iconUrl) {
		f.IconURL = iconUrl
	} else {
		f.addError(validation.InvalidUrlErrTemplateString, `footer iconUrl`, iconUrl)
	}
	return f
}

/*
SetText takes a string and sets the FooterBuilder's text to that value. It then returns the pointer to the FooterBuilder. The discord
API limits FooterBuilder values to 2048 characters, so this function will do nothing if len(val) > 2048
(This function fails silently)
*/
func (f *FooterBuilder) SetText(val string) *FooterBuilder {
	if len(val) <= validation.UpperCharLimit {
		f.Text = val
	} else {
		f.addError(validation.CharacterCountExceedsLimitLongErrTemplateString, `footer text`, validation.UpperCharLimit, len(val))
	}
	return f
}

/*
SetProxyIconURL takes an image address string prefixed with https:// / http:// / attachment:// and adds it to the FooterBuilder
(if the string does not start with one of these, no URL will be added). It then returns the pointer to the FooterBuilder
structure
(This function fails silently)
*/
func (f *FooterBuilder) SetProxyIconURL(proxyIconUrl string) *FooterBuilder {
	if validation.CheckValidIconURL(proxyIconUrl) {
		f.ProxyIconURL = proxyIconUrl
	} else {
		f.addError(validation.InvalidUrlErrTemplateString, `footer proxyIconUrl`, proxyIconUrl)
	}
	return f
}
