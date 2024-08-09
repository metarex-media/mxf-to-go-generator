//  Copyright ©2019-2024  Mr MXF   info@mrmxf.com
//  BSD-3-Clause License  https://opensource.org/license/bsd-3-clause/

package main

import (
	"fmt"
	"os"
)

func main() {

	// read every register value and extract the bases from them
	xmlReg, bases, err := baseSetUp("./registers/Types.xml", []string{"./registers/Elements.xml", "./registers/Groups.xml", "./registers/Labels.xml", "./registers/Types.xml",
		"./dev/elements-dev.xml", "./dev/groups-dev.xml",
		"./dev/elements-mrx-dev.xml", "./dev/groups-mrx-dev.xml"})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	f, _ := os.Create("./generated/basetypes.go")
	// base types is an extra file to decode the base types, e.g. uint16
	lengths, err := encodeBaseTypes(bases, f)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// encode the registers as go code
	// encode the base elements
	ftyp, _ := os.Create("./generated/types.go")
	err = encodeTypes(xmlReg, lengths, "./registers/Types.xml", ftyp)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fgr, _ := os.Create("./generated/groups.go")
	err = encodeGroups(xmlReg, lengths, []string{"./registers/Groups.xml", "./dev/groups-mrx-dev.xml", "./dev/groups-dev.xml"}, fgr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	flabel, _ := os.Create("./generated/labels.go")
	err = encodeLabels("./registers/Labels.xml", flabel)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fess, _ := os.Create("./generated/essence.go")
	err = encodeEssence([]string{"./registers/Essence.xml", "./dev/essence-dev.xml"}, fess)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fel, _ := os.Create("./generated/elements.go")
	err = encodeElements(xmlReg, lengths, "./registers/Elements.xml", fel)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	feldev, _ := os.Create("./generated/elements-dev.go")
	// encode the deve elements
	err = encodeElements(xmlReg, lengths, "./dev/elements-dev.xml", feldev)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	felmdev, _ := os.Create("./generated/elements-mrx-dev.go")
	err = encodeElements(xmlReg, lengths, "./dev/elements-mrx-dev.xml", felmdev)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
