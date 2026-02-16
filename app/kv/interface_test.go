package kv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKVBasic(t *testing.T) {
	kv := KV{}
	err := kv.Open()
	assert.Nil(t, err)
	defer kv.Close()

  // データの保存
	updated, err := kv.Set([]byte("k1"), []byte("v1"))
	assert.True(t, updated && err == nil)

	// データの取得
	val, ok, err := kv.Get([]byte("k1"))
	assert.True(t, string(val) == "v1" && ok && err == nil)

	// 存在しないデータの取得
	_, ok, err = kv.Get([]byte("xxx"))
	assert.True(t, !ok && err == nil)

	// 存在しないデータの削除
	updated, err = kv.Del([]byte("xxx"))
	assert.True(t, !updated && err == nil)

	// データの削除
	updated, err = kv.Del([]byte("k1"))
	assert.True(t, updated && err == nil)

	// データの削除の確認
	_, ok, err = kv.Get([]byte("xxx"))
	assert.True(t, !ok && err == nil)
}
