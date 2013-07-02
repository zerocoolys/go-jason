package json

import (
// "bytes"
// "io"
// "io/ioutil"
// "strconv"
// "strings"
)

var (
	buffer []byte
)

type ContentBuilder interface {
	StartObject(name string) (builder ContentBuilder)
	EndObject() (builder ContentBuilder)

	// Field(name string, value string) (builder ContentBuilder)

	// StartArray(name string) (builder ContentBuilder)
	// EndArray() (builder ContentBuilder)

	// Value(value interface{}) (builder ContentBuilder)
}

type XContentBuilder struct {
	maps      map[string]interface{}
	fieldName string
	stack     []interface{}
	status    []int //
}

var (
	ENUM_STATUS_OBJECT = 1
	ENUM_STATUS_ARRAY  = 2
	ENUM_STATUS_FIELD  = 3
	ENUM_STATUS_VALUE  = 4
)

func (b *XContentBuilder) StartObject(name string) (xBuilder ContentBuilder) {
	if b.maps == nil {
		b.maps = make(map[string]interface{}, 16)
	}

	pushStatus(name, ENUM_STATUS_OBJECT, b)
	return b
}

// func (b *XContentBuilder)

func (b *XContentBuilder) EndObject() (xBuilder ContentBuilder) {
	si := len(b.status)
	if si == 0 || b.status[si-1] != ENUM_STATUS_OBJECT {
		return b
	}

	popStatus(b)
	return b
}

func pushStatus(name string, status int, b *XContentBuilder) {
	i := len(b.stack)
	si := len(b.status)

	b.stack = make([]interface{}, i+1)
	b.status = make([]int, si+1)

	b.stack[i] = name
	b.status[si] = status
}

func popStatus(b *XContentBuilder) {
	i := len(b.stack)
	si := len(b.status)
	b.stack = b.stack[:i-1]
	b.status = b.status[:si-1]

}
