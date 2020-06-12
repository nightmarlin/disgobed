# Disgobed Validation

A validation module that checks that your embeds are valid before sending them!
Doesn't require disgobed itself to run!

```go
package myPackage

import (
    [...]
    `log`

    v `github.com/Nightmarlin/disgobed/validation`
    `github.com/andersfylling/disgord`
)

[...]

    var myMessage disgord.Message
    var myMessageEmbed disgord.Embed

    // Do some stuff to myMessage and myMessageEmbed

    errs := v.ValidateEmbed(&myMessageEmbed, &myMessage)

    if errs == nil {
        disgord.client.CreateMessage(context.Background(), channelID,
            &disgord.CreateMessageParams{
                Embed: &res
            })
    } else {
        for _, v := *errs {
            log.Error(v)
        }
    }

[...]

```