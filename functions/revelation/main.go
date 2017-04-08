package main

import (
	"encoding/json"

	"github.com/apex/go-apex"

	"github.com/shiimaxx/revelation"
)

type message struct {
	Hello string `json:"hello"`
}

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		var postItems []revelation.PostItem

		postItems, err := revelation.Random()
		if err != nil {
			return nil, err
		}

		err = revelation.ToSlack(postItems)
		if err != nil {
			return nil, err
		}

		return postItems, nil
	})
}
