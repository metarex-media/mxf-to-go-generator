//  Copyright ©2019-2024  Mr MXF   info@mrmxf.com
//  BSD-3-Clause License  https://opensource.org/license/bsd-3-clause/

package main

import (
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
	"sync"
	"time"
)

type LabelsRegister struct {
	Entries entry `xml:"Entries,omitempty"`
}

type entry struct {
	Entry []entryinfo `xml:"Entry,omitempty"`
}

type Extension struct {
	MetaDefinition metaDefinitions `xml:"MetaDefinitions,omitempty"`
	SchemeID       string          `xml:"SchemeID,omitempty"`
	SchemeURI      string          `xml:"SchemeURI,omitempty"`
	Other          interface{}     `xml:",any"`
}

type metaDefinitions struct {
	// Entry []typeDefinitionInteger `xml:"TypeDefinitionInteger,omitempty"`
	Else []typeDefinitionInteger `xml:",any"`
}

type typeDefinitionInteger struct {
	Symbol         string `xml:"Symbol,omitempty"`
	Identification string `xml:"Identification,omitempty"`
	Name           string `xml:"Name,omitempty"`
	Size           int    `xml:"Size,omitempty"`
	RenamedType    string `xml:"RenamedType,omitempty"`
	ElementType    string `xml:"ElementType,omitempty" yaml:"ElementType,omitempty"`
	Members        member `xml:"Members,omitempty"`
}

type member struct {
	Names []string `xml:"Name"`
	Type  []string `xml:"Type"`
}

type entryinfo struct {
	Register         string   `xml:"Register,omitempty" yaml:"Register,omitempty"`
	UL               string   `xml:"UL,omitempty" yaml:"UL,omitempty"`
	NamespaceName    string   `xml:"NamespaceName,omitempty" yaml:"NamespaceName,omitempty"`
	Kind             string   `xml:"Kind,omitempty" yaml:"Kind,omitempty"`
	Name             string   `xml:"Name,omitempty" yaml:"Name,omitempty"`
	Symbol           string   `xml:"Symbol,omitempty" yaml:"Symbol,omitempty"`
	Definition       string   `xml:"Definition,omitempty" yaml:"Definition,omitempty"`
	DefiningDocument string   `xml:"DefiningDocument,omitempty" yaml:"DefiningDocument,omitempty"`
	IsDeprecated     bool     `xml:"IsDeprecated,omitempty" yaml:"IsDeprecated,omitempty"`
	Contents         contents `xml:"Contents,omitempty" yaml:"Contents,omitempty"`
	Facets           facets   `xml:"Facets,omitempty" yaml:"Facets,omitempty"`
	TypeKind         string   `xml:"TypeKind,omitempty" yaml:"TypeKind,omitempty"`
	BaseType         string   `xml:"BaseType,omitempty" yaml:"BaseType,omitempty"`
	Type             string   `xml:"Type,omitempty" yaml:"Type,omitempty"`
	Parent           string   `xml:"Parent,omitempty" yaml:"Parent,omitempty"`
	TypeSize         int      `xml:"TypeSize,omitempty" yaml:"TypeSize,omitempty"`
	KLVSyntax        string   `xml:"KLVSyntax,omitempty" yaml:"KLVSyntax,omitempty"`
	Base             bool
	// baseDecodetype   string
}

type contents struct {
	Record []record `xml:"Record,omitempty"`
}

type record struct {
	UL         string `xml:"UL,omitempty" yaml:"UL,omitempty"`
	IsOptional string `xml:"IsOptional,omitempty" yaml:"IsOptional,omitempty"`
	LocalTag   string `xml:"LocalTag,omitempty"` // local tags are unique as well
}

type facets struct {
	Facet []facet `xml:"Facet,omitempty"`
}

type facet struct {
	Type         string `xml:"Type,omitempty" yaml:"Type,omitempty"`
	Value        string `xml:"Value,omitempty" yaml:"Value,omitempty"`
	Symbol       string `xml:"Symbol,omitempty" yaml:"Symbol,omitempty"`
	Major        string `xml:"Major,omitempty" yaml:"Major,omitempty"`
	Applications string `xml:"Applications,omitempty" yaml:"Applications,omitempty"`
	IsDeprecated bool   `xml:"IsDeprecated,omitempty" yaml:"IsDeprecated,omitempty"`
}

type syncmap struct {
	mapper map[string]int
	sync   *sync.Mutex
}

var decode = `func Decode%v(value []byte) (any, error){

	%v
}
`

var encode = `func Encode%v(value %v) ([]byte, error){

	%v
}
`

var tnow = time.Now

// encodeBaseTypes generates the register base types as go code
// giving the decode function and the length of the input required for that decode.
func encodeBaseTypes(bases map[string]entryinfo, destination io.Writer) (syncmap, error) {

	// update to use the master commit tag or something
	topComment := "// Copyright ©2019-2024  Mr MXF   info@mrmxf.com\n// BSD-3-Clause License  https://opensource.org/license/bsd-3-clause/\n//\n"
	topComment += fmt.Sprintf("// Package mxf2go was made using the smpte registers (https://registry.smpte-ra.org) on %v\n// if this seems out of date you will want to regenerated from the smpte library, to ensure compatibility\n", tnow())
	topComment += "package mxf2go\nimport (\n    \"encoding/binary\"\n    \"fmt\"\n)\n"
	_, err := destination.Write([]byte(topComment))

	if err != nil {
		return syncmap{}, err
	}

	keys := make([]string, len(bases))
	i := 0
	for k := range bases {
		keys[i] = k
		i++
	}
	slices.Sort(keys)
	// generate the lengths of the functions
	var mu sync.Mutex
	lengths := syncmap{mapper: make(map[string]int), sync: &mu}

	for _, k := range keys {
		v := bases[k]
		// generate the decode fucntions for the base types
		t := string(v.Register[0]) + strings.ReplaceAll(v.Symbol, "[]", "_")

		contents := "\n" // fmt.Sprintf("type %v struct {\n    %v\n} \n\n", t, strings.ToLower(v.Symbol)))
		fill := typeToFilling(k, lengths)
		contents += fmt.Sprintf(decode, t, fill)
		// make an encode here
		if v.TypeKind != "Character" {
			fillEn := typeToFillingEncode(k)
			contents += fmt.Sprintf(encode, t, strings.ToLower(v.Symbol), fillEn)
		}
		contents += "\n"
		_, err = destination.Write([]byte(contents))

		if err != nil {
			return syncmap{}, err
		}
	}
	lengths.mapper["TStrongReference"] = 0
	lengths.mapper["TWeakReference"] = 0
	// Add strong reference work around
	strongRef := "type TStrongReference []byte\ntype TWeakReference []byte\n"
	strongRef += "func DecodeTStrongReference(value []byte) (any, error)  {\n      field , _ := DecodeTUInt8Array(value)\n"
	strongRef += "	result := TStrongReference(field.(TUInt8Array))\n    return result, nil\n}"

	_, err = destination.Write([]byte(strongRef + marshall))
	if err != nil {
		return syncmap{}, err
	}

	return lengths, nil
}

var marshall = `
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
	//generate the name

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

`

// encodeTypes save the types register as go code
func encodeTypes(xmlReg map[string]entryinfo, lengths syncmap, regLocation string, destination io.Writer) error {

	typeBytes, err := os.ReadFile(regLocation)
	if err != nil {
		return err
	}
	var typeRegister LabelsRegister

	err = xml.Unmarshal(typeBytes, &typeRegister)
	if err != nil {
		return err
	}

	// add the package and import
	header := "// Copyright ©2019-2024  Mr MXF   info@mrmxf.com\n// BSD-3-Clause License  https://opensource.org/license/bsd-3-clause/\n"
	header += "package mxf2go\n\n"
	header += "import (\n"
	header += "    \"encoding/binary\"\n"
	header += "	\"unicode/utf16\"\n"
	header += "	\"unicode/utf8\"\n"
	header += ")\n\n"

	_, err = destination.Write([]byte(header))

	if err != nil {
		return err
	}

	var locker sync.Mutex
	bufferLock := &locker

	var wait sync.WaitGroup
	waiter := &wait

	for _, reger := range typeRegister.Entries.Entry {

		// if this the entry is valid type and not a base type
		if reger.BaseType != "" || reger.TypeKind != "" {

			// concurrently run everything to account for some decode functions
			// using values that are generated later, this allows the register
			// to not be sorted into types first
			var outError error
			waiter.Add(1)
			reg := reger // remove the pointer
			go func() {
				defer wait.Done()

				if tp, ok := xmlReg[reg.BaseType]; ok && tp.Base {

					// if there something points to a rename
					res := typeGenerator(reg, tp, xmlReg, lengths)
					bufferLock.Lock()
					_, outError = destination.Write([]byte(res))
					bufferLock.Unlock()
				} else {

					// if there's a record send it through we no base type as that is not used
					if reg.TypeKind == "Record" {
						res := typeGenerator(reg, entryinfo{}, xmlReg, lengths)
						bufferLock.Lock()
						_, outError = destination.Write([]byte(res))
						bufferLock.Unlock()
					} else if !xmlReg[reg.UL].Base {
						// fmt.Println(xmlReg[reg.BaseType], reg.Symbol, reg)
						res := typeGenerator(reg, xmlReg[reg.BaseType], xmlReg, lengths)
						bufferLock.Lock()
						_, outError = destination.Write([]byte(res))
						bufferLock.Unlock()
					}
				}

			}()

			if outError != nil {
				return outError
			}
		}
	}

	// wait until every thing has been written to the buffer
	waiter.Wait()
	_, err = destination.Write([]byte(hook))
	// flush everything to the file

	return err
}

var hook = `
func EncodeTUTF16StringArray(value TUTF16StringArray) ([]byte, error) {
	parts := utf16.Encode([]rune(string(value)))
	result, _ := EncodeTUInt16Array([]uint16(parts))
	return result, nil
}`

func encodeElements(xmlReg map[string]entryinfo, lengths syncmap, regLocation string, destination io.Writer) error {

	typeBytes, err := os.ReadFile(regLocation)

	if err != nil {
		return err
	}

	var elementsRegister LabelsRegister
	err = xml.Unmarshal(typeBytes, &elementsRegister)
	if err != nil {
		return err
	}

	_, err = destination.Write([]byte("// Copyright ©2019-2024  Mr MXF   info@mrmxf.com\n// BSD-3-Clause License  https://opensource.org/license/bsd-3-clause/\npackage mxf2go\n\n"))
	if err != nil {
		return err
	}

	for _, reg := range elementsRegister.Entries.Entry {

		if reg.Type != "" {
			res := typeGenerator(reg, xmlReg[reg.Type], xmlReg, lengths)
			_, err = destination.Write([]byte(res))
			if err != nil {
				return err
			}

		} else if reg.Kind != "NODE" {
			// fudge an any value in for the moment with elements with no type
			res := typeGenerator(reg, xmlReg["urn:smpte:ul:060e2b34.01040101.04100300.00000000"], xmlReg, lengths)
			_, err = destination.Write([]byte(res))
			if err != nil {
				return err
			}
		}
	}

	return err
}

func encodeGroups(xmlReg map[string]entryinfo, lengths syncmap, regLocation []string, destination io.Writer) error {
	var groupRegister LabelsRegister
	groupRegister.Entries = entry{}

	for _, reg := range regLocation {
		typeBytes, err := os.ReadFile(reg)
		if err != nil {
			return err
		}
		// fmt.Println(e)
		var groupRegisterInternal LabelsRegister
		err = xml.Unmarshal(typeBytes, &groupRegister)
		if err != nil {
			return err
		}
		//	fmt.Println(e, len(g3.Entries.Entry))
		groupRegister.Entries.Entry = append(groupRegister.Entries.Entry, groupRegisterInternal.Entries.Entry...)

	}

	// generate package and structs to be used
	top := "// Copyright ©2019-2024  Mr MXF   info@mrmxf.com\n// BSD-3-Clause License  https://opensource.org/license/bsd-3-clause/\n"
	top += "package mxf2go\n\n"
	top += "import (\n"
	top += "    \"encoding/binary\"\n"
	top += ")\n\n"

	top += "// Group contains the register properties of a group,\n"
	top += "// as well as the function to decode the bytes into a go value.\n"
	top += "type Group struct {\n"
	top += "	UL     string\n"
	top += "	IsOpt  bool\n"
	top += "	Length int\n"
	top += "	Decode func([]byte) (any, error)\n"
	top += "}\n\n"
	top += "// GroupID contains the name of a group and\n"
	top += "// all the possible fields it contains, which can be found with their\n"
	top += "// UL in the format\n"
	top += "// \"urn:smpte:ul:00000000.00000000.00000000.00000000\"\n"
	top += "type GroupID struct {\n"
	top += "	Name   string\n"
	top += "	Group  map[string]Group\n"
	top += "}\n\n"

	_, err := destination.Write([]byte(top + berCode))
	if err != nil {
		return err
	}

	endMap := "\n\n// Groups is a map of the groups in the groups register.\n// Values are found with their UL which takes the format\n// \"urn:smpte:ul:00000000.00000000.00000000.00000000\".\n//\n// Each group contains the individual fields and decode functions.\nvar Groups = map[string]GroupID {\n"
	tagMap := "\n\n// ShortHandLookUp contain a map of the short hand hex bytes as\n// declared  in the register and their corresponding UL.\nvar ShortHandLookUp = map[string]string {\n"

	tagMatches := make(map[string]bool)

	// generate the individual groups
	// and the complete lookup table of every group
	for _, reg := range groupRegister.Entries.Entry {

		// Extract every leaf value
		// generating a map vairable for each type
		if reg.Kind != "NODE" {
			res, generatedVarName := groupsExtractor(reg, xmlReg, lengths)
			_, err = destination.Write([]byte(res))
			if err != nil {
				return err
			}
			endMap += fmt.Sprintf("\"%v\" : {Name : \"%v\", Group:%v}, \n", reg.UL, string(generatedVarName[1:]), generatedVarName)

			for _, rec := range reg.Contents.Record {

				if rec.LocalTag != "" && !tagMatches[rec.LocalTag] {
					tagMap += fmt.Sprintf("\"%v\" : \"%v\", \n", rec.LocalTag, rec.UL)
					tagMatches[rec.LocalTag] = true
				}

			}
		}

	}
	endMap += "}"
	tagMap += "}"

	// add the endmap at the end

	_, err = destination.Write([]byte(endMap + tagMap))
	if err != nil {
		return err
	}

	return nil
}

var berCode = `
func BEREncode(number int) []byte {

switch {
case number < 127:
	return []byte{byte(number)}
case number < 0xff:
	return []byte{0x81, byte(number)}
case number < 0xffff:
	full := make([]byte, 2)
	binary.BigEndian.PutUint16(full, uint16(number))
	return append([]byte{0x82}, full...)
case number < 0xffffffff:
	full := make([]byte, 4)
	binary.BigEndian.PutUint32(full, uint32(number))
	return append([]byte{0x84}, full...)
default:
	full := make([]byte, 8)
	binary.BigEndian.PutUint64(full, uint64(number))
	return append([]byte{0x88}, full...)
}

}

`

func encodeEssence(regLocation []string, destination io.Writer) error {

	var essenceRegister LabelsRegister
	essenceRegister.Entries = entry{}

	for _, reg := range regLocation {
		typeBytes, err := os.ReadFile(reg)
		if err != nil {
			return err
		}
		// fmt.Println(e)
		var groupRegisterInternal LabelsRegister
		err = xml.Unmarshal(typeBytes, &essenceRegister)
		if err != nil {
			return err
		}
		//	fmt.Println(e, len(g3.Entries.Entry))
		essenceRegister.Entries.Entry = append(essenceRegister.Entries.Entry, groupRegisterInternal.Entries.Entry...)

	}

	top := "// Copyright ©2019-2024  Mr MXF   info@mrmxf.com\n// BSD-3-Clause License  https://opensource.org/license/bsd-3-clause/\n"
	top += "package mxf2go\n\n"

	top += "// EssenceInformation is the register layout of the essence information\n"
	top += "type EssenceInformation struct {\n"
	top += "	UL               string   `xml:\"UL,omitempty\"`\n"
	top += "	Name             string   `xml:\"Name,omitempty\"`\n"
	top += "	Symbol           string   `xml:\"Symbol,omitempty\"`\n"
	top += "	Definition       string   `xml:\"Definition,omitempty\"`\n"
	top += "	DefiningDocument string   `xml:\"DefiningDocument,omitempty\"`\n"
	top += "	IsDeprecated     bool     `xml:\"IsDeprecated,omitempty\"`\n"
	top += "}\n\n"
	top += "// EssenceLookUp is a map of the essence in the essence register.\n"
	top += "// Values are found with their UL which takes the format\n"
	top += "// \"urn:smpte:ul:00000000.00000000.00000000.00000000\"\n"
	top += "var EssenceLookUp = map[string]EssenceInformation{\n"

	_, err := destination.Write([]byte(top))
	if err != nil {
		return err
	}

	for _, reg := range essenceRegister.Entries.Entry {

		if reg.Kind != "NODE" {
			reg.Definition = strings.ReplaceAll(reg.Definition, "\n", "")
			// variable += fmt.Sprintf("    %v %v = %v\n", strings.ReplaceAll(entry.Facets.Facet[0].Symbol, " ", "_"), name, entry.Facets.Facet[0].Value)
			_, err = destination.Write([]byte(fmt.Sprintf("    \"%v\" : { UL: \"%v\",  Name: \"%v\", Symbol: \"%v\",  Definition: \"%v\", DefiningDocument:\"%v\", IsDeprecated: %v},\n",
				reg.UL, reg.UL, reg.Name, reg.Symbol, reg.Definition, reg.DefiningDocument, reg.IsDeprecated)))

			if err != nil {
				return err
			}
		}

	}

	_, err = destination.Write([]byte("}\n"))
	if err != nil {
		return err
	}

	return err
}

func encodeLabels(regLocation string, destination io.Writer) error {
	typeBytes, err := os.ReadFile(regLocation)

	if err != nil {
		return err
	}
	// fmt.Println(e)
	var labelLayout LabelsRegister
	err = xml.Unmarshal(typeBytes, &labelLayout)

	if err != nil {
		return err
	}

	header := "// Copyright ©2019-2024  Mr MXF   info@mrmxf.com\n// BSD-3-Clause License  https://opensource.org/license/bsd-3-clause/\n"
	header += "package mxf2go\n\n"
	header += "// LabelInformation is the register layout of the label information\n"
	header += "type LabelInformation struct {\n"
	header += "	UL               string   `xml:\"UL,omitempty\"`\n"
	header += "	Name             string   `xml:\"Name,omitempty\"`\n"
	header += "	Symbol           string   `xml:\"Symbol,omitempty\"`\n"
	header += "	Definition       string   `xml:\"Definition,omitempty\"`\n"
	header += "	DefiningDocument string   `xml:\"DefiningDocument,omitempty\"`\n"
	header += "	IsDeprecated     bool     `xml:\"IsDeprecated,omitempty\"`\n"
	header += "}\n"
	header += "const (\n"

	_, err = destination.Write([]byte(header))
	if err != nil {
		return err
	}

	start := 0
	// find the start position for entries
	for i, reg := range labelLayout.Entries.Entry {
		if reg.Kind != "NODE" {
			_, err = destination.Write([]byte(fmt.Sprintf("    %v = iota \n", reg.Symbol)))
			if err != nil {
				return err
			}

			start = i

			break
		}
	}
	for _, reg := range labelLayout.Entries.Entry[start+1:] {

		if reg.Kind != "NODE" {
			// variable += fmt.Sprintf("    %v %v = %v\n", strings.ReplaceAll(entry.Facets.Facet[0].Symbol, " ", "_"), name, entry.Facets.Facet[0].Value)

			_, err = destination.Write([]byte(fmt.Sprintf("    %v\n", reg.Symbol)))
			if err != nil {
				return err
			}
		}

	}

	_, err = destination.Write([]byte(")\n// LabelsRegister is an array of the labels in the labels register.\n// Values are found with the positional constants e.g. mxf2go.MXF_GC_IAData_Frame_Wrapped\nvar LabelsRegister = []LabelInformation{\n"))
	if err != nil {
		return err
	}

	for _, reg := range labelLayout.Entries.Entry[start:] {

		if reg.Kind != "NODE" {
			reg.Definition = strings.ReplaceAll(reg.Definition, "\n", "")
			// variable += fmt.Sprintf("    %v %v = %v\n", strings.ReplaceAll(entry.Facets.Facet[0].Symbol, " ", "_"), name, entry.Facets.Facet[0].Value)
			_, err = destination.Write([]byte(fmt.Sprintf("    { UL: \"%v\",  Name: \"%v\", Symbol: \"%v\",  Definition: \"%v\", DefiningDocument:\"%v\", IsDeprecated: %v},\n",
				reg.UL, reg.Name, reg.Symbol, reg.Definition, reg.DefiningDocument, reg.IsDeprecated)))

			if err != nil {
				return err
			}

		}

	}

	_, err = destination.Write([]byte("}\n\n// LabelsLookUp is a map of the labels in the labels register.\n// Values are found with their UL which takes the format\n// \"urn:smpte:ul:00000000.00000000.00000000.00000000\"\nvar LabelsLookUp = map[string]LabelInformation{\n"))
	if err != nil {
		return err
	}

	for _, reg := range labelLayout.Entries.Entry[start:] {

		if reg.Kind != "NODE" {
			reg.Definition = strings.ReplaceAll(reg.Definition, "\n", "")
			// variable += fmt.Sprintf("    %v %v = %v\n", strings.ReplaceAll(entry.Facets.Facet[0].Symbol, " ", "_"), name, entry.Facets.Facet[0].Value)
			_, err := destination.Write([]byte(fmt.Sprintf("    \"%v\" : { UL: \"%v\",  Name: \"%v\", Symbol: \"%v\",  Definition: \"%v\", DefiningDocument:\"%v\", IsDeprecated: %v},\n",
				reg.UL, reg.UL, reg.Name, reg.Symbol, reg.Definition, reg.DefiningDocument, reg.IsDeprecated)))
			if err != nil {
				return err
			}
		}

	}
	_, err = destination.Write([]byte("}\n"))
	if err != nil {
		return err
	}

	return nil
}

// TypeGenerator extracts the type of an entry,
// then generates the go equivalent, wrapping any
func typeGenerator(entry, dest entryinfo, types map[string]entryinfo, lengths syncmap) string {

	// update the symbol to
	if entry.Symbol == "type" {
		entry.Symbol = "Type"
	}

	var generatedVariable string

	// replace the symbols go doesn't liek in variable names
	variableName := strings.ReplaceAll(entry.Symbol, " ", "_")
	variableName = strings.ReplaceAll(variableName, "-", "_")
	variableName = string(entry.Register[0]) + variableName
	/*
	   func EncodeTLocalTagType(value TLocalTagType) ([]byte, error) {
	   	result, _ := EncodeTUInt16(uint16(value))
	   	return result, nil
	   }*/
	// the switch statement generates the type and any accompanying decode methods
	switch entry.TypeKind {
	case "Record":

		// record is a struct that can have any number of fields
		generatedVariable = "struct {\n"

		// get the strut fields
		for _, f := range entry.Facets.Facet {
			var targetType string
			// get the base type of this type
			base := findBaseOg(types[f.Type], types)
			if base.Base {
				targetType = strings.ToLower(types[f.Type].Symbol)
			} else {
				// if not then find that decode field.Val = field.Val(in[start:end])
				targetName := strings.ReplaceAll(types[f.Type].Symbol, "-", "_")
				targetType = string(types[f.Type].Register[0]) + strings.ReplaceAll(targetName, " ", "_")
			}
			// update with name and type
			generatedVariable += fmt.Sprintf("    %v %v\n", strings.ReplaceAll(f.Symbol, " ", "_"), targetType)
		}
		generatedVariable += "}\n"

		// make the decode bit here
		generatedVariable += decodeMaker(entry.Facets.Facet, variableName, lengths, types)

	case "Enumeration":
		// enum is a type where enumerated values have a meaning
		var targetType string
		if !dest.Base {
			targetType = string(dest.Register[0]) + dest.Symbol
		} else {
			targetType = strings.ToLower(dest.Symbol) // this is as the base is UInt16
		}

		generatedVariable = fmt.Sprintf("%v\n\n", targetType)

		//	}
		// generate the enumerated values
		if len(entry.Facets.Facet) > 0 {

			// if it's string make the name
			if entry.Facets.Facet[0].Symbol != "" {
				// if there are fields make the string and marshall text functions

				generatedVariable += fmt.Sprintf("func (mt %v) MarshalText() ([]byte, error) {\n", variableName) // for use with the son and yaml marshallers
				generatedVariable += "    return []byte(mt.String()), nil\n}\n\n"

				if types[entry.BaseType].TypeKind == "String" {
					generatedVariable += fmt.Sprintf("func (s %v)String() string {\n\n     switch string(s) {\n", variableName)
					for _, f := range entry.Facets.Facet {
						generatedVariable += fmt.Sprintf("    case \"%v\":\n        return \"%v\" \n", f.Value, strings.ReplaceAll(f.Symbol, " ", "_"))
						// generatedVariable += fmt.Sprintf("    %v_%v %v = []rune(\"%v\")\n", variableName, strings.ReplaceAll(f.Symbol, " ", "_"), variableName, f.Value)
					}
				} else {
					generatedVariable += fmt.Sprintf("func (s %v)String() string {\n\n     switch s {\n", variableName)
					for _, f := range entry.Facets.Facet {
						generatedVariable += fmt.Sprintf("    case %v:\n        return \"%v\" \n", f.Value, strings.ReplaceAll(f.Symbol, " ", "_"))
						// generatedVariable += fmt.Sprintf("    %v_%v  %v = %v \n", variableName, strings.ReplaceAll(f.Symbol, " ", "_"), variableName, f.Value)
					}
				}
				generatedVariable += "    default:\n        return \"invalid value\"\n"
				generatedVariable += "    }\n}"
			}

		}

		// add the decode method
		generatedVariable += "\n" + decodeWrapMaker(entry, variableName, lengths, types)

	case "VariableArray", "Set":
		// set is a type of array where the size and length are declared before decoding
		// variable arrays also do this, if they are not an array of a base value. e.g. []UID

		var targetType string
		if !dest.Base {
			targetType = string(dest.Register[0]) + dest.Symbol
		} else {
			targetType = strings.ToLower(dest.Symbol)
		}

		generatedVariable = fmt.Sprintf("[]%v\n", targetType)
		// generate decode function
		generatedVariable += arrDecodeMaker(entry, variableName, fmt.Sprintf("[]%v", targetType), lengths, types)

		if entry.Symbol == "UTF16StringArray" { // if of this type replace the characters with strings as the UL does not line uo with its type
			// manually intervene with this one
			generatedVariable = `string
	func DecodeTUTF16StringArray(value []byte) (any, error) {
		field, _ := DecodeTUTF16String(value)
		result := TUTF16StringArray(field.(string))

		return result, nil
	}`
		}
		// variable data decode will just loop until its all gone
	case "FixedArray":

		targetType := ""
		if !dest.Base {
			targetType = string(dest.Register[0]) + dest.Symbol
		} else {
			targetType = strings.ToLower(dest.Symbol)
		}
		// array of fixed size [int]type#
		//	symb = string(dest.Register[0]) + dest.Symbol
		generatedVariable = fmt.Sprintf("[%v]%v\n", entry.TypeSize, targetType)
		generatedVariable += arrDecodeMaker(entry, variableName, fmt.Sprintf("[%v]%v", entry.TypeSize, targetType), lengths, types)
		// update these to be structs
		// make an for loop type size

	case "WeakReference":

		// variable = "\r var " + name + "=" + string(dest.Register[0]) + dest.Symbol
		generatedVariable = "TWeakReference"
	case "StrongReference":
		generatedVariable = "TStrongReference"
	case "Rename":
		// this is a remapping of the name to whatever the basetype is
		if !dest.Base {
			generatedVariable = string(dest.Register[0]) + dest.Symbol
		} else {
			generatedVariable = strings.ToLower(dest.Symbol)
		}
		// then add the decode, which just wraps whatever it is renaming
		generatedVariable += "\n" + decodeWrapMaker(entry, variableName, lengths, types)

	case "String": // replace with array of characters
		generatedVariable = "[]rune"
		// choose a method to decode, this will be based on their basetype string
		generatedVariable += stringDecodeMaker(entry, types, variableName)
	default:

		// default is a rename style with no generated decode
		if !dest.Base {

			generatedVariable = string(dest.Register[0]) + dest.Symbol
		} else {
			generatedVariable = strings.ToLower(dest.Symbol)
		}
	}

	// reset the word type to Type, making it Golang friendly
	if generatedVariable == "type" {
		generatedVariable = "Type"
	}

	return fmt.Sprintf("type %v %v\n", variableName, generatedVariable)
}

func groupsExtractor(entry entryinfo, types map[string]entryinfo, lengths syncmap) (string, string) {

	// groups is a map of UL and the decode functino for that type
	variable := "= map[string]Group {\n"
	var structVariable string
	// inherit all the fields from parent groups
	parentUL := entry.Parent
	groupName := strings.ReplaceAll(entry.Symbol, " ", "_")
	groupName = strings.ReplaceAll(groupName, "-", "_")

	// if entry.KLVSyntax == "53" {
	//		shorthand := 53
	//}

	encoder := fmt.Sprintf("func (g *%v%vStruct)Encode() ([]byte, error){\n    var result,field ,BERField []byte\n\n", string(entry.Register[0]), groupName)

	if entry.KLVSyntax == "53" || entry.KLVSyntax == "06 53" {
		encoder = fmt.Sprintf("func (g *%v%vStruct)Encode(primer *Primer) ([]byte, error){\n    var result,field ,BERField []byte\n\n", string(entry.Register[0]), groupName)
		// encoder = fmt.Sprintf("func (G *%v%vStruct)Encode(dynamicTag *uint16, dynamicTags map[string][]byte) ([]byte, error){\n    var result,field ,BERField []byte\n\n", string(entry.Register[0]), groupName)
	}

	for parentUL != "" {

		parentEntry := types[parentUL]
		// add the map contents of the parent
		mapvariable, structVar := mapMaker(parentEntry.Contents.Record, types, lengths)
		encoder += encodeMaker(parentEntry.Contents.Record, types, entry.KLVSyntax)

		variable += mapvariable
		structVariable += structVar
		// update the parent and repeat until no more parents
		parentUL = parentEntry.Parent
	}

	encoder += encodeMaker(entry.Contents.Record, types, entry.KLVSyntax)

	// finally add the map contents of the original group
	mapvariable, structVar := mapMaker(entry.Contents.Record, types, lengths)
	structVariable += structVar
	variable += mapvariable
	variable += "}"

	if len(structVariable) == 0 {
		encoder = fmt.Sprintf("func (g *%v%vStruct)Encode() ([]byte, error){\n    var result []byte\n\n", string(entry.Register[0]), groupName)
	}

	structVariable = fmt.Sprintf("type %v%vStruct struct {\n%v\n}", string(entry.Register[0]), groupName, structVariable)

	if len(entry.Contents.Record) != 0 {
		encodeMaker(entry.Contents.Record, types, entry.KLVSyntax)
	}

	hexb, _ := hex.DecodeString(strings.ReplaceAll(entry.UL[13:], ".", ""))
	// fmt.Printf("%032x, %v\n", hexb, strings.ReplaceAll(f.UL[13:], ".", ""))

	if entry.KLVSyntax == "53" || entry.KLVSyntax == "06 53" {
		encoder += fmt.Sprintf("    header := []byte{%v,%v,%v,%v,%v, 83,%v, %v,%v,%v,%v,%v,%v,%v,%v,%v}\n",
			hexb[0], hexb[1], hexb[2], hexb[3], hexb[4], hexb[6], hexb[7], hexb[8],
			hexb[9], hexb[10], hexb[11], hexb[12], hexb[13], hexb[14], hexb[15])
	} else {
		encoder += fmt.Sprintf("    header := []byte{%v,%v,%v,%v,%v, 2,%v, 1,%v,%v,%v,%v,%v,%v,%v,%v}\n",
			hexb[0], hexb[1], hexb[2], hexb[3], hexb[4], hexb[6], hexb[8],
			hexb[9], hexb[10], hexb[11], hexb[12], hexb[13], hexb[14], hexb[15])
	}
	// `    header := []byte{ 53}

	encoder += `
	totalLength := BEREncode(len(result))
	header = append(header, totalLength...)

	result = append(header, result...)
	return result, nil
}
`
	if len(entry.Contents.Record) == 0 && entry.Parent == "" {
		encoder = "\n"
	}

	return fmt.Sprintf("var %v%v %v\n%v\n%v\n", string(entry.Register[0]), groupName, variable, structVariable, encoder), string(entry.Register[0]) + groupName
}

func mapMaker(records []record, types map[string]entryinfo, lengths syncmap) (string, string) {
	var variable string

	// set up  struct here
	// then a decode

	var structFields string
	for _, f := range records {

		mapElement := types[f.UL]
		if mapElement.Type == "" {
			variable += fmt.Sprintf("    \"%v\" : 	{UL:\"%v\", Decode: DecodeTany},\n ", f.UL, types[f.UL].Symbol)
			if f.IsOptional == "false" || types[f.UL].Symbol == "InstanceID" {
				structFields += fmt.Sprintf("    %v any // IsOptional:%v\n", types[f.UL].Symbol, f.IsOptional)
			}
		} else {
			mapElement = types[mapElement.Type]
			//	mapElementName := string(mapElement.Register[0]) + strings.ReplaceAll(mapElement.Symbol, " ", "_")

			end := ",\n"
			//	{UL: "EPackageID", IsOpt: false, Decode: EUMIDVideor2{}.Decode}
			// find bass type here
			mapElementName, typeSymbol, mpaSymbol, _, _ := findBase(mapElement, types, "Elements")

			if mpaSymbol != "" {
				mapElementName = mpaSymbol
			}

			// manually update the references afterwards as these are not finished on their approach
			if strings.Contains(mapElementName, "StrongReference") && !strings.Contains(mapElementName, "StrongReferenceSet") && !strings.Contains(mapElementName, "StrongReferenceVector") {
				mapElementName = "TStrongReference"
			} else if strings.Contains(mapElementName, "Reference") && !strings.Contains(mapElementName, "ReferenceSet") && !strings.Contains(mapElementName, "ReferenceVector") {
				mapElementName = "TWeakReference"
			}

			// only update the structs if they aren't optional
			if f.IsOptional != "true" || types[f.UL].Symbol == "InstanceID" {
				if strings.ToLower(string(mapElementName[1:])) == typeSymbol {
					structFields += fmt.Sprintf("    %v %v // IsOptional:%v \n", types[f.UL].Symbol, typeSymbol, f.IsOptional)
				} else {
					structFields += fmt.Sprintf("    %v %v // IsOptional:%v\n", types[f.UL].Symbol, mapElementName, f.IsOptional)
				}

			}

			variable += fmt.Sprintf("    \"%v\" : 	{UL:\"%v\", Length : %v ,Decode: Decode%v}%v ", f.UL, types[f.UL].Symbol, lengths.mapper[mapElementName], mapElementName, end)
		}

	}
	return variable, structFields
}

func encodeMaker(records []record, types map[string]entryinfo, lengthform string) string {

	var encoderFill string

	for _, f := range records {
		if f.IsOptional != "true" || types[f.UL].Symbol == "InstanceID" {
			mapElement := types[f.UL]
			if mapElement.Type == "" {

				encoderFill += fmt.Sprintf("    field, _ = EncodeTany(g.%v)\n", types[f.UL].Symbol)

			} else {
				mapElement = types[mapElement.Type]

				// mapElementName := string(mapElement.Register[0]) + strings.ReplaceAll(mapElement.Symbol, " ", "_")

				//	{UL: "EPackageID", IsOpt: false, Decode: EUMIDVideor2{}.Decode}
				// find bass type here
				mapElementName, _, mpaSymbol, _, _ := findBase(mapElement, types, "Elements")

				if mpaSymbol != "" {
					mapElementName = mpaSymbol
				}

				// manually update the references afterwards as these are not finished on their approach
				if strings.Contains(mapElementName, "StrongReference") && !strings.Contains(mapElementName, "StrongReferenceSet") && !strings.Contains(mapElementName, "StrongReferenceVector") {
					mapElementName = "TStrongReference"
				} else if strings.Contains(mapElementName, "Reference") && !strings.Contains(mapElementName, "ReferenceSet") && !strings.Contains(mapElementName, "ReferenceVector") {
					mapElementName = "TWeakReference"
				}

				// the utilise this as a string filler bit
				//	fmt.Println(mapElementName, mpaSymbol, tsim)
				encoderFill += fmt.Sprintf("    field, _ = Encode%v(g.%v)\n", mapElementName, types[f.UL].Symbol)

			}

			// fmt.Printf("%032x, %v\n", hexb, strings.ReplaceAll(f.UL[13:], ".", ""))

			// if shortform == "53" o r 06 53
			hexb, _ := hex.DecodeString(strings.ReplaceAll(f.UL[13:], ".", ""))
			// f.Shortag encoderfill += the orher
			if lengthform == "53" || lengthform == "06 53" {
				// if there is a predefined tag than use that
				// otherwise use a dynamic system
				if len(f.LocalTag) == 4 {
					//	fmt.Println(f.LocalTag, "LOCAL TAG", f)
					hexbLocalTag, _ := hex.DecodeString(f.LocalTag)
					encoderFill += fmt.Sprintf("    result = append(result,[]byte{%v,%v}...)\n",
						hexbLocalTag[0], hexbLocalTag[1])
					// encoderFill += fmt.Sprintf("    dynamicTags[string([]byte{%v,%v})]= %v\n", hexbLocalTag[0], hexbLocalTag[1], hexToByte(hexb))
					encoderFill += fmt.Sprintf("    primer.AddEntry(%v,[]byte{%v,%v})\n", hexToByte(hexb), hexbLocalTag[0], hexbLocalTag[1])
					encoderFill += "    BERField = binary.BigEndian.AppendUint16([]byte{}, uint16(len(field)))\n"
				} else {

					codestring := strings.ReplaceAll(f.UL[13:], ".", "")
					//		encoderFill += fmt.Sprintf("    code%v := binary.BigEndian.AppendUint16([]byte{}, *dynamicTag)\n    *dynamicTag--\n", codestring)
					//		encoderFill += fmt.Sprintf("    dynamicTags[string(code%v)]= %v\n", codestring, hexToByte(hexb))
					encoderFill += fmt.Sprintf("    code%v := primer.AddEntry(%v,[]byte{})\n", codestring, hexToByte(hexb))
					encoderFill += fmt.Sprintf("    result = append(result,code%v...)\n", codestring)
					encoderFill += "    BERField = binary.BigEndian.AppendUint16([]byte{}, uint16(len(field)))\n"
				}
			} else {

				encoderFill += fmt.Sprintf("    result = append(result,%v...)\n", hexToByte(hexb))
				encoderFill += "    BERField = BEREncode(len(field))\n"
			}

			encoderFill += "    result = append(result, BERField...)\n"
			encoderFill += "    result = append(result, field...)\n\n"
		}
	}

	/*func (G *GLinkEncryptionKeyQueryIDResponseStruct) Encode() ([]byte, error) {
	result := []byte{}

	//add the key first
	result = append(result, []byte{0x53}...)
	field, _ := EncodeTUInt32(G.ASMRequestID)
	BERField := BEREncode(len(field))
	result = append(result, BERField...)
	result = append(result, field...)

	field1, _ := EncodeTUInt8(G.ASMKeyPresentFlag)
	result = append(result, field1...)

	field2, _ := EncodeTUInt32(uint32(G.ASMResponse))
	result = append(result, field2...)

	// generate the header and length after everything has been written
	// 6th byte is 02 for the moment
	// 8th is 1
	header := []byte{}
	totalLength := BEREncode(len(result))
	header = append(header, totalLength...)

	result = append(header, result...)*/
	return encoderFill
}

func hexToByte(hexb []byte) string {
	return fmt.Sprintf("[]byte{%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v}",
		hexb[0], hexb[1], hexb[2], hexb[3], hexb[4], hexb[5], hexb[6], hexb[7], hexb[8],
		hexb[9], hexb[10], hexb[11], hexb[12], hexb[13], hexb[14], hexb[15])
}
