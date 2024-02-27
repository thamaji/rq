rq
====

http request library for golang

## Install
```
$ go get github.com/thamaji/rq
```

## Example
```go
package main

import (
	"errors"
	"log"
	"os"
	"path"

	"github.com/thamaji/rq"
)

func main() {
	client := rq.NewClient(rq.BaseURL("http://localhost"), rq.Verbose(os.Stdout))

	userID := "user1"
	user := struct {
		ID    string `json:"id"`
		Email string `json:"email"`
	}{}

	// GET: http://localhost/api/users/user1
	if err := client.Get(path.Join("/api/users", userID)).FetchJSON(&user); err != nil {
		if errors.Is(err, rq.ErrNotFound) {
			log.Println("user does not exist")
			return
		}
		log.Fatalln(err)
	}

	// GET: http://localhost/api/users?id=user1
	if err := client.Get("/api/users", rq.Query("id", userID)).FetchJSON(&user); err != nil {
		log.Fatalln(err)
	}

	// POST: http://localhost/api/users
	if err := client.Post("/api/users", rq.BodyJSON(&user)).FetchJSON(&user); err != nil {
		log.Fatalln(err)
	}
}
```