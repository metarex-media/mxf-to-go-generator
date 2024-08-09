// Copyright ©2019-2024  Mr MXF   info@mrmxf.com
// BSD-3-Clause License  https://opensource.org/license/bsd-3-clause/
package mxf2go

import (
	"encoding/binary"
	"unicode/utf16"
	"unicode/utf8"
)

type TFilmType uint8

func (mt TFilmType) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TFilmType) String() string {

	switch s {
	case 0:
		return "FtNull"
	case 1:
		return "Ft35MM"
	case 2:
		return "Ft16MM"
	case 3:
		return "Ft8MM"
	case 4:
		return "Ft65MM"
	default:
		return "invalid value"
	}
}
func DecodeTFilmType(value []byte) (any, error) {
	field, _ := DecodeTUInt8(value[0:1])
	result := TFilmType(field.(uint8))
	return result, nil
}
func EncodeTFilmType(value TFilmType) ([]byte, error) {
	result, _ := EncodeTUInt8(uint8(value))
	return result, nil
}

type TPositionType int64

func DecodeTPositionType(value []byte) (any, error) {
	field, _ := DecodeTInt64(value[0:8])
	result := TPositionType(field.(int64))
	return result, nil
}
func EncodeTPositionType(value TPositionType) ([]byte, error) {
	result, _ := EncodeTInt64(int64(value))
	return result, nil
}

type TRGBAComponentKind uint8

func (mt TRGBAComponentKind) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TRGBAComponentKind) String() string {

	switch s {
	case 48:
		return "CompNone"
	case 65:
		return "CompAlpha"
	case 66:
		return "CompBlue"
	case 70:
		return "CompFill"
	case 71:
		return "CompGreen"
	case 80:
		return "CompPalette"
	case 82:
		return "CompRed"
	case 0:
		return "CompNull"
	case 114:
		return "CompRedLSBs"
	case 103:
		return "CompGreenLSBs"
	case 98:
		return "CompBlueLSBs"
	case 97:
		return "CompAlphaLSBs"
	case 85:
		return "CompColorDifferenceU"
	case 86:
		return "CompColorDifferenceV"
	case 87:
		return "CompComposite"
	case 88:
		return "CompNonCoSitedLuma"
	case 89:
		return "CompLuma"
	case 90:
		return "CompDepth"
	case 117:
		return "CompColorDifferenceULSBs"
	case 118:
		return "CompColorDifferenceVLSBs"
	case 119:
		return "CompCompositeLSBs"
	case 120:
		return "CompNonCoSitedLumaLSBs"
	case 121:
		return "CompLumaLSBs"
	case 122:
		return "CompDepthLSBs"
	case 216:
		return "CompColorX"
	case 217:
		return "CompColorY"
	case 218:
		return "CompColorZ"
	default:
		return "invalid value"
	}
}
func DecodeTRGBAComponentKind(value []byte) (any, error) {
	field, _ := DecodeTUInt8(value[0:1])
	result := TRGBAComponentKind(field.(uint8))
	return result, nil
}
func EncodeTRGBAComponentKind(value TRGBAComponentKind) ([]byte, error) {
	result, _ := EncodeTUInt8(uint8(value))
	return result, nil
}

type TLengthType int64

func DecodeTLengthType(value []byte) (any, error) {
	field, _ := DecodeTInt64(value[0:8])
	result := TLengthType(field.(int64))
	return result, nil
}
func EncodeTLengthType(value TLengthType) ([]byte, error) {
	result, _ := EncodeTInt64(int64(value))
	return result, nil
}

type TJPEGTableIDType int32

func DecodeTJPEGTableIDType(value []byte) (any, error) {
	field, _ := DecodeTInt32(value[0:4])
	result := TJPEGTableIDType(field.(int32))
	return result, nil
}
func EncodeTJPEGTableIDType(value TJPEGTableIDType) ([]byte, error) {
	result, _ := EncodeTInt32(int32(value))
	return result, nil
}

type TLocalTagType uint16

func DecodeTLocalTagType(value []byte) (any, error) {
	field, _ := DecodeTUInt16(value[0:2])
	result := TLocalTagType(field.(uint16))
	return result, nil
}
func EncodeTLocalTagType(value TLocalTagType) ([]byte, error) {
	result, _ := EncodeTUInt16(uint16(value))
	return result, nil
}

type TReferenceType uint8

func (mt TReferenceType) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TReferenceType) String() string {

	switch s {
	case 0:
		return "RefLimitMinimum"
	case 1:
		return "RefLimitMaximum"
	case 2:
		return "RefMinimum"
	case 3:
		return "RefMaximum"
	case 4:
		return "RefEnumvalue"
	default:
		return "invalid value"
	}
}
func DecodeTReferenceType(value []byte) (any, error) {
	field, _ := DecodeTUInt8(value[0:1])
	result := TReferenceType(field.(uint8))
	return result, nil
}
func EncodeTReferenceType(value TReferenceType) ([]byte, error) {
	result, _ := EncodeTUInt8(uint8(value))
	return result, nil
}

type TS309M uint64

func DecodeTS309M(value []byte) (any, error) {
	field, _ := DecodeTUInt64(value[0:8])
	result := TS309M(field.(uint64))
	return result, nil
}
func EncodeTS309M(value TS309M) ([]byte, error) {
	result, _ := EncodeTUInt64(uint64(value))
	return result, nil
}

type TAlphaTransparencyType uint8

func (mt TAlphaTransparencyType) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TAlphaTransparencyType) String() string {

	switch s {
	case 0:
		return "MinValueTransparent"
	case 1:
		return "MaxValueTransparent"
	default:
		return "invalid value"
	}
}
func DecodeTAlphaTransparencyType(value []byte) (any, error) {
	field, _ := DecodeTUInt8(value[0:1])
	result := TAlphaTransparencyType(field.(uint8))
	return result, nil
}
func EncodeTAlphaTransparencyType(value TAlphaTransparencyType) ([]byte, error) {
	result, _ := EncodeTUInt8(uint8(value))
	return result, nil
}

type TS12M uint64

func DecodeTS12M(value []byte) (any, error) {
	field, _ := DecodeTUInt64(value[0:8])
	result := TS12M(field.(uint64))
	return result, nil
}
func EncodeTS12M(value TS12M) ([]byte, error) {
	result, _ := EncodeTUInt64(uint64(value))
	return result, nil
}

type TFieldNumber uint8

func (mt TFieldNumber) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TFieldNumber) String() string {

	switch s {
	case 0:
		return "UnspecifiedField"
	case 1:
		return "FieldOne"
	case 2:
		return "FieldTwo"
	default:
		return "invalid value"
	}
}
func DecodeTFieldNumber(value []byte) (any, error) {
	field, _ := DecodeTUInt8(value[0:1])
	result := TFieldNumber(field.(uint8))
	return result, nil
}
func EncodeTFieldNumber(value TFieldNumber) ([]byte, error) {
	result, _ := EncodeTUInt8(uint8(value))
	return result, nil
}

type TMicroTime197001010000 uint64

func DecodeTMicroTime197001010000(value []byte) (any, error) {
	field, _ := DecodeTUInt64(value[0:8])
	result := TMicroTime197001010000(field.(uint64))
	return result, nil
}
func EncodeTMicroTime197001010000(value TMicroTime197001010000) ([]byte, error) {
	result, _ := EncodeTUInt64(uint64(value))
	return result, nil
}

type TFix32Dec3 int32

func DecodeTFix32Dec3(value []byte) (any, error) {
	field, _ := DecodeTInt32(value[0:4])
	result := TFix32Dec3(field.(int32))
	return result, nil
}
func EncodeTFix32Dec3(value TFix32Dec3) ([]byte, error) {
	result, _ := EncodeTInt32(int32(value))
	return result, nil
}

type TElectroSpatialFormulation uint8

func (mt TElectroSpatialFormulation) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TElectroSpatialFormulation) String() string {

	switch s {
	case 0:
		return "ElectroSpatialFormulation_Default"
	case 1:
		return "ElectroSpatialFormulation_TwoChannelMode"
	case 2:
		return "ElectroSpatialFormulation_SingleChannelMode"
	case 3:
		return "ElectroSpatialFormulation_PrimarySecondaryMode"
	case 4:
		return "ElectroSpatialFormulation_StereophonicMode"
	case 7:
		return "ElectroSpatialFormulation_SingleChannelDoubleSamplingFrequencyMode"
	case 8:
		return "ElectroSpatialFormulation_StereoLeftChannelDoubleSamplingFrequencyMode"
	case 9:
		return "ElectroSpatialFormulation_StereoRightChannelDoubleSamplingFrequencyMode"
	case 15:
		return "ElectroSpatialFormulation_MultiChannelMode"
	default:
		return "invalid value"
	}
}
func DecodeTElectroSpatialFormulation(value []byte) (any, error) {
	field, _ := DecodeTUInt8(value[0:1])
	result := TElectroSpatialFormulation(field.(uint8))
	return result, nil
}
func EncodeTElectroSpatialFormulation(value TElectroSpatialFormulation) ([]byte, error) {
	result, _ := EncodeTUInt8(uint8(value))
	return result, nil
}

type TCompactAdIDIdentifierType uint32

func DecodeTCompactAdIDIdentifierType(value []byte) (any, error) {
	field, _ := DecodeTUInt32(value[0:4])
	result := TCompactAdIDIdentifierType(field.(uint32))
	return result, nil
}
func EncodeTCompactAdIDIdentifierType(value TCompactAdIDIdentifierType) ([]byte, error) {
	result, _ := EncodeTUInt32(uint32(value))
	return result, nil
}

type TPhaseFrameType int32

func DecodeTPhaseFrameType(value []byte) (any, error) {
	field, _ := DecodeTInt32(value[0:4])
	result := TPhaseFrameType(field.(int32))
	return result, nil
}
func EncodeTPhaseFrameType(value TPhaseFrameType) ([]byte, error) {
	result, _ := EncodeTInt32(int32(value))
	return result, nil
}

type TEmphasisType uint8

func (mt TEmphasisType) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TEmphasisType) String() string {

	switch s {
	case 0:
		return "Emphasis_Unknown"
	case 1:
		return "Emphasis_Reserved0"
	case 2:
		return "Emphasis_Reserved1"
	case 3:
		return "Emphasis_Reserved2"
	case 4:
		return "Emphasis_None"
	case 5:
		return "Emphasis_Reserved3"
	case 6:
		return "Emphasis_15and50"
	case 7:
		return "Emphasis_ITU"
	default:
		return "invalid value"
	}
}
func DecodeTEmphasisType(value []byte) (any, error) {
	field, _ := DecodeTUInt8(value[0:1])
	result := TEmphasisType(field.(uint8))
	return result, nil
}
func EncodeTEmphasisType(value TEmphasisType) ([]byte, error) {
	result, _ := EncodeTUInt8(uint8(value))
	return result, nil
}

type TtechnicalAttributeUInt8StrongReference TStrongReference
type TtechnicalAttributeUInt16StrongReference TStrongReference
type TtechnicalAttributeUInt32StrongReference TStrongReference
type TUInt32Array []uint32

func EncodeTUInt32Array(value TUInt32Array) ([]byte, error) {
	result := []byte{}
	for _, val := range value {
		field, _ := EncodeTUInt32(uint32(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTUInt32Array(value []byte) (any, error) {
	arrayCount := len(value) / 4
	var result TUInt32Array = make([]uint32, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTUInt32(value[i*4 : i*4+4])
		result[i] = field.(uint32) // else
	}
	return result, nil
}

type TtechnicalAttributeUInt64StrongReference TStrongReference
type TtechnicalAttributeFloatStrongReference TStrongReference
type TtechnicalAttributeRationalStrongReference TStrongReference
type TAuxBitsModeType uint8

func (mt TAuxBitsModeType) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TAuxBitsModeType) String() string {

	switch s {
	case 0:
		return "AuxBitsMode_NotDefined"
	case 1:
		return "AuxBitsMode_MainAudioSampleData"
	case 2:
		return "AuxBitsMode_SingleCoordinationSignal"
	case 3:
		return "AuxBitsMode_UserDefined"
	case 4:
		return "AuxBitsMode_Reserved0"
	case 5:
		return "AuxBitsMode_Reserved1"
	case 6:
		return "AuxBitsMode_Reserved2"
	case 7:
		return "AuxBitsMode_Reserved3"
	default:
		return "invalid value"
	}
}
func DecodeTAuxBitsModeType(value []byte) (any, error) {
	field, _ := DecodeTUInt8(value[0:1])
	result := TAuxBitsModeType(field.(uint8))
	return result, nil
}
func EncodeTAuxBitsModeType(value TAuxBitsModeType) ([]byte, error) {
	result, _ := EncodeTUInt8(uint8(value))
	return result, nil
}

type TtechnicalAttributeAnyURIStrongReference TStrongReference
type TtechnicalAttributeBooleanStrongReference TStrongReference
type TaudioFormatExtendedStrongReference TStrongReference
type TaspectRatioStrongReference TStrongReference
type TChannelStatusModeArray []TChannelStatusModeType

func EncodeTChannelStatusModeArray(value TChannelStatusModeArray) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	binary.BigEndian.PutUint32(result[4:8], 1)
	for _, val := range value {
		field, _ := EncodeTUInt8(uint8(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTChannelStatusModeArray(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TChannelStatusModeArray = make([]TChannelStatusModeType, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTUInt8(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TChannelStatusModeType(field.(uint8)) // else
	}
	return result, nil
}

type TheightStrongReference TStrongReference
type TUUID [16]uint8

func EncodeTUUID(value TUUID) ([]byte, error) {
	result := []byte{}
	for _, val := range value {
		field, _ := EncodeTUInt8(uint8(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTUUID(value []byte) (any, error) {
	var result TUUID = [16]uint8{}
	for i := 0; i < 16; i++ {
		field, _ := DecodeTUInt8(value[i*1 : i*1+1])
		result[i] = field.(uint8) // else
	}
	return result, nil
}

type TChannelStatusModeType uint8

func (mt TChannelStatusModeType) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TChannelStatusModeType) String() string {

	switch s {
	case 0:
		return "ChannelStatusMode_None"
	case 1:
		return "ChannelStatusMode_Minimum"
	case 2:
		return "ChannelStatusMode_Standard"
	case 3:
		return "ChannelStatusMode_Fixed"
	case 4:
		return "ChannelStatusMode_Stream"
	case 5:
		return "ChannelStatusMode_Essence"
	default:
		return "invalid value"
	}
}
func DecodeTChannelStatusModeType(value []byte) (any, error) {
	field, _ := DecodeTUInt8(value[0:1])
	result := TChannelStatusModeType(field.(uint8))
	return result, nil
}
func EncodeTChannelStatusModeType(value TChannelStatusModeType) ([]byte, error) {
	result, _ := EncodeTUInt8(uint8(value))
	return result, nil
}

type TUserDataModeArray []TUserDataModeType

func EncodeTUserDataModeArray(value TUserDataModeArray) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	binary.BigEndian.PutUint32(result[4:8], 1)
	for _, val := range value {
		field, _ := EncodeTUInt8(uint8(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTUserDataModeArray(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TUserDataModeArray = make([]TUserDataModeType, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTUInt8(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TUserDataModeType(field.(uint8)) // else
	}
	return result, nil
}

type TUInt8Array3 [3]uint8

func EncodeTUInt8Array3(value TUInt8Array3) ([]byte, error) {
	result := []byte{}
	for _, val := range value {
		field, _ := EncodeTUInt8(uint8(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTUInt8Array3(value []byte) (any, error) {
	var result TUInt8Array3 = [3]uint8{}
	for i := 0; i < 3; i++ {
		field, _ := DecodeTUInt8(value[i*1 : i*1+1])
		result[i] = field.(uint8) // else
	}
	return result, nil
}

type TUInt8Array6 [6]uint8

func EncodeTUInt8Array6(value TUInt8Array6) ([]byte, error) {
	result := []byte{}
	for _, val := range value {
		field, _ := EncodeTUInt8(uint8(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTUInt8Array6(value []byte) (any, error) {
	var result TUInt8Array6 = [6]uint8{}
	for i := 0; i < 6; i++ {
		field, _ := DecodeTUInt8(value[i*1 : i*1+1])
		result[i] = field.(uint8) // else
	}
	return result, nil
}

type TwidthStrongReference TStrongReference
type TtrackStrongReference TStrongReference
type TS352M [4]uint8

func EncodeTS352M(value TS352M) ([]byte, error) {
	result := []byte{}
	for _, val := range value {
		field, _ := EncodeTUInt8(uint8(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTS352M(value []byte) (any, error) {
	var result TS352M = [4]uint8{}
	for i := 0; i < 4; i++ {
		field, _ := DecodeTUInt8(value[i*1 : i*1+1])
		result[i] = field.(uint8) // else
	}
	return result, nil
}

type TcaptioningStrongReference TStrongReference
type TUserDataModeType uint8

func (mt TUserDataModeType) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TUserDataModeType) String() string {

	switch s {
	case 0:
		return "UserDataMode_NotDefined"
	case 1:
		return "UserDataMode_192BitBlockStructure"
	case 2:
		return "UserDataMode_AES18"
	case 3:
		return "UserDataMode_UserDefined"
	case 4:
		return "UserDataMode_IEC"
	case 5:
		return "UserDataMode_Metadata"
	case 6:
		return "UserDataMode_Reserved0"
	case 7:
		return "UserDataMode_Reserved1"
	case 8:
		return "UserDataMode_Reserved2"
	case 9:
		return "UserDataMode_Reserved3"
	case 10:
		return "UserDataMode_Reserved4"
	case 11:
		return "UserDataMode_Reserved5"
	case 12:
		return "UserDataMode_Reserved6"
	case 13:
		return "UserDataMode_Reserved7"
	case 14:
		return "UserDataMode_Reserved8"
	case 15:
		return "UserDataMode_Reserved9"
	default:
		return "invalid value"
	}
}
func DecodeTUserDataModeType(value []byte) (any, error) {
	field, _ := DecodeTUInt8(value[0:1])
	result := TUserDataModeType(field.(uint8))
	return result, nil
}
func EncodeTUserDataModeType(value TUserDataModeType) ([]byte, error) {
	result, _ := EncodeTUInt8(uint8(value))
	return result, nil
}

type TsubtitlingStrongReference TStrongReference
type TS330M_Spatial [12]uint8

func EncodeTS330M_Spatial(value TS330M_Spatial) ([]byte, error) {
	result := []byte{}
	for _, val := range value {
		field, _ := EncodeTUInt8(uint8(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTS330M_Spatial(value []byte) (any, error) {
	var result TS330M_Spatial = [12]uint8{}
	for i := 0; i < 12; i++ {
		field, _ := DecodeTUInt8(value[i*1 : i*1+1])
		result[i] = field.(uint8) // else
	}
	return result, nil
}

type TancillaryDataStrongReference TStrongReference
type TBoolean uint8

func (mt TBoolean) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TBoolean) String() string {

	switch s {
	case 0:
		return "False"
	case 1:
		return "True"
	default:
		return "invalid value"
	}
}
func DecodeTBoolean(value []byte) (any, error) {
	field, _ := DecodeTUInt8(value[0:1])
	result := TBoolean(field.(uint8))
	return result, nil
}
func EncodeTBoolean(value TBoolean) ([]byte, error) {
	result, _ := EncodeTUInt8(uint8(value))
	return result, nil
}

type TRational struct {
	Numerator   int32
	Denominator int32
}

func DecodeTRational(value []byte) (any, error) {
	result := TRational{}
	var field any
	field, _ = DecodeTInt32(value[0:4])
	result.Numerator = field.(int32) // else
	field, _ = DecodeTInt32(value[4:8])
	result.Denominator = field.(int32) // else
	return result, nil
}
func EncodeTRational(value TRational) ([]byte, error) {
	result := []byte{}
	field0, _ := EncodeTInt32(int32(value.Numerator)) // else
	result = append(result, field0...)
	field1, _ := EncodeTInt32(int32(value.Denominator)) // else
	result = append(result, field1...)
	return result, nil
}

type TaudioProgrammeStrongReference TStrongReference
type TExtUMID [32]uint8

func EncodeTExtUMID(value TExtUMID) ([]byte, error) {
	result := []byte{}
	for _, val := range value {
		field, _ := EncodeTUInt8(uint8(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTExtUMID(value []byte) (any, error) {
	var result TExtUMID = [32]uint8{}
	for i := 0; i < 32; i++ {
		field, _ := DecodeTUInt8(value[i*1 : i*1+1])
		result[i] = field.(uint8) // else
	}
	return result, nil
}

type TGeographicCoordinate struct {
	Latitude  TFix32Dec3
	Longitude TFix32Dec3
}

func DecodeTGeographicCoordinate(value []byte) (any, error) {
	result := TGeographicCoordinate{}
	var field any
	field, _ = DecodeTFix32Dec3(value[0:4])
	result.Latitude = field.(TFix32Dec3) // else
	field, _ = DecodeTFix32Dec3(value[4:8])
	result.Longitude = field.(TFix32Dec3) // else
	return result, nil
}
func EncodeTGeographicCoordinate(value TGeographicCoordinate) ([]byte, error) {
	result := []byte{}
	field0, _ := EncodeTFix32Dec3(TFix32Dec3(value.Latitude)) // else
	result = append(result, field0...)
	field1, _ := EncodeTFix32Dec3(TFix32Dec3(value.Longitude)) // else
	result = append(result, field1...)
	return result, nil
}

type TaudioContentStrongReference TStrongReference
type TaudioObjectStrongReference TStrongReference
type TUInt8Array16 [16]uint8

func EncodeTUInt8Array16(value TUInt8Array16) ([]byte, error) {
	result := []byte{}
	for _, val := range value {
		field, _ := EncodeTUInt8(uint8(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTUInt8Array16(value []byte) (any, error) {
	var result TUInt8Array16 = [16]uint8{}
	for i := 0; i < 16; i++ {
		field, _ := DecodeTUInt8(value[i*1 : i*1+1])
		result[i] = field.(uint8) // else
	}
	return result, nil
}

type TSignalStandardType uint8

func (mt TSignalStandardType) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TSignalStandardType) String() string {

	switch s {
	case 0:
		return "SignalStandard_None"
	case 1:
		return "SignalStandard_ITU601"
	case 2:
		return "SignalStandard_ITU1358"
	case 3:
		return "SignalStandard_SMPTE347M"
	case 4:
		return "SignalStandard_SMPTE274M"
	case 5:
		return "SignalStandard_SMPTE296M"
	case 6:
		return "SignalStandard_SMPTE349M"
	case 7:
		return "SignalStandard_SMPTE428_1"
	default:
		return "invalid value"
	}
}
func DecodeTSignalStandardType(value []byte) (any, error) {
	field, _ := DecodeTUInt8(value[0:1])
	result := TSignalStandardType(field.(uint8))
	return result, nil
}
func EncodeTSignalStandardType(value TSignalStandardType) ([]byte, error) {
	result, _ := EncodeTUInt8(uint8(value))
	return result, nil
}

type TUTF16String []rune

func DecodeTUTF16String(value []byte) (any, error) {
	parts := make([]uint16, len(value)/2)
	for i := 0; i < len(value); i += 2 {
		parts[i/2] = binary.BigEndian.Uint16(value[i : i+2])
	}
	runes := utf16.Decode(parts)
	return string(runes), nil
}

func EncodeTUTF16String(value TUTF16String) ([]byte, error) {
	parts := utf16.Encode([]rune(string(value)))
	result, _ := EncodeTUInt16Array([]uint16(parts))
	return result, nil
}

type TaudioPackFormatStrongReference TStrongReference
type TaudioChannelFormatStrongReference TStrongReference
type TUInt16Array []uint16

func EncodeTUInt16Array(value TUInt16Array) ([]byte, error) {
	result := []byte{}
	for _, val := range value {
		field, _ := EncodeTUInt16(uint16(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTUInt16Array(value []byte) (any, error) {
	arrayCount := len(value) / 2
	var result TUInt16Array = make([]uint16, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTUInt16(value[i*2 : i*2+2])
		result[i] = field.(uint16) // else
	}
	return result, nil
}

type TaudioBlockFormatStrongReference TStrongReference
type TISO7 []rune

func DecodeTISO7(value []byte) (any, error) {
	runes := make([]rune, len(value))
	for i, char := range value {
		runes[i] = rune(char)
	}
	return string(runes), nil
}

func EncodeTISO7(value TISO7) ([]byte, error) {
	result := []byte(string(value))
	return result, nil
}

type TaudioStreamFormatStrongReference TStrongReference
type TJ2KComponentSizing struct {
	Ssiz  uint8
	XRSiz uint8
	YRSiz uint8
}

func DecodeTJ2KComponentSizing(value []byte) (any, error) {
	result := TJ2KComponentSizing{}
	var field any
	field, _ = DecodeTUInt8(value[0:1])
	result.Ssiz = field.(uint8) // else
	field, _ = DecodeTUInt8(value[1:2])
	result.XRSiz = field.(uint8) // else
	field, _ = DecodeTUInt8(value[2:3])
	result.YRSiz = field.(uint8) // else
	return result, nil
}
func EncodeTJ2KComponentSizing(value TJ2KComponentSizing) ([]byte, error) {
	result := []byte{}
	field0, _ := EncodeTUInt8(uint8(value.Ssiz)) // else
	result = append(result, field0...)
	field1, _ := EncodeTUInt8(uint8(value.XRSiz)) // else
	result = append(result, field1...)
	field2, _ := EncodeTUInt8(uint8(value.YRSiz)) // else
	result = append(result, field2...)
	return result, nil
}

type TScanningDirectionType uint8

func (mt TScanningDirectionType) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TScanningDirectionType) String() string {

	switch s {
	case 0:
		return "ScanningDirection_LeftToRightTopToBottom"
	case 1:
		return "ScanningDirection_RightToLeftTopToBottom"
	case 2:
		return "ScanningDirection_LeftToRightBottomToTop"
	case 3:
		return "ScanningDirection_RightToLeftBottomToTop"
	case 4:
		return "ScanningDirection_TopToBottomLeftToRight"
	case 5:
		return "ScanningDirection_TopToBottomRightToLeft"
	case 6:
		return "ScanningDirection_BottomToTopLeftToRight"
	case 7:
		return "ScanningDirection_BottomToTopRightToLeft"
	default:
		return "invalid value"
	}
}
func DecodeTScanningDirectionType(value []byte) (any, error) {
	field, _ := DecodeTUInt8(value[0:1])
	result := TScanningDirectionType(field.(uint8))
	return result, nil
}
func EncodeTScanningDirectionType(value TScanningDirectionType) ([]byte, error) {
	result, _ := EncodeTUInt8(uint8(value))
	return result, nil
}

type TVC2WaveletArray []TVC2WaveletType

func EncodeTVC2WaveletArray(value TVC2WaveletArray) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	binary.BigEndian.PutUint32(result[4:8], 1)
	for _, val := range value {
		field, _ := EncodeTUInt8(uint8(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTVC2WaveletArray(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TVC2WaveletArray = make([]TVC2WaveletType, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTUInt8(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TVC2WaveletType(field.(uint8)) // else
	}
	return result, nil
}

type TaudioTrackFormatStrongReference TStrongReference
type TaudioTrackUIDStrongReference TStrongReference
type TUUIDArray []TUUID

func EncodeTUUIDArray(value TUUIDArray) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	binary.BigEndian.PutUint32(result[4:8], 16)
	for _, val := range value {
		field, _ := EncodeTUUID(TUUID(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTUUIDArray(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TUUIDArray = make([]TUUID, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTUUID(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = field.(TUUID) // else
	}
	return result, nil
}

type TUTF8String []rune

func DecodeTUTF8String(value []byte) (any, error) {
	runes := make([]rune, len(value))
	i := 0
	for len(value) > 0 {
		r, size := utf8.DecodeRune(value)
		runes[i] = r
		value = value[size:]
		i++
	}
	return string(runes), nil
}

func EncodeTUTF8String(value TUTF8String) ([]byte, error) {
	result := []byte{}
	for _, char := range value {
		runeBytes := make([]byte, 4)
		write := utf8.EncodeRune(runeBytes, char)
		result = append(result, runeBytes[:write]...)
	}
	return result, nil
}

type TProductVersionType struct {
	Major      uint16
	Minor      uint16
	Tertiary   uint16
	PatchLevel uint16
	BuildType  TProductReleaseType
}

func DecodeTProductVersionType(value []byte) (any, error) {
	result := TProductVersionType{}
	var field any
	field, _ = DecodeTUInt16(value[0:2])
	result.Major = field.(uint16) // else
	field, _ = DecodeTUInt16(value[2:4])
	result.Minor = field.(uint16) // else
	field, _ = DecodeTUInt16(value[4:6])
	result.Tertiary = field.(uint16) // else
	field, _ = DecodeTUInt16(value[6:8])
	result.PatchLevel = field.(uint16) // else
	field, _ = DecodeTUInt8(value[8:9])
	result.BuildType = TProductReleaseType(field.(uint8)) // else
	return result, nil
}
func EncodeTProductVersionType(value TProductVersionType) ([]byte, error) {
	result := []byte{}
	field0, _ := EncodeTUInt16(uint16(value.Major)) // else
	result = append(result, field0...)
	field1, _ := EncodeTUInt16(uint16(value.Minor)) // else
	result = append(result, field1...)
	field2, _ := EncodeTUInt16(uint16(value.Tertiary)) // else
	result = append(result, field2...)
	field3, _ := EncodeTUInt16(uint16(value.PatchLevel)) // else
	result = append(result, field3...)
	field4, _ := EncodeTUInt8(uint8(value.BuildType)) // else
	result = append(result, field4...)
	return result, nil
}

type TVersionType struct {
	VersionMajor int8
	VersionMinor int8
}

func DecodeTVersionType(value []byte) (any, error) {
	result := TVersionType{}
	var field any
	field, _ = DecodeTInt8(value[0:1])
	result.VersionMajor = field.(int8) // else
	field, _ = DecodeTInt8(value[1:2])
	result.VersionMinor = field.(int8) // else
	return result, nil
}
func EncodeTVersionType(value TVersionType) ([]byte, error) {
	result := []byte{}
	field0, _ := EncodeTInt8(int8(value.VersionMajor)) // else
	result = append(result, field0...)
	field1, _ := EncodeTInt8(int8(value.VersionMinor)) // else
	result = append(result, field1...)
	return result, nil
}

type TToleranceModeType uint8

func (mt TToleranceModeType) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TToleranceModeType) String() string {

	switch s {
	case 0:
		return "ToleranceMode_Estimated"
	case 1:
		return "ToleranceMode_Assumed"
	case 2:
		return "ToleranceMode_Precise"
	case 3:
		return "ToleranceMode_Window"
	case 4:
		return "ToleranceMode_Interpolated"
	default:
		return "invalid value"
	}
}
func DecodeTToleranceModeType(value []byte) (any, error) {
	field, _ := DecodeTUInt8(value[0:1])
	result := TToleranceModeType(field.(uint8))
	return result, nil
}
func EncodeTToleranceModeType(value TToleranceModeType) ([]byte, error) {
	result, _ := EncodeTUInt8(uint8(value))
	return result, nil
}

type TContentScanningType uint8

func (mt TContentScanningType) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TContentScanningType) String() string {

	switch s {
	case 0:
		return "ContentScanning_NotKnown"
	case 1:
		return "ContentScanning_Progressive"
	case 2:
		return "ContentScanning_Interlace"
	case 3:
		return "ContentScanning_Mixed"
	default:
		return "invalid value"
	}
}
func DecodeTContentScanningType(value []byte) (any, error) {
	field, _ := DecodeTUInt8(value[0:1])
	result := TContentScanningType(field.(uint8))
	return result, nil
}
func EncodeTContentScanningType(value TContentScanningType) ([]byte, error) {
	result, _ := EncodeTUInt8(uint8(value))
	return result, nil
}

type TIDRefStrongReference TStrongReference
type TTitleAlignmentType uint8

func (mt TTitleAlignmentType) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TTitleAlignmentType) String() string {

	switch s {
	case 0:
		return "TitleAlignment_Left"
	case 1:
		return "TitleAlignment_Center"
	case 2:
		return "TitleAlignment_Right"
	default:
		return "invalid value"
	}
}
func DecodeTTitleAlignmentType(value []byte) (any, error) {
	field, _ := DecodeTUInt8(value[0:1])
	result := TTitleAlignmentType(field.(uint8))
	return result, nil
}
func EncodeTTitleAlignmentType(value TTitleAlignmentType) ([]byte, error) {
	result, _ := EncodeTUInt8(uint8(value))
	return result, nil
}

type TaudioBlockMatrixCoefficientStrongReference TStrongReference
type TtimeStrongReference TStrongReference
type TAVCContentScanningType uint8

func (mt TAVCContentScanningType) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TAVCContentScanningType) String() string {

	switch s {
	case 0:
		return "AVCContentScanning_NotKnown"
	case 1:
		return "AVCContentScanning_ProgressiveFramePicture"
	case 2:
		return "AVCContentScanning_InterlaceFieldPicture"
	case 3:
		return "AVCContentScanning_InterlaceFramePicture"
	case 4:
		return "AVCContentScanning_Interlace_FrameFieldPicture"
	default:
		return "invalid value"
	}
}
func DecodeTAVCContentScanningType(value []byte) (any, error) {
	field, _ := DecodeTUInt8(value[0:1])
	result := TAVCContentScanningType(field.(uint8))
	return result, nil
}
func EncodeTAVCContentScanningType(value TAVCContentScanningType) ([]byte, error) {
	result, _ := EncodeTUInt8(uint8(value))
	return result, nil
}

type TvideoNoiseFilterStrongReference TStrongReference
type TFileDescriptorStrongReference TStrongReference
type TmetadataFormatStrongReference TStrongReference
type TtimecodeFormatStrongReference TStrongReference
type TaudienceStrongReference TStrongReference
type TeventStrongReference TStrongReference
type TMPEG4VisualCodedContentType uint8

func (mt TMPEG4VisualCodedContentType) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TMPEG4VisualCodedContentType) String() string {

	switch s {
	case 0:
		return "MPEG4VisualCodedContent_Unknown"
	case 1:
		return "MPEG4VisualCodedContent_Progressive"
	case 2:
		return "MPEG4VisualCodedContent_Interlaced"
	case 3:
		return "MPEG4VisualCodedContent_Mixed"
	default:
		return "invalid value"
	}
}
func DecodeTMPEG4VisualCodedContentType(value []byte) (any, error) {
	field, _ := DecodeTUInt8(value[0:1])
	result := TMPEG4VisualCodedContentType(field.(uint8))
	return result, nil
}
func EncodeTMPEG4VisualCodedContentType(value TMPEG4VisualCodedContentType) ([]byte, error) {
	result, _ := EncodeTUInt8(uint8(value))
	return result, nil
}

type TawardStrongReference TStrongReference
type TaffiliationStrongReference TStrongReference
type TfilterStrongReference TStrongReference
type TRationalArray []TRational

func EncodeTRationalArray(value TRationalArray) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	binary.BigEndian.PutUint32(result[4:8], 8)
	for _, val := range value {
		field, _ := EncodeTRational(TRational(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTRationalArray(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TRationalArray = make([]TRational, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTRational(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = field.(TRational) // else
	}
	return result, nil
}

type TreferenceScreenStrongReference TStrongReference
type TVC2WaveletType uint8

func (mt TVC2WaveletType) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TVC2WaveletType) String() string {

	switch s {
	case 0:
		return "VC2Wavelet_DeslauriersDubuc_9_7"
	case 1:
		return "VC2Wavelet_LeGall_5_3"
	case 2:
		return "VC2Wavelet_DeslauriersDubuc_13_7"
	case 3:
		return "VC2Wavelet_HaarNoShift"
	case 4:
		return "VC2Wavelet_HaarSingleShiftPerLevel"
	case 5:
		return "VC2Wavelet_FidelityFilter"
	case 6:
		return "VC2Wavelet_Daubechies_9_7_IntegerApproximation"
	default:
		return "invalid value"
	}
}
func DecodeTVC2WaveletType(value []byte) (any, error) {
	field, _ := DecodeTUInt8(value[0:1])
	result := TVC2WaveletType(field.(uint8))
	return result, nil
}
func EncodeTVC2WaveletType(value TVC2WaveletType) ([]byte, error) {
	result, _ := EncodeTUInt8(uint8(value))
	return result, nil
}

type TaudioContentDialogueStrongReference TStrongReference
type TColorCorrectionFilterWheelSettingType uint8

func (mt TColorCorrectionFilterWheelSettingType) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TColorCorrectionFilterWheelSettingType) String() string {

	switch s {
	case 0:
		return "ColorCorrectionFilterWheelSetting_CrossEffectFilter"
	case 1:
		return "ColorCorrectionFilterWheelSetting_CCFilter3200K"
	case 2:
		return "ColorCorrectionFilterWheelSetting_CCFilter4300K"
	case 3:
		return "ColorCorrectionFilterWheelSetting_CCFilter6300K"
	case 4:
		return "ColorCorrectionFilterWheelSetting_CCFilter5600K"
	default:
		return "invalid value"
	}
}
func DecodeTColorCorrectionFilterWheelSettingType(value []byte) (any, error) {
	field, _ := DecodeTUInt8(value[0:1])
	result := TColorCorrectionFilterWheelSettingType(field.(uint8))
	return result, nil
}
func EncodeTColorCorrectionFilterWheelSettingType(value TColorCorrectionFilterWheelSettingType) ([]byte, error) {
	result, _ := EncodeTUInt8(uint8(value))
	return result, nil
}

type TaudioObjectInteractionStrongReference TStrongReference
type TGeographicCoordinateArray []TGeographicCoordinate

func EncodeTGeographicCoordinateArray(value TGeographicCoordinateArray) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	binary.BigEndian.PutUint32(result[4:8], 8)
	for _, val := range value {
		field, _ := EncodeTGeographicCoordinate(TGeographicCoordinate(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTGeographicCoordinateArray(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TGeographicCoordinateArray = make([]TGeographicCoordinate, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTGeographicCoordinate(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = field.(TGeographicCoordinate) // else
	}
	return result, nil
}

type TaudioBlockDivergenceStrongReference TStrongReference
type TASMLEKeyIDMapping struct {
	ASMLEKeyID       uint32
	ASMKey           TUInt8Array16
	ASMExpireTime    uint32
	ASMAttributeData uint64
}

func DecodeTASMLEKeyIDMapping(value []byte) (any, error) {
	result := TASMLEKeyIDMapping{}
	var field any
	field, _ = DecodeTUInt32(value[0:4])
	result.ASMLEKeyID = field.(uint32) // else
	field, _ = DecodeTUInt8Array16(value[4:20])
	result.ASMKey = field.(TUInt8Array16) // else
	field, _ = DecodeTUInt32(value[20:24])
	result.ASMExpireTime = field.(uint32) // else
	field, _ = DecodeTUInt64(value[24:32])
	result.ASMAttributeData = field.(uint64) // else
	return result, nil
}
func EncodeTASMLEKeyIDMapping(value TASMLEKeyIDMapping) ([]byte, error) {
	result := []byte{}
	field0, _ := EncodeTUInt32(uint32(value.ASMLEKeyID)) // else
	result = append(result, field0...)
	field1, _ := EncodeTUInt8Array16(TUInt8Array16(value.ASMKey)) // else
	result = append(result, field1...)
	field2, _ := EncodeTUInt32(uint32(value.ASMExpireTime)) // else
	result = append(result, field2...)
	field3, _ := EncodeTUInt64(uint64(value.ASMAttributeData)) // else
	result = append(result, field3...)
	return result, nil
}

type TImageSensorReadoutModeType uint8

func (mt TImageSensorReadoutModeType) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TImageSensorReadoutModeType) String() string {

	switch s {
	case 0:
		return "ImageSensorReadoutMode_InterlacedField"
	case 1:
		return "ImageSensorReadoutMode_InterlacedFrame"
	case 2:
		return "ImageSensorReadoutMode_ProgressiveFrame"
	default:
		return "invalid value"
	}
}
func DecodeTImageSensorReadoutModeType(value []byte) (any, error) {
	field, _ := DecodeTUInt8(value[0:1])
	result := TImageSensorReadoutModeType(field.(uint8))
	return result, nil
}
func EncodeTImageSensorReadoutModeType(value TImageSensorReadoutModeType) ([]byte, error) {
	result, _ := EncodeTUInt8(uint8(value))
	return result, nil
}

type TaudioBlockZoneExclusionStrongReference TStrongReference
type TS309MArray []TS309M

func EncodeTS309MArray(value TS309MArray) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	binary.BigEndian.PutUint32(result[4:8], 8)
	for _, val := range value {
		field, _ := EncodeTS309M(TS309M(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTS309MArray(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TS309MArray = make([]TS309M, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTS309M(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = field.(TS309M) // else
	}
	return result, nil
}

type TaudioBlockPositionStrongReference TStrongReference
type TaudioBlockJumpPositionStrongReference TStrongReference
type TClassDefinitionWeakReference TWeakReference
type TColorPrimary struct {
	X uint16
	Y uint16
}

func DecodeTColorPrimary(value []byte) (any, error) {
	result := TColorPrimary{}
	var field any
	field, _ = DecodeTUInt16(value[0:2])
	result.X = field.(uint16) // else
	field, _ = DecodeTUInt16(value[2:4])
	result.Y = field.(uint16) // else
	return result, nil
}
func EncodeTColorPrimary(value TColorPrimary) ([]byte, error) {
	result := []byte{}
	field0, _ := EncodeTUInt16(uint16(value.X)) // else
	result = append(result, field0...)
	field1, _ := EncodeTUInt16(uint16(value.Y)) // else
	result = append(result, field1...)
	return result, nil
}

type TContainerDefinitionWeakReference TWeakReference
type TfilterSettingStrongReference TStrongReference
type TAutoWhiteBalanceModeType uint8

func (mt TAutoWhiteBalanceModeType) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TAutoWhiteBalanceModeType) String() string {

	switch s {
	case 0:
		return "AutoWhiteBalanceMode_PresetWhiteBalanceSetup"
	case 1:
		return "AutoWhiteBalanceMode_AutomaticWhiteBalanceSetup"
	case 2:
		return "AutoWhiteBalanceMode_HoldWhiteBalanceSetup"
	case 3:
		return "AutoWhiteBalanceMode_OnePushWhiteBalanceSetup"
	default:
		return "invalid value"
	}
}
func DecodeTAutoWhiteBalanceModeType(value []byte) (any, error) {
	field, _ := DecodeTUInt8(value[0:1])
	result := TAutoWhiteBalanceModeType(field.(uint8))
	return result, nil
}
func EncodeTAutoWhiteBalanceModeType(value TAutoWhiteBalanceModeType) ([]byte, error) {
	result, _ := EncodeTUInt8(uint8(value))
	return result, nil
}

type TDataDefinitionWeakReference TWeakReference
type TExtUMIDArray []TExtUMID

func EncodeTExtUMIDArray(value TExtUMIDArray) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	binary.BigEndian.PutUint32(result[4:8], 32)
	for _, val := range value {
		field, _ := EncodeTExtUMID(TExtUMID(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTExtUMIDArray(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TExtUMIDArray = make([]TExtUMID, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTExtUMID(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = field.(TExtUMID) // else
	}
	return result, nil
}

type TreferenceScreenCentrePositionStrongReference TStrongReference
type TDefinitionObjectWeakReference TWeakReference
type TInterpolationDefinitionWeakReference TWeakReference
type TAutoFocusSensingAreaSettingType uint8

func (mt TAutoFocusSensingAreaSettingType) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TAutoFocusSensingAreaSettingType) String() string {

	switch s {
	case 0:
		return "AutoFocusSensingAreaSetting_ManualFocusMode"
	case 1:
		return "AutoFocusSensingAreaSetting_CenterSensitiveAutoFocusMode"
	case 2:
		return "AutoFocusSensingAreaSetting_FullScreenSensingAutoFocusMode"
	case 3:
		return "AutoFocusSensingAreaSetting_MultiSpotSensingAutoFocusMode"
	case 4:
		return "AutoFocusSensingAreaSetting_SingleSpotSensingAutoFocusMode"
	default:
		return "invalid value"
	}
}
func DecodeTAutoFocusSensingAreaSettingType(value []byte) (any, error) {
	field, _ := DecodeTUInt8(value[0:1])
	result := TAutoFocusSensingAreaSettingType(field.(uint8))
	return result, nil
}
func EncodeTAutoFocusSensingAreaSettingType(value TAutoFocusSensingAreaSettingType) ([]byte, error) {
	result, _ := EncodeTUInt8(uint8(value))
	return result, nil
}

type TJ2KExtendedCapabilities struct {
	Pcap  uint32
	Ccapi TUInt16Array
}

func DecodeTJ2KExtendedCapabilities(value []byte) (any, error) {
	result := TJ2KExtendedCapabilities{}
	var field any
	field, _ = DecodeTUInt32(value[0:4])
	result.Pcap = field.(uint32) // else
	field, _ = DecodeTUInt16Array(value[4:])
	result.Ccapi = field.(TUInt16Array) // else
	return result, nil
}
func EncodeTJ2KExtendedCapabilities(value TJ2KExtendedCapabilities) ([]byte, error) {
	result := []byte{}
	field0, _ := EncodeTUInt32(uint32(value.Pcap)) // else
	result = append(result, field0...)
	field1, _ := EncodeTUInt16Array(TUInt16Array(value.Ccapi)) // else
	result = append(result, field1...)
	return result, nil
}

type TPackageWeakReference TWeakReference
type TreferenceScreenWidthStrongReference TStrongReference
type TOperationDefinitionWeakReference TWeakReference
type TgainInteractionRangeStrongReference TStrongReference
type THEVCCodedContentType uint8

func (mt THEVCCodedContentType) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s THEVCCodedContentType) String() string {

	switch s {
	case 0:
		return "HEVCCodedContent_NotKnown"
	case 1:
		return "HEVCCodedContent_ProgressiveFramePicture"
	case 2:
		return "HEVCCodedContent_InterlacedFieldPicture"
	default:
		return "invalid value"
	}
}
func DecodeTHEVCCodedContentType(value []byte) (any, error) {
	field, _ := DecodeTUInt8(value[0:1])
	result := THEVCCodedContentType(field.(uint8))
	return result, nil
}
func EncodeTHEVCCodedContentType(value THEVCCodedContentType) ([]byte, error) {
	result, _ := EncodeTUInt8(uint8(value))
	return result, nil
}

type TParameterDefinitionWeakReference TWeakReference
type TpositionInteractionRangeStrongReference TStrongReference
type TUInt8Array []uint8

func EncodeTUInt8Array(value TUInt8Array) ([]byte, error) {
	result := []byte{}
	for _, val := range value {
		field, _ := EncodeTUInt8(uint8(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTUInt8Array(value []byte) (any, error) {
	arrayCount := len(value) / 1
	var result TUInt8Array = make([]uint8, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTUInt8(value[i*1 : i*1+1])
		result[i] = field.(uint8) // else
	}
	return result, nil
}

type TTypeDefinitionWeakReference TWeakReference
type TJ2KComponentSizingArray []TJ2KComponentSizing

func EncodeTJ2KComponentSizingArray(value TJ2KComponentSizingArray) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	binary.BigEndian.PutUint32(result[4:8], 3)
	for _, val := range value {
		field, _ := EncodeTJ2KComponentSizing(TJ2KComponentSizing(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTJ2KComponentSizingArray(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TJ2KComponentSizingArray = make([]TJ2KComponentSizing, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTJ2KComponentSizing(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = field.(TJ2KComponentSizing) // else
	}
	return result, nil
}

type TaudioBlockZoneStrongReference TStrongReference
type TPluginDefinitionWeakReference TWeakReference
type TTIFFByteOrderType uint16

func (mt TTIFFByteOrderType) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TTIFFByteOrderType) String() string {

	switch s {
	case 7777:
		return "TIFFByteOrder_BigEndian"
	case 7373:
		return "TIFFByteOrder_LittleEndian"
	default:
		return "invalid value"
	}
}
func DecodeTTIFFByteOrderType(value []byte) (any, error) {
	field, _ := DecodeTUInt16(value[0:2])
	result := TTIFFByteOrderType(field.(uint16))
	return result, nil
}
func EncodeTTIFFByteOrderType(value TTIFFByteOrderType) ([]byte, error) {
	result, _ := EncodeTUInt16(uint16(value))
	return result, nil
}

type TCodecDefinitionWeakReference TWeakReference
type TUInt8Array12 [12]uint8

func EncodeTUInt8Array12(value TUInt8Array12) ([]byte, error) {
	result := []byte{}
	for _, val := range value {
		field, _ := EncodeTUInt8(uint8(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTUInt8Array12(value []byte) (any, error) {
	var result TUInt8Array12 = [12]uint8{}
	for i := 0; i < 12; i++ {
		field, _ := DecodeTUInt8(value[i*1 : i*1+1])
		result[i] = field.(uint8) // else
	}
	return result, nil
}

type TThreeColorPrimaries [3]TColorPrimary

func EncodeTThreeColorPrimaries(value TThreeColorPrimaries) ([]byte, error) {
	result := []byte{}
	for _, val := range value {
		field, _ := EncodeTColorPrimary(TColorPrimary(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTThreeColorPrimaries(value []byte) (any, error) {
	var result TThreeColorPrimaries = [3]TColorPrimary{}
	for i := 0; i < 3; i++ {
		field, _ := DecodeTColorPrimary(value[i*4 : i*4+4])
		result[i] = field.(TColorPrimary) // else
	}
	return result, nil
}

type TPropertyDefinitionWeakReference TWeakReference
type TTypeDefinitionExtendibleEnumerationWeakReference TWeakReference
type TContentStorageStrongReference TStrongReference
type TidentifierStrongReferenceSet []TidentifierStrongReference

func EncodeTidentifierStrongReferenceSet(value TidentifierStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTidentifierStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TidentifierStrongReferenceSet = make([]TidentifierStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TidentifierStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TInt32Array []int32

func EncodeTInt32Array(value TInt32Array) ([]byte, error) {
	result := []byte{}
	for _, val := range value {
		field, _ := EncodeTInt32(int32(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTInt32Array(value []byte) (any, error) {
	arrayCount := len(value) / 4
	var result TInt32Array = make([]int32, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTInt32(value[i*4 : i*4+4])
		result[i] = field.(int32) // else
	}
	return result, nil
}

type TDictionaryStrongReference TStrongReference
type TTIFFCompressionKindType uint16

func (mt TTIFFCompressionKindType) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TTIFFCompressionKindType) String() string {

	switch s {
	case 1:
		return "TIFFCompressionKind_NoCompression"
	case 7:
		return "TIFFCompressionKind_LosslessHuffmanJPEG"
	default:
		return "invalid value"
	}
}
func DecodeTTIFFCompressionKindType(value []byte) (any, error) {
	field, _ := DecodeTUInt16(value[0:2])
	result := TTIFFCompressionKindType(field.(uint16))
	return result, nil
}
func EncodeTTIFFCompressionKindType(value TTIFFCompressionKindType) ([]byte, error) {
	result, _ := EncodeTUInt16(uint16(value))
	return result, nil
}

type TUInt32Set []uint32

func EncodeTUInt32Set(value TUInt32Set) ([]byte, error) {
	result := []byte{}
	for _, val := range value {
		field, _ := EncodeTUInt32(uint32(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTUInt32Set(value []byte) (any, error) {
	arrayCount := len(value) / 4
	var result TUInt32Set = make([]uint32, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTUInt32(value[i*4 : i*4+4])
		result[i] = field.(uint32) // else
	}
	return result, nil
}

type TEssenceDescriptorStrongReference TStrongReference
type TRIFFChunkStrongReference TStrongReference
type TInt64Array []int64

func EncodeTInt64Array(value TInt64Array) ([]byte, error) {
	result := []byte{}
	for _, val := range value {
		field, _ := EncodeTInt64(int64(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTInt64Array(value []byte) (any, error) {
	arrayCount := len(value) / 8
	var result TInt64Array = make([]int64, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTInt64(value[i*8 : i*8+8])
		result[i] = field.(int64) // else
	}
	return result, nil
}

type TNetworkLocatorStrongReference TStrongReference
type TTIFFPhotometricInterpretationKindType uint16

func (mt TTIFFPhotometricInterpretationKindType) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TTIFFPhotometricInterpretationKindType) String() string {

	switch s {
	case 32803:
		return "TIFFPhotometricInterpretationKind_ColorFilterArray"
	case 34892:
		return "TIFFPhotometricInterpretationKind_LinearRaw"
	default:
		return "invalid value"
	}
}
func DecodeTTIFFPhotometricInterpretationKindType(value []byte) (any, error) {
	field, _ := DecodeTUInt16(value[0:2])
	result := TTIFFPhotometricInterpretationKindType(field.(uint16))
	return result, nil
}
func EncodeTTIFFPhotometricInterpretationKindType(value TTIFFPhotometricInterpretationKindType) ([]byte, error) {
	result, _ := EncodeTUInt16(uint16(value))
	return result, nil
}

type TtitleStrongReferenceSet []TtitleStrongReference

func EncodeTtitleStrongReferenceSet(value TtitleStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTtitleStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TtitleStrongReferenceSet = make([]TtitleStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TtitleStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TOperationGroupStrongReference TStrongReference
type TASMLEKeyIDMappingSet []TASMLEKeyIDMapping

func EncodeTASMLEKeyIDMappingSet(value TASMLEKeyIDMappingSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	binary.BigEndian.PutUint32(result[4:8], 32)
	for _, val := range value {
		field, _ := EncodeTASMLEKeyIDMapping(TASMLEKeyIDMapping(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTASMLEKeyIDMappingSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TASMLEKeyIDMappingSet = make([]TASMLEKeyIDMapping, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTASMLEKeyIDMapping(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = field.(TASMLEKeyIDMapping) // else
	}
	return result, nil
}

type TSegmentStrongReference TStrongReference
type TDataValue []uint8

func EncodeTDataValue(value TDataValue) ([]byte, error) {
	result := []byte{}
	for _, val := range value {
		field, _ := EncodeTUInt8(uint8(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTDataValue(value []byte) (any, error) {
	arrayCount := len(value) / 1
	var result TDataValue = make([]uint8, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTUInt8(value[i*1 : i*1+1])
		result[i] = field.(uint8) // else
	}
	return result, nil
}

type TalternativeTitleStrongReferenceSet []TalternativeTitleStrongReference

func EncodeTalternativeTitleStrongReferenceSet(value TalternativeTitleStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTalternativeTitleStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TalternativeTitleStrongReferenceSet = make([]TalternativeTitleStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TalternativeTitleStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TDescriptiveFrameworkStrongReference TStrongReference
type TMetaDefinitionStrongReference TStrongReference
type TKLVDataDefinitionStrongReference TStrongReference
type TTaggedValueDefinitionStrongReference TStrongReference
type TPositionArray []uint8

func EncodeTPositionArray(value TPositionArray) ([]byte, error) {
	result := []byte{}
	for _, val := range value {
		field, _ := EncodeTUInt8(uint8(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTPositionArray(value []byte) (any, error) {
	arrayCount := len(value) / 1
	var result TPositionArray = make([]uint8, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTUInt8(value[i*1 : i*1+1])
		result[i] = field.(uint8) // else
	}
	return result, nil
}

type TentityStrongReferenceSet []TentityStrongReference

func EncodeTentityStrongReferenceSet(value TentityStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTentityStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TentityStrongReferenceSet = make([]TentityStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TentityStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TDescriptiveObjectStrongReference TStrongReference
type TSourceClipStrongReference TStrongReference
type TCueWordsStrongReference TStrongReference
type TParameterDefinitionStrongReferenceSet []TParameterDefinitionStrongReference

func EncodeTParameterDefinitionStrongReferenceSet(value TParameterDefinitionStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTParameterDefinitionStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TParameterDefinitionStrongReferenceSet = make([]TParameterDefinitionStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TParameterDefinitionStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TParameterDefinitionWeakReferenceSet []TParameterDefinitionWeakReference

func EncodeTParameterDefinitionWeakReferenceSet(value TParameterDefinitionWeakReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTUInt8Array(TUInt8Array(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTParameterDefinitionWeakReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TParameterDefinitionWeakReferenceSet = make([]TParameterDefinitionWeakReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTUInt8Array(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TParameterDefinitionWeakReference(field.(TUInt8Array)) // else
	}
	return result, nil
}

type TGeographicAreaStrongReference TStrongReference
type TSourceReferenceStrongReference TStrongReference
type TDataDefinitionWeakReferenceSet []TDataDefinitionWeakReference

func EncodeTDataDefinitionWeakReferenceSet(value TDataDefinitionWeakReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTUInt8Array(TUInt8Array(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTDataDefinitionWeakReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TDataDefinitionWeakReferenceSet = make([]TDataDefinitionWeakReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTUInt8Array(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TDataDefinitionWeakReference(field.(TUInt8Array)) // else
	}
	return result, nil
}

type TGeographicPolygonStrongReference TStrongReference
type TSubDescriptorStrongReference TStrongReference
type TClassDefinitionStrongReference TStrongReference
type TsubjectStrongReferenceSet []TsubjectStrongReference

func EncodeTsubjectStrongReferenceSet(value TsubjectStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTsubjectStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TsubjectStrongReferenceSet = make([]TsubjectStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TsubjectStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TRGBAComponent struct {
	Code          TRGBAComponentKind
	ComponentSize uint8
}

func DecodeTRGBAComponent(value []byte) (any, error) {
	result := TRGBAComponent{}
	var field any
	field, _ = DecodeTUInt8(value[0:1])
	result.Code = TRGBAComponentKind(field.(uint8)) // else
	field, _ = DecodeTUInt8(value[1:2])
	result.ComponentSize = field.(uint8) // else
	return result, nil
}
func EncodeTRGBAComponent(value TRGBAComponent) ([]byte, error) {
	result := []byte{}
	field0, _ := EncodeTUInt8(uint8(value.Code)) // else
	result = append(result, field0...)
	field1, _ := EncodeTUInt8(uint8(value.ComponentSize)) // else
	result = append(result, field1...)
	return result, nil
}

type TMetaDictionaryStrongReference TStrongReference
type TUInt8Array8 [8]uint8

func EncodeTUInt8Array8(value TUInt8Array8) ([]byte, error) {
	result := []byte{}
	for _, val := range value {
		field, _ := EncodeTUInt8(uint8(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTUInt8Array8(value []byte) (any, error) {
	var result TUInt8Array8 = [8]uint8{}
	for i := 0; i < 8; i++ {
		field, _ := DecodeTUInt8(value[i*1 : i*1+1])
		result[i] = field.(uint8) // else
	}
	return result, nil
}

type TPluginDefinitionWeakReferenceSet []TPluginDefinitionWeakReference

func EncodeTPluginDefinitionWeakReferenceSet(value TPluginDefinitionWeakReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTUInt8Array(TUInt8Array(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTPluginDefinitionWeakReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TPluginDefinitionWeakReferenceSet = make([]TPluginDefinitionWeakReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTUInt8Array(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TPluginDefinitionWeakReference(field.(TUInt8Array)) // else
	}
	return result, nil
}

type TCodecDefinitionStrongReference TStrongReference
type TPrefaceStrongReference TStrongReference
type TNameValueStrongReference TStrongReference
type TComponentStrongReference TStrongReference
type TPluginDefinitionStrongReferenceSet []TPluginDefinitionStrongReference

func EncodeTPluginDefinitionStrongReferenceSet(value TPluginDefinitionStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTPluginDefinitionStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TPluginDefinitionStrongReferenceSet = make([]TPluginDefinitionStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TPluginDefinitionStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TdescriptionStrongReferenceSet []TdescriptionStrongReference

func EncodeTdescriptionStrongReferenceSet(value TdescriptionStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTdescriptionStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TdescriptionStrongReferenceSet = make([]TdescriptionStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TdescriptionStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TTextBasedObjectStrongReference TStrongReference
type TContainerDefinitionStrongReference TStrongReference
type TEntryStrongReference TStrongReference
type TLocatorStrongReferenceVector []TLocatorStrongReference

func EncodeTLocatorStrongReferenceVector(value TLocatorStrongReferenceVector) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTLocatorStrongReferenceVector(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TLocatorStrongReferenceVector = make([]TLocatorStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TLocatorStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TPropertyDefinitionWeakReferenceSet []TPropertyDefinitionWeakReference

func EncodeTPropertyDefinitionWeakReferenceSet(value TPropertyDefinitionWeakReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTUInt8Array(TUInt8Array(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTPropertyDefinitionWeakReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TPropertyDefinitionWeakReferenceSet = make([]TPropertyDefinitionWeakReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTUInt8Array(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TPropertyDefinitionWeakReference(field.(TUInt8Array)) // else
	}
	return result, nil
}

type TControlPointStrongReference TStrongReference
type TRegisterAdministrationStrongReference TStrongReference
type TDateStruct struct {
	Year  int16
	Month uint8
	Day   uint8
}

func DecodeTDateStruct(value []byte) (any, error) {
	result := TDateStruct{}
	var field any
	field, _ = DecodeTInt16(value[0:2])
	result.Year = field.(int16) // else
	field, _ = DecodeTUInt8(value[2:3])
	result.Month = field.(uint8) // else
	field, _ = DecodeTUInt8(value[3:4])
	result.Day = field.(uint8) // else
	return result, nil
}
func EncodeTDateStruct(value TDateStruct) ([]byte, error) {
	result := []byte{}
	field0, _ := EncodeTInt16(int16(value.Year)) // else
	result = append(result, field0...)
	field1, _ := EncodeTUInt8(uint8(value.Month)) // else
	result = append(result, field1...)
	field2, _ := EncodeTUInt8(uint8(value.Day)) // else
	result = append(result, field2...)
	return result, nil
}

type TDataDefinitionStrongReference TStrongReference
type TEntryAdministrationStrongReference TStrongReference
type TdateStrongReferenceSet []TdateStrongReference

func EncodeTdateStrongReferenceSet(value TdateStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTdateStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TdateStrongReferenceSet = make([]TdateStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TdateStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TPackageMarkerStrongReference TStrongReference
type TTypeDefinitionExtendibleEnumerationWeakReferenceSet []TTypeDefinitionExtendibleEnumerationWeakReference

func EncodeTTypeDefinitionExtendibleEnumerationWeakReferenceSet(value TTypeDefinitionExtendibleEnumerationWeakReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTUInt8Array(TUInt8Array(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTTypeDefinitionExtendibleEnumerationWeakReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TTypeDefinitionExtendibleEnumerationWeakReferenceSet = make([]TTypeDefinitionExtendibleEnumerationWeakReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTUInt8Array(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TTypeDefinitionExtendibleEnumerationWeakReference(field.(TUInt8Array)) // else
	}
	return result, nil
}

type TPropertyDefinitionStrongReferenceSet []TPropertyDefinitionStrongReference

func EncodeTPropertyDefinitionStrongReferenceSet(value TPropertyDefinitionStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTPropertyDefinitionStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TPropertyDefinitionStrongReferenceSet = make([]TPropertyDefinitionStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TPropertyDefinitionStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TEssenceDataStrongReference TStrongReference
type TApplicationPluginObjectStrongReference TStrongReference
type TTrackStrongReferenceVector []TTrackStrongReference

func EncodeTTrackStrongReferenceVector(value TTrackStrongReferenceVector) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTTrackStrongReferenceVector(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TTrackStrongReferenceVector = make([]TTrackStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TTrackStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TIdentificationStrongReference TStrongReference
type TtechnicalAttributeInt64StrongReference TStrongReference
type TInterpolationDefinitionStrongReference TStrongReference
type TtypeStrongReferenceSet []TtypeStrongReference

func EncodeTtypeStrongReferenceSet(value TtypeStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTtypeStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TtypeStrongReferenceSet = make([]TtypeStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TtypeStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TExtensionSchemeStrongReference TStrongReference
type TTimeStruct struct {
	Hour     uint8
	Minute   uint8
	Second   uint8
	Fraction uint8
}

func DecodeTTimeStruct(value []byte) (any, error) {
	result := TTimeStruct{}
	var field any
	field, _ = DecodeTUInt8(value[0:1])
	result.Hour = field.(uint8) // else
	field, _ = DecodeTUInt8(value[1:2])
	result.Minute = field.(uint8) // else
	field, _ = DecodeTUInt8(value[2:3])
	result.Second = field.(uint8) // else
	field, _ = DecodeTUInt8(value[3:4])
	result.Fraction = field.(uint8) // else
	return result, nil
}
func EncodeTTimeStruct(value TTimeStruct) ([]byte, error) {
	result := []byte{}
	field0, _ := EncodeTUInt8(uint8(value.Hour)) // else
	result = append(result, field0...)
	field1, _ := EncodeTUInt8(uint8(value.Minute)) // else
	result = append(result, field1...)
	field2, _ := EncodeTUInt8(uint8(value.Second)) // else
	result = append(result, field2...)
	field3, _ := EncodeTUInt8(uint8(value.Fraction)) // else
	result = append(result, field3...)
	return result, nil
}

type TLocatorStrongReference TStrongReference
type TOperationDefinitionWeakReferenceVector []TOperationDefinitionWeakReference

func EncodeTOperationDefinitionWeakReferenceVector(value TOperationDefinitionWeakReferenceVector) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTUInt8Array(TUInt8Array(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTOperationDefinitionWeakReferenceVector(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TOperationDefinitionWeakReferenceVector = make([]TOperationDefinitionWeakReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTUInt8Array(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TOperationDefinitionWeakReference(field.(TUInt8Array)) // else
	}
	return result, nil
}

type TTypeDefinitionStrongReferenceSet []TTypeDefinitionStrongReference

func EncodeTTypeDefinitionStrongReferenceSet(value TTypeDefinitionStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTTypeDefinitionStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TTypeDefinitionStrongReferenceSet = make([]TTypeDefinitionStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TTypeDefinitionStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TSegmentStrongReferenceVector []TSegmentStrongReference

func EncodeTSegmentStrongReferenceVector(value TSegmentStrongReferenceVector) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTSegmentStrongReferenceVector(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TSegmentStrongReferenceVector = make([]TSegmentStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TSegmentStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TOrganizationGlobalReferenceSet []TOrganizationGlobalReference

func EncodeTOrganizationGlobalReferenceSet(value TOrganizationGlobalReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTUInt8Array(TUInt8Array(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTOrganizationGlobalReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TOrganizationGlobalReferenceSet = make([]TOrganizationGlobalReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTUInt8Array(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TOrganizationGlobalReference(field.(TUInt8Array)) // else
	}
	return result, nil
}

type TSourceReferenceStrongReferenceVector []TSourceReferenceStrongReference

func EncodeTSourceReferenceStrongReferenceVector(value TSourceReferenceStrongReferenceVector) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTSourceReferenceStrongReferenceVector(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TSourceReferenceStrongReferenceVector = make([]TSourceReferenceStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TSourceReferenceStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TPackageStrongReference TStrongReference
type TTaggedValueStrongReferenceVector []TTaggedValueStrongReference

func EncodeTTaggedValueStrongReferenceVector(value TTaggedValueStrongReferenceVector) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTTaggedValueStrongReferenceVector(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TTaggedValueStrongReferenceVector = make([]TTaggedValueStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TTaggedValueStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TlanguageStrongReferenceSet []TlanguageStrongReference

func EncodeTlanguageStrongReferenceSet(value TlanguageStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTlanguageStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TlanguageStrongReferenceSet = make([]TlanguageStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TlanguageStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TTrackStrongReference TStrongReference
type TKLVDataDefinitionStrongReferenceSet []TKLVDataDefinitionStrongReference

func EncodeTKLVDataDefinitionStrongReferenceSet(value TKLVDataDefinitionStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTKLVDataDefinitionStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TKLVDataDefinitionStrongReferenceSet = make([]TKLVDataDefinitionStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TKLVDataDefinitionStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TTypeDefinitionWeakReferenceVector []TTypeDefinitionWeakReference

func EncodeTTypeDefinitionWeakReferenceVector(value TTypeDefinitionWeakReferenceVector) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTUInt8Array(TUInt8Array(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTTypeDefinitionWeakReferenceVector(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TTypeDefinitionWeakReferenceVector = make([]TTypeDefinitionWeakReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTUInt8Array(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TTypeDefinitionWeakReference(field.(TUInt8Array)) // else
	}
	return result, nil
}

type TKLVDataStrongReferenceVector []TKLVDataStrongReference

func EncodeTKLVDataStrongReferenceVector(value TKLVDataStrongReferenceVector) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTKLVDataStrongReferenceVector(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TKLVDataStrongReferenceVector = make([]TKLVDataStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TKLVDataStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TParameterStrongReferenceVector []TParameterStrongReference

func EncodeTParameterStrongReferenceVector(value TParameterStrongReferenceVector) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTParameterStrongReferenceVector(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TParameterStrongReferenceVector = make([]TParameterStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TParameterStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TLocationGlobalReferenceSet []TLocationGlobalReference

func EncodeTLocationGlobalReferenceSet(value TLocationGlobalReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTUInt8Array(TUInt8Array(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTLocationGlobalReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TLocationGlobalReferenceSet = make([]TLocationGlobalReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTUInt8Array(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TLocationGlobalReference(field.(TUInt8Array)) // else
	}
	return result, nil
}

type TDataDefinitionWeakReferenceVector []TDataDefinitionWeakReference

func EncodeTDataDefinitionWeakReferenceVector(value TDataDefinitionWeakReferenceVector) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTUInt8Array(TUInt8Array(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTDataDefinitionWeakReferenceVector(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TDataDefinitionWeakReferenceVector = make([]TDataDefinitionWeakReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTUInt8Array(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TDataDefinitionWeakReference(field.(TUInt8Array)) // else
	}
	return result, nil
}

type TrightsStrongReferenceSet []TrightsStrongReference

func EncodeTrightsStrongReferenceSet(value TrightsStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTrightsStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TrightsStrongReferenceSet = make([]TrightsStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TrightsStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TcoverageStrongReferenceSet []TcoverageStrongReference

func EncodeTcoverageStrongReferenceSet(value TcoverageStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTcoverageStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TcoverageStrongReferenceSet = make([]TcoverageStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TcoverageStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TOperationDefinitionStrongReference TStrongReference
type TTaggedValueDefinitionStrongReferenceSet []TTaggedValueDefinitionStrongReference

func EncodeTTaggedValueDefinitionStrongReferenceSet(value TTaggedValueDefinitionStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTTaggedValueDefinitionStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TTaggedValueDefinitionStrongReferenceSet = make([]TTaggedValueDefinitionStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TTaggedValueDefinitionStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TParticipantGlobalReferenceVector []TParticipantGlobalReference

func EncodeTParticipantGlobalReferenceVector(value TParticipantGlobalReferenceVector) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTUInt8Array(TUInt8Array(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTParticipantGlobalReferenceVector(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TParticipantGlobalReferenceVector = make([]TParticipantGlobalReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTUInt8Array(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TParticipantGlobalReference(field.(TUInt8Array)) // else
	}
	return result, nil
}

type TParameterStrongReference TStrongReference
type TAS_12_DescriptiveObjectStrongReference TStrongReference
type TFileDescriptorStrongReferenceVector []TFileDescriptorStrongReference

func EncodeTFileDescriptorStrongReferenceVector(value TFileDescriptorStrongReferenceVector) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTFileDescriptorStrongReferenceVector(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TFileDescriptorStrongReferenceVector = make([]TFileDescriptorStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TFileDescriptorStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TcoreMetadataStrongReference TStrongReference
type TratingStrongReferenceSet []TratingStrongReference

func EncodeTratingStrongReferenceSet(value TratingStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTratingStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TratingStrongReferenceSet = make([]TratingStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TratingStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TParameterDefinitionStrongReference TStrongReference
type TDescriptiveObjectStrongReferenceSet []TDescriptiveObjectStrongReference

func EncodeTDescriptiveObjectStrongReferenceSet(value TDescriptiveObjectStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTDescriptiveObjectStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TDescriptiveObjectStrongReferenceSet = make([]TDescriptiveObjectStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TDescriptiveObjectStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TOrganizationGlobalReferenceVector []TOrganizationGlobalReference

func EncodeTOrganizationGlobalReferenceVector(value TOrganizationGlobalReferenceVector) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTUInt8Array(TUInt8Array(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTOrganizationGlobalReferenceVector(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TOrganizationGlobalReferenceVector = make([]TOrganizationGlobalReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTUInt8Array(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TOrganizationGlobalReference(field.(TUInt8Array)) // else
	}
	return result, nil
}

type TRIFFChunkStrongReferenceVector []TRIFFChunkStrongReference

func EncodeTRIFFChunkStrongReferenceVector(value TRIFFChunkStrongReferenceVector) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTRIFFChunkStrongReferenceVector(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TRIFFChunkStrongReferenceVector = make([]TRIFFChunkStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TRIFFChunkStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TmetadataSchemaInformationStrongReference TStrongReference
type TPluginDefinitionStrongReference TStrongReference
type TcustomRelationStrongReferenceSet []TcustomRelationStrongReference

func EncodeTcustomRelationStrongReferenceSet(value TcustomRelationStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTcustomRelationStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TcustomRelationStrongReferenceSet = make([]TcustomRelationStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TcustomRelationStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TversionStrongReference TStrongReference
type TNameValueStrongReferenceSet []TNameValueStrongReference

func EncodeTNameValueStrongReferenceSet(value TNameValueStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTNameValueStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TNameValueStrongReferenceSet = make([]TNameValueStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TNameValueStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TDescriptiveObjectStrongReferenceVector []TDescriptiveObjectStrongReference

func EncodeTDescriptiveObjectStrongReferenceVector(value TDescriptiveObjectStrongReferenceVector) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTDescriptiveObjectStrongReferenceVector(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TDescriptiveObjectStrongReferenceVector = make([]TDescriptiveObjectStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TDescriptiveObjectStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TLocationGlobalReferenceVector []TLocationGlobalReference

func EncodeTLocationGlobalReferenceVector(value TLocationGlobalReferenceVector) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTUInt8Array(TUInt8Array(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTLocationGlobalReferenceVector(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TLocationGlobalReferenceVector = make([]TLocationGlobalReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTUInt8Array(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TLocationGlobalReference(field.(TUInt8Array)) // else
	}
	return result, nil
}

type TPropertyDefinitionStrongReference TStrongReference
type TbasicRelationStrongReferenceSet []TbasicRelationStrongReference

func EncodeTbasicRelationStrongReferenceSet(value TbasicRelationStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTbasicRelationStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TbasicRelationStrongReferenceSet = make([]TbasicRelationStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TbasicRelationStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TApplicationPluginObjectStrongReferenceSet []TApplicationPluginObjectStrongReference

func EncodeTApplicationPluginObjectStrongReferenceSet(value TApplicationPluginObjectStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTApplicationPluginObjectStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TApplicationPluginObjectStrongReferenceSet = make([]TApplicationPluginObjectStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TApplicationPluginObjectStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TTaggedValueStrongReference TStrongReference
type TParticipantGlobalReference TWeakReference
type TSubDescriptorStrongReferenceVector []TSubDescriptorStrongReference

func EncodeTSubDescriptorStrongReferenceVector(value TSubDescriptorStrongReferenceVector) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTSubDescriptorStrongReferenceVector(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TSubDescriptorStrongReferenceVector = make([]TSubDescriptorStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TSubDescriptorStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TpublicationHistoryStrongReference TStrongReference
type TOrganizationGlobalReference TWeakReference
type TTypeDefinitionStrongReference TStrongReference
type TplanningStrongReference TStrongReference
type TLocationGlobalReference TWeakReference
type TformatStrongReferenceSet []TformatStrongReference

func EncodeTformatStrongReferenceSet(value TformatStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTformatStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TformatStrongReferenceSet = make([]TformatStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TformatStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TExtensionSchemeStrongReferenceSet []TExtensionSchemeStrongReference

func EncodeTExtensionSchemeStrongReferenceSet(value TExtensionSchemeStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTExtensionSchemeStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TExtensionSchemeStrongReferenceSet = make([]TExtensionSchemeStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TExtensionSchemeStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TtypeGroupStrongReference TStrongReference
type TApplicationPluginObjectGlobalReference TWeakReference
type TformatGroupStrongReference TStrongReference
type TRegisterEntryStrongReferenceVector []TEntryStrongReference

func EncodeTRegisterEntryStrongReferenceVector(value TRegisterEntryStrongReferenceVector) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTRegisterEntryStrongReferenceVector(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TRegisterEntryStrongReferenceVector = make([]TEntryStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TEntryStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TKLVDataStrongReference TStrongReference
type TDescriptiveMarkerGlobalReference TWeakReference
type TMetaDefinitionStrongReferenceSet []TMetaDefinitionStrongReference

func EncodeTMetaDefinitionStrongReferenceSet(value TMetaDefinitionStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTMetaDefinitionStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TMetaDefinitionStrongReferenceSet = make([]TMetaDefinitionStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TMetaDefinitionStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TpartStrongReferenceSet []TpartStrongReference

func EncodeTpartStrongReferenceSet(value TpartStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTpartStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TpartStrongReferenceSet = make([]TpartStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TpartStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TstatusGroupStrongReference TStrongReference
type TspatialStrongReference TStrongReference
type TComponentStrongReferenceVector []TComponentStrongReference

func EncodeTComponentStrongReferenceVector(value TComponentStrongReferenceVector) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTComponentStrongReferenceVector(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TComponentStrongReferenceVector = make([]TComponentStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TComponentStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TaddressStrongReference TStrongReference
type TParticipantGlobalReferenceSet []TParticipantGlobalReference

func EncodeTParticipantGlobalReferenceSet(value TParticipantGlobalReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTUInt8Array(TUInt8Array(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTParticipantGlobalReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TParticipantGlobalReferenceSet = make([]TParticipantGlobalReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTUInt8Array(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TParticipantGlobalReference(field.(TUInt8Array)) // else
	}
	return result, nil
}

type TAS_11_Audio_Track_Layout_Enum uint8

func (mt TAS_11_Audio_Track_Layout_Enum) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TAS_11_Audio_Track_Layout_Enum) String() string {

	switch s {
	case 0:
		return "Layout_EBU_R_48_1a"
	case 1:
		return "Layout_EBU_R_48_1b"
	case 2:
		return "Layout_EBU_R_48_1c"
	case 3:
		return "Layout_EBU_R_48_2a"
	case 4:
		return "Layout_EBU_R_48_2b"
	case 5:
		return "Layout_EBU_R_48_2c"
	case 6:
		return "Layout_EBU_R_48_3a"
	case 7:
		return "Layout_EBU_R_48_3b"
	case 8:
		return "Layout_EBU_R_48_4a"
	case 9:
		return "Layout_EBU_R_48_4b"
	case 10:
		return "Layout_EBU_R_48_4c"
	case 11:
		return "Layout_EBU_R_48_5a"
	case 12:
		return "Layout_EBU_R_48_5b"
	case 13:
		return "Layout_EBU_R_48_6a"
	case 14:
		return "Layout_EBU_R_48_6b"
	case 15:
		return "Layout_EBU_R_48_7a"
	case 16:
		return "Layout_EBU_R_48_7b"
	case 17:
		return "Layout_EBU_R_48_8a"
	case 18:
		return "Layout_EBU_R_48_8b"
	case 19:
		return "Layout_EBU_R_48_8c"
	case 20:
		return "Layout_EBU_R_48_9a"
	case 21:
		return "Layout_EBU_R_48_9b"
	case 22:
		return "Layout_EBU_R_48_10a"
	case 23:
		return "Layout_EBU_R_48_11a"
	case 24:
		return "Layout_EBU_R_48_11b"
	case 25:
		return "Layout_EBU_R_48_11c"
	case 26:
		return "Layout_EBU_R_123_2a"
	case 27:
		return "Layout_EBU_R_123_4a"
	case 28:
		return "Layout_EBU_R_123_4b"
	case 29:
		return "Layout_EBU_R_123_4c"
	case 30:
		return "Layout_EBU_R_123_8a"
	case 31:
		return "Layout_EBU_R_123_8b"
	case 32:
		return "Layout_EBU_R_123_8c"
	case 33:
		return "Layout_EBU_R_123_8d"
	case 34:
		return "Layout_EBU_R_123_8e"
	case 35:
		return "Layout_EBU_R_123_8f"
	case 36:
		return "Layout_EBU_R_123_8g"
	case 37:
		return "Layout_EBU_R_123_8h"
	case 38:
		return "Layout_EBU_R_123_8i"
	case 39:
		return "Layout_EBU_R_123_12a"
	case 40:
		return "Layout_EBU_R_123_12b"
	case 41:
		return "Layout_EBU_R_123_12c"
	case 42:
		return "Layout_EBU_R_123_12d"
	case 43:
		return "Layout_EBU_R_123_12e"
	case 44:
		return "Layout_EBU_R_123_12f"
	case 45:
		return "Layout_EBU_R_123_12g"
	case 46:
		return "Layout_EBU_R_123_12h"
	case 47:
		return "Layout_EBU_R_123_16a"
	case 48:
		return "Layout_EBU_R_123_16b"
	case 49:
		return "Layout_EBU_R_123_16c"
	case 50:
		return "Layout_EBU_R_123_16d"
	case 51:
		return "Layout_EBU_R_123_16e"
	case 52:
		return "Layout_EBU_R_123_16f"
	case 255:
		return "Layout_Undefined"
	default:
		return "invalid value"
	}
}
func DecodeTAS_11_Audio_Track_Layout_Enum(value []byte) (any, error) {
	field, _ := DecodeTUInt8(value[0:1])
	result := TAS_11_Audio_Track_Layout_Enum(field.(uint8))
	return result, nil
}
func EncodeTAS_11_Audio_Track_Layout_Enum(value TAS_11_Audio_Track_Layout_Enum) ([]byte, error) {
	result, _ := EncodeTUInt8(uint8(value))
	return result, nil
}

type TtemporalStrongReference TStrongReference
type TrationalStrongReference TStrongReference
type TtextualAnnotationStrongReferenceSet []TtextualAnnotationStrongReference

func EncodeTtextualAnnotationStrongReferenceSet(value TtextualAnnotationStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTtextualAnnotationStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TtextualAnnotationStrongReferenceSet = make([]TtextualAnnotationStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TtextualAnnotationStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TpartMetadataStrongReference TStrongReference
type TAS_11_Captions_Type_Enum uint8

func (mt TAS_11_Captions_Type_Enum) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TAS_11_Captions_Type_Enum) String() string {

	switch s {
	case 0:
		return "Captions_Hard_of_Hearing"
	case 1:
		return "Captions_Translation"
	default:
		return "invalid value"
	}
}
func DecodeTAS_11_Captions_Type_Enum(value []byte) (any, error) {
	field, _ := DecodeTUInt8(value[0:1])
	result := TAS_11_Captions_Type_Enum(field.(uint8))
	return result, nil
}
func EncodeTAS_11_Captions_Type_Enum(value TAS_11_Captions_Type_Enum) ([]byte, error) {
	result, _ := EncodeTUInt8(uint8(value))
	return result, nil
}

type TcoordinatesStrongReference TStrongReference
type TmediumStrongReference TStrongReference
type TControlPointStrongReferenceVector []TControlPointStrongReference

func EncodeTControlPointStrongReferenceVector(value TControlPointStrongReferenceVector) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTControlPointStrongReferenceVector(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TControlPointStrongReferenceVector = make([]TControlPointStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TControlPointStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TloudnessMetadataStrongReference TStrongReference
type TcoverageStrongReference TStrongReference
type TidentifierStrongReference TStrongReference
type TpackageInfoStrongReference TStrongReference
type TdateTypeStrongReferenceSet []TdateTypeStrongReference

func EncodeTdateTypeStrongReferenceSet(value TdateTypeStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTdateTypeStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TdateTypeStrongReferenceSet = make([]TdateTypeStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TdateTypeStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TtitleStrongReference TStrongReference
type TcodecStrongReference TStrongReference
type TalternativeTitleStrongReference TStrongReference
type TpublicationMediumStrongReference TStrongReference
type TdimensionStrongReference TStrongReference
type TaudioMXFLookupStrongReference TStrongReference
type TIdentificationStrongReferenceVector []TIdentificationStrongReference

func EncodeTIdentificationStrongReferenceVector(value TIdentificationStrongReferenceVector) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTIdentificationStrongReferenceVector(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TIdentificationStrongReferenceVector = make([]TIdentificationStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TIdentificationStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TentityStrongReference TStrongReference
type TlocatorStrongReference TStrongReference
type TpublicationChannelStrongReference TStrongReference
type TdateStrongReference TStrongReference
type TobjectTypeStrongReferenceSet []TobjectTypeStrongReference

func EncodeTobjectTypeStrongReferenceSet(value TobjectTypeStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTobjectTypeStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TobjectTypeStrongReferenceSet = make([]TobjectTypeStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TobjectTypeStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type ThashStrongReference TStrongReference
type TtypeStrongReference TStrongReference
type TsubjectStrongReference TStrongReference
type TcontainerFormatStrongReference TStrongReference
type TgenreStrongReferenceSet []TgenreStrongReference

func EncodeTgenreStrongReferenceSet(value TgenreStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTgenreStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TgenreStrongReferenceSet = make([]TgenreStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TgenreStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TdescriptionStrongReference TStrongReference
type TaudioBlockMatrixStrongReference TStrongReference
type TtextualAnnotationStrongReference TStrongReference
type TtargetAudienceStrongReferenceSet []TtargetAudienceStrongReference

func EncodeTtargetAudienceStrongReferenceSet(value TtargetAudienceStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTtargetAudienceStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TtargetAudienceStrongReferenceSet = make([]TtargetAudienceStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TtargetAudienceStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TtargetAudienceStrongReference TStrongReference
type TregionStrongReference TStrongReference
type TClassDefinitionStrongReferenceSet []TClassDefinitionStrongReference

func EncodeTClassDefinitionStrongReferenceSet(value TClassDefinitionStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTClassDefinitionStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TClassDefinitionStrongReferenceSet = make([]TClassDefinitionStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TClassDefinitionStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TpublicationServiceStrongReference TStrongReference
type TcustomRelationStrongReference TStrongReference
type TbasicRelationStrongReference TStrongReference
type TlocationStrongReference TStrongReference
type TregionStrongReferenceSet []TregionStrongReference

func EncodeTregionStrongReferenceSet(value TregionStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTregionStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TregionStrongReferenceSet = make([]TregionStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TregionStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TperiodOfTimeStrongReference TStrongReference
type TformatStrongReference TStrongReference
type TdepartmentStrongReference TStrongReference
type TcontactStrongReference TStrongReference
type TpartStrongReference TStrongReference
type TlocationStrongReferenceSet []TlocationStrongReference

func EncodeTlocationStrongReferenceSet(value TlocationStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTlocationStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TlocationStrongReferenceSet = make([]TlocationStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TlocationStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TpublicationEventStrongReference TStrongReference
type TdateTypeStrongReference TStrongReference
type TorganizationStrongReference TStrongReference
type TcompoundNameStrongReference TStrongReference
type TlanguageStrongReference TStrongReference
type TobjectTypeStrongReference TStrongReference
type TCodecDefinitionStrongReferenceSet []TCodecDefinitionStrongReference

func EncodeTCodecDefinitionStrongReferenceSet(value TCodecDefinitionStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTCodecDefinitionStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TCodecDefinitionStrongReferenceSet = make([]TCodecDefinitionStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TCodecDefinitionStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TcountryTypeStrongReference TStrongReference
type TroleStrongReference TStrongReference
type TrightsStrongReference TStrongReference
type TperiodOfTimeStrongReferenceSet []TperiodOfTimeStrongReference

func EncodeTperiodOfTimeStrongReferenceSet(value TperiodOfTimeStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTperiodOfTimeStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TperiodOfTimeStrongReferenceSet = make([]TperiodOfTimeStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TperiodOfTimeStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TdetailsStrongReference TStrongReference
type TgenreStrongReference TStrongReference
type TvideoFormatStrongReference TStrongReference
type TaudioFormatStrongReference TStrongReference
type TContainerDefinitionStrongReferenceSet []TContainerDefinitionStrongReference

func EncodeTContainerDefinitionStrongReferenceSet(value TContainerDefinitionStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTContainerDefinitionStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TContainerDefinitionStrongReferenceSet = make([]TContainerDefinitionStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TContainerDefinitionStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TratingStrongReference TStrongReference
type TbasicLinkStrongReference TStrongReference
type TcontactStrongReferenceSet []TcontactStrongReference

func EncodeTcontactStrongReferenceSet(value TcontactStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTcontactStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TcontactStrongReferenceSet = make([]TcontactStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TcontactStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TsigningFormatStrongReference TStrongReference
type TdataFormatStrongReference TStrongReference
type TtechnicalAttributeInt8StrongReference TStrongReference
type TpublicationEventStrongReferenceSet []TpublicationEventStrongReference

func EncodeTpublicationEventStrongReferenceSet(value TpublicationEventStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTpublicationEventStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TpublicationEventStrongReferenceSet = make([]TpublicationEventStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TpublicationEventStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TtechnicalAttributeInt16StrongReference TStrongReference
type TtechnicalAttributeInt32StrongReference TStrongReference
type TimageFormatStrongReference TStrongReference
type TDataDefinitionStrongReferenceSet []TDataDefinitionStrongReference

func EncodeTDataDefinitionStrongReferenceSet(value TDataDefinitionStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTDataDefinitionStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TDataDefinitionStrongReferenceSet = make([]TDataDefinitionStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TDataDefinitionStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TtechnicalAttributeStringStrongReference TStrongReference
type TorganizationStrongReferenceSet []TorganizationStrongReference

func EncodeTorganizationStrongReferenceSet(value TorganizationStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTorganizationStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TorganizationStrongReferenceSet = make([]TorganizationStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TorganizationStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TaudioBlockPositionStrongReferenceSet []TaudioBlockPositionStrongReference

func EncodeTaudioBlockPositionStrongReferenceSet(value TaudioBlockPositionStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTaudioBlockPositionStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TaudioBlockPositionStrongReferenceSet = make([]TaudioBlockPositionStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TaudioBlockPositionStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TProductReleaseType uint8

func (mt TProductReleaseType) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TProductReleaseType) String() string {

	switch s {
	case 0:
		return "VersionUnknown"
	case 1:
		return "VersionReleased"
	case 2:
		return "VersionDebug"
	case 3:
		return "VersionPatched"
	case 4:
		return "VersionBeta"
	case 5:
		return "VersionPrivateBuild"
	default:
		return "invalid value"
	}
}
func DecodeTProductReleaseType(value []byte) (any, error) {
	field, _ := DecodeTUInt8(value[0:1])
	result := TProductReleaseType(field.(uint8))
	return result, nil
}
func EncodeTProductReleaseType(value TProductReleaseType) ([]byte, error) {
	result, _ := EncodeTUInt8(uint8(value))
	return result, nil
}

type TOperationDefinitionStrongReferenceSet []TOperationDefinitionStrongReference

func EncodeTOperationDefinitionStrongReferenceSet(value TOperationDefinitionStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTOperationDefinitionStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TOperationDefinitionStrongReferenceSet = make([]TOperationDefinitionStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TOperationDefinitionStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TroleStrongReferenceSet []TroleStrongReference

func EncodeTroleStrongReferenceSet(value TroleStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTroleStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TroleStrongReferenceSet = make([]TroleStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TroleStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TPackageStrongReferenceSet []TPackageStrongReference

func EncodeTPackageStrongReferenceSet(value TPackageStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTPackageStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TPackageStrongReferenceSet = make([]TPackageStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TPackageStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TaudioPackFormatStrongReferenceSet []TaudioPackFormatStrongReference

func EncodeTaudioPackFormatStrongReferenceSet(value TaudioPackFormatStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTaudioPackFormatStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TaudioPackFormatStrongReferenceSet = make([]TaudioPackFormatStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TaudioPackFormatStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TEssenceDataStrongReferenceSet []TEssenceDataStrongReference

func EncodeTEssenceDataStrongReferenceSet(value TEssenceDataStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTEssenceDataStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TEssenceDataStrongReferenceSet = make([]TEssenceDataStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TEssenceDataStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TInterpolationDefinitionStrongReferenceSet []TInterpolationDefinitionStrongReference

func EncodeTInterpolationDefinitionStrongReferenceSet(value TInterpolationDefinitionStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTInterpolationDefinitionStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TInterpolationDefinitionStrongReferenceSet = make([]TInterpolationDefinitionStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TInterpolationDefinitionStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TbasicLinkStrongReferenceSet []TbasicLinkStrongReference

func EncodeTbasicLinkStrongReferenceSet(value TbasicLinkStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTbasicLinkStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TbasicLinkStrongReferenceSet = make([]TbasicLinkStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TbasicLinkStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TfilterSettingStrongReferenceSet []TfilterSettingStrongReference

func EncodeTfilterSettingStrongReferenceSet(value TfilterSettingStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTfilterSettingStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TfilterSettingStrongReferenceSet = make([]TfilterSettingStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TfilterSettingStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TaudioObjectInteractionStrongReferenceSet []TaudioObjectInteractionStrongReference

func EncodeTaudioObjectInteractionStrongReferenceSet(value TaudioObjectInteractionStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTaudioObjectInteractionStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TaudioObjectInteractionStrongReferenceSet = make([]TaudioObjectInteractionStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TaudioObjectInteractionStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TcompoundNameStrongReferenceSet []TcompoundNameStrongReference

func EncodeTcompoundNameStrongReferenceSet(value TcompoundNameStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTcompoundNameStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TcompoundNameStrongReferenceSet = make([]TcompoundNameStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TcompoundNameStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TgainInteractionRangeStrongReferenceSet []TgainInteractionRangeStrongReference

func EncodeTgainInteractionRangeStrongReferenceSet(value TgainInteractionRangeStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTgainInteractionRangeStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TgainInteractionRangeStrongReferenceSet = make([]TgainInteractionRangeStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TgainInteractionRangeStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TUKDPP_Audio_Description_Type_Enum uint8

func (mt TUKDPP_Audio_Description_Type_Enum) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TUKDPP_Audio_Description_Type_Enum) String() string {

	switch s {
	case 0:
		return "AD_Control_data_Narration"
	case 1:
		return "AD_Mix"
	default:
		return "invalid value"
	}
}
func DecodeTUKDPP_Audio_Description_Type_Enum(value []byte) (any, error) {
	field, _ := DecodeTUInt8(value[0:1])
	result := TUKDPP_Audio_Description_Type_Enum(field.(uint8))
	return result, nil
}
func EncodeTUKDPP_Audio_Description_Type_Enum(value TUKDPP_Audio_Description_Type_Enum) ([]byte, error) {
	result, _ := EncodeTUInt8(uint8(value))
	return result, nil
}

type TaudioChannelFormatStrongReferenceSet []TaudioChannelFormatStrongReference

func EncodeTaudioChannelFormatStrongReferenceSet(value TaudioChannelFormatStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTaudioChannelFormatStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TaudioChannelFormatStrongReferenceSet = make([]TaudioChannelFormatStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TaudioChannelFormatStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TdetailsStrongReferenceSet []TdetailsStrongReference

func EncodeTdetailsStrongReferenceSet(value TdetailsStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTdetailsStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TdetailsStrongReferenceSet = make([]TdetailsStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TdetailsStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TFadeType uint8

func (mt TFadeType) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TFadeType) String() string {

	switch s {
	case 0:
		return "FadeNone"
	case 1:
		return "FadeLinearAmp"
	case 2:
		return "FadeLinearPower"
	default:
		return "invalid value"
	}
}
func DecodeTFadeType(value []byte) (any, error) {
	field, _ := DecodeTUInt8(value[0:1])
	result := TFadeType(field.(uint8))
	return result, nil
}
func EncodeTFadeType(value TFadeType) ([]byte, error) {
	result, _ := EncodeTUInt8(uint8(value))
	return result, nil
}

type TtechnicalAttributeUInt64StrongReferenceSet []TtechnicalAttributeUInt64StrongReference

func EncodeTtechnicalAttributeUInt64StrongReferenceSet(value TtechnicalAttributeUInt64StrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTtechnicalAttributeUInt64StrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TtechnicalAttributeUInt64StrongReferenceSet = make([]TtechnicalAttributeUInt64StrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TtechnicalAttributeUInt64StrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TUKDPP_Sign_Language_Enum uint8

func (mt TUKDPP_Sign_Language_Enum) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TUKDPP_Sign_Language_Enum) String() string {

	switch s {
	case 0:
		return "Sign_Language_BSL_British_Sign_Language"
	case 1:
		return "Sign_Language_BSL_Makaton"
	default:
		return "invalid value"
	}
}
func DecodeTUKDPP_Sign_Language_Enum(value []byte) (any, error) {
	field, _ := DecodeTUInt8(value[0:1])
	result := TUKDPP_Sign_Language_Enum(field.(uint8))
	return result, nil
}
func EncodeTUKDPP_Sign_Language_Enum(value TUKDPP_Sign_Language_Enum) ([]byte, error) {
	result, _ := EncodeTUInt8(uint8(value))
	return result, nil
}

type TtechnicalAttributeFloatStrongReferenceSet []TtechnicalAttributeFloatStrongReference

func EncodeTtechnicalAttributeFloatStrongReferenceSet(value TtechnicalAttributeFloatStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTtechnicalAttributeFloatStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TtechnicalAttributeFloatStrongReferenceSet = make([]TtechnicalAttributeFloatStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TtechnicalAttributeFloatStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TUKDPP_Signing_Present_Enum uint8

func (mt TUKDPP_Signing_Present_Enum) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TUKDPP_Signing_Present_Enum) String() string {

	switch s {
	case 0:
		return "Signing_Yes"
	case 1:
		return "Signing_No"
	case 2:
		return "Signing_Signer_only"
	default:
		return "invalid value"
	}
}
func DecodeTUKDPP_Signing_Present_Enum(value []byte) (any, error) {
	field, _ := DecodeTUInt8(value[0:1])
	result := TUKDPP_Signing_Present_Enum(field.(uint8))
	return result, nil
}
func EncodeTUKDPP_Signing_Present_Enum(value TUKDPP_Signing_Present_Enum) ([]byte, error) {
	result, _ := EncodeTUInt8(uint8(value))
	return result, nil
}

type TStrongReferenceAS_07_DMS_Device TStrongReference
type TaudioBlockFormatStrongReferenceSet []TaudioBlockFormatStrongReference

func EncodeTaudioBlockFormatStrongReferenceSet(value TaudioBlockFormatStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTaudioBlockFormatStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TaudioBlockFormatStrongReferenceSet = make([]TaudioBlockFormatStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TaudioBlockFormatStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TcountryTypeStrongReferenceSet []TcountryTypeStrongReference

func EncodeTcountryTypeStrongReferenceSet(value TcountryTypeStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTcountryTypeStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TcountryTypeStrongReferenceSet = make([]TcountryTypeStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TcountryTypeStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TtechnicalAttributeRationalStrongReferenceSet []TtechnicalAttributeRationalStrongReference

func EncodeTtechnicalAttributeRationalStrongReferenceSet(value TtechnicalAttributeRationalStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTtechnicalAttributeRationalStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TtechnicalAttributeRationalStrongReferenceSet = make([]TtechnicalAttributeRationalStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TtechnicalAttributeRationalStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TUKDPP_3D_Type_Enum uint8

func (mt TUKDPP_3D_Type_Enum) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TUKDPP_3D_Type_Enum) String() string {

	switch s {
	case 0:
		return "ThreeD_Side_by_side"
	case 1:
		return "ThreeD_Dual"
	case 2:
		return "ThreeD_Left_eye_only"
	case 3:
		return "ThreeD_Right_eye_only"
	default:
		return "invalid value"
	}
}
func DecodeTUKDPP_3D_Type_Enum(value []byte) (any, error) {
	field, _ := DecodeTUInt8(value[0:1])
	result := TUKDPP_3D_Type_Enum(field.(uint8))
	return result, nil
}
func EncodeTUKDPP_3D_Type_Enum(value TUKDPP_3D_Type_Enum) ([]byte, error) {
	result, _ := EncodeTUInt8(uint8(value))
	return result, nil
}

type TUKDPP_PSE_Pass_Enum uint8

func (mt TUKDPP_PSE_Pass_Enum) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TUKDPP_PSE_Pass_Enum) String() string {

	switch s {
	case 0:
		return "PSE_Yes"
	case 1:
		return "PSE_No"
	case 2:
		return "PSE_Not_tested"
	default:
		return "invalid value"
	}
}
func DecodeTUKDPP_PSE_Pass_Enum(value []byte) (any, error) {
	field, _ := DecodeTUInt8(value[0:1])
	result := TUKDPP_PSE_Pass_Enum(field.(uint8))
	return result, nil
}
func EncodeTUKDPP_PSE_Pass_Enum(value TUKDPP_PSE_Pass_Enum) ([]byte, error) {
	result, _ := EncodeTUInt8(uint8(value))
	return result, nil
}

type TTapeFormatType uint8

func (mt TTapeFormatType) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TTapeFormatType) String() string {

	switch s {
	case 0:
		return "TapeFormatNull"
	case 1:
		return "BetacamFormat"
	case 2:
		return "BetacamSPFormat"
	case 3:
		return "VHSFormat"
	case 4:
		return "SVHSFormat"
	case 5:
		return "EightMillimeterFormat"
	case 6:
		return "Hi8Format"
	default:
		return "invalid value"
	}
}
func DecodeTTapeFormatType(value []byte) (any, error) {
	field, _ := DecodeTUInt8(value[0:1])
	result := TTapeFormatType(field.(uint8))
	return result, nil
}
func EncodeTTapeFormatType(value TTapeFormatType) ([]byte, error) {
	result, _ := EncodeTUInt8(uint8(value))
	return result, nil
}

type TStrongReferenceSetAS_07_DMS_Device []TStrongReferenceAS_07_DMS_Device

func EncodeTStrongReferenceSetAS_07_DMS_Device(value TStrongReferenceSetAS_07_DMS_Device) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTStrongReferenceSetAS_07_DMS_Device(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TStrongReferenceSetAS_07_DMS_Device = make([]TStrongReferenceAS_07_DMS_Device, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TStrongReferenceAS_07_DMS_Device(field.(TStrongReference)) // else
	}
	return result, nil
}

type TaudioStreamFormatStrongReferenceSet []TaudioStreamFormatStrongReference

func EncodeTaudioStreamFormatStrongReferenceSet(value TaudioStreamFormatStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTaudioStreamFormatStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TaudioStreamFormatStrongReferenceSet = make([]TaudioStreamFormatStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TaudioStreamFormatStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TaudioFormatStrongReferenceSet []TaudioFormatStrongReference

func EncodeTaudioFormatStrongReferenceSet(value TaudioFormatStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTaudioFormatStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TaudioFormatStrongReferenceSet = make([]TaudioFormatStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TaudioFormatStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TStrongReferenceAS_07_DMS_Identifier TStrongReference
type TtechnicalAttributeAnyURIStrongReferenceSet []TtechnicalAttributeAnyURIStrongReference

func EncodeTtechnicalAttributeAnyURIStrongReferenceSet(value TtechnicalAttributeAnyURIStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTtechnicalAttributeAnyURIStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TtechnicalAttributeAnyURIStrongReferenceSet = make([]TtechnicalAttributeAnyURIStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TtechnicalAttributeAnyURIStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TUKDPP_Audio_Loudness_Standard_Enum uint8

func (mt TUKDPP_Audio_Loudness_Standard_Enum) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TUKDPP_Audio_Loudness_Standard_Enum) String() string {

	switch s {
	case 0:
		return "Loudness_None"
	case 1:
		return "Loudness_EBU_R_128"
	default:
		return "invalid value"
	}
}
func DecodeTUKDPP_Audio_Loudness_Standard_Enum(value []byte) (any, error) {
	field, _ := DecodeTUInt8(value[0:1])
	result := TUKDPP_Audio_Loudness_Standard_Enum(field.(uint8))
	return result, nil
}
func EncodeTUKDPP_Audio_Loudness_Standard_Enum(value TUKDPP_Audio_Loudness_Standard_Enum) ([]byte, error) {
	result, _ := EncodeTUInt8(uint8(value))
	return result, nil
}

type TVideoSignalType uint8

func (mt TVideoSignalType) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TVideoSignalType) String() string {

	switch s {
	case 0:
		return "VideoSignalNull"
	case 1:
		return "NTSCSignal"
	case 2:
		return "PALSignal"
	default:
		return "invalid value"
	}
}
func DecodeTVideoSignalType(value []byte) (any, error) {
	field, _ := DecodeTUInt8(value[0:1])
	result := TVideoSignalType(field.(uint8))
	return result, nil
}
func EncodeTVideoSignalType(value TVideoSignalType) ([]byte, error) {
	result, _ := EncodeTUInt8(uint8(value))
	return result, nil
}

type TvideoFormatStrongReferenceSet []TvideoFormatStrongReference

func EncodeTvideoFormatStrongReferenceSet(value TvideoFormatStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTvideoFormatStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TvideoFormatStrongReferenceSet = make([]TvideoFormatStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TvideoFormatStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TpositionInteractionRangeStrongReferenceSet []TpositionInteractionRangeStrongReference

func EncodeTpositionInteractionRangeStrongReferenceSet(value TpositionInteractionRangeStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTpositionInteractionRangeStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TpositionInteractionRangeStrongReferenceSet = make([]TpositionInteractionRangeStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TpositionInteractionRangeStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TtechnicalAttributeBooleanStrongReferenceSet []TtechnicalAttributeBooleanStrongReference

func EncodeTtechnicalAttributeBooleanStrongReferenceSet(value TtechnicalAttributeBooleanStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTtechnicalAttributeBooleanStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TtechnicalAttributeBooleanStrongReferenceSet = make([]TtechnicalAttributeBooleanStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TtechnicalAttributeBooleanStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TStrongReferenceSetAS_07_DMS_Identifier []TStrongReferenceAS_07_DMS_Identifier

func EncodeTStrongReferenceSetAS_07_DMS_Identifier(value TStrongReferenceSetAS_07_DMS_Identifier) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTStrongReferenceSetAS_07_DMS_Identifier(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TStrongReferenceSetAS_07_DMS_Identifier = make([]TStrongReferenceAS_07_DMS_Identifier, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TStrongReferenceAS_07_DMS_Identifier(field.(TStrongReference)) // else
	}
	return result, nil
}

type TaudioTrackFormatStrongReferenceSet []TaudioTrackFormatStrongReference

func EncodeTaudioTrackFormatStrongReferenceSet(value TaudioTrackFormatStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTaudioTrackFormatStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TaudioTrackFormatStrongReferenceSet = make([]TaudioTrackFormatStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TaudioTrackFormatStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TTapeCaseType uint8

func (mt TTapeCaseType) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TTapeCaseType) String() string {

	switch s {
	case 0:
		return "TapeCaseNull"
	case 1:
		return "ThreeFourthInchVideoTape"
	case 2:
		return "VHSVideoTape"
	case 3:
		return "EightMillimeterVideoTape"
	case 4:
		return "BetacamVideoTape"
	case 5:
		return "CompactCassette"
	case 6:
		return "DATCartridge"
	case 7:
		return "NagraAudioTape"
	default:
		return "invalid value"
	}
}
func DecodeTTapeCaseType(value []byte) (any, error) {
	field, _ := DecodeTUInt8(value[0:1])
	result := TTapeCaseType(field.(uint8))
	return result, nil
}
func EncodeTTapeCaseType(value TTapeCaseType) ([]byte, error) {
	result, _ := EncodeTUInt8(uint8(value))
	return result, nil
}

type TAPP_TimecodeTypeEnum uint16

func (mt TAPP_TimecodeTypeEnum) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TAPP_TimecodeTypeEnum) String() string {

	switch s {
	case 1:
		return "TC_VITC"
	case 2:
		return "TC_LTC"
	case 3:
		return "TC_Both"
	default:
		return "invalid value"
	}
}
func DecodeTAPP_TimecodeTypeEnum(value []byte) (any, error) {
	field, _ := DecodeTUInt16(value[0:2])
	result := TAPP_TimecodeTypeEnum(field.(uint16))
	return result, nil
}
func EncodeTAPP_TimecodeTypeEnum(value TAPP_TimecodeTypeEnum) ([]byte, error) {
	result, _ := EncodeTUInt16(uint16(value))
	return result, nil
}

type TaudioFormatExtendedStrongReferenceSet []TaudioFormatExtendedStrongReference

func EncodeTaudioFormatExtendedStrongReferenceSet(value TaudioFormatExtendedStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTaudioFormatExtendedStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TaudioFormatExtendedStrongReferenceSet = make([]TaudioFormatExtendedStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TaudioFormatExtendedStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TimageFormatStrongReferenceSet []TimageFormatStrongReference

func EncodeTimageFormatStrongReferenceSet(value TimageFormatStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTimageFormatStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TimageFormatStrongReferenceSet = make([]TimageFormatStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TimageFormatStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TColorSitingType uint8

func (mt TColorSitingType) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TColorSitingType) String() string {

	switch s {
	case 0:
		return "CoSiting"
	case 1:
		return "HorizontalMidpoint"
	case 2:
		return "ThreeTap"
	case 3:
		return "Quincunx"
	case 4:
		return "Rec601"
	case 5:
		return "LineAlternating"
	case 6:
		return "VerticalMidpoint"
	case 255:
		return "UnknownSiting"
	default:
		return "invalid value"
	}
}
func DecodeTColorSitingType(value []byte) (any, error) {
	field, _ := DecodeTUInt8(value[0:1])
	result := TColorSitingType(field.(uint8))
	return result, nil
}
func EncodeTColorSitingType(value TColorSitingType) ([]byte, error) {
	result, _ := EncodeTUInt8(uint8(value))
	return result, nil
}

type TeventStrongReferenceSet []TeventStrongReference

func EncodeTeventStrongReferenceSet(value TeventStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTeventStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TeventStrongReferenceSet = make([]TeventStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TeventStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TaudioTrackUIDStrongReferenceSet []TaudioTrackUIDStrongReference

func EncodeTaudioTrackUIDStrongReferenceSet(value TaudioTrackUIDStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTaudioTrackUIDStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TaudioTrackUIDStrongReferenceSet = make([]TaudioTrackUIDStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TaudioTrackUIDStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TaspectRatioStrongReferenceSet []TaspectRatioStrongReference

func EncodeTaspectRatioStrongReferenceSet(value TaspectRatioStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTaspectRatioStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TaspectRatioStrongReferenceSet = make([]TaspectRatioStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TaspectRatioStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TsigningFormatStrongReferenceSet []TsigningFormatStrongReference

func EncodeTsigningFormatStrongReferenceSet(value TsigningFormatStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTsigningFormatStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TsigningFormatStrongReferenceSet = make([]TsigningFormatStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TsigningFormatStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TdataFormatStrongReferenceSet []TdataFormatStrongReference

func EncodeTdataFormatStrongReferenceSet(value TdataFormatStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTdataFormatStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TdataFormatStrongReferenceSet = make([]TdataFormatStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TdataFormatStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TtechnicalAttributeInt64StrongReferenceSet []TtechnicalAttributeInt64StrongReference

func EncodeTtechnicalAttributeInt64StrongReferenceSet(value TtechnicalAttributeInt64StrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTtechnicalAttributeInt64StrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TtechnicalAttributeInt64StrongReferenceSet = make([]TtechnicalAttributeInt64StrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TtechnicalAttributeInt64StrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TEditHintType uint8

func (mt TEditHintType) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TEditHintType) String() string {

	switch s {
	case 0:
		return "NoEditHint"
	case 1:
		return "Proportional"
	case 2:
		return "RelativeLeft"
	case 3:
		return "RelativeRight"
	case 4:
		return "RelativeFixed"
	default:
		return "invalid value"
	}
}
func DecodeTEditHintType(value []byte) (any, error) {
	field, _ := DecodeTUInt8(value[0:1])
	result := TEditHintType(field.(uint8))
	return result, nil
}
func EncodeTEditHintType(value TEditHintType) ([]byte, error) {
	result, _ := EncodeTUInt8(uint8(value))
	return result, nil
}

type TheightStrongReferenceSet []TheightStrongReference

func EncodeTheightStrongReferenceSet(value TheightStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTheightStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TheightStrongReferenceSet = make([]TheightStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TheightStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TtechnicalAttributeStringStrongReferenceSet []TtechnicalAttributeStringStrongReference

func EncodeTtechnicalAttributeStringStrongReferenceSet(value TtechnicalAttributeStringStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTtechnicalAttributeStringStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TtechnicalAttributeStringStrongReferenceSet = make([]TtechnicalAttributeStringStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TtechnicalAttributeStringStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TaudienceStrongReferenceSet []TaudienceStrongReference

func EncodeTaudienceStrongReferenceSet(value TaudienceStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTaudienceStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TaudienceStrongReferenceSet = make([]TaudienceStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TaudienceStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TtimeStrongReferenceSet []TtimeStrongReference

func EncodeTtimeStrongReferenceSet(value TtimeStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTtimeStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TtimeStrongReferenceSet = make([]TtimeStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TtimeStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TaffiliationStrongReferenceSet []TaffiliationStrongReference

func EncodeTaffiliationStrongReferenceSet(value TaffiliationStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTaffiliationStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TaffiliationStrongReferenceSet = make([]TaffiliationStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TaffiliationStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TtimecodeFormatStrongReferenceSet []TtimecodeFormatStrongReference

func EncodeTtimecodeFormatStrongReferenceSet(value TtimecodeFormatStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTtimecodeFormatStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TtimecodeFormatStrongReferenceSet = make([]TtimecodeFormatStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TtimecodeFormatStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TtechnicalAttributeUInt8StrongReferenceSet []TtechnicalAttributeUInt8StrongReference

func EncodeTtechnicalAttributeUInt8StrongReferenceSet(value TtechnicalAttributeUInt8StrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTtechnicalAttributeUInt8StrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TtechnicalAttributeUInt8StrongReferenceSet = make([]TtechnicalAttributeUInt8StrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TtechnicalAttributeUInt8StrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TfilterStrongReferenceSet []TfilterStrongReference

func EncodeTfilterStrongReferenceSet(value TfilterStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTfilterStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TfilterStrongReferenceSet = make([]TfilterStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TfilterStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TtechnicalAttributeUInt32StrongReferenceSet []TtechnicalAttributeUInt32StrongReference

func EncodeTtechnicalAttributeUInt32StrongReferenceSet(value TtechnicalAttributeUInt32StrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTtechnicalAttributeUInt32StrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TtechnicalAttributeUInt32StrongReferenceSet = make([]TtechnicalAttributeUInt32StrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TtechnicalAttributeUInt32StrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TwidthStrongReferenceSet []TwidthStrongReference

func EncodeTwidthStrongReferenceSet(value TwidthStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTwidthStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TwidthStrongReferenceSet = make([]TwidthStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TwidthStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TmediumStrongReferenceSet []TmediumStrongReference

func EncodeTmediumStrongReferenceSet(value TmediumStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTmediumStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TmediumStrongReferenceSet = make([]TmediumStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TmediumStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TtechnicalAttributeInt32StrongReferenceSet []TtechnicalAttributeInt32StrongReference

func EncodeTtechnicalAttributeInt32StrongReferenceSet(value TtechnicalAttributeInt32StrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTtechnicalAttributeInt32StrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TtechnicalAttributeInt32StrongReferenceSet = make([]TtechnicalAttributeInt32StrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TtechnicalAttributeInt32StrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TtechnicalAttributeInt16StrongReferenceSet []TtechnicalAttributeInt16StrongReference

func EncodeTtechnicalAttributeInt16StrongReferenceSet(value TtechnicalAttributeInt16StrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTtechnicalAttributeInt16StrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TtechnicalAttributeInt16StrongReferenceSet = make([]TtechnicalAttributeInt16StrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TtechnicalAttributeInt16StrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TaudioBlockMatrixCoefficientStrongReferenceSet []TaudioBlockMatrixCoefficientStrongReference

func EncodeTaudioBlockMatrixCoefficientStrongReferenceSet(value TaudioBlockMatrixCoefficientStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTaudioBlockMatrixCoefficientStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TaudioBlockMatrixCoefficientStrongReferenceSet = make([]TaudioBlockMatrixCoefficientStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TaudioBlockMatrixCoefficientStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TawardStrongReferenceSet []TawardStrongReference

func EncodeTawardStrongReferenceSet(value TawardStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTawardStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TawardStrongReferenceSet = make([]TawardStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TawardStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TPulldownDirectionType uint8

func (mt TPulldownDirectionType) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TPulldownDirectionType) String() string {

	switch s {
	case 0:
		return "TapeToFilmSpeed"
	case 1:
		return "FilmToTapeSpeed"
	default:
		return "invalid value"
	}
}
func DecodeTPulldownDirectionType(value []byte) (any, error) {
	field, _ := DecodeTUInt8(value[0:1])
	result := TPulldownDirectionType(field.(uint8))
	return result, nil
}
func EncodeTPulldownDirectionType(value TPulldownDirectionType) ([]byte, error) {
	result, _ := EncodeTUInt8(uint8(value))
	return result, nil
}

type TaudioContentStrongReferenceSet []TaudioContentStrongReference

func EncodeTaudioContentStrongReferenceSet(value TaudioContentStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTaudioContentStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TaudioContentStrongReferenceSet = make([]TaudioContentStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TaudioContentStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TaudioProgrammeStrongReferenceSet []TaudioProgrammeStrongReference

func EncodeTaudioProgrammeStrongReferenceSet(value TaudioProgrammeStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTaudioProgrammeStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TaudioProgrammeStrongReferenceSet = make([]TaudioProgrammeStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TaudioProgrammeStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TancillaryDataStrongReferenceSet []TancillaryDataStrongReference

func EncodeTancillaryDataStrongReferenceSet(value TancillaryDataStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTancillaryDataStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TancillaryDataStrongReferenceSet = make([]TancillaryDataStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TancillaryDataStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TsubtitlingStrongReferenceSet []TsubtitlingStrongReference

func EncodeTsubtitlingStrongReferenceSet(value TsubtitlingStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTsubtitlingStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TsubtitlingStrongReferenceSet = make([]TsubtitlingStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TsubtitlingStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TaudioObjectStrongReferenceSet []TaudioObjectStrongReference

func EncodeTaudioObjectStrongReferenceSet(value TaudioObjectStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTaudioObjectStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TaudioObjectStrongReferenceSet = make([]TaudioObjectStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TaudioObjectStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TtrackStrongReferenceSet []TtrackStrongReference

func EncodeTtrackStrongReferenceSet(value TtrackStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTtrackStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TtrackStrongReferenceSet = make([]TtrackStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TtrackStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TtechnicalAttributeInt8StrongReferenceSet []TtechnicalAttributeInt8StrongReference

func EncodeTtechnicalAttributeInt8StrongReferenceSet(value TtechnicalAttributeInt8StrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTtechnicalAttributeInt8StrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TtechnicalAttributeInt8StrongReferenceSet = make([]TtechnicalAttributeInt8StrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TtechnicalAttributeInt8StrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TPulldownKindType uint8

func (mt TPulldownKindType) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TPulldownKindType) String() string {

	switch s {
	case 0:
		return "TwoThreePD"
	case 1:
		return "PALPD"
	case 2:
		return "OneToOneNTSC"
	case 3:
		return "OneToOnePAL"
	case 4:
		return "VideoTapNTSC"
	case 5:
		return "OneToOneHDSixty"
	case 6:
		return "TwentyFourToSixtyPD"
	case 7:
		return "TwoToOnePD"
	default:
		return "invalid value"
	}
}
func DecodeTPulldownKindType(value []byte) (any, error) {
	field, _ := DecodeTUInt8(value[0:1])
	result := TPulldownKindType(field.(uint8))
	return result, nil
}
func EncodeTPulldownKindType(value TPulldownKindType) ([]byte, error) {
	result, _ := EncodeTUInt8(uint8(value))
	return result, nil
}

type TTCSource uint8

func (mt TTCSource) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TTCSource) String() string {

	switch s {
	case 0:
		return "TimecodeLTC"
	case 1:
		return "TimecodeVITC"
	default:
		return "invalid value"
	}
}
func DecodeTTCSource(value []byte) (any, error) {
	field, _ := DecodeTUInt8(value[0:1])
	result := TTCSource(field.(uint8))
	return result, nil
}
func EncodeTTCSource(value TTCSource) ([]byte, error) {
	result, _ := EncodeTUInt8(uint8(value))
	return result, nil
}

type TEdgeType uint8

func (mt TEdgeType) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TEdgeType) String() string {

	switch s {
	case 0:
		return "EtNull"
	case 1:
		return "EtKeycode"
	case 2:
		return "EtEdgenum4"
	case 3:
		return "EtEdgenum5"
	case 8:
		return "EtHeaderSize"
	default:
		return "invalid value"
	}
}
func DecodeTEdgeType(value []byte) (any, error) {
	field, _ := DecodeTUInt8(value[0:1])
	result := TEdgeType(field.(uint8))
	return result, nil
}
func EncodeTEdgeType(value TEdgeType) ([]byte, error) {
	result, _ := EncodeTUInt8(uint8(value))
	return result, nil
}

type TLayoutType uint8

func (mt TLayoutType) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TLayoutType) String() string {

	switch s {
	case 0:
		return "FullFrame"
	case 1:
		return "SeparateFields"
	case 2:
		return "OneField"
	case 3:
		return "MixedFields"
	case 4:
		return "SegmentedFrame"
	default:
		return "invalid value"
	}
}
func DecodeTLayoutType(value []byte) (any, error) {
	field, _ := DecodeTUInt8(value[0:1])
	result := TLayoutType(field.(uint8))
	return result, nil
}
func EncodeTLayoutType(value TLayoutType) ([]byte, error) {
	result, _ := EncodeTUInt8(uint8(value))
	return result, nil
}

type TmetadataFormatStrongReferenceSet []TmetadataFormatStrongReference

func EncodeTmetadataFormatStrongReferenceSet(value TmetadataFormatStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTmetadataFormatStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TmetadataFormatStrongReferenceSet = make([]TmetadataFormatStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TmetadataFormatStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TaudioBlockZoneStrongReferenceSet []TaudioBlockZoneStrongReference

func EncodeTaudioBlockZoneStrongReferenceSet(value TaudioBlockZoneStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTaudioBlockZoneStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TaudioBlockZoneStrongReferenceSet = make([]TaudioBlockZoneStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TaudioBlockZoneStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TcaptioningStrongReferenceSet []TcaptioningStrongReference

func EncodeTcaptioningStrongReferenceSet(value TcaptioningStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTcaptioningStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TcaptioningStrongReferenceSet = make([]TcaptioningStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TcaptioningStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TtechnicalAttributeUInt16StrongReferenceSet []TtechnicalAttributeUInt16StrongReference

func EncodeTtechnicalAttributeUInt16StrongReferenceSet(value TtechnicalAttributeUInt16StrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTtechnicalAttributeUInt16StrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TtechnicalAttributeUInt16StrongReferenceSet = make([]TtechnicalAttributeUInt16StrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TtechnicalAttributeUInt16StrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TIDRefStrongReferenceSet []TIDRefStrongReference

func EncodeTIDRefStrongReferenceSet(value TIDRefStrongReferenceSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTStrongReference(TStrongReference(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTIDRefStrongReferenceSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TIDRefStrongReferenceSet = make([]TIDRefStrongReference, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTStrongReference(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = TIDRefStrongReference(field.(TStrongReference)) // else
	}
	return result, nil
}

type TAUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 TUInt8Array8
}

func DecodeTAUID(value []byte) (any, error) {
	result := TAUID{}
	var field any
	field, _ = DecodeTUInt32(value[0:4])
	result.Data1 = field.(uint32) // else
	field, _ = DecodeTUInt16(value[4:6])
	result.Data2 = field.(uint16) // else
	field, _ = DecodeTUInt16(value[6:8])
	result.Data3 = field.(uint16) // else
	field, _ = DecodeTUInt8Array8(value[8:16])
	result.Data4 = field.(TUInt8Array8) // else
	return result, nil
}
func EncodeTAUID(value TAUID) ([]byte, error) {
	result := []byte{}
	field0, _ := EncodeTUInt32(uint32(value.Data1)) // else
	result = append(result, field0...)
	field1, _ := EncodeTUInt16(uint16(value.Data2)) // else
	result = append(result, field1...)
	field2, _ := EncodeTUInt16(uint16(value.Data3)) // else
	result = append(result, field2...)
	field3, _ := EncodeTUInt8Array8(TUInt8Array8(value.Data4)) // else
	result = append(result, field3...)
	return result, nil
}

type TLocalTagEntry struct {
	LocalTag TLocalTagType
	UID      TAUID
}

func DecodeTLocalTagEntry(value []byte) (any, error) {
	result := TLocalTagEntry{}
	var field any
	field, _ = DecodeTLocalTagType(value[0:2])
	result.LocalTag = field.(TLocalTagType) // else
	field, _ = DecodeTAUID(value[2:18])
	result.UID = field.(TAUID) // else
	return result, nil
}
func EncodeTLocalTagEntry(value TLocalTagEntry) ([]byte, error) {
	result := []byte{}
	field0, _ := EncodeTLocalTagType(TLocalTagType(value.LocalTag)) // else
	result = append(result, field0...)
	field1, _ := EncodeTAUID(TAUID(value.UID)) // else
	result = append(result, field1...)
	return result, nil
}

type TTimeStamp struct {
	Date TDateStruct
	Time TTimeStruct
}

func DecodeTTimeStamp(value []byte) (any, error) {
	result := TTimeStamp{}
	var field any
	field, _ = DecodeTDateStruct(value[0:4])
	result.Date = field.(TDateStruct) // else
	field, _ = DecodeTTimeStruct(value[4:8])
	result.Time = field.(TTimeStruct) // else
	return result, nil
}
func EncodeTTimeStamp(value TTimeStamp) ([]byte, error) {
	result := []byte{}
	field0, _ := EncodeTDateStruct(TDateStruct(value.Date)) // else
	result = append(result, field0...)
	field1, _ := EncodeTTimeStruct(TTimeStruct(value.Time)) // else
	result = append(result, field1...)
	return result, nil
}

type TPackageIDType struct {
	SMPTELabel   TUInt8Array12
	Length       uint8
	InstanceHigh uint8
	InstanceMid  uint8
	InstanceLow  uint8
	Material     TAUID
}

func DecodeTPackageIDType(value []byte) (any, error) {
	result := TPackageIDType{}
	var field any
	field, _ = DecodeTUInt8Array12(value[0:12])
	result.SMPTELabel = field.(TUInt8Array12) // else
	field, _ = DecodeTUInt8(value[12:13])
	result.Length = field.(uint8) // else
	field, _ = DecodeTUInt8(value[13:14])
	result.InstanceHigh = field.(uint8) // else
	field, _ = DecodeTUInt8(value[14:15])
	result.InstanceMid = field.(uint8) // else
	field, _ = DecodeTUInt8(value[15:16])
	result.InstanceLow = field.(uint8) // else
	field, _ = DecodeTAUID(value[16:32])
	result.Material = field.(TAUID) // else
	return result, nil
}
func EncodeTPackageIDType(value TPackageIDType) ([]byte, error) {
	result := []byte{}
	field0, _ := EncodeTUInt8Array12(TUInt8Array12(value.SMPTELabel)) // else
	result = append(result, field0...)
	field1, _ := EncodeTUInt8(uint8(value.Length)) // else
	result = append(result, field1...)
	field2, _ := EncodeTUInt8(uint8(value.InstanceHigh)) // else
	result = append(result, field2...)
	field3, _ := EncodeTUInt8(uint8(value.InstanceMid)) // else
	result = append(result, field3...)
	field4, _ := EncodeTUInt8(uint8(value.InstanceLow)) // else
	result = append(result, field4...)
	field5, _ := EncodeTAUID(TAUID(value.Material)) // else
	result = append(result, field5...)
	return result, nil
}

type TPluginCategoryType TAUID

func DecodeTPluginCategoryType(value []byte) (any, error) {
	field, _ := DecodeTAUID(value[0:16])
	result := TPluginCategoryType(field.(TAUID))
	return result, nil
}
func EncodeTPluginCategoryType(value TPluginCategoryType) ([]byte, error) {
	result, _ := EncodeTAUID(TAUID(value))
	return result, nil
}

type TOperationCategoryType TAUID

func DecodeTOperationCategoryType(value []byte) (any, error) {
	field, _ := DecodeTAUID(value[0:16])
	result := TOperationCategoryType(field.(TAUID))
	return result, nil
}
func EncodeTOperationCategoryType(value TOperationCategoryType) ([]byte, error) {
	result, _ := EncodeTAUID(TAUID(value))
	return result, nil
}

type TTransferCharacteristicType TAUID

func DecodeTTransferCharacteristicType(value []byte) (any, error) {
	field, _ := DecodeTAUID(value[0:16])
	result := TTransferCharacteristicType(field.(TAUID))
	return result, nil
}
func EncodeTTransferCharacteristicType(value TTransferCharacteristicType) ([]byte, error) {
	result, _ := EncodeTAUID(TAUID(value))
	return result, nil
}

type TS327M TDataValue

func DecodeTS327M(value []byte) (any, error) {
	field, _ := DecodeTDataValue(value)
	result := TS327M(field.(TDataValue))
	return result, nil
}
func EncodeTS327M(value TS327M) ([]byte, error) {
	result, _ := EncodeTDataValue(TDataValue(value))
	return result, nil
}

type TIndexEntry TDataValue

func DecodeTIndexEntry(value []byte) (any, error) {
	field, _ := DecodeTDataValue(value)
	result := TIndexEntry(field.(TDataValue))
	return result, nil
}
func EncodeTIndexEntry(value TIndexEntry) ([]byte, error) {
	result, _ := EncodeTDataValue(TDataValue(value))
	return result, nil
}

type TST2109PayloadSeries TDataValue

func DecodeTST2109PayloadSeries(value []byte) (any, error) {
	field, _ := DecodeTDataValue(value)
	result := TST2109PayloadSeries(field.(TDataValue))
	return result, nil
}
func EncodeTST2109PayloadSeries(value TST2109PayloadSeries) ([]byte, error) {
	result, _ := EncodeTDataValue(TDataValue(value))
	return result, nil
}

type TBiM TDataValue

func DecodeTBiM(value []byte) (any, error) {
	field, _ := DecodeTDataValue(value)
	result := TBiM(field.(TDataValue))
	return result, nil
}
func EncodeTBiM(value TBiM) ([]byte, error) {
	result, _ := EncodeTDataValue(TDataValue(value))
	return result, nil
}

type TS312M TDataValue

func DecodeTS312M(value []byte) (any, error) {
	field, _ := DecodeTDataValue(value)
	result := TS312M(field.(TDataValue))
	return result, nil
}
func EncodeTS312M(value TS312M) ([]byte, error) {
	result, _ := EncodeTDataValue(TDataValue(value))
	return result, nil
}

type TDeltaEntry TDataValue

func DecodeTDeltaEntry(value []byte) (any, error) {
	field, _ := DecodeTDataValue(value)
	result := TDeltaEntry(field.(TDataValue))
	return result, nil
}
func EncodeTDeltaEntry(value TDeltaEntry) ([]byte, error) {
	result, _ := EncodeTDataValue(TDataValue(value))
	return result, nil
}

type TRGBALayout [8]TRGBAComponent

func EncodeTRGBALayout(value TRGBALayout) ([]byte, error) {
	result := []byte{}
	for _, val := range value {
		field, _ := EncodeTRGBAComponent(TRGBAComponent(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTRGBALayout(value []byte) (any, error) {
	var result TRGBALayout = [8]TRGBAComponent{}
	for i := 0; i < 8; i++ {
		field, _ := DecodeTRGBAComponent(value[i*2 : i*2+2])
		result[i] = field.(TRGBAComponent) // else
	}
	return result, nil
}

type TLocalTagEntryBatch []TLocalTagEntry

func EncodeTLocalTagEntryBatch(value TLocalTagEntryBatch) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	binary.BigEndian.PutUint32(result[4:8], 18)
	for _, val := range value {
		field, _ := EncodeTLocalTagEntry(TLocalTagEntry(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTLocalTagEntryBatch(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TLocalTagEntryBatch = make([]TLocalTagEntry, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTLocalTagEntry(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = field.(TLocalTagEntry) // else
	}
	return result, nil
}

type TDeltaEntryArray []TDeltaEntry

func EncodeTDeltaEntryArray(value TDeltaEntryArray) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTDeltaEntry(TDeltaEntry(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTDeltaEntryArray(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TDeltaEntryArray = make([]TDeltaEntry, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTDeltaEntry(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = field.(TDeltaEntry) // else
	}
	return result, nil
}

type TAUIDArray []TAUID

func EncodeTAUIDArray(value TAUIDArray) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	binary.BigEndian.PutUint32(result[4:8], 16)
	for _, val := range value {
		field, _ := EncodeTAUID(TAUID(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTAUIDArray(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TAUIDArray = make([]TAUID, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTAUID(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = field.(TAUID) // else
	}
	return result, nil
}

type TIndexEntryArray []TIndexEntry

func EncodeTIndexEntryArray(value TIndexEntryArray) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	if len(value) != 0 { // else leave as the default of 0
		binary.BigEndian.PutUint32(result[4:8], uint32(len(value[0])))
	}
	for _, val := range value {
		field, _ := EncodeTIndexEntry(TIndexEntry(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTIndexEntryArray(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TIndexEntryArray = make([]TIndexEntry, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTIndexEntry(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = field.(TIndexEntry) // else
	}
	return result, nil
}

type TAUIDSet []TAUID

func EncodeTAUIDSet(value TAUIDSet) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	binary.BigEndian.PutUint32(result[4:8], 16)
	for _, val := range value {
		field, _ := EncodeTAUID(TAUID(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTAUIDSet(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TAUIDSet = make([]TAUID, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTAUID(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = field.(TAUID) // else
	}
	return result, nil
}

type TUMID TPackageIDType

func DecodeTUMID(value []byte) (any, error) {
	field, _ := DecodeTPackageIDType(value[0:32])
	result := TUMID(field.(TPackageIDType))
	return result, nil
}
func EncodeTUMID(value TUMID) ([]byte, error) {
	result, _ := EncodeTPackageIDType(TPackageIDType(value))
	return result, nil
}

type TViewingEnvironmentType TAUID

func DecodeTViewingEnvironmentType(value []byte) (any, error) {
	field, _ := DecodeTAUID(value[0:16])
	result := TViewingEnvironmentType(field.(TAUID))
	return result, nil
}
func EncodeTViewingEnvironmentType(value TViewingEnvironmentType) ([]byte, error) {
	result, _ := EncodeTAUID(TAUID(value))
	return result, nil
}

type TUsageType TAUID

func DecodeTUsageType(value []byte) (any, error) {
	field, _ := DecodeTAUID(value[0:16])
	result := TUsageType(field.(TAUID))
	return result, nil
}
func EncodeTUsageType(value TUsageType) ([]byte, error) {
	result, _ := EncodeTAUID(TAUID(value))
	return result, nil
}

type TCodingEquationsType TAUID

func DecodeTCodingEquationsType(value []byte) (any, error) {
	field, _ := DecodeTAUID(value[0:16])
	result := TCodingEquationsType(field.(TAUID))
	return result, nil
}
func EncodeTCodingEquationsType(value TCodingEquationsType) ([]byte, error) {
	result, _ := EncodeTAUID(TAUID(value))
	return result, nil
}

type TAutoExposureModeType TAUID

func DecodeTAutoExposureModeType(value []byte) (any, error) {
	field, _ := DecodeTAUID(value[0:16])
	result := TAutoExposureModeType(field.(TAUID))
	return result, nil
}
func EncodeTAutoExposureModeType(value TAutoExposureModeType) ([]byte, error) {
	result, _ := EncodeTAUID(TAUID(value))
	return result, nil
}

type TColorPrimariesType TAUID

func DecodeTColorPrimariesType(value []byte) (any, error) {
	field, _ := DecodeTAUID(value[0:16])
	result := TColorPrimariesType(field.(TAUID))
	return result, nil
}
func EncodeTColorPrimariesType(value TColorPrimariesType) ([]byte, error) {
	result, _ := EncodeTAUID(TAUID(value))
	return result, nil
}

type TUMIDArray []TUMID

func EncodeTUMIDArray(value TUMIDArray) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result[0:4], uint32(len(value)))
	binary.BigEndian.PutUint32(result[4:8], 32)
	for _, val := range value {
		field, _ := EncodeTUMID(TUMID(val))
		result = append(result, field...)
	}
	return result, nil
}
func DecodeTUMIDArray(value []byte) (any, error) {
	count, _ := DecodeTUInt32(value[0:4])
	arrayCount := int(count.(uint32))
	fieldLen, _ := DecodeTUInt32(value[4:8])
	fieldSize := int(fieldLen.(uint32))
	var result TUMIDArray = make([]TUMID, arrayCount)
	for i := 0; i < arrayCount; i++ {
		field, _ := DecodeTUMID(value[8+i*fieldSize : 8+i*fieldSize+fieldSize])
		result[i] = field.(TUMID) // else
	}
	return result, nil
}

type TUTF16StringArray string

func DecodeTUTF16StringArray(value []byte) (any, error) {
	field, _ := DecodeTUTF16String(value)
	result := TUTF16StringArray(field.(string))

	return result, nil
}

type TAS_07_DMS_IdentifierRoleCode TUTF8String

func (mt TAS_07_DMS_IdentifierRoleCode) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TAS_07_DMS_IdentifierRoleCode) String() string {

	switch string(s) {
	case "Main":
		return "Main"
	case "Additional":
		return "Additional"
	case "GSP":
		return "GSP"
	default:
		return "invalid value"
	}
}
func DecodeTAS_07_DMS_IdentifierRoleCode(value []byte) (any, error) {
	field, _ := DecodeTUTF8String(value)
	result := TAS_07_DMS_IdentifierRoleCode(field.(TUTF8String))
	return result, nil
}
func EncodeTAS_07_DMS_IdentifierRoleCode(value TAS_07_DMS_IdentifierRoleCode) ([]byte, error) {
	result, _ := EncodeTUTF8String(TUTF8String(value))
	return result, nil
}

type TISO639 TISO7

func DecodeTISO639(value []byte) (any, error) {
	field, _ := DecodeTISO7(value)
	result := TISO639(field.(TISO7))
	return result, nil
}
func EncodeTISO639(value TISO639) ([]byte, error) {
	result, _ := EncodeTISO7(TISO7(value))
	return result, nil
}

type TPublishingMediumCode TISO7

func DecodeTPublishingMediumCode(value []byte) (any, error) {
	field, _ := DecodeTISO7(value)
	result := TPublishingMediumCode(field.(TISO7))
	return result, nil
}
func EncodeTPublishingMediumCode(value TPublishingMediumCode) ([]byte, error) {
	result, _ := EncodeTISO7(TISO7(value))
	return result, nil
}

type TGammaCode TISO7

func DecodeTGammaCode(value []byte) (any, error) {
	field, _ := DecodeTISO7(value)
	result := TGammaCode(field.(TISO7))
	return result, nil
}
func EncodeTGammaCode(value TGammaCode) ([]byte, error) {
	result, _ := EncodeTISO7(TISO7(value))
	return result, nil
}

type TObjectTypeCode TISO7

func DecodeTObjectTypeCode(value []byte) (any, error) {
	field, _ := DecodeTISO7(value)
	result := TObjectTypeCode(field.(TISO7))
	return result, nil
}
func EncodeTObjectTypeCode(value TObjectTypeCode) ([]byte, error) {
	result, _ := EncodeTISO7(TISO7(value))
	return result, nil
}

type TRFC2152 TISO7

func DecodeTRFC2152(value []byte) (any, error) {
	field, _ := DecodeTISO7(value)
	result := TRFC2152(field.(TISO7))
	return result, nil
}
func EncodeTRFC2152(value TRFC2152) ([]byte, error) {
	result, _ := EncodeTISO7(TISO7(value))
	return result, nil
}

type TContractLineCode TISO7

func DecodeTContractLineCode(value []byte) (any, error) {
	field, _ := DecodeTISO7(value)
	result := TContractLineCode(field.(TISO7))
	return result, nil
}
func EncodeTContractLineCode(value TContractLineCode) ([]byte, error) {
	result, _ := EncodeTISO7(TISO7(value))
	return result, nil
}

type TCanonicalEIDRServiceIdentifierType TCanonicalDOINameType

func DecodeTCanonicalEIDRServiceIdentifierType(value []byte) (any, error) {
	field, _ := DecodeTCanonicalDOINameType(value)
	result := TCanonicalEIDRServiceIdentifierType(field.(TCanonicalDOINameType))
	return result, nil
}
func EncodeTCanonicalEIDRServiceIdentifierType(value TCanonicalEIDRServiceIdentifierType) ([]byte, error) {
	result, _ := EncodeTCanonicalDOINameType(TCanonicalDOINameType(value))
	return result, nil
}

type TISO3166_Country TISO7

func DecodeTISO3166_Country(value []byte) (any, error) {
	field, _ := DecodeTISO7(value)
	result := TISO3166_Country(field.(TISO7))
	return result, nil
}
func EncodeTISO3166_Country(value TISO3166_Country) ([]byte, error) {
	result, _ := EncodeTISO7(TISO7(value))
	return result, nil
}

type TCanonicalFullAdIDIdentifierType TUTF16String

func DecodeTCanonicalFullAdIDIdentifierType(value []byte) (any, error) {
	field, _ := DecodeTUTF16String(value)
	result := TCanonicalFullAdIDIdentifierType(field.(TUTF16String))
	return result, nil
}
func EncodeTCanonicalFullAdIDIdentifierType(value TCanonicalFullAdIDIdentifierType) ([]byte, error) {
	result, _ := EncodeTUTF16String(TUTF16String(value))
	return result, nil
}

type TAS_07_DMS_IdentifierTypeCode TUTF8String

func (mt TAS_07_DMS_IdentifierTypeCode) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TAS_07_DMS_IdentifierTypeCode) String() string {

	switch string(s) {
	case "UUID":
		return "UUID"
	case "UMID":
		return "UMID"
	case "UL":
		return "UL"
	case "Other":
		return "Other"
	default:
		return "invalid value"
	}
}
func DecodeTAS_07_DMS_IdentifierTypeCode(value []byte) (any, error) {
	field, _ := DecodeTUTF8String(value)
	result := TAS_07_DMS_IdentifierTypeCode(field.(TUTF8String))
	return result, nil
}
func EncodeTAS_07_DMS_IdentifierTypeCode(value TAS_07_DMS_IdentifierTypeCode) ([]byte, error) {
	result, _ := EncodeTUTF8String(TUTF8String(value))
	return result, nil
}

type TSamplingStructureCode TISO7

func DecodeTSamplingStructureCode(value []byte) (any, error) {
	field, _ := DecodeTISO7(value)
	result := TSamplingStructureCode(field.(TISO7))
	return result, nil
}
func EncodeTSamplingStructureCode(value TSamplingStructureCode) ([]byte, error) {
	result, _ := EncodeTISO7(TISO7(value))
	return result, nil
}

type TUTCmilliseconds TISO7

func DecodeTUTCmilliseconds(value []byte) (any, error) {
	field, _ := DecodeTISO7(value)
	result := TUTCmilliseconds(field.(TISO7))
	return result, nil
}
func EncodeTUTCmilliseconds(value TUTCmilliseconds) ([]byte, error) {
	result, _ := EncodeTISO7(TISO7(value))
	return result, nil
}

type TCanonicalDOINameType TUTF16String

func DecodeTCanonicalDOINameType(value []byte) (any, error) {
	field, _ := DecodeTUTF16String(value)
	result := TCanonicalDOINameType(field.(TUTF16String))
	return result, nil
}
func EncodeTCanonicalDOINameType(value TCanonicalDOINameType) ([]byte, error) {
	result, _ := EncodeTUTF16String(TUTF16String(value))
	return result, nil
}

type TContractTypeCode TISO7

func DecodeTContractTypeCode(value []byte) (any, error) {
	field, _ := DecodeTISO7(value)
	result := TContractTypeCode(field.(TISO7))
	return result, nil
}
func EncodeTContractTypeCode(value TContractTypeCode) ([]byte, error) {
	result, _ := EncodeTISO7(TISO7(value))
	return result, nil
}

type TAS_07_DMS_DataDescriptionCode TUTF8String

func (mt TAS_07_DMS_DataDescriptionCode) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (s TAS_07_DMS_DataDescriptionCode) String() string {

	switch string(s) {
	case "Caption":
		return "Caption"
	case "TimedText":
		return "TimedText"
	case "Graphic":
		return "Graphic"
	case "RelatedDocument":
		return "RelatedDocument"
	case "SupplementaryMetadata":
		return "SupplementaryMetadata"
	case "AssociatedMaterial":
		return "AssociatedMaterial"
	case "Trailer":
		return "Trailer"
	case "QC":
		return "QC"
	case "Other":
		return "Other"
	default:
		return "invalid value"
	}
}
func DecodeTAS_07_DMS_DataDescriptionCode(value []byte) (any, error) {
	field, _ := DecodeTUTF8String(value)
	result := TAS_07_DMS_DataDescriptionCode(field.(TUTF8String))
	return result, nil
}
func EncodeTAS_07_DMS_DataDescriptionCode(value TAS_07_DMS_DataDescriptionCode) ([]byte, error) {
	result, _ := EncodeTUTF8String(TUTF8String(value))
	return result, nil
}

type TISO639_Ext TISO7

func DecodeTISO639_Ext(value []byte) (any, error) {
	field, _ := DecodeTISO7(value)
	result := TISO639_Ext(field.(TISO7))
	return result, nil
}
func EncodeTISO639_Ext(value TISO639_Ext) ([]byte, error) {
	result, _ := EncodeTISO7(TISO7(value))
	return result, nil
}

type TISO3166_Region TISO7

func DecodeTISO3166_Region(value []byte) (any, error) {
	field, _ := DecodeTISO7(value)
	result := TISO3166_Region(field.(TISO7))
	return result, nil
}
func EncodeTISO3166_Region(value TISO3166_Region) ([]byte, error) {
	result, _ := EncodeTISO7(TISO7(value))
	return result, nil
}

type TJobFunctionCode TISO7

func DecodeTJobFunctionCode(value []byte) (any, error) {
	field, _ := DecodeTISO7(value)
	result := TJobFunctionCode(field.(TISO7))
	return result, nil
}
func EncodeTJobFunctionCode(value TJobFunctionCode) ([]byte, error) {
	result, _ := EncodeTISO7(TISO7(value))
	return result, nil
}

type TISO_639_2_Language_Code TUTF16String

func DecodeTISO_639_2_Language_Code(value []byte) (any, error) {
	field, _ := DecodeTUTF16String(value)
	result := TISO_639_2_Language_Code(field.(TUTF16String))
	return result, nil
}
func EncodeTISO_639_2_Language_Code(value TISO_639_2_Language_Code) ([]byte, error) {
	result, _ := EncodeTUTF16String(TUTF16String(value))
	return result, nil
}

type TCanonicalEIDRIdentifierType TCanonicalDOINameType

func DecodeTCanonicalEIDRIdentifierType(value []byte) (any, error) {
	field, _ := DecodeTCanonicalDOINameType(value)
	result := TCanonicalEIDRIdentifierType(field.(TCanonicalDOINameType))
	return result, nil
}
func EncodeTCanonicalEIDRIdentifierType(value TCanonicalEIDRIdentifierType) ([]byte, error) {
	result, _ := EncodeTCanonicalDOINameType(TCanonicalDOINameType(value))
	return result, nil
}

func EncodeTUTF16StringArray(value TUTF16StringArray) ([]byte, error) {
	parts := utf16.Encode([]rune(string(value)))
	result, _ := EncodeTUInt16Array([]uint16(parts))
	return result, nil
}
