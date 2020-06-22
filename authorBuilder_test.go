package disgobed

import (
	"fmt"
	"testing"

	"github.com/Nightmarlin/disgobed/validation"
	"github.com/andersfylling/disgord"
	"github.com/maxatome/go-testdeep/td"
)

/*
TestNewAuthor tests that the default author object is as it should be
*/
func TestNewAuthor(tt *testing.T) {
	t := td.NewT(tt)

	var (
		got  AuthorBuilder
		want AuthorBuilder
	)

	t.Log(`1. test NewAuthor() returns appropriate value`)
	t.Log(` - create author struct`)

	want = AuthorBuilder{
		EmbedAuthor: &disgord.EmbedAuthor{
			URL:          "",
			Name:         "",
			IconURL:      "",
			ProxyIconURL: "",
		},
		Errors: nil,
	}

	t.Log(` - run test`)

	got = *NewAuthor()
	t.Cmp(got, want)

	t.Log(`NewAuthor() test complete`)
}

func TestAuthor_Finalize(tt *testing.T) {
	t := td.NewT(tt)

	var (
		gotAuthor  *disgord.EmbedAuthor
		gotErrors  *[]error
		wantAuthor *disgord.EmbedAuthor
		wantErrors *[]error
	)

	t.Log(`1. test Finalize() on empty author struct`)
	t.Log(` - create discordgo author struct and expected error struct`)
	wantErrors = nil
	wantAuthor = &disgord.EmbedAuthor{
		URL:          "",
		Name:         "",
		IconURL:      "",
		ProxyIconURL: "",
	}

	t.Log(` - run test`)

	gotAuthor, gotErrors = NewAuthor().Finalize()
	t.Cmp(gotErrors, wantErrors)
	t.Cmp(gotAuthor, wantAuthor)

	t.Log(`2. test Finalize() on author struct with iconUrl`)
	t.Log(` - create iconUrl, expected author and expected errors`)
	var testUrl = `https://github.com/Nightmarlin`
	wantErrors = nil
	wantAuthor = &disgord.EmbedAuthor{
		URL:          "",
		Name:         "",
		IconURL:      testUrl,
		ProxyIconURL: "",
	}

	t.Log(` - run test`)
	gotAuthor, gotErrors = NewAuthor().SetIconURL(testUrl).Finalize()

	t.Cmp(gotErrors, wantErrors)
	t.Cmp(gotAuthor, wantAuthor)

	t.Log(`3. test correct error generation on incorrect iconUrl type`)
	t.Log(` - create iconUrl, expected author and expected errors`)
	testUrl = `aka.ms/ps7`
	wantErrors = &[]error{
		fmt.Errorf(validation.InvalidUrlErrTemplateString, "author iconUrl", testUrl),
	}
	wantAuthor = &disgord.EmbedAuthor{
		URL:          "",
		Name:         "",
		IconURL:      "",
		ProxyIconURL: "",
	}

	t.Log(` - run test`)
	gotAuthor, gotErrors = NewAuthor().SetIconURL(testUrl).Finalize()

	t.Cmp(gotErrors, wantErrors)
	t.Cmp(gotAuthor, wantAuthor)

	t.Log(`AuthorBuilder.Finalize() test complete`)
}
