package kv

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEntryEncode(t *testing.T) {
	ent := Entry{key: []byte("k1"), val: []byte("xxx")}
	// 2文字 + 3文字 + false + k1 + xxx
	data := []byte{2, 0, 0, 0, 3, 0, 0, 0, 0, 'k', '1', 'x', 'x', 'x'}

	assert.Equal(t, data, ent.Encode())

	decoded := Entry{}
	err := decoded.Decode(bytes.NewBuffer(data))
	assert.Nil(t, err)
	assert.Equal(t, ent, decoded)

	ent = Entry{key: []byte("k1"), deleted: true}
	// 2文字 + 0文字 + true + k1
	data = []byte{2, 0, 0, 0, 0, 0, 0, 0, 1, 'k', '1'}

	assert.Equal(t, data, ent.Encode())

	decoded = Entry{}
	err = decoded.Decode(bytes.NewBuffer(data))
	assert.Nil(t, err)
	assert.Equal(t, ent, decoded)
}
