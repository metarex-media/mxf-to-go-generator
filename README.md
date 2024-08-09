# MXF Golang Gen

This repo generates a go code version of the [SMPTE RA registers][smpte-reg].
Running MXF golang gen will generate a folder of go code,
made from the registers, as the package mxf2go.
Checkout the generated pure go code,
as a standalone repo [here][m2g].

Before embarking on using this repo, please make sure you are familiar with the
technical concepts of the [Material Exchange Format (MXF)][MXFspec].

## Contents

- [Installation](#installation)
- [Example](#example)
- [Design Notes](#design-notes)
  - [Types](#types)
  - [Elements](#elements)
  - [Labels](#labels)
  - [Essence](#essence)
- [Additional code](#additional-code)
- [Adding custom fields](#adding-custom-fields)
  
## Installation

Make sure you have the latest version of Go from the [official golang
source][g1] installed. Then run the following command to compile the program.

```cmd
go build
```

This tool only runs on the command line.

## Example

To run MXF golang gen you use the following command,
which will generate the register go code.

```cmd
./MXFGen
```

Check out the `./generated` folder after running it for the freshly
written go code. At line 4 of `./generated/basetypes.go`
the time the file was written to is commented, check the
comment and make sure it matches with when you ran the program.

## Design Notes

Each register has to fit into the go design ethos,
this involved making some design choices to allow the truest
handling of the registers as go code.

The following registers and their designs are given below:

- [Types](#types)
- [Elements](#elements)
- [Labels](#labels)
- [Essence](#essence)
- [Groups](#groups)

The pure go code versions that are generated
can be found in this [repo][m2g].

### Types

The type register is contains all the types used by mxf, these
are all built from the `BasicTypes` Node in the register.

The register to go types as matched as the following:

- records - structs
- fixed - array (of fixed length)
- variable - array slice
- string - []rune
- strong and weak references - []byte
- any other type is a direct translation to the go type. e.g. float (register) to float32 (go)
because the types share the same name.

Some types wrap other register types before reaching the go type,
leading to some types wrapping round another type register value several times,
before reaching the go base.
For example TPackageStrongReference is of type TStrongReference, which is of go type []byte.

Accompanying encode functions and decode functions are also generated. To allow
the types to be easily handled, and for the other go registers to fully utilise
the types. For example the Groups register uses the decode and encode functions
of these types, for generating the header bytes.

The encode functions **all** follow the layout of

```go
Encode{TypeName}(value {TypeName})([]byte, error)
```

And the decode functions have the following layout.

```go
Decode{TypeName}(value []byte)(any, error)
```

The types are returned as an any to allow the generic handling
of MXF byte streams.

Any enumerated type also has a generated `String` and `MarshallTest` function,
to allow the values to be human readable. When parsed as JSON and
YAML outputs.

e.g. a type of AS_11_Captions_Type_Enum has the following generated code.

```go
type TAS_11_Captions_Type_Enum uint8

func (mt TAS_11_Captions_Type_Enum) MarshalText() ([]byte, error) {
    return []byte(mt.String()), nil
}

func (s TAS_11_Captions_Type_Enum)String() string {

     switch s {
    case 0:
        return "Captions_Hard_of_Hearing" 
    case 1:
        return "Captions_Translation" 
    default:
        return "invalid value"
    }
}
```

### Elements

The elements register wraps a MXF type from the type register.
The generated go file wraps the types from the types.go file
as an element, creating a new type.

Elements with no declared type are stored
as an any type e.g. `type EDOI any`.

No encode or decode functions are made for the elements.

### Labels

The labels register is an array of the labels used in MXF.
This is directly translated into an array and a map, where the
labels UL is the map key. These labels in go, contains the register
entry for the label as in the `labels.xml` register file.

The Universal Label (UL) of a label
`urn:smpte:ul:00000000.00000000.00000000.00000000` is the map key. The UL is
the default UL given in the register, where any `7f` bytes, which means a byte
that can have multiple values, is kept as 7f in their UL, rather than including
all the values that may be present.

### Essence

Essence contains the description for the essence (the data contents) within
an mxf file, the essence register contains all of the essence entries.
This is directly translated into a map, where the
essence UL is the map key. These essence in go, contains the register
entry for the essence as in the `essence.xml` register file.

The Universal Label (UL) of an essence
`urn:smpte:ul:00000000.00000000.00000000.00000000` is the map key. The UL is
the default UL given in the register, where any `7f` bytes, which means a byte
that can have multiple values, is kept as 7f in their UL, rather than including
all the values that may be present.

### Groups

The groups register contains the groups that are used in MXF headers and partitions.
These groups are made up of several records, which can be inherited from parent groups.
Which translate in go as the group being a struct, and the records it
contains are the struct fields.

These groups then have an accompanying encode method, which encode the whole
mxf group as a Key Length Value (KLV). The encode method is designed to be self contained, where
once the bytes have been generated, the group bytes can be used straight an mxf file.

```go
type Example struct {
   ExampleField any
} 

func (e E)Encode(primer *Primer) ([]byte, error){
   // contents here
}
```

The decoding for groups works slightly differently, as you have to find out
what group the byte stream is. Furthermore each field may not be in the same
order in the group as expected.
All groups are placed in the `Groups` look up table. The Universal Label (UL) of a group
`urn:smpte:ul:00000000.00000000.00000000.00000000` is the map key. The UL is
the default UL given in the register, where any `7f` bytes, which means a byte
that can have multiple values, is kept as 7f in their UL, rather than including
all the values that may be present.

This group then has another lookup table for every possible field in the group,
with the UL of the field being the key. This lookup table then contains the decode
function, for the stream of bytes to convert it into that field as a go type.

### Additional code

Some additional code is required for the base types in the register to work,
as the base types, float, int etc need their decode functions to be built to
be referenced by the other types. So these are coded in basetypes.go
as they have a different format to the other types in the register.

Additional code is added to some types for marshalling the bytes as
text, so when they are called by json and yaml parsers the formatted version is
used instead of the byte values.

The following types have the MarshalText function:

- UUID
- PackageIDType
- TAUID
- StrongReference
- WeakReference

The internal workings of MXF require a map of shorthand keys,
that relate to the UL of some fields within groups.
This map is called the "Primer" and has been included as
part of the generated code as an object that the encoders use.
This can be used like so, where a new primer is made
and then used to encode the preface group in the demo below.

```go
// make the new primer
primer := mxf2go.NewPrimer()
pre := mxf2go.GPrefaceStruct{} // fill in your own fields
// generate the preface bytes
prefaceBytes, _ := pre.Encode(primer)
```

It is important that you use the same Primer object
for every encode function, to accurately catalogue every
UL used in the MXF file.

## Adding custom fields

If you want to add custom fields to the registers
then you can modify the values in the `./register` folder.

This can be done to add private fields to the MXF register,
or for any other purpose.

Either by modifying values currently in the register or
by adding whole new entries to the array.
Then when you rerun the program the go code will reflect these updates.

[g1]:   https://go.dev/doc/install                "Golang Installation"
[smpte-reg]: https://registry.smpte-ra.org/apps/pages/published/ "SMPTE registers"
[m2g]: https://github.com/metarex-media/mxf-to-go "MXFToGo code base"
[mxfspec]: https://pub.smpte.org/doc/377/
