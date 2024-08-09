//  Copyright ©2019-2024  Mr MXF   info@mrmxf.com
//  BSD-3-Clause License  https://opensource.org/license/bsd-3-clause/

package main

import (
	"fmt"
	"strings"
	"time"
)

// decodeMaker builds the decode and encode functions for a struct
func decodeMaker(fields []facet, targetUL string, lengths syncmap, types map[string]entryinfo) string {
	header := fmt.Sprintf("func Decode%v(value []byte) (any, error){\n", targetUL)
	result := fmt.Sprintf("   result := %v{}\n    var field any\n", targetUL)
	encoder := fmt.Sprintf("func Encode%v(value %v) ([]byte, error){\n    result := []byte{}\n    ", targetUL, targetUL)
	// extract each field in the order they were declared
	// adding the decode values into the sync map at the end
	totalLength := 0
	for i, f := range fields {

		targetType, typeSymbol, wrapType, start, end := findBase(types[f.Type], types, "struct")
		footer := fmt.Sprintf("    result.%v = %v%vfield.(%v)%v // else\n", strings.ReplaceAll(f.Symbol, " ", "_"), wrapType, start, typeSymbol, end)
		lengths.sync.Lock()
		length, ok := lengths.mapper[targetType]
		lengths.sync.Unlock()
		count := 0
		// wait here until the lengths have been calculated elsewhere
		for !ok {
			//	fmt.Println(lengths, targetType)
			time.Sleep(time.Millisecond * 10)
			lengths.sync.Lock()
			length, ok = lengths.mapper[targetType]
			lengths.sync.Unlock()
			if count == 250 {
				// just relay a message sometimes fmt.Println(targetType)
				break
			}
			count++
		}

		encoder += fmt.Sprintf("    field%v, _ := Encode%v(%v(value.%v)) // else\n", i, targetType, typeSymbol, strings.ReplaceAll(f.Symbol, " ", "_"))
		encoder += fmt.Sprintf("    result = append(result, field%v...)\n", i)

		if length == 0 { // replace with just fill it up
			result += fmt.Sprintf("    field , _ = Decode%v(value[%v:])\n", targetType, totalLength)
		} else {
			result += fmt.Sprintf("    field , _ = Decode%v(value[%v:%v])\n", targetType, totalLength, totalLength+length)
		}

		result += footer
		totalLength += length
	}
	encoder += "    return result, nil\n}\n"
	result += "    return result, nil\n}\n"

	lengths.sync.Lock()
	lengths.mapper[targetUL] = totalLength
	lengths.sync.Unlock()

	return header + result + encoder

}

// decodeWrapMaker wraps a decode function with its own type
// creating a chain of decodes
func decodeWrapMaker(target entryinfo, targetUL string, lengths syncmap, types map[string]entryinfo) string {
	header := fmt.Sprintf("func Decode%v(value []byte) (any, error){\n", targetUL)

	totalLength := 0

	targetType, typeSymbol, _, _, _ := findBase(target, types, "notstruct")
	footer := fmt.Sprintf("    result := %v(field.(%v)) \n", targetUL, typeSymbol)
	lengths.sync.Lock()
	length, ok := lengths.mapper[targetType]
	lengths.sync.Unlock()
	count := 0
	for !ok {
		//	fmt.Println(lengths, targetType)
		time.Sleep(time.Millisecond * 10)
		lengths.sync.Lock()
		length, ok = lengths.mapper[targetType]
		lengths.sync.Unlock()
		if count == 250 {
			break
		}
		count++
	}

	result := fmt.Sprintf("    field , _ := Decode%v(value[%v:%v])\n", targetType, totalLength, totalLength+length)
	if length == 0 { // replace with just fill it up
		result = fmt.Sprintf("    field , _ := Decode%v(value)\n", targetType)
	}

	result += footer
	// totalLength += length

	result += "    return result, nil\n}\n"
	lengths.sync.Lock()
	lengths.mapper[targetUL] = length
	lengths.sync.Unlock()

	decodeTest := `func Encode%v(value %v) ([]byte, error) {
	result, _ := Encode%v(%v(value))
	return result, nil
}`
	result += fmt.Sprintf(decodeTest, targetUL, targetUL, targetType, typeSymbol)
	return header + result

}

func arrDecodeMaker(varType entryinfo, targetSymbol, arr string, lengths syncmap, types map[string]entryinfo) string {
	// variable arrays require a length uint32 and a count uint32
	header := fmt.Sprintf("func Decode%v(value []byte) (any, error){\n", targetSymbol)
	encodeString := fmt.Sprintf("func Encode%v(value %v) ([]byte, error) {\n", targetSymbol, targetSymbol)
	var result string
	var loopCount any = varType.TypeSize

	targetType, typeSymbol, wrapType, start, end := findBase(varType, types, "false")
	extract := fmt.Sprintf("         result[i] = %v%vfield.(%v)%v // else\n", wrapType, start, typeSymbol, end)

	// length is the length of the array and is either known as an int
	// or calculated as part of the decode function and is a named variable
	var length any

	lengths.sync.Lock()
	length, ok := lengths.mapper[targetType]
	lengths.sync.Unlock()
	count := 0
	// check for the length of the decode function in the array
	// this is run concurrently
	for !ok {
		if count == 250 {
			break
		}
		count++

		time.Sleep(time.Millisecond * 10)
		lengths.sync.Lock()
		length, ok = lengths.mapper[targetType]
		lengths.sync.Unlock()
	}

	// variable arrays and sets need the bit at the top telling them the length of the array
	switch {
	case (varType.TypeKind == "VariableArray" || varType.TypeKind == "Set") && !types[varType.BaseType].Base: // == "Set" { // do some variable array set up
		result += "   count , _  := DecodeTUInt32(value[0:4])\n"
		result += "   arrayCount := int(count.(uint32))\n"
		result += "   fieldLen, _  := DecodeTUInt32(value[4:8])\n"
		result += "   fieldSize := int(fieldLen.(uint32))\n"
		length = "fieldSize"
		loopCount = "arrayCount"
		if varType.TypeSize == 0 { // if a variable array make one the desired length
			result += fmt.Sprintf("   var result %v = make(%v, arrayCount)\n", targetSymbol, arr)
		} else {
			result += fmt.Sprintf("   var result %v = %v{}\n", targetSymbol, arr)
		}

	case varType.TypeSize == 0: // if a variable array of a known size
		lenmid := length
		if lenmid == 0 { // prevent 0 errors by replacing with 1
			lenmid = 1
		}
		result += fmt.Sprintf("   arrayCount := len(value)/%v\n", lenmid)
		loopCount = "arrayCount"
		result += fmt.Sprintf("   var result %v = make(%v, arrayCount)\n", targetSymbol, arr)
	default:
		result += fmt.Sprintf("   var result %v = %v{}\n", targetSymbol, arr)
	}

	result += fmt.Sprintf("    for i:=0; i < %v; i++ {\n", loopCount)
	if (varType.TypeKind == "VariableArray" || varType.TypeKind == "Set") && !types[varType.BaseType].Base {
		// includes the offset for including the count of objects in the arrray
		result += fmt.Sprintf("        field , _ := Decode%v(value[8+i*%v:8+i*%v+%v])\n", targetType, length, length, length)
		arrayCountString := `    result := make([]byte, 8)
		binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))`
		lengths.sync.Lock()
		if lengths.mapper[targetType] == 0 {
			arrayCountString += `
	if len(value) != 0 {  // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
			`
		} else {

			arrayCountString += fmt.Sprintf(`
		binary.BigEndian.PutUint32(result[4:8], %v)
		`, lengths.mapper[targetType])
		}
		lengths.sync.Unlock()
		encodeString += arrayCountString
	} else {
		result += fmt.Sprintf("        field , _ := Decode%v(value[i*%v:i*%v+%v])\n", targetType, length, length, length)
		encodeString += "result := []byte{}\n"
	}

	result += extract
	result += "    }\n return result, nil}\n"

	if length != "fieldSize" { // if the length isn't variable update as an integer
		lengths.sync.Lock()
		lengths.mapper[targetSymbol] = length.(int) * varType.TypeSize
		lengths.sync.Unlock()
	} else {
		lengths.sync.Lock()
		lengths.mapper[targetSymbol] = 0
		lengths.sync.Unlock()
	}

	some :=
		`for _, val := range value {
		field, _ := Encode%v(%v(val))
		result = append(result, field...)
	}
	return result, nil
}
`
	encodeString += fmt.Sprintf(some, targetType, typeSymbol)
	header = encodeString + header

	return header + result
}

func stringDecodeMaker(entry entryinfo, xmlReg map[string]entryinfo, variableName string) string {
	generatedVariable := fmt.Sprintf("\nfunc Decode%v(value []byte) (any, error){\n", variableName)
	genEncoder := fmt.Sprintf("\nfunc Encode%v(value %v) ([]byte, error){\n", variableName, variableName)
	stringType := xmlReg[entry.BaseType].Symbol
	// decode the string on the base character, there shouldn't be the need to decode just on the character

	switch stringType {
	case "UTF8Character":
		generatedVariable += "    runes := make([]rune, len(value))\n    i := 0\n    for len(value) > 0 {\n        r, size := utf8.DecodeRune(value)\n        runes[i] = r\n"
		generatedVariable += "        value = value[size:]\n        i++\n    }\n"
		genEncoder += "     result := []byte{}\n    for _, char := range value {\n           runeBytes := make([]byte,4)\n        write := utf8.EncodeRune(runeBytes,char)\n"
		genEncoder += "        result = append(result, runeBytes[:write]...)\n    }\n"
	case "Character": // "UTF16String":
		generatedVariable += "    parts := make([]uint16, len(value)/2)\n    for i := 0; i< len(value); i += 2 {\n        parts[i/2] = binary.BigEndian.Uint16(value[i:i+2])\n       }\n    runes := utf16.Decode(parts)\n"
		genEncoder += "parts := utf16.Encode([]rune(string(value)))\n    result, _  := EncodeTUInt16Array([]uint16(parts))\n"
	default: // default is just ascii
		generatedVariable += "    runes := make([]rune, len(value))\n    for i, char := range value {\n        runes[i] = rune(char)\n    }\n"
		genEncoder += "result := []byte(string(value))\n"
	}

	// ISO7 := "    runes := make([]rune, len(value))\n    for i, c := range g {        \nrunes[i] = rune(c)    \n}\n"
	generatedVariable += "    return string(runes) , nil\n}\n"
	genEncoder += "    return result , nil\n}\n"

	return generatedVariable + genEncoder
}
