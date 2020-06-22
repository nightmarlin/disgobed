/*
Package disgobed wraps the discordgo embed with helper functions to facilitate easier construction.
Note that all methods in this module act ByReference, directly changing the embed they are called on, instead of
creating and returning a new embed
*/
package disgobed

import (
	"fmt"
	"time"

	"github.com/Nightmarlin/disgobed/validation"
	"github.com/andersfylling/disgord"
)

/*
EmbedBuilder wraps the disgord.EmbedBuilder type and adds features. Never create it directly, instead use the NewEmbed function

	embed := NewEmbed()

and call the methods to set the properties, allowing for chains that look like this

	embed := NewEmbed()
		.SetTitle(`example`)
		.SetDescription(`test`)
		.SetURL(`example.com`)
		.Finalize()

for healthy embedment!
*/
type EmbedBuilder struct {
	*disgord.Embed
	Errors *[]error
}

/*
ValidateEmbed returns whether or not discord is likely accept the embed attached to it. If discord is unlikely to
accept the embed, it returns a list of reasons why. If msg is not nil, the checker will also validate `attachment://`
urls
*/
func (e *EmbedBuilder) Validate(msg *disgord.Message) *[]error {
	toCheck, errs := e.Finalize()
	if errs != nil { // Make use of builtin err checking
		return errs // Short-Circuit and dont run expensive checks if we already have errors
	}
	return validation.ValidateEmbed(toCheck, msg)
}

/*
Finalize strips away the extra functions and returns the wrapped type. It should always be called before an embed is
sent. Finalize will also purge the error cache!
*/
func (e *EmbedBuilder) Finalize() (*disgord.Embed, *[]error) {
	defer func(e *EmbedBuilder) { e.Errors = nil }(e)
	return e.Embed, e.Errors
}

/*
addError takes a message string and adds it to the error slice stored in EmbedBuilder. If the pointer is nil a new error slice
is created. This function takes the same inputs as fmt.Sprintf
*/
func (e *EmbedBuilder) addError(format string, values ...interface{}) {
	if e.Errors == nil {
		e.Errors = &[]error{}
	}
	*e.Errors = append(*e.Errors, fmt.Errorf(format, values...))
}

/*
addRawError takes a pre-existing error and adds it to the stored slice. If the pointer is nil a new error slice is
created.
*/
func (e *EmbedBuilder) addRawError(err error) {
	if e.Errors == nil {
		e.Errors = &[]error{}
	}
	*e.Errors = append(*e.Errors, err)
}

/*
addAllRawErrors takes a pre-existing error slice and adds it to the stored slice. If the pointer is nil a new error
slice is created.
*/
func (e *EmbedBuilder) addAllRawErrors(errs *[]error) {
	if errs == nil {
		return
	}
	for _, err := range *errs {
		e.addRawError(err)
	}
}

/*
NewEmbed creates and returns an empty embed
*/
func NewEmbed() *EmbedBuilder {
	res := &EmbedBuilder{
		Embed:  &disgord.Embed{},
		Errors: nil,
	}
	return res
}

/*
SetTitle edits the embed's title and returns the pointer to the embed. The discord API limits embed titles to 256
characters, so this function will do nothing if len(title) > 256
(This function fails silently)
*/
func (e *EmbedBuilder) SetTitle(title string) *EmbedBuilder {
	if len(title) <= validation.LowerCharLimit {
		e.Title = title
	} else {
		e.addError(validation.CharacterCountExceedsLimitErrTemplateString, `embed title`, validation.LowerCharLimit, len(title), title)
	}
	return e
}

/*
SetDescription edits the embed's description and returns the pointer to the embed. The discord API limits embed
descriptions to 2048 characters, so this function will do nothing if len(desc) > 2048
(This function fails silently)
*/
func (e *EmbedBuilder) SetDescription(desc string) *EmbedBuilder {
	if len(desc) <= validation.UpperCharLimit {
		e.Description = desc
	} else {
		e.addError(validation.CharacterCountExceedsLimitLongErrTemplateString, `embed description`, validation.UpperCharLimit, len(desc))
	}
	return e
}

/*
SetURL edits the embed's main URL and returns the pointer to the embed
*/
func (e *EmbedBuilder) SetURL(url string) *EmbedBuilder {
	e.URL = url
	return e
}

/*
SetColor edits the embed's highlight colour and returns the pointer to the embed.
Color values must be between 0 and 16777215 otherwise the change will not be registered
(This function fails silently)
*/
func (e *EmbedBuilder) SetColor(color int) *EmbedBuilder {
	if color >= 0 && color < validation.MaxColorValue {
		e.Color = color
	} else {
		e.addError(validation.ValueNotBetweenErrTemplateString, `embed color`, color, 0, validation.MaxColorValue)
	}
	return e
}

/*
SetCurrentTimestamp sets the embed's timestamp to the current UTC time in the appropriate discord format and returns
the pointer to the embed
*/
func (e *EmbedBuilder) SetCurrentTimestamp() *EmbedBuilder {
	utcTime := disgord.Time{Time: time.Now().UTC()}
	return e.setRawTimestamp(utcTime)
}

/*
SetCustomTimestamp sets the embed's timestamp to that specified by the time.Time structure passed to it.
The value stored is the corresponding UTC time in the appropriate discord format.
SetCustomTimestamp returns the pointer to the embed
*/
func (e *EmbedBuilder) SetCustomTimestamp(t time.Time) *EmbedBuilder {
	utcTime := disgord.Time{Time: t.UTC()}
	return e.setRawTimestamp(utcTime)
}

/*
Sets the timestamp string to the argument and returns the pointer to the embed. Was exposed but the potential for error
was too high, so has since been replaced with SetCustomTimestamp(t time.Time)
*/
func (e *EmbedBuilder) setRawTimestamp(timestamp disgord.Time) *EmbedBuilder {
	e.Timestamp = timestamp
	return e
}

/*
InlineAllFields sets the Inline property on all currently attached fields to true and returns the pointer to the embed
*/
func (e *EmbedBuilder) InlineAllFields() *EmbedBuilder {
	for _, f := range e.Fields {
		f.Inline = true
	}
	return e
}

/*
OutlineAllFields sets the Inline property on all currently attached fields to false and returns the pointer to the embed
*/
func (e *EmbedBuilder) OutlineAllFields() *EmbedBuilder {
	for _, f := range e.Fields {
		f.Inline = false
	}
	return e
}

/*
AddFields takes N FieldBuilder structures and adds them to the embed, then returns the pointer to the embed.
Note that FieldBuilder structures are `Finalize`d once added and should not be changed after being added.
The discord API limits embeds to having 25 Fields, so this function will add the first items from the list until that
limit is reached
(This function fails silently)
*/
func (e *EmbedBuilder) AddFields(fields ...*FieldBuilder) *EmbedBuilder {
	for _, f := range fields {
		e.AddField(f)
	}
	return e
}

/*
AddRawFields takes N disgord.EmbedField structures and adds them to the embed, then returns the pointer to the
embed. The discord API limits embeds to having 25 Fields, so this function will add the first items from the list until
that limit is reached
(This function fails silently)
*/
func (e *EmbedBuilder) AddRawFields(fields ...*disgord.EmbedField) *EmbedBuilder {
	for _, f := range fields {
		e.AddRawField(f)
	}
	return e
}

/*
AddField takes a FieldBuilder structure and adds it to the embed, then returns the pointer to the embed.
Note that the FieldBuilder structure is `Finalize`d once added and should not be changed after being added.
The discord API limits embeds to having 25 Fields, so this function will not add any fields if the limit has already
been reached. All errors are propagated to the main embed
(This function fails silently)
*/
func (e *EmbedBuilder) AddField(field *FieldBuilder) *EmbedBuilder {
	res, errs := field.Finalize()
	e.addAllRawErrors(errs)
	return e.AddRawField(res)
}

/*
AddRawField takes a disgord.EmbedField structure and adds it to the embed, then returns the pointer to the
embed. The discord API limits embeds to having 25 Fields, so this function will not add any fields if the limit has
already been reached
(This function fails silently)
*/
func (e *EmbedBuilder) AddRawField(field *disgord.EmbedField) *EmbedBuilder {
	if len(e.Fields) < validation.MaxFieldCount {
		e.Fields = append(e.Fields, field)
	} else {
		e.addError(validation.FieldLimitReachedErrTemplateString, field.Name, validation.MaxFieldCount)
	}
	return e
}

/*
SetAuthor takes an AuthorBuilder structure and sets the embed's author field to it, then returns the pointer to the embed.
Note that the AuthorBuilder structure is `Finalize`d once added and should not be changed after being added. All errors are
propagated to the main embed
*/
func (e *EmbedBuilder) SetAuthor(author *AuthorBuilder) *EmbedBuilder {
	res, errs := author.Finalize()
	e.addAllRawErrors(errs)
	return e.SetRawAuthor(res)
}

/*
SetRawAuthor takes a disgord.EmbedAuthor and sets the embed's author field to it, then returns the pointer to
the embed
*/
func (e *EmbedBuilder) SetRawAuthor(author *disgord.EmbedAuthor) *EmbedBuilder {
	e.Author = author
	return e
}

/*
SetThumbnail takes a ThumbnailBuilder structure and sets the embed's thumbnail field to it, then returns the pointer to the
embed. Note that the ThumbnailBuilder structure is `Finalize`d once added and should not be changed after being added
*/
func (e *EmbedBuilder) SetThumbnail(thumb *ThumbnailBuilder) *EmbedBuilder {
	res, errs := thumb.Finalize()
	e.addAllRawErrors(errs)
	return e.SetRawThumbnail(res)
}

/*
SetRawThumbnail takes a disgord.EmbedThumbnail and sets the embed's thumbnail field to it, then returns the
pointer to the embed
*/
func (e *EmbedBuilder) SetRawThumbnail(thumb *disgord.EmbedThumbnail) *EmbedBuilder {
	e.Thumbnail = thumb
	return e
}

/*
SetProvider allows you to set the provider of an embed. It will then return the pointer to the embed.
See the providerBuilder.go docs for some extra information
*/
func (e *EmbedBuilder) SetProvider(provider *ProviderBuilder) *EmbedBuilder {
	res, errs := provider.Finalize()
	if errs != nil { // This should never run
		e.addAllRawErrors(errs)
	}
	return e.SetRawProvider(res)
}

/*
SetRawProvider allows you to set the disgord.EmbedProvider of an embed.
It will then return the pointer to the embed.
See the providerBuilder.go docs for some extra information
*/
func (e *EmbedBuilder) SetRawProvider(provider *disgord.EmbedProvider) *EmbedBuilder {
	e.Provider = provider
	return e
}

/*
SetFooter sets the embed's footer property to the FooterBuilder passed to it, then returns the pointer to the embed.
Note that the FooterBuilder structure is `Finalize`d once added and should not be changed after being added. FooterBuilder errors
will be propagated into the embed struct
*/
func (e *EmbedBuilder) SetFooter(footer *FooterBuilder) *EmbedBuilder {
	res, errs := footer.Finalize()
	e.addAllRawErrors(errs)
	return e.SetRawFooter(res)
}

/*
SetRawFooter takes a disgord.EmbedThumbnail and sets the embed's thumbnail field to it, then returns the
pointer to the embed
*/
func (e *EmbedBuilder) SetRawFooter(footer *disgord.EmbedFooter) *EmbedBuilder {
	e.Footer = footer
	return e
}

/*
SetVideo sets the embed's video property to the VideoBuilder passed to it, then returns the pointer to the embed.
Note that the VideoBuilder structure is `Finalize`d once added and should not be changed after being added
*/
func (e *EmbedBuilder) SetVideo(vid *VideoBuilder) *EmbedBuilder {
	res, errs := vid.Finalize()
	e.addAllRawErrors(errs)
	return e.SetRawVideo(res)
}

/*
SetRawVideo takes a disgord.EmbedVideo and sets the embed's thumbnail field to it, then returns the pointer to
the embed
*/
func (e *EmbedBuilder) SetRawVideo(vid *disgord.EmbedVideo) *EmbedBuilder {
	e.Video = vid
	return e
}

/*
SetImage sets the embed's image property to the ImageBuilder passed to it, then returns the pointer to the embed.
Note that the ImageBuilder structure is `Finalize`d once added and should not be changed after being added. ImageBuilder errors
will be propagated into the embed struct
*/
func (e *EmbedBuilder) SetImage(img *ImageBuilder) *EmbedBuilder {
	res, errs := img.Finalize()
	e.addAllRawErrors(errs)
	return e.SetRawImage(res)
}

/*
SetRawImage takes a disgord.EmbedImage and sets the embed's image field to it, then returns the pointer to the
embed
*/
func (e *EmbedBuilder) SetRawImage(img *disgord.EmbedImage) *EmbedBuilder {
	e.Image = img
	return e
}

/*
SetType checks if the embed type passed to it is valid. If it is, it sets the embed's type to that, otherwise it does
nothing. It then returns the pointer to the embed
(This function fails silently)
*/
func (e *EmbedBuilder) SetType(embedType string) *EmbedBuilder {
	if validation.CheckTypeValid(embedType) {
		e.Type = embedType
	} else {
		e.addError(validation.InvalidEmbedTypeErrTemplateString, embedType)
	}
	return e
}
