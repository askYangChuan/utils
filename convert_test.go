package utils

import (
	"fmt"
	"testing"
)

type child1 struct {
	C uint64
	D uint
	E uint8
	F uint32
}

type child struct {
	I32 int32
	U32 uint32
	ArrayInt [3]int
	ArrayUint16 [2]uint16
}

type Father struct {
	I8 int8
	U8 uint8
	I16 int16
	S child
	I64 int64
	U64 uint64
}

func TestStruct2Byte(t *testing.T) {
	a := Father{}
	c := Father{}
	c.I8 = 8
	c.U8 = 8
	c.I16 = 16
	c.I64 = 64
	c.U64 = 64
	c.S.I32 = 32
	c.S.U32 = 32
	c.S.ArrayInt[0] = 32
	c.S.ArrayInt[1] = 32
	c.S.ArrayInt[2] = 32
	c.S.ArrayUint16[0] = 16
	c.S.ArrayUint16[1] = 16


	fmt.Println(c)
	b, err := StructToByte(&c)
	fmt.Println(b)
	fmt.Println(err)
	n, err := ByteToStruct(b, &a)

	fmt.Println(n)
	fmt.Println(err)
	fmt.Println(a)
}
