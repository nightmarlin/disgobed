package disgobed

import (
	"fmt"

	"github.com/Nightmarlin/disgobed/validation"
	"github.com/andersfylling/disgord"
)

/*
FieldBuilder wraps the disgord.EmbedField type and adds features
*/
type FieldBuilder struct {
	*disgord.EmbedField
	Errors *[]error
}

/*
Finalize strips away the extra functions and returns the wrapped type. It should always be called before a field is
added. Finalize will also purge the error cache!
*/
func (f *FieldBuilder) Finalize() (*disgord.EmbedField, *[]error) {
	defer func(f *FieldBuilder) { f.Errors = nil }(f)
	return f.EmbedField, f.Errors
}

/*
NewField creates and returns a new empty field object
*/
func NewField() *FieldBuilder {
	return &FieldBuilder{
		EmbedField: &disgord.EmbedField{},
		Errors:     nil,
	}
}

/*
addError takes a message string and adds it to the error slice stored in FieldBuilder. If the pointer is nil a new error slice
is created. This function takes the same inputs as fmt.Sprintf
*/
func (f *FieldBuilder) addError(format string, values ...interface{}) {
	if f.Errors == nil {
		f.Errors = &[]error{}
	}
	*f.Errors = append(*f.Errors, fmt.Errorf(format, values...))
}

/*
SetInline sets whether the field is inline or not then returns the pointer to the FieldBuilder
*/
func (f *FieldBuilder) SetInline(isInline bool) *FieldBuilder {
	f.Inline = isInline
	return f
}

/*
SetName sets the name of the field then returns the pointer to the FieldBuilder. The discord API limits FieldBuilder names to 256
characters, so this function will do nothing if len(name) > 256. FieldBuilder names must also not be empty, so this function
will do nothing if name == ``
(This function fails silently)
*/
func (f *FieldBuilder) SetName(name string) *FieldBuilder {
	if len(name) <= validation.LowerCharLimit {
		if name == `` {
			f.addError(validation.ValueIsEmptyErrString, `field name`)
		} else {
			f.Name = name
		}
	} else {
		f.addError(validation.CharacterCountExceedsLimitErrTemplateString, `field name`, validation.LowerCharLimit, len(name), name)
	}
	return f
}

/*
SetValue sets the value of the field then returns the pointer to the FieldBuilder. The discord API limits FieldBuilder values to 1024
characters, so this function will do nothing if len(name) > 1024. FieldBuilder values must not be empty, so this function will
do nothing if val == ``
(This function fails silently)
*/
func (f *FieldBuilder) SetValue(val string) *FieldBuilder {
	if len(val) <= validation.MiddleCharLimit {
		if val == `` {
			f.addError(validation.ValueIsEmptyErrString, `field value`)
		} else {
			f.Value = val
		}
	} else {
		f.addError(validation.CharacterCountExceedsLimitLongErrTemplateString, `field value`, validation.MiddleCharLimit, len(val))
	}
	return f
}
