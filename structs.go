package cachebench

import (
	"bytes"
	"encoding/gob"
)

func init() {
	gob.Register(SomeStruct{})
}

type SomeStruct struct {
	ID         int
	I          int
	D          int
	B          int
	T1         int
	T2         int
	Time       int64
	StringID   string
	Name1      string
	Name2      string
	StringTime string
	Type       string
	Status     string
	S          string
}

func (m *SomeStruct) Encode() ([]byte, error) {
	buf := make([]byte, 0)
	bb := bytes.NewBuffer(buf)

	enc := gob.NewEncoder(bb)
	err := enc.Encode(m)
	if err != nil {
		return nil, err
	}
	return bb.Bytes(), nil
}

func (m *SomeStruct) Decode(data []byte) error {
	bb := bytes.NewBuffer(data)
	dec := gob.NewDecoder(bb)
	return dec.Decode(m)
}
