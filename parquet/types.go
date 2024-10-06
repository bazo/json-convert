package parquet

import (
	"json-convert/utils"
	"time"
)

type Line struct {
	Int      int
	Int8     int8
	Int16    int16
	Int32    int32
	Int64    int64
	String   string
	Bool     bool
	SString  []string
	SInt     []int
	SInt8    []int8
	SInt16   []int16
	SInt32   []int32
	SInt64   []int64
	SFloat32 []float32
	SFloat64 []float64
	SBool    []bool
	Created  time.Time
}

type Another struct {
	Image string `parquet:"image"`
}

type ParquetLine struct {
	Int      int       `parquet:"int"`
	Int8     int8      `parquet:"int8"`
	Int16    int16     `parquet:"int16"`
	Int32    int32     `parquet:"int32"`
	Int64    int64     `parquet:"int64"`
	String   string    `parquet:"string"`
	Bool     bool      `parquet:"bool"`
	SString  []string  `parquet:"sstring"`
	SInt     []int     `parquet:"sint"`
	SInt8    []int8    `parquet:"sint8"`
	SInt16   []int16   `parquet:"sint16"`
	SInt32   []int32   `parquet:"sint32"`
	SInt64   []int64   `parquet:"sint64"`
	SFloat32 []float32 `parquet:"sfloat32"`
	SFloat64 []float64 `parquet:"sfloat64"`
	SBool    []bool    `parquet:"sbool"`
	Created  int64     `parquet:"created"`
}

func (l *Line) ToParquet() *ParquetLine {
	return &ParquetLine{
		Int:      l.Int,
		Int8:     l.Int8,
		Int16:    l.Int16,
		Int32:    l.Int32,
		Int64:    l.Int64,
		String:   l.String,
		Bool:     l.Bool,
		SString:  l.SString,
		SInt:     l.SInt,
		SInt8:    l.SInt8,
		SInt16:   l.SInt16,
		SInt32:   l.SInt32,
		SInt64:   l.SInt64,
		SFloat32: l.SFloat32,
		SFloat64: l.SFloat64,
		SBool:    l.SBool,
		Created:  utils.TimeToMillis(l.Created),
	}
}
