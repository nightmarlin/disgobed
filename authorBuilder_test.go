package disgobed

import (
	"fmt"
	"testing"

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

	t.Log(`2. test Finalize() on author struct with url`)
	t.Log(` - create url, expected author and expected errors`)
	var testUrl = `https://github.com/Nightmarlin`
	wantErrors = nil
	wantAuthor = &disgord.EmbedAuthor{
		URL:          testUrl,
		Name:         "",
		IconURL:      "",
		ProxyIconURL: "",
	}

	t.Log(` - run test`)
	gotAuthor, gotErrors = NewAuthor().SetURL(testUrl).Finalize()

	t.Cmp(gotErrors, wantErrors)
	t.Cmp(gotAuthor, wantAuthor)

	t.Log(`3. test correct error generation on incorrect url type`)
	t.Log(` - create url, expected author and expected errors`)
	wantErrors = &[]error{
		fmt.Errorf(``),
	}

	t.Log(`AuthorBuilder.Finalize() test complete`)
}
