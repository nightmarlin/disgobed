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
