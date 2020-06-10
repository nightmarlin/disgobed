# DiscordGo Embed Wrapper

[![Go Report Card](https://goreportcard.com/badge/github.com/Nightmarlin/discordgoembedwrapper)](https://goreportcard.com/report/github.com/Nightmarlin/discordgoembedwrapper)

## About

This module wraps the discordgo embed structure with a set of helper functions.
This allows for the easy construction and sending of DiscordGo Embeds in a more idiomatic way.

This library has been written to comply with the specification at the
[API docs](https://discord.com/developers/docs/resources/channel#embed-object-embed-structure)

## How to use

If using go modules:

```
require github.com/Nightmarlin/discordgoembedwrapper
```

Else

```
go get github.com/Nightmarlin/discordgoembedwrapper
```

Then whenever you want to send an embedded message:

```go
package mypackage

import (
    embeds `github.com/Nightmarlin/discordgoembedwrapper`
)

[...]
  res := embeds.NewEmbed(). // Generate new Embed
    SetType(embeds.RichEmbedType). //
    SetTitle(`Test Embed`).
    SetDescription(`A very interesting text embed`).
    SetThumbnail(`https://upload.wikimedia.org/wikipedia/commons/thumb/5/5a/DOM-model.svg/1024px-DOM-model.svg.png`).
    Finalize()

  session.ChannelMessageSendEmbed(channelid, res)
[...]
```