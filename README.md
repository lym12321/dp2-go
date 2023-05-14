## What's this ?
> A Golang module for [dp2](https://github.com/DigitalPlatform/dp2).

## How to use ?
```
go get github.com/lym12321/dp2-go
```

## For Example
```go
package main

import (
	"fmt"

	"github.com/lym12321/dp2-go/dp2"
)

func main() {
	server := "http://your_dp2_server/dp2library/rest"
	username := "your_username"
	password := "your_password"
	var Session string
	var api dp2.Method = dp2.SystemConfig{ServiceUrl: server, Dp2Username: username, Dp2Password: password, Session: &Session}
	// do something...
}

```