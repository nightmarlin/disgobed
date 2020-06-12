package disgobed

import (
	"strings"
)

func checkValidIconURL(url string) bool {
	return strings.HasPrefix(url, `https://`) || // Prefer https:// or attachment:// over http://
		strings.HasPrefix(url, `attachment://`) ||
		strings.HasPrefix(url, `http://`)
}

const (
	// The lowest embed property character limit
	lowerCharLimit = 256

	// The middle embed property character limit
	middleCharLimit = 1024

	// The upper embed property character limit
	upperCharLimit = 2048

	// The maximum total number of characters in an embed
	maxTotalCharLimit = 6000

	// The maximum number of embed fields
	maxFieldCount = 25

	// Largest acceptable colour value
	maxColorValue = 16777215
)

const (
	// [Type Property] '[Value]' does not start with "http://" | "https://" | "attachment://"
	invalidUrlErrTemplateString = `%v '%v' does not start with "http://" | "https://" | "attachment://"`

	// [Type] height '[Value]' or [Type] width '[Value]' is less than or equal to 0
	invalidHWErrTemplateString = `%v height '%v' or %v width '%v' is less than or equal to 0`

	// [Type Property] exceeds [Limit]: length = [Length] | '[Value]'
	characterCountExceedsLimitErrTemplateString = `%v exceeds %v characters: length = %v | '%v'`

	// [Type Property] exceeds [Limit]: length = [Length]
	characterCountExceedsLimitLongErrTemplateString = `%v exceeds %v characters: length = %v`

	// adding field '[FieldName]' would cause field count to exceed [Limit]
	fieldLimitReachedErrTemplateString = `adding field '%v' would cause field count to exceed %v`

	// embed type '[Type]' is not one of "rich" | "image" | "video" | "gifv" | "link" | "article"
	invalidEmbedTypeErrTemplateString = `embed type '%v' is not one of "rich" | "image" | "video" | "gifv" | "link" | "article"`

	// [Type Property] [Value] is not between [LowerLimit] and [UpperLimit]
	valueNotBetweenErrTemplateString = `%v '%v' is not between $v and %v`

	// [Type Property] should not be empty if set
	valueIsEmptyErrString = `$v should not be empty if set`
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
checkTypeValid checks if the embed type is one of the pre-determined constants and returns true if it is
*/
func checkTypeValid(embedType string) bool {
	return embedType == RichEmbedType ||
		embedType == ImageEmbedType ||
		embedType == VideoEmbedType ||
		embedType == GifvEmbedType ||
		embedType == LinkEmbedType ||
		embedType == ArticleEmbedType
}
