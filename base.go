//  Copyright ©2019-2024  Mr MXF   info@mrmxf.com
//  BSD-3-Clause License  https://opensource.org/license/bsd-3-clause/

package main

import (
	"strings"
)

func findBaseOg(starter entryinfo, register map[string]entryinfo) entryinfo {

	// if the base is a record then it is the start type
	if starter.Base || starter.TypeKind == "Record" {
		return starter
	}

	mid := register[starter.UL]
	var base bool

	// loop through the register chain to find the start point
	for !base {
		if mid.Base || mid.TypeKind == "Record" {
			base = true
		}

		mid = register[mid.BaseType]
	}

	return mid

}

func findBase(starter entryinfo, register map[string]entryinfo, registerType string) (targetType, typeSymbol, wrapType, start, end string) {

	// I want to get the entry info of the base,
	// as the closest to a base Stype as possible
	// return type, wrapper and symbols

	// switch statement

	startUL := starter.BaseType
	if registerType == "struct" { // options are to stop arrays being skipped out

		startUL = starter.UL

	}

	if starter.Register == "Elements" {

		// if it has no type then it is any
		startUL = starter.Type
		if startUL == "" {
			targetType = "Tany"
			return

		}
	} else if registerType == "Elements" {
		startUL = starter.UL
	}

	n := strings.ReplaceAll(register[startUL].Symbol, "-", "_")
	targetType = string(register[startUL].Register[0]) + strings.ReplaceAll(n, " ", "_")

	// var extract string
	//	var length int
	// var wrap, start, end string

	if register[startUL].TypeKind == "Record" || register[startUL].TypeKind == "FixedArray" { /*||
		register[starter.BaseType].TypeKind == "WeakReference" || register[starter.BaseType].TypeKind == "StrongReference" {*/

		typeSymbol = targetType

	} else {

		baseSymbol := startUL

		if register[baseSymbol].Base {
			//	length = lengths[register[symbol].Symbol]
			typeSymbol = strings.ToLower(register[baseSymbol].Symbol)
		} else {

			for !register[baseSymbol].Base {

				mid := register[baseSymbol]

				typeSymbol = mid.Symbol
				if mid.BaseType != "" {
					baseSymbol = mid.BaseType

				}

				baseFound := false

				// if the target already has its own decode function, do not search ay further
				switch {
				case mid.Base || mid.TypeKind == "Record" || mid.Register == "Groups" || mid.TypeKind == "Rename":

					if mid.Base {

						typeSymbol = strings.ToLower(typeSymbol)
					} else {
						typeSymbol = string(mid.Register[0]) + typeSymbol

					}

					baseFound = true

					//	} else if mid.TypeKind == "Rename" {
					//		typeSymbol = string(mid.Register[0]) + typeSymbol
				case mid.TypeKind == "Enumeration":
					wrapType = targetType
					start, end = "(", ")"
					typeSymbol = strings.ToLower(register[mid.BaseType].Symbol)
					targetType = string(register[mid.BaseType].Register[0]) + strings.ReplaceAll(register[mid.BaseType].Symbol, " ", "_")

				case mid.TypeKind == "WeakReference":
					wrapType = targetType
					start, end = "(", ")"
					targetType = "TUInt8Array"
					typeSymbol = "TUInt8Array"

					baseFound = true
				case mid.TypeKind == "StrongReference":
					wrapType = targetType
					start, end = "(", ")"
					targetType = "TStrongReference"
					typeSymbol = "TStrongReference"

					baseFound = true
					// } //else if mid.TypeKind == "String" {
					//	typeSymbol = strings.ToLower(register[mid.BaseType].Symbol)
					//	targetType = string(register[mid.BaseType].Register[0]) + strings.ReplaceAll(register[mid.BaseType].Symbol, " ", "_")
				default: // just make a generic filler

					typeSymbol = string(mid.Register[0]) + typeSymbol

					//	length = lengths[rsymbl]
				}

				if baseFound {
					break
				}
			}
		}

	}

	return
}
