//  Copyright ©2019-2024  Mr MXF   info@mrmxf.com
//  BSD-3-Clause License  https://opensource.org/license/bsd-3-clause/

package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestElements(t *testing.T) {
	// Read set up etc

	// read every register value and extract the bases from them
	setUpFiles := []string{"./registers/Elements.xml", "./registers/Groups.xml", "./registers/Labels.xml", "./registers/Types.xml",
		"./dev/elements-dev.xml", "./dev/groups-dev.xml",
		"./dev/elements-mrx-dev.xml", "./dev/groups-mrx-dev.xml"}
	xmlReg, bases, btErr := baseSetUp("./registers/Types.xml", setUpFiles)

	Convey("Checking that the base set up runs without error", t, func() {
		Convey(fmt.Sprintf("Running with the %v", setUpFiles), func() {
			Convey("no errors occur during set up", func() {
				So(btErr, ShouldBeNil)
			})
		})
	})

	/////////////////
	// baseTypes
	////////////////
	baseBuf := bytes.NewBuffer([]byte{})
	tnow = func() time.Time { return time.Time{} }
	// base types is an extra file to decode the base types, e.g. uint16
	lengths, btErr := encodeBaseTypes(bases, baseBuf)
	// fbase, _ := os.Create("./testdata/testRegisters/basetypes-out")
	// fbase.Write(baseBuf.Bytes())
	expected, expecErr := os.ReadFile("./testdata/testRegisters/basetypes-out")

	Convey("Checking that the base types are correctly generated", t, func() {
		Convey("encoding the standard base types and comparing to ./testdata/testRegisters/basetypes-out", func() {
			Convey("tno error is returned and the file matches exactly", func() {
				So(btErr, ShouldBeNil)
				So(expecErr, ShouldBeNil)
				So(string(expected), ShouldResemble, baseBuf.String())
			})
		})
	})

	/////////////////
	// Elements
	////////////////

	eleBuf := bytes.NewBuffer([]byte{})
	eleErr := encodeElements(xmlReg, lengths, "./testdata/testRegisters/elements.xml", eleBuf)
	expected, expecErr = os.ReadFile("./testdata/testRegisters/elements-out")

	Convey("Checking that the elements fields and transformed into go", t, func() {
		Convey("Encoding ./testdata/testRegisters/elements.xml as go code", func() {
			Convey("no errors are generated and the output matches", func() {
				So(eleErr, ShouldBeNil)
				So(expecErr, ShouldBeNil)
				So(string(expected), ShouldResemble, eleBuf.String())
			})
		})
	})

	/////////////////
	// Labels
	////////////////
	labelBuf := bytes.NewBuffer([]byte{})
	labelErr := encodeLabels("./testdata/testRegisters/labels.xml", labelBuf)
	// fLabel, _ := os.Create("./testdata/testRegisters/labels-out")
	// fLabel.Write(labelBuf.Bytes())
	expected, expecErr = os.ReadFile("./testdata/testRegisters/labels-out")
	Convey("Checking that the labels fields and transformed into go", t, func() {
		Convey("Encoding ./testdata/testRegisters/labels.xml as go code", func() {
			Convey("no errors are generated and the output matches", func() {
				So(labelErr, ShouldBeNil)
				So(expecErr, ShouldBeNil)
				So(string(expected), ShouldResemble, labelBuf.String())
			})
		})
	})

	/////////////////
	// Essence
	////////////////

	essBuf := bytes.NewBuffer([]byte{})
	essErr := encodeEssence([]string{"./testdata/testRegisters/essence.xml"}, essBuf)
	//	fEss, _ := os.Create("./testdata/testRegisters/essence-out")
	//	fEss.Write(essBuf.Bytes())
	expected, expecErr = os.ReadFile("./testdata/testRegisters/essence-out")
	Convey("Checking that the labels fields and transformed into go", t, func() {
		Convey("Encoding ./testdata/testRegisters/essence.xml as go code", func() {
			Convey("no errors are generated and the output matches", func() {
				So(labelErr, ShouldBeNil)
				So(essErr, ShouldBeNil)
				So(string(expected), ShouldResemble, essBuf.String())
			})
		})
	})

	//////////////
	// Types
	/////////////

	inFiles := []string{"./testdata/testRegisters/types/types-enum.xml",
		"./testdata/testRegisters/types/types-int.xml",
		"./testdata/testRegisters/types/types-ref.xml",
		"./testdata/testRegisters/types/types-string.xml",
		"./testdata/testRegisters/types/types-struct.xml",
		"./testdata/testRegisters/types/types-arr.xml",
		"./testdata/testRegisters/types/types-fixedarr.xml"}
	out := []string{"./testdata/testRegisters/types/types-enum-out",
		"./testdata/testRegisters/types/types-int-out",
		"./testdata/testRegisters/types/types-ref-out",
		"./testdata/testRegisters/types/types-string-out",
		"./testdata/testRegisters/types/types-struct-out",
		"./testdata/testRegisters/types/types-arr-out",
		"./testdata/testRegisters/types/types-fixedarr-out"}
	for i, in := range inFiles {

		typeBuf := bytes.NewBuffer([]byte{})
		typeErr := encodeTypes(xmlReg, lengths, in, typeBuf)

		expected, expecErr = os.ReadFile(out[i])
		Convey("Checking that the individual type fields and transformed into go", t, func() {
			Convey(fmt.Sprintf("Encoding %s as go code", in), func() {
				Convey("no errors are generated and the output matches", func() {
					So(labelErr, ShouldBeNil)
					So(typeErr, ShouldBeNil)
					So(string(expected), ShouldResemble, typeBuf.String())
				})
			})
		})
	}

	/////////////////
	// Groups
	////////////////

	groupBuf := bytes.NewBuffer([]byte{})
	groupErr := encodeGroups(xmlReg, lengths, []string{"./testdata/testRegisters/groups.xml"}, groupBuf)
	//	fgroup, _ := os.Create("./testdata/testRegisters/groups-out")
	//	fgroup.Write(groupBuf.Bytes())
	expected, expecErr = os.ReadFile("./testdata/testRegisters/groups-out")
	Convey("Checking that the individual type fields and transformed into go", t, func() {
		Convey(fmt.Sprintf("Encoding %s as go code", "./testdata/testRegisters/groups.xml"), func() {
			Convey("no errors are generated and the output matches", func() {
				So(labelErr, ShouldBeNil)
				So(groupErr, ShouldBeNil)
				So(string(expected), ShouldResemble, groupBuf.String())
			})
		})
	})

}

func NewPrimer() *Primer {
	start := uint16(0xffff)
	return &Primer{
		dynamicTag: &start,
		Tags:       make(map[string][]byte),
	}
}

type Primer struct {
	dynamicTag *uint16
	Tags       map[string][]byte
}

// AddEntry adds an
func (p *Primer) AddEntry(id, shorthand []byte) []byte {
	sId := string(id)

	// if its already added do not decrement
	if _, ok := p.Tags[sId]; ok {
		return []byte{}
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
