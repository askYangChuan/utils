package utils

import (
	"encoding/binary"
	"fmt"
	"reflect"
)

func byteToStructDoSlice(b []byte, v reflect.Value, endian binary.ByteOrder) (int, error) {
	var offset int = 0
	for i := 0; i < v.Len(); i++ {
		v1 := v.Index(i)
		n, err :=  byteToStructField(b[offset:], v1, endian)
		if err != nil {
			return 0, err
		}
		offset += n
	}
	return offset, nil
}

func byteToStructField(b []byte, v reflect.Value, endian binary.ByteOrder) (int, error) {
	var offset int = 0
	var tmp int = 0
	var err error = nil
	switch v.Kind() {
	case reflect.Int8:
		v.SetInt(int64(b[0]))
		offset++
	case reflect.Uint8:
		v.SetUint(uint64(b[0]))
		offset++
	case reflect.Int16:
		v.SetInt(int64(endian.Uint16(b)))
		offset += 2
	case reflect.Uint16:
		v.SetUint(uint64(endian.Uint16(b)))
		offset += 2
	case reflect.Int, reflect.Int32:
		v.SetInt(int64(endian.Uint32(b)))
		offset += 4
	case reflect.Uint, reflect.Uint32:
		v.SetUint(uint64(endian.Uint32(b)))
		offset += 4
	case reflect.Int64:
		v.SetInt(int64(endian.Uint64(b)))
		offset += 8
	case reflect.Uint64:
		v.SetUint(endian.Uint64(b))
		offset += 8
	case reflect.Array:
		v = v.Slice(0, v.Len())
		fallthrough
	case reflect.Slice:
		tmp, err = byteToStructDoSlice(b, v, endian)
		offset += tmp
	case reflect.Struct:
		tmp, err = byteToStruct(b, v, endian)
		offset += tmp
	default:
		return 0, fmt.Errorf("not support member %v", v.Kind())
	}
	return offset, err
}

func byteToStruct(b []byte, v reflect.Value, endian binary.ByteOrder) (int, error) {
	var offset int = 0
	for i := 0; i < v.NumField(); i++ {
		v1 := v.Field(i)
		n, err := byteToStructField(b[offset:], v1, endian)
		if err != nil {
			return 0, err
		}
		offset += n

	}
	return offset, nil
}

func ByteToStruct(b []byte, r interface{}, args ...bool) (int, error){
	v := reflect.ValueOf(r)

	if v.Kind() != reflect.Ptr {
		return 0, fmt.Errorf("receiver is not a ptr")
	}
	v = v.Elem()
	if v.Kind() != reflect.Struct {
		return 0, fmt.Errorf("receiver is not a struct")
	}

	var endian binary.ByteOrder = binary.BigEndian
	if(len(args) > 0) {
		if args[0] == false {
			endian = binary.LittleEndian
		}
	}

	return byteToStruct(b, v, endian)
}