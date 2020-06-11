package discordgoembedwrapper

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

/*
Field wraps the discordgo.MessageEmbedField type and adds features
*/
type Field struct {
	*discordgo.MessageEmbedField
	Errors *[]error
}

/*
Finalize strips away the extra functions and returns the wrapped type. It should always be called before a field is
added. Finalize will also purge the error cache!
*/
func (f *Field) Finalize() (*discordgo.MessageEmbedField, *[]error) {
	defer func(f *Field) { f.Errors = nil }(f)
	return f.MessageEmbedField, f.Errors
}

/*
NewField creates and returns a new empty field object
*/
func NewField() *Field {
	return &Field{
		MessageEmbedField: &discordgo.MessageEmbedField{},
		Errors:            nil,
	}
}

/*
addError takes a message string and adds it to the error slice stored in Field. If the pointer is nil a new error slice
is created. This function takes the same inputs as fmt.Sprintf
*/
func (f *Field) addError(format string, values ...interface{}) {
	if f.Errors == nil {
		f.Errors = &[]error{}
	}
	*f.Errors = append(*f.Errors, fmt.Errorf(format, values...))
}

/*
SetInline sets whether the field is inline or not then returns the pointer to the Field
*/
func (f *Field) SetInline(isInline bool) *Field {
	f.Inline = isInline
	return f
}

/*
SetName sets the name of the field then returns the pointer to the Field. The discord API limits Field names to 256
characters, so this function will do nothing if len(name) > 256
(This function fails silently)
*/
func (f *Field) SetName(name string) *Field {
	if len(name) <= 256 {
		f.Name = name
	} else {
		f.addError(`field name exceeds 256 characters: len(name) = %v | '%v'`, len(name), name)
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
	} else {
		f.addError(`field value exceeds 1024 characters: len(val) = %v`, len(val))
	}
	return f
}
