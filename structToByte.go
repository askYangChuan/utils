package utils

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
)

func structToByteDoSlice(v reflect.Value, endian binary.ByteOrder, b *bytes.Buffer) error{
	for i := 0; i < v.Len(); i++ {
		v1 := v.Index(i)
		err := structToByteField(v1, endian, b)
		if err != nil {
			return err
		}
	}
	return nil
}

func structToByteField(v reflect.Value, endian binary.ByteOrder, b *bytes.Buffer) error {
	var err error = nil
	temp := make([]byte, 8)

	switch v.Kind() {
	case reflect.Int8:
		r := v.Interface().(int8)
		b.WriteByte(byte(r))
	case reflect.Uint8:
		r := v.Interface().(uint8)
		b.WriteByte(byte(r))
	case reflect.Int16:
		r := v.Interface().(int16)
		endian.PutUint16(temp, uint16(r))
		b.Write(temp[:2])
	case reflect.Uint16:
		r := v.Interface().(uint16)
		endian.PutUint16(temp, r)
		b.Write(temp[:2])
	case reflect.Int32:
		r := v.Interface().(int32)
		endian.PutUint32(temp, uint32(r))
		b.Write(temp[:4])
	case reflect.Uint32:
		r := v.Interface().(uint32)
		endian.PutUint32(temp, r)
		b.Write(temp[:4])
	case reflect.Int:
		r := v.Interface().(int)
		endian.PutUint32(temp, uint32(r))
		b.Write(temp[:4])
	case reflect.Uint:
		r := v.Interface().(uint)
		endian.PutUint32(temp, uint32(r))
		b.Write(temp[:4])
	case reflect.Int64:
		r := v.Interface().(int64)
		endian.PutUint64(temp, uint64(r))
		b.Write(temp[:8])
	case reflect.Uint64:
		r := v.Interface().(uint64)
		endian.PutUint64(temp, r)
		b.Write(temp[:8])
	case reflect.Array:
		v = v.Slice(0, v.Len())
		fallthrough
	case reflect.Slice:
		err = structToByteDoSlice(v, endian, b)
	case reflect.Struct:
		err = structToByte(v, endian, b)
	default:
		return fmt.Errorf("unknown member kind %v", v.Kind())
	}

	return err
}

func structToByte(v reflect.Value, endian binary.ByteOrder, b *bytes.Buffer) (error) {
	var err error
	switch v.Kind() {
	case reflect.Struct:
	default:
		return fmt.Errorf("not a struct")
	}

	for i := 0; i < v.NumField(); i++ {
		err = structToByteField(v.Field(i), endian, b)
		if err != nil {
			return err
		}
	}

	return err
}

/* true and default is BigEndian, false is littleEndian */
func StructToByte(origin interface{}, args ...bool) ([]byte, error) {
	v := reflect.ValueOf(origin)

	switch v.Kind() {
	case reflect.Ptr:
		v = v.Elem()
	case reflect.Struct:
	default:
		return nil, fmt.Errorf("not ptr or struct")
	}

	var b bytes.Buffer
	var endian binary.ByteOrder = binary.BigEndian
	if(len(args) > 0) {
		if args[0] == false {
			endian = binary.LittleEndian
		}
	}
	err := structToByte(v, endian, &b)
	return b.Bytes(), err
}








