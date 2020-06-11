package disgobed

import (
	"reflect"
	"testing"

	"github.com/bwmarrin/discordgo"
)

/*
Determines whether two author structures are equal by value
*/
func authorsEqual(a1, a2 Author) bool {
	// Checks the message embed authors are the same
	var underlyingAuthorsEqual = reflect.DeepEqual(*a1.MessageEmbedAuthor, *a2.MessageEmbedAuthor)

	// Checks that errors is nil for both or not nil for both
	var errorsBothNilOrNonNil = (a1.Errors == nil && a2.Errors == nil) || (a1.Errors != nil && a2.Errors != nil)

	// If both aren't nil, checks that both slices are equal
	var errorsSlicesEqual = true
	if a1.Errors != nil && errorsBothNilOrNonNil {
		errorsSlicesEqual = reflect.DeepEqual(*a1.Errors, *a2.Errors)
	}

	// If any failed, return false
	return underlyingAuthorsEqual && errorsBothNilOrNonNil && errorsSlicesEqual
}

/*
TestNewAuthor tests that the default author object is as it should be
*/
func TestNewAuthor(t *testing.T) {
	want := Author{
		MessageEmbedAuthor: &discordgo.MessageEmbedAuthor{
			URL:          "",
			Name:         "",
			IconURL:      "",
			ProxyIconURL: "",
		},
		Errors: nil,
	}
	if got := *NewAuthor(); !authorsEqual(want, got) {
		t.Errorf(`1. GetDevState() = %v, wanted %v`, got, want)
	}
}
