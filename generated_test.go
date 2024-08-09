//  Copyright ©2019-2024  Mr MXF   info@mrmxf.com
//  BSD-3-Clause License  https://opensource.org/license/bsd-3-clause/

package main

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"testing"

	mxf2go "github.com/metarex-media/mxf-to-go-generator/generated"
	. "github.com/smartystreets/goconvey/convey"
)

// test the generated files here to see they match the behaviour etc

func TestLabels(t *testing.T) {

	/*
		does it match this workflow

		// getting all the register information of a label, from knowing the labels symbol (as go)
		labelWouldLike := generatedmrx.LabelsRegister[TransferCharacteristic_Gamma_2_6]

		// getting the label from a known univeral label
		sometarget := generatedmrx.LabelsLookUp["urn:smpte:ul:060e2b34.04010101.01010101.01010000"]

	*/

	labelInt := []int{
		mxf2go.TransferCharacteristic_Gamma_2_6,
		mxf2go.J2K_4KIMF_SingleTileLossyProfile_M11S7,
		mxf2go.SMPTE2035MultiChannelProgramCombo11f,
		mxf2go.MXFGCAVCNALUnitStreamWithVideoStream5SIDCustomStripeWrapped,
		// throw some random ones in there to weed out anything poor
		rand.Intn(len(mxf2go.LabelsRegister)),
		rand.Intn(len(mxf2go.LabelsRegister)),
		rand.Intn(len(mxf2go.LabelsRegister)),
	}

	for _, l := range labelInt {
		arrLabel := mxf2go.LabelsRegister[l]

		mapLabel := mxf2go.LabelsLookUp[arrLabel.UL]

		Convey("Checking that the labels can be used as intended", t, func() {
			Convey(fmt.Sprintf("Getting the label at position %v and comparing it to the label of the same UL", l), func() {
				Convey("The array label matches the map label", func() {
					So(arrLabel, ShouldResemble, mapLabel)
				})
			})
		})

	}

}

func TestEncodeGroups(t *testing.T) {

	type Encoder interface {
		Encode(primer *mxf2go.Primer) ([]byte, error)
	}

	idid := mxf2go.TUUID([16]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16})

	// test different structs with interesting fields
	identifier := mxf2go.GIdentificationStruct{InstanceID: idid, ApplicationSupplierName: []rune("metarex.media"), ApplicationName: []rune("An example slice of code"),
		ApplicationVersionString: []rune("0.0.1"),
		GenerationID:             mxf2go.TAUID{Data1: 3242, Data2: 232, Data3: 492, Data4: mxf2go.TUInt8Array8{60, 18}}}

	pre := mxf2go.GPrefaceStruct{InstanceID: idid,
		FileLastModified: mxf2go.TTimeStamp{Date: mxf2go.TDateStruct{Year: 2000, Month: 12, Day: 31},
			Time: mxf2go.TTimeStruct{Hour: 23, Minute: 59, Second: 59, Fraction: 12}},
		ContentStorageObject: mxf2go.TStrongReference([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}),
		FormatVersion:        mxf2go.TVersionType{VersionMajor: 1, VersionMinor: 12},
		IdentificationList:   mxf2go.TIdentificationStrongReferenceVector{[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}},
	}

	cd := mxf2go.GClassDefinitionStruct{
		MetaDefinitionName: []rune("A test of the definition name"),
		ParentClass:        []byte{1, 2, 3, 4, 5, 6, 7, 8},
		IsConcrete:         0,
	}

	primeEncode := []Encoder{&identifier, &pre, &cd}
	expectedOutcome := []string{"./testdata/identification", "./testdata/preface", "./testdata/classdef"}

	for i, e := range primeEncode {
		p := mxf2go.NewPrimer()
		gotB, err := e.Encode(p)

		expecB, ferr := os.ReadFile(expectedOutcome[i])
		//	f.Write(gotB)

		Convey("Checking that the groups can encode into the correct bytes", t, func() {
			Convey(fmt.Sprintf("Encoding a struct of type %v into mxf bytes", reflect.TypeOf(e)), func() {
				Convey("the encoding is complete without errors and match the expected bytes", func() {
					So(ferr, ShouldBeNil)
					So(err, ShouldBeNil)
					So(gotB, ShouldResemble, expecB)
				})
			})
		})

	}

	type EncodeOther interface {
		Encode() ([]byte, error)
	}

	b := mxf2go.GBadRequestResponseStruct{ASMBadRequestCopy: []uint8{3, 35, 224, 53, 2, 3, 3, 96}, ASMResponse: 5}

	encode := []EncodeOther{&b}
	expectedOutcome2 := []string{"./testdata/badrequest"}

	for i, e := range encode {

		gotB, err := e.Encode()
		expecB, ferr := os.ReadFile(expectedOutcome2[i])

		Convey("Checking that the groups can encode into the correct bytes", t, func() {
			Convey(fmt.Sprintf("Encoding a struct of type %v into mxf bytes", reflect.TypeOf(e)), func() {
				Convey("the encoding is complete without errors and match the expected bytes", func() {
					So(ferr, ShouldBeNil)
					So(err, ShouldBeNil)
					So(gotB, ShouldResemble, expecB)
				})
			})
		})
	}
}

func TestDecodeGroups(t *testing.T) {

	files := []string{"./testdata/identification", "./testdata/preface", "./testdata/classdef"}
	// cheat the primer
	types := []Encoder{&mxf2go.GIdentificationStruct{},
		&mxf2go.GPrefaceStruct{},
		&mxf2go.GClassDefinitionStruct{}}

	expected := []map[string]any{
		{"ApplicationName": "An example slice of code", "ApplicationProductID": mxf2go.TAUID{Data1: 0, Data2: 0, Data3: 0, Data4: [8]uint8{0, 0, 0, 0, 0, 0, 0, 0}}, "ApplicationSupplierName": "metarex.media", "ApplicationVersionString": "0.0.1", "GenerationID": mxf2go.TAUID{Data1: 3242, Data2: 232, Data3: 492, Data4: [8]uint8{60, 18, 0, 0, 0, 0, 0, 0}}, "InstanceID": mxf2go.TUUID{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}},
		{"ContentStorageObject": mxf2go.TStrongReference{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, "DescriptiveSchemes": mxf2go.TAUIDSet{}, "EssenceContainers": mxf2go.TAUIDSet{}, "FileLastModified": mxf2go.TTimeStamp{Date: mxf2go.TDateStruct{Year: 2000, Month: 12, Day: 31}, Time: mxf2go.TTimeStruct{Hour: 23, Minute: 59, Second: 59, Fraction: 12}}, "FormatVersion": mxf2go.TVersionType{VersionMajor: 1, VersionMinor: 12}, "IdentificationList": mxf2go.TIdentificationStrongReferenceVector{mxf2go.TIdentificationStrongReference{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}}, "InstanceID": mxf2go.TUUID{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, "OperationalPattern": mxf2go.TAUID{Data1: 0, Data2: 0, Data3: 0, Data4: [8]uint8{0, 0, 0, 0, 0, 0, 0, 0}}},
		{"InstanceID": mxf2go.TUUID{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, "IsConcrete": mxf2go.TBoolean(0), "MetaDefinitionIdentification": mxf2go.TAUID{Data1: 0, Data2: 0, Data3: 0, Data4: [8]uint8{0, 0, 0, 0, 0, 0, 0, 0}}, "MetaDefinitionName": "A test of the definition name", "ParentClass": mxf2go.TWeakReference{1, 2, 3, 4, 5, 6, 7, 8}},
	}

	for i, file := range files {
		fileByt, ferr := os.ReadFile(file)
		results, err := decodeWrapper(fileByt, types[i])

		Convey("Checking that the groups can decode all their fields, when the fields are fed decode functions", t, func() {
			Convey(fmt.Sprintf("Feeding byte stream segments into %v", reflect.TypeOf(types[i])), func() {
				Convey("no errors are found and the output map matches the expected", func() {
					So(ferr, ShouldBeNil)
					So(err, ShouldBeNil)
					So(results, ShouldEqual, expected[i])
				})
			})
		})
	}
}

type Encoder interface {
	Encode(primer *mxf2go.Primer) ([]byte, error)
}

// decode wrapper extracts the group as a map string any
// It extracts the group name (the first 16 bytes)
// and gets the group map with the decode options in.
//
// Then it loops through the byte stream getting the shorthand keys
// and comparing them to the primer, then once it has the field it
// feeds the bytes into the decode function. The any object that is returned
// is placed in the map.
func decodeWrapper(groupBytes []byte, group Encoder) (map[string]any, error) {
	title := byteToHexHeader(groupBytes[:16])
	groups, ok := mxf2go.Groups[title]

	if !ok {
		return nil, fmt.Errorf("no group found for the key %v", title)
	}

	_, keyLength := BerDecode(groupBytes[16:])
	pos := 16 + keyLength

	end := len(groupBytes)

	// mock
	p := mxf2go.NewPrimer()
	_, _ = group.Encode(p)
	outMap := make(map[string]any)
	for pos < end {

		// All fields in this test are 2 length tag and length
		// so we do not need to calcualte the ber lengths

		fieldKey := groupBytes[pos : pos+2]
		var g mxf2go.Group
		for k, v := range p.Tags {
			if string(v) == string(fieldKey) {
				g = groups.Group[byteToHex([]byte(k))]
			}
		}

		if g.UL == "" {
			return nil, fmt.Errorf("no field found for the shorthand key of %v", fieldKey)
		}

		// 2 byte length of length
		length := binary.BigEndian.Uint16(groupBytes[pos+2 : pos+4])

		out, err := g.Decode(groupBytes[pos+4 : pos+4+int(length)])
		if err != nil {
			return nil, err
		}
		outMap[g.UL] = out

		pos += int(length) + 4
	}

	return outMap, nil
}

func byteToHexHeader(hexb []byte) string {
	return fmt.Sprintf("urn:smpte:ul:%02x%02x%02x%02x.%02x7f%02x%02x.%02x%02x%02x%02x.%02x%02x%02x%02x",
		hexb[0], hexb[1], hexb[2], hexb[3], hexb[4], hexb[6], hexb[7], hexb[8],
		hexb[9], hexb[10], hexb[11], hexb[12], hexb[13], hexb[14], hexb[15])
}

func byteToHex(hexb []byte) string {
	return fmt.Sprintf("urn:smpte:ul:%02x%02x%02x%02x.%02x%02x%02x%02x.%02x%02x%02x%02x.%02x%02x%02x%02x",
		hexb[0], hexb[1], hexb[2], hexb[3], hexb[4], hexb[5], hexb[6], hexb[7], hexb[8],
		hexb[9], hexb[10], hexb[11], hexb[12], hexb[13], hexb[14], hexb[15])
}

// BerDecode decodes BERencoded lengths up to 9 bytes long
// including the indentifier byte.
func BerDecode(num []byte) (length int, encodeLength int) {

	if len(num) == 0 {
		return 0, 0
	}
	// mxf doesn;t exceed a length of 9
	// which is 1 giving the length and 8 bytes
	start := num[0]
	if start < 0x7f {
		return int(start), 1
	} else {

		// take the 4 lsbf for the length
		length := 0x0f & start

		if length > 8 {
			return 0, 0
		}

		complete := make([]byte, 8)
		// lengthproxy := int(length)
		postion := 7

		if int(length) > len(num)-1 {
			length = uint8(len(num) - 1)
		}

		for lengthproxy := int(length); lengthproxy > 0; lengthproxy-- {
			complete[postion] = num[lengthproxy]
			postion--
			// lengthproxy--
		}
		order := binary.BigEndian
		// 8 is the identifier
		return int(order.Uint64(complete)), int(length + 1)
	}

}
