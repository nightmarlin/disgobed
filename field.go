package discordgoembedwrapper

import (
	"github.com/bwmarrin/discordgo"
)

/*
Field wraps the discordgo.MessageEmbedField type and adds features
*/
type Field struct {
	*discordgo.MessageEmbedField
}

/*
Finalize strips away the extra functions and returns the wrapped type
*/
func (f *Field) Finalize() *discordgo.MessageEmbedField {
	return f.MessageEmbedField
}

/*
SetInline sets whether the field is inline or not then returns the pointer to the Field
*/
func (f *Field) SetInline(isInline bool) *Field {
	f.Inline = isInline
	return f
}

/*
SetName sets the name of the field then returns the pointer to the Field. The discord API limits Field names to 256 characters, so this function will do
nothing if len(name) > 256
(This function fails silently)
*/
func (f *Field) SetName(name string) *Field {
	if len(name) <= 256 {
		f.Name = name
	}
	return f
}

/*
SetValue sets the value of the field then returns the pointer to the Field. The discord API limits Field values to 1024
characters, so this function will do nothing if len(name) > 1024
(This function fails silently)
*/
func (f *Field) SetValue(val string) *Field {
	if len(val) <= 1024 {
		f.Value = val
	}
	return f
}
