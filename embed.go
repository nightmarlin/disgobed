// Package discordgoembedwrapper Wraps the discordgo embed with helper functions to facilitate easier construction
package discordgoembedwrapper

import (
	"errors"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

/*
Embed wraps the discordgo.MessageEmbed type and adds features. Never create it directly, instead use the NewEmbed function

	embed := NewEmbed()

and call the methods to set the properties, allowing for chains that look like this

	embed := NewEmbed()
		.SetTitle(`example`)
		.SetDescription(`test`)
		.SetURL(`example.com`)
		.Finalize()

for healthy embedment!

Note that all methods in this module act ByReference, and as such change the embed they are called on, instead of
creating and returning a new embed
*/
type Embed struct {
	*discordgo.MessageEmbed
	Errors *[]error
}

/*
Finalize strips away the extra functions and returns the wrapped type. It should always be called before an embed is
sent. Finalize will also purge the error cache!
*/
func (e *Embed) Finalize() (*discordgo.MessageEmbed, *[]error) {
	defer func(e *Embed) { e.Errors = nil }(e)
	return e.MessageEmbed, e.Errors
}

/*
addError takes a message string and adds it to the error slice stored in Embed. If the pointer is nil a new error slice
is created. this function takes the same inputs as fmt.Sprintf
*/
func (e *Embed) addError(format string, values ...interface{}) {
	if e.Errors == nil {
		e.Errors = &[]error{}
	}
	*e.Errors = append(*e.Errors, errors.New(fmt.Sprintf(format, values...)))
}

/*
addRawError takes a pre-existing error and adds it to the stored slice. If the pointer is nil a new error slice is
created.
*/
func (e *Embed) addRawError(err error) {
	if e.Errors == nil {
		e.Errors = &[]error{}
	}
	*e.Errors = append(*e.Errors, err)
}

/*
NewEmbed creates and returns an empty embed
*/
func NewEmbed() *Embed {
	res := &Embed{
		MessageEmbed: &discordgo.MessageEmbed{},
		Errors:       nil,
	}
	return res
}

/*
SetTitle edits the embed's title and returns the pointer to the embed. The discord API limits embed titles to 256
characters, so this function will do nothing if len(title) > 256
(This function fails silently)
*/
func (e *Embed) SetTitle(title string) *Embed {
	if len(title) <= 256 {
		e.Title = title
	} else {
		e.addError(`embed title exceeds 256 characters: len(title) = %v | %v`, len(title), title)
	}
	return e
}

/*
SetDescription edits the embed's description and returns the pointer to the embed. The discord API limits embed
descriptions to 2048 characters, so this function will do nothing if len(desc) > 2048
(This function fails silently)
*/
func (e *Embed) SetDescription(desc string) *Embed {
	if len(desc) <= 2048 {
		e.Description = desc
	} else {
		e.addError(`embed description exceeds 2048 characters: len(desc) = %v`, len(desc))
	}
	return e
}

/*
SetURL edits the embed's main URL and returns the pointer to the embed
*/
func (e *Embed) SetURL(url string) *Embed {
	e.URL = url
	return e
}

/*
SetColor edits the embed's highlight colour and returns the pointer to the embed.
Color values must be between 0 and 16777215 otherwise the change will not be registered
(This function fails silently)
*/
func (e *Embed) SetColor(color int) *Embed {
	if color >= 0 && color < 16777215 {
		e.Color = color
	} else {
		e.addError(`color '%v' is not between 0 and 16777215`, color)
	}
	return e
}

/*
SetCurrentTimestamp sets the embed's timestamp to the current UTC time in the appropriate discord format and returns
the pointer to the embed
*/
func (e *Embed) SetCurrentTimestamp() *Embed {
	utcTime := time.Now().UTC().Format(`2006-01-02T15:04:05.000Z`)
	return e.setRawTimestamp(utcTime)
}

/*
SetCustomTimestamp sets the embed's timestamp to that specified by the time.Time structure passed to it.
The value stored is the corresponding UTC time in the appropriate discord format.
SetCustomTimestamp returns the pointer to the embed
*/
func (e *Embed) SetCustomTimestamp(t time.Time) *Embed {
	utcTime := t.UTC().Format(`2006-01-02T15:04:05.000Z`)
	return e.setRawTimestamp(utcTime)
}

/*
Sets the timestamp string to the argument and returns the pointer to the embed. Was exposed but the potential for error
was too high, so has since been replaced with SetCustomTimestamp(t time.Time)
*/
func (e *Embed) setRawTimestamp(timestamp string) *Embed {
	e.Timestamp = timestamp
	return e
}

/*
InlineAllFields sets the Inline property on all currently attached fields to true and returns the pointer to the embed
*/
func (e *Embed) InlineAllFields() *Embed {
	for _, f := range e.Fields {
		f.Inline = true
	}
	return e
}

/*
OutlineAllFields sets the Inline property on all currently attached fields to false and returns the pointer to the embed
*/
func (e *Embed) OutlineAllFields() *Embed {
	for _, f := range e.Fields {
		f.Inline = false
	}
	return e
}

/*
AddFields takes N Field structures and adds them to the embed, then returns the pointer to the embed.
Note that Field structures are `Finalize`d once added and should not be changed after being added.
The discord API limits embeds to having 25 Fields, so this function will add the first items from the list until that
limit is reached
(This function fails silently)
*/
func (e *Embed) AddFields(fields ...*Field) *Embed {
	for _, f := range fields {
		e.AddField(f)
	}
	return e
}

/*
AddRawFields takes N discordgo.MessageEmbedField structures and adds them to the embed, then returns the pointer to the
embed. The discord API limits embeds to having 25 Fields, so this function will add the first items from the list until
that limit is reached
(This function fails silently)
*/
func (e *Embed) AddRawFields(fields ...*discordgo.MessageEmbedField) *Embed {
	for _, f := range fields {
		e.AddRawField(f)
	}
	return e
}

/*
AddField takes a Field structure and adds it to the embed, then returns the pointer to the embed.
Note that the Field structure is `Finalize`d once added and should not be changed after being added.
The discord API limits embeds to having 25 Fields, so this function will not add any fields if the limit has already
been reached. All errors are propagated to the main embed
(This function fails silently)
*/
func (e *Embed) AddField(field *Field) *Embed {
	res, errs := field.Finalize()
	if errs != nil {
		for _, err := range *errs {
			e.addRawError(err)
		}
	}
	return e.AddRawField(res)
}

/*
AddRawField takes a discordgo.MessageEmbedField structure and adds it to the embed, then returns the pointer to the
embed. The discord API limits embeds to having 25 Fields, so this function will not add any fields if the limit has
already been reached
(This function fails silently)
*/
func (e *Embed) AddRawField(field *discordgo.MessageEmbedField) *Embed {
	if len(e.Fields) < 25 {
		e.Fields = append(e.Fields, field)
	} else {
		e.addError(`adding field '%v' would cause field count to exceed 25`, field.Name)
	}
	return e
}

/*
SetAuthor takes an Author structure and sets the embed's author field to it, then returns the pointer to the embed.
Note that the Author structure is `Finalize`d once added and should not be changed after being added. All errors are
propagated to the main embed
*/
func (e *Embed) SetAuthor(author *Author) *Embed {
	res, errs := author.Finalize()
	if errs != nil {
		for _, err := range *errs {
			e.addRawError(err)
		}
	}

	return e.SetRawAuthor(res)
}

/*
SetRawAuthor takes a discordgo.MessageEmbedAuthor and sets the embed's author field to it, then returns the pointer to
the embed
*/
func (e *Embed) SetRawAuthor(author *discordgo.MessageEmbedAuthor) *Embed {
	e.Author = author
	return e
}

/*
SetThumbnail takes a Thumbnail structure and sets the embed's thumbnail field to it, then returns the pointer to the
embed. Note that the Thumbnail structure is `Finalize`d once added and should not be changed after being added
*/
func (e *Embed) SetThumbnail(thumb *Thumbnail) *Embed {
	return e.SetRawThumbnail(thumb.Finalize())
}

/*
SetRawThumbnail takes a discordgo.MessageEmbedThumbnail and sets the embed's thumbnail field to it, then returns the
pointer to the embed
*/
func (e *Embed) SetRawThumbnail(thumb *discordgo.MessageEmbedThumbnail) *Embed {
	e.Thumbnail = thumb
	return e
}

/*
SetProvider allows you to set the provider of an embed. It will then return the pointer to the embed.
See the provider.go docs for some extra information
*/
func (e *Embed) SetProvider(provider *Provider) *Embed {
	return e.SetRawProvider(provider.Finalize())
}

/*
SetRawProvider allows you to set the discordgo.MessageEmbedProvider of an embed.
It will then return the pointer to the embed.
See the provider.go docs for some extra information
*/
func (e *Embed) SetRawProvider(provider *discordgo.MessageEmbedProvider) *Embed {
	e.Provider = provider
	return e
}

/*
SetFooter sets the embed's footer property to the Footer passed to it, then returns the pointer to the embed.
Note that the Footer structure is `Finalize`d once added and should not be changed after being added
*/
func (e *Embed) SetFooter(footer *Footer) *Embed {
	return e.SetRawFooter(footer.Finalize())
}

/*
SetRawFooter takes a discordgo.MessageEmbedThumbnail and sets the embed's thumbnail field to it, then returns the
pointer to the embed
*/
func (e *Embed) SetRawFooter(footer *discordgo.MessageEmbedFooter) *Embed {
	e.Footer = footer
	return e
}

/*
SetVideo sets the embed's video property to the Video passed to it, then returns the pointer to the embed.
Note that the Video structure is `Finalize`d once added and should not be changed after being added
*/
func (e *Embed) SetVideo(vid *Video) *Embed {
	return e.SetRawVideo(vid.Finalize())
}

/*
SetRawVideo takes a discordgo.MessageEmbedVideo and sets the embed's thumbnail field to it, then returns the pointer to
the embed
*/
func (e *Embed) SetRawVideo(vid *discordgo.MessageEmbedVideo) *Embed {
	e.Video = vid
	return e
}

/*
SetImage sets the embed's image property to the Image passed to it, then returns the pointer to the embed.
Note that the Image structure is `Finalize`d once added and should not be changed after being added
*/
func (e *Embed) SetImage(img *Image) *Embed {
	return e.SetRawImage(img.Finalize())
}

/*
SetRawImage takes a discordgo.MessageEmbedImage and sets the embed's image field to it, then returns the pointer to the
embed
*/
func (e *Embed) SetRawImage(img *discordgo.MessageEmbedImage) *Embed {
	e.Image = img
	return e
}

/*
SetType checks if the embed type passed to it is valid. If it is, it sets the embed's type to that, otherwise it does
nothing. It then returns the pointer to the embed
(This function fails silently)
*/
func (e *Embed) SetType(embedType string) *Embed {
	if checkTypeValid(embedType) {
		e.Type = embedType
	} else {
		e.addError(`embed type '%v' is not one of "rich" | "image" | "video" | "gifv" | "link" | "article"`,
			embedType)
	}
	return e
}
