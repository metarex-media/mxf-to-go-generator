//  Copyright ©2019-2024  Mr MXF   info@mrmxf.com
//  BSD-3-Clause License  https://opensource.org/license/bsd-3-clause/

package main

import (
	"encoding/xml"
	"os"
	"strings"
)

func baseSetUp(typeRegisterFile string, registers []string) (xmlReg map[string]entryinfo, bases map[string]entryinfo, err error) {
	// registers := []string{"./registers/Elements.xml", "./registers/Groups.xml", "./registers/Labels.xml", "./registers/Types.xml"}

	xmlReg = make(map[string]entryinfo)

	// generate a register map of every value
	for _, r := range registers {
		b, errr := os.ReadFile(r)

		if errr != nil {
			err = errr
			return
		}

		var g LabelsRegister
		err = xml.Unmarshal(b, &g)
		if err != nil {
			return
		}

		for _, reg := range g.Entries.Entry {
			xmlReg[reg.UL] = reg
		}
	}

	typeBytes, errr := os.ReadFile(typeRegisterFile)
	if errr != nil {
		err = errr
		return
	}

	var typeRegister LabelsRegister
	err = xml.Unmarshal(typeBytes, &typeRegister)
	if err != nil {
		return
	}

	// /fmt.Println(e)
	bases = make(map[string]entryinfo)

	for i, reg := range typeRegister.Entries.Entry {
		if reg.Type == "" && reg.BaseType == "" && len(reg.Facets.Facet) == 0 && reg.Kind == "LEAF" {
			reg.Base = true

			// replace symbols with go terms
			switch reg.Symbol {
			case "Float":
				reg.Symbol = "float32"
			case "Indirect", "Opaque":

				reg.Symbol = "any"
			case "Stream":
				reg.Symbol = "byte"
			}

			// update the register with the base tag
			xmlReg[reg.UL] = reg
			bases[reg.Symbol] = reg
			typeRegister.Entries.Entry[i] = reg
		}
	}
	return
}

// typeToFilling is used for the basic type decodes
func typeToFilling(targetType string, lengths syncmap) string {

	switch strings.ToLower(targetType) {
	case "[]byte", "any", "byte":
		lengths.mapper["T"+targetType] = 0
		return "\n return value , nil"
	case "uint64":
		lengths.mapper["T"+targetType] = 8
		return "\n return binary.BigEndian.Uint64(value[0:8]) , nil"
	case "uint32":
		lengths.mapper["T"+targetType] = 4
		return "\n return binary.BigEndian.Uint32(value[0:4]) , nil"
	case "uint16":
		lengths.mapper["T"+targetType] = 2
		return "\n return binary.BigEndian.Uint16(value[0:2]) , nil"
	case "uint8":
		lengths.mapper["T"+targetType] = 1
		return "\n return value[0] , nil"
	case "int64":
		lengths.mapper["T"+targetType] = 8
		return "\n return int64(binary.BigEndian.Uint64(value[0:8])) , nil"
	case "int32":
		lengths.mapper["T"+targetType] = 4
		return "\n return int32(binary.BigEndian.Uint32(value[0:4])) , nil"
	case "int16":
		lengths.mapper["T"+targetType] = 2
		return "\n return int16(binary.BigEndian.Uint16(value[0:2])) , nil"
	case "int8":
		lengths.mapper["T"+targetType] = 1
		return "\n return int8(value[0]) , nil"
	case "char":
		return "\n return rune(value[0]) , nil"
	case "rune":
		return "\n return string(value) , nil"
	default:
		return "\n return nil , nil"
		// runes give utf16 encoding etc I'll leave as a single one for the moment
	}

}

// typeToFilling is used for the basic type decodes
func typeToFillingEncode(targetType string) string {

	switch strings.ToLower(targetType) {
	case "[]byte", "any", "byte":
		// lengths.mapper["T"+targetType] = 0
		return "\n return []byte{} , nil"
	case "uint64":
		fill := `		result := make([]byte, 8)
		binary.BigEndian.PutUint64(result, value)
		return result, nil`
		return fill
	case "uint32":
		fill := `		result := make([]byte, 4)
		binary.BigEndian.PutUint32(result, value)
		return result, nil`
		return fill
	case "uint16":
		fill := `		result := make([]byte, 2)
		binary.BigEndian.PutUint16(result, value)
		return result, nil`
		return fill
	case "uint8":
		//	lengths.mapper["T"+targetType] = 1
		fill := `     result := []byte{value}
				return result, nil`
		return fill
	case "int64":
		// lengths.mapper["T"+targetType] = 8
		fill := `		result := make([]byte, 8)
		binary.BigEndian.PutUint64(result, uint64(value))
		return result, nil`
		return fill
	case "int32":
		// lengths.mapper["T"+targetType] = 4
		fill := `		result := make([]byte, 4)
		binary.BigEndian.PutUint32(result, uint32(value))
		return result, nil`
		return fill
	case "int16":
		// lengths.mapper["T"+targetType] = 2
		fill := `		result := make([]byte, 2)
		binary.BigEndian.PutUint16(result, uint16(value))
		return result, nil`
		return fill
	case "int8":
		//	lengths.mapper["T"+targetType] = 1
		fill := `     result := []byte{uint8(value)}
		return result, nil`
		return fill
	case "char":
		return "\n return []byte{} , nil" // "\n return rune(value[0]) , nil"
	case "rune":
		return "\n return []byte{} , nil" // "\n return string(value) , nil"
	default:
		return "\n return []byte{} , nil" // "\n return nil , nil"
		// runes give utf16 encoding etc I'll leave as a single one for the moment
	}

}
