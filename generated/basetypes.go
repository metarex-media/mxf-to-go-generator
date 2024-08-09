// Copyright ©2019-2024  Mr MXF   info@mrmxf.com
// BSD-3-Clause License  https://opensource.org/license/bsd-3-clause/
//
// Package mxf2go was made using the smpte registers (https://registry.smpte-ra.org) on 2024-08-09 14:47:09.754847092 +0100 BST m=+0.315049359
// if this seems out of date you will want to regenerated from the smpte library, to ensure compatibility
package mxf2go

import (
	"encoding/binary"
	"fmt"
)

func DecodeTChar(value []byte) (any, error) {

	return rune(value[0]), nil
}

func DecodeTCharacter(value []byte) (any, error) {

	return nil, nil
}

func DecodeTInt16(value []byte) (any, error) {

	return int16(binary.BigEndian.Uint16(value[0:2])), nil
}
func EncodeTInt16(value int16) ([]byte, error) {

	result := make([]byte, 2)
	binary.BigEndian.PutUint16(result, uint16(value))
	return result, nil
}

func DecodeTInt32(value []byte) (any, error) {

	return int32(binary.BigEndian.Uint32(value[0:4])), nil
}
func EncodeTInt32(value int32) ([]byte, error) {

	result := make([]byte, 4)
	binary.BigEndian.PutUint32(result, uint32(value))
	return result, nil
}

func DecodeTInt64(value []byte) (any, error) {

	return int64(binary.BigEndian.Uint64(value[0:8])), nil
}
func EncodeTInt64(value int64) ([]byte, error) {

	result := make([]byte, 8)
	binary.BigEndian.PutUint64(result, uint64(value))
	return result, nil
}

func DecodeTInt8(value []byte) (any, error) {

	return int8(value[0]), nil
}
func EncodeTInt8(value int8) ([]byte, error) {

	result := []byte{uint8(value)}
	return result, nil
}

func DecodeTUInt16(value []byte) (any, error) {

	return binary.BigEndian.Uint16(value[0:2]), nil
}
func EncodeTUInt16(value uint16) ([]byte, error) {

	result := make([]byte, 2)
	binary.BigEndian.PutUint16(result, value)
	return result, nil
}

func DecodeTUInt32(value []byte) (any, error) {

	return binary.BigEndian.Uint32(value[0:4]), nil
}
func EncodeTUInt32(value uint32) ([]byte, error) {

	result := make([]byte, 4)
	binary.BigEndian.PutUint32(result, value)
	return result, nil
}

func DecodeTUInt64(value []byte) (any, error) {

	return binary.BigEndian.Uint64(value[0:8]), nil
}
func EncodeTUInt64(value uint64) ([]byte, error) {

	result := make([]byte, 8)
	binary.BigEndian.PutUint64(result, value)
	return result, nil
}

func DecodeTUInt8(value []byte) (any, error) {

	return value[0], nil
}
func EncodeTUInt8(value uint8) ([]byte, error) {

	result := []byte{value}
	return result, nil
}

func DecodeTUTF8Character(value []byte) (any, error) {

	return nil, nil
}

func DecodeTany(value []byte) (any, error) {

	return value, nil
}
func EncodeTany(value any) ([]byte, error) {

	return []byte{}, nil
}

func DecodeTbyte(value []byte) (any, error) {

	return value, nil
}
func EncodeTbyte(value byte) ([]byte, error) {

	return []byte{}, nil
}

func DecodeTfloat32(value []byte) (any, error) {

	return nil, nil
}
func EncodeTfloat32(value float32) ([]byte, error) {

	return []byte{}, nil
}

type TStrongReference []byte
type TWeakReference []byte

func DecodeTStrongReference(value []byte) (any, error) {
	field, _ := DecodeTUInt8Array(value)
	result := TStrongReference(field.(TUInt8Array))
	return result, nil
}
func EncodeTStrongReference(value TStrongReference) ([]byte, error) {

	result := []byte{}
	for _, val := range value {
		field, _ := EncodeTUInt8(uint8(val))
		result = append(result, field...)
	}
	return result, nil
}

func (mt TStrongReference) MarshalText() ([]byte, error) {
	var fullUUID string
	for _, uid := range mt {
		fullUUID += fmt.Sprintf("%02x", uid)
	}
	return []byte(fullUUID), nil
}

func DecodeTWeakReference(value []byte) (any, error) {
	field, _ := DecodeTUInt8Array(value)
	result := TWeakReference(field.(TUInt8Array))
	return result, nil
}

func EncodeTWeakReference(value TWeakReference) ([]byte, error) {

	result := []byte{}
	for _, val := range value {
		field, _ := EncodeTUInt8(uint8(val))
		result = append(result, field...)
	}
	return result, nil
}

func (mt TWeakReference) MarshalText() ([]byte, error) {
	var fullUUID string
	for _, uid := range mt {
		fullUUID += fmt.Sprintf("%02x", uid)
	}
	return []byte(fullUUID), nil
}

func (mt TUUID) MarshalText() ([]byte, error) {
	var fullUUID string
	for _, uid := range mt {
		fullUUID += fmt.Sprintf("%02x", uid)
	}
	return []byte(fullUUID), nil
}

func (mt TPackageIDType) MarshalText() ([]byte, error) {
	var fullUUID string

	for _, uid := range mt.SMPTELabel {
		fullUUID += fmt.Sprintf("%02x", uid)
	}
	// generate the name

	fullUUID += fmt.Sprintf("%02x", mt.Length)
	fullUUID += fmt.Sprintf("%02x", mt.InstanceHigh)
	fullUUID += fmt.Sprintf("%02x", mt.InstanceMid)
	fullUUID += fmt.Sprintf("%02x", mt.InstanceLow)

	materialString, _ := mt.Material.MarshalText()
	fullUUID += string(materialString)
	return []byte(fullUUID), nil
}

func (mt TAUID) MarshalText() ([]byte, error) {
	var fullUUID string

	fullUUID += fmt.Sprintf("%08x", mt.Data1)
	fullUUID += fmt.Sprintf("%04x", mt.Data2)
	fullUUID += fmt.Sprintf("%04x", mt.Data3)

	for _, uid := range mt.Data4 {
		fullUUID += fmt.Sprintf("%02x", uid)
	}

	return []byte(fullUUID), nil
}

// Primer is the go body of the primer to be used for encoding
type Primer struct {
	dynamicTag *uint16
	// Tags is a map of the 16 byte key and 2 byte shorthand
	// to get the key bytes just run the []byte(key)
	Tags map[string][]byte
}

// AddEntry adds an entry to the primer, it checks the key
// has not already been added to prevent duplicates.
// if the short hand key does not have a length of 2 then it
// is not used.
func (p *Primer) AddEntry(id, shorthand []byte) []byte {
	sId := string(id)

	// if its already added do not decrement
	// and return the value already used
	if tag, ok := p.Tags[sId]; ok {
		return tag
	}

	if len(shorthand) != 2 {
		code := binary.BigEndian.AppendUint16([]byte{}, *p.dynamicTag)
		*p.dynamicTag--
		p.Tags[sId] = code
		return code
	} else {
		p.Tags[sId] = shorthand
		return shorthand
	}
}

// NewPrimer generates a new primer for encoding groups
func NewPrimer() *Primer {
	start := uint16(0xffff)
	return &Primer{
		dynamicTag: &start,
		Tags:       make(map[string][]byte),
	}
}
