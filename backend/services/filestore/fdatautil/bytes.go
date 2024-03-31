package fdatautil

import (
	"bytes"
	"io"
)

type BytesFdata struct {
	bytes []byte
}

func NewFromBytes(bytes []byte) *BytesFdata {
	return &BytesFdata{
		bytes: bytes,
	}
}

func (b *BytesFdata) AsBytes() ([]byte, error) {
	return b.bytes, nil
}

func (b *BytesFdata) AsReader() (io.Reader, error) {
	return bytes.NewReader(b.bytes), nil
}

func (b *BytesFdata) Close() error {
	b.bytes = nil

	return nil
}
