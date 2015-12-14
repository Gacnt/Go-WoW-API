# Go-WoW-API
An API for World of Warcraft

Currently in the works

# Installation

`go get github.com/Gacnt/Go-WoW-API`

# To Use:

```
import "github.com/Gacnt/Go-WoW-API"

var wow = wowapi.NewAPI("apikey", "locale")

playerItems := wow.GetItems("Username", "Server")

log.Println(playerItems)
```

The above example will get all the items the player is currently wearing in game, to see the response visit the API in depth [here](https://godoc.org/github.com/Gacnt/Go-WoW-API)
