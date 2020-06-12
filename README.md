# Disgobed &mdash; A Disgord Embed Wrapper

[![Go Report Card](https://goreportcard.com/badge/github.com/Nightmarlin/disgobed)](https://goreportcard.com/report/github.com/Nightmarlin/disgobed)

## About

This module wraps the Disgord embed structure with a set of helper functions.
This allows for the easy construction and sending of Disgord Embeds in a more idiomatic way.

This library has been written to comply with the specification at the
[API docs](https://discord.com/developers/docs/resources/channel#embed-object-embed-structure).
&amp; by that I mean: I spent way too long looking at how exactly Discord likes their embeds,
then worked out a set of ways to validate that!

## How to use

If using go modules:

```
require github.com/Nightmarlin/disgobed
```

Else

```
go get github.com/Nightmarlin/disgobed
```

Then whenever you want to send an embedded message:

```go
package mypackage

import (
    `github.com/Nightmarlin/disgobed`
    `github.com/andersfylling/disgord`
)

[...]
  res, errs := disgobed.NewEmbed(). // Generate new Embed
    SetType(embeds.RichEmbedType).
    SetTitle(`Test Embed`).
    SetDescription(`A very interesting text embed`).
    SetThumbnail(disgobed.NewThumbnail().
      SetURL(`https://upload.wikimedia.org/wikipedia/commons/thumb/5/5a/DOM-model.svg/1024px-DOM-model.svg.png`).
      SetHW(917, 886)).
    Finalize()

  if errs == nil {
    client.CreateMessage(context.Background(), channelID, &disgord.CreateMessageParams{
      Embed: &res
    }
  }
[...]
```

## Interesting Information

`Finalize()` is a really important function! The [Embed](./embed.go) struct caches all errors that
may occur, due to validation failures or other reasons - this means that you have to actively check
for errors. `Finalize()` makes this easy by returning the cached errors or nil as well as the final
embed. You might find you still want to send an embed that is invalid…
