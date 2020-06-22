package disgobed

import (
	"testing"

	"github.com/andersfylling/disgord"
	"github.com/maxatome/go-testdeep/td"
)

/*
TestNewAuthor tests that the default author object is as it should be
*/
func TestNewEmbed(tt *testing.T) {
	t := td.NewT(tt)

	var (
		gotEmbed   *disgord.Embed
		gotErrors  *[]error
		wantEmbed  *disgord.Embed
		wantErrors *[]error
	)

	t.Log(`1. test NewEmbed() returns appropriate value`)
	t.Log(` - create expected return structures`)
	wantEmbed = &disgord.Embed{}
	wantErrors = nil

	t.Log(` - run test`)
	gotEmbed, gotErrors = NewEmbed().Finalize()
	t.Cmp(gotEmbed, wantEmbed)
	t.Cmp(gotErrors, wantErrors)
}
