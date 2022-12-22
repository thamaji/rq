rq
====

http request library for golang

## Install
```
$ go get github.com/thamaji/rq
```

## Example
```
package main

import (
	"log"

	"github.com/thamaji/rq"
)

func main() {
	userID := "user1"
	user := struct {
		ID    string `json:"id"`
		Email string `json:"email"`
	}{}

	// GET: http://localhost/api/users/user1
	if err := rq.Get(rq.URL("http://localhost/api/users", userID)).FetchJSON(&user); err != nil {
		log.Fatalln(err)
	}

	// GET: http://localhost/api/users?id=user1
	if err := rq.Get("http://localhost/api/users", rq.Query("id", userID)).FetchJSON(&user); err != nil {
		log.Fatalln(err)
	}

	// POST: http://localhost/api/users
	if err := rq.Post("http://localhost/api/users", rq.BodyJSON(&user)).FetchJSON(&user); err != nil {
		log.Fatalln(err)
	}
}
```