package validation

import (
	"github.com/andersfylling/disgord"
)

/*
ValidateEmbed returns whether or not discord is likely accept the embed attached to it. If discord is unlikely to
accept the embed, it returns a list of reasons why. If msg is not nil, the checker will also validate `attachment://`
urls
*/
func ValidateEmbed(embed *disgord.Embed, msg *disgord.Message) *[]error {
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
	 *   6) EmbedBuilder types are limited to one of "rich" | "image" | "video" | "gifv" | "link" | "article"
	 *   7) The following fields must not be empty when present
	 *      | footer text
	 *      | field name
	 *      | field value
	 *   8) (optional with presence of msg) All `attachment://` urls should reference attached items
	 *   9) (optional with presence of msg) Message content cannot exceed 2000 characters
	 */
	return nil
}
