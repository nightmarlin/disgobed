package validation

import (
	"strings"
)

// CheckValidIconURL checks that discord will accept the given url in the restricted fields
func CheckValidIconURL(url string) bool {
	for _, pfx := range acceptablePrefixes {
		if strings.HasPrefix(url, pfx) {
			return true
		}
	}
	return false
}

var (
	// A list of the prefixes specific discord URLs will accept
	acceptablePrefixes = [3]string{
		`https://`, // Prefer `https://`
		`attachment://`,
		`http://`,
	}
)

const (
	// LowerCharLimit is the lowest embed property character limit
	LowerCharLimit = 256

	// MiddleCharLimit is the middle embed property character limit
	MiddleCharLimit = 1024

	// UpperCharLimit is the upper embed property character limit
	UpperCharLimit = 2048

	// MaxTotalCharLimit is the maximum total number of characters in an embed
	MaxTotalCharLimit = 6000

	// MaxFieldCount is the maximum number of embed fields
	MaxFieldCount = 25

	// MaxColorValue is the largest acceptable colour value
	MaxColorValue = 16777215
)

const (
	// InvalidUrlErrTemplateString : [Type Property] '[Value]' does not start with "http://" | "https://" | "attachment://"
	InvalidUrlErrTemplateString = `%v '%v' does not start with "http://" | "https://" | "attachment://"`

	// InvalidHWErrTemplateString :  [Type] height '[Value]' or [Type] width '[Value]' is less than or equal to 0
	InvalidHWErrTemplateString = `%v height '%v' or %v width '%v' is less than or equal to 0`

	// CharacterCountExceedsLimitErrTemplateString : [Type Property] exceeds [Limit]: length = [Length] | '[Value]'
	CharacterCountExceedsLimitErrTemplateString = `%v exceeds %v characters: length = %v | '%v'`

	// CharacterCountExceedsLimitLongErrTemplateString : [Type Property] exceeds [Limit]: length = [Length]
	CharacterCountExceedsLimitLongErrTemplateString = `%v exceeds %v characters: length = %v`

	// FieldLimitReachedErrTemplateString : adding field '[FieldName]' would cause field count to exceed [Limit]
	FieldLimitReachedErrTemplateString = `adding field '%v' would cause field count to exceed %v`

	// InvalidEmbedTypeErrTemplateString : embed type '[Type]' is not one of "rich" | "image" | "video" | "gifv" | "link" | "article"
	InvalidEmbedTypeErrTemplateString = `embed type '%v' is not one of "rich" | "image" | "video" | "gifv" | "link" | "article"`

	// ValueNotBetweenErrTemplateString : [Type Property] [Value] is not between [LowerLimit] and [UpperLimit]
	ValueNotBetweenErrTemplateString = `%v '%v' is not between %v and %v`

	// ValueIsEmptyErrString : [Type Property] should not be empty if set
	ValueIsEmptyErrString = `%v should not be empty if set`
)

const (
	// RichEmbedType describes a rich embed - generally ignored by clients
	RichEmbedType = `rich`

	// ImageEmbedType describes an image embed - generally ignored by clients
	ImageEmbedType = `image`

	// VideoEmbedType describes a video embed - generally ignored by clients
	VideoEmbedType = `video`

	// GifvEmbedType describes a gifv embed - generally ignored by clients
	GifvEmbedType = `gifv`

	// LinkEmbedType describes a link embed - generally ignored by clients
	LinkEmbedType = `link`

	// ArticleEmbedType describes an article embed - generally ignored by clients
	ArticleEmbedType = `article`
)

/*
CheckTypeValid checks if the embed type is one of the pre-determined constants and returns true if it is
*/
func CheckTypeValid(embedType string) bool {
	return embedType == RichEmbedType ||
		embedType == ImageEmbedType ||
		embedType == VideoEmbedType ||
		embedType == GifvEmbedType ||
		embedType == LinkEmbedType ||
		embedType == ArticleEmbedType
}
