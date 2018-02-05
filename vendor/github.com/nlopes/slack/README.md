Slack API in Go [![GoDoc](https://godoc.org/github.com/nlopes/slack?status.svg)](https://godoc.org/github.com/nlopes/slack) [![Build Status](https://travis-ci.org/nlopes/slack.svg)](https://travis-ci.org/nlopes/slack)
===============

[![Join the chat at https://gitter.im/go-slack/Lobby](https://badges.gitter.im/go-slack/Lobby.svg)](https://gitter.im/go-slack/Lobby?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

This library supports most if not all of the `api.slack.com` REST
calls, as well as the Real-Time Messaging protocol over websocket, in
a fully managed way.

## Change log

### v0.1.0 - May 28, 2017

This is released before adding context support.
As the used context package is the one from Go 1.7 this will be the last
compatible with Go < 1.7.

Please check [0.1.0](https://github.com/nlopes/slack/releases/tag/v0.1.0)

### CHANGELOG.md

As of this version a [CHANGELOG.md](https://github.com/nlopes/slack/blob/master/CHANGELOG.md) is available. Please visit it for updates.

## Installing

### *go get*

    $ go get -u github.com/nlopes/slack

## Example

### Getting all groups

```golang
import (
	"fmt"

	"github.com/nlopes/slack"
)

func main() ***REMOVED***
	api := slack.New("YOUR_TOKEN_HERE")
	// If you set debugging, it will log all requests to the console
	// Useful when encountering issues
	// api.SetDebug(true)
	groups, err := api.GetGroups(false)
	if err != nil ***REMOVED***
		fmt.Printf("%s\n", err)
		return
	***REMOVED***
	for _, group := range groups ***REMOVED***
		fmt.Printf("ID: %s, Name: %s\n", group.ID, group.Name)
	***REMOVED***
***REMOVED***
```

### Getting User Information

```golang
import (
    "fmt"

    "github.com/nlopes/slack"
)

func main() ***REMOVED***
    api := slack.New("YOUR_TOKEN_HERE")
    user, err := api.GetUserInfo("U023BECGF")
    if err != nil ***REMOVED***
	    fmt.Printf("%s\n", err)
	    return
***REMOVED***
    fmt.Printf("ID: %s, Fullname: %s, Email: %s\n", user.ID, user.Profile.RealName, user.Profile.Email)
***REMOVED***
```

## Minimal RTM usage:

See https://github.com/nlopes/slack/blob/master/examples/websocket/websocket.go


## Contributing

You are more than welcome to contribute to this project.  Fork and
make a Pull Request, or create an Issue if you see any problem.

## License

BSD 2 Clause license