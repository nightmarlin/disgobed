package validation

import (
	"testing"

	"github.com/maxatome/go-testdeep/td"
)

/*
TestNewAuthor tests that the default author object is as it should be
*/
func TestCheckValidIconUrl(tt *testing.T) {
	t := td.NewT(tt)

	t.Log(`setting up map`)
	var urls = map[string]bool{
		`yes`:                       false,
		`attachment://video1.mp4`:   true,
		`ssl://generic.archive.org`: false,
		`https://cdn.discordapp.com/chan/msg/name.jpg`: true,
		`localhost:8080//img1.png`:                     false,
		`http://unsecure.server.me/song1.mp3`:          true,
		``:                                             false,
	}

	for input, want := range urls {
		t.Logf(` - testing url '%v'`, input)
		t.Cmp(CheckValidIconURL(input), want)
	}
}

func TestCheckTypeValid(tt *testing.T) {
	t := td.NewT(tt)

	t.Log(`setting up map`)
	var types = map[string]bool{
		`yes`:          false,
		`rich`:         true,
		``:             false,
		`video`:        true,
		`link`:         true,
		`meme`:         false,
		`image`:        true,
		`awesomeThing`: false,
	}

	for input, want := range types {
		t.Logf(` - testing type '%v'`, input)
		t.Cmp(CheckTypeValid(input), want)
	}
}
