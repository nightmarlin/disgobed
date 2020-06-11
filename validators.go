package discordgoembedwrapper

import (
	"github.com/bwmarrin/discordgo"
)

/*
ValidateEmbed returns whether or not discord is likely accept the embed attached to it. If discord is unlikely to
accept the embed, it returns a list of reasons why. If msg is not nil, the checker will also validate `attachment://`
urls
*/
func (e *Embed) ValidateEmbed(msg *discordgo.Message) (bool, *[]error) {
	toCheck, errs := e.Finalize()
	if errs != nil { // Make use of builtin err checking
		return false, errs // Short-Circuit and dont run expensive checks if we already have errors
	}
	return ValidateEmbed(*toCheck, msg)
}

/*
ValidateEmbed returns whether or not discord is likely accept the embed attached to it. If discord is unlikely to
accept the embed, it returns a list of reasons why. If msg is not nil, the checker will also validate `attachment://`
urls
*/
func ValidateEmbed(embed discordgo.MessageEmbed, msg *discordgo.Message) (bool, *[]error) {
	// TODO: Write this
	/* To Check:
	 *   1) The characters in all title, description, field.name, field.value, footer.text, and author.name fields must
	 *      not exceed 6000 characters in total
	 *   2) The following fields are limited to 256 characters
	 *      | embed title
	 *      | field name
	 *      | author name
	 *   3) Field.value is limited to 1024 characters
	 *   4) The following fields are limited to 1024 characters
	 *      | embed description
	 *      | footer text
	 *   5) An embed can have a maximum of 25 attached fields
	 *   6) Embed types are limited to one of "rich" | "image" | "video" | "gifv" | "link" | "article"
	 *   7) The following fields must not be empty when present
	 *      | footer text
	 *      | field name
	 *      | field value
	 *   8) (optional with presence of msg) All `attachment://` urls should reference attached items
	 *   9) (optional with presence of msg) Message content cannot exceed 2000 characters
	 */
	return true, nil
}
