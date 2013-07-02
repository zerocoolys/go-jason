package main

import (
	"json"
)

func main() {
	var builder json.ContentBuilder = new(json.XContentBuilder)
	builder.StartObject("hello")
}
