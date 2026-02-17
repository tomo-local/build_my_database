package kv

import (
	"encoding/binary"
	"io"
)

type Entry struct {
	key []byte
	val []byte
}

func (ent *Entry) Encode() []byte {
	// 初期のデータ構造を定義 (key)4byte+（val)4byte+keyサイズ+valサイズ
	data := make([]byte, 4+4+len(ent.key)+len(ent.val))
	// keyの文字数
  binary.LittleEndian.PutUint32(data[0:4], uint32(len(ent.key)))
	// valの文字数
  binary.LittleEndian.PutUint32(data[4:8], uint32(len(ent.val)))
	//
	copy(data[8:], ent.key)
	copy(data[8+len(ent.key):], ent.val)
	return data
}

func (ent *Entry) Decode(r io.Reader) error {
	// 最初のkeyとvalのサイズを格納する変数を定義
	var header [8]byte
	// [8]bytesまで、データを読み込んでねということ
	if _, err := io.ReadFull(r, header[:]); err != nil {
		return err
	}

	// keyの文字数
	keyLen := int(binary.LittleEndian.Uint32(header[0:4]))
	// valの文字数
	valLen := int(binary.LittleEndian.Uint32(header[4:8]))

	data := make([]byte, keyLen+valLen)
	if _, err := io.ReadFull(r, data); err != nil {
		return err
	}

	// 最初からkeyLenの長さまで取得
	ent.key = data[:keyLen]
	// keyLenの長さから最後まで取得
	ent.val = data[keyLen:]
	return nil
}
