package types

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
	Embedded Embedded
	Created  time.Time
}

type Embedded struct {
	Number        int64   `parquet:"name=number, type=INT64"`
	Height        int64   `parquet:"name=height, type=INT64"`
	AnotherStruct Another `parquet:"name=another_struct"`
}

type Another struct {
	Image string `parquet:"name=image, type=BYTE_ARRAY, encoding=PLAIN_DICTIONARY"`
}

type ParquetLine struct {
	Int      int       `parquet:"name=int, type=INT32"`
	Int8     int8      `parquet:"name=int8, type=INT32"`
	Int16    int16     `parquet:"name=int16, type=INT32"`
	Int32    int32     `parquet:"name=int32, type=INT32"`
	Int64    int64     `parquet:"name=int64, type=INT64"`
	String   string    `parquet:"name=string, type=BYTE_ARRAY, encoding=PLAIN_DICTIONARY"`
	Bool     bool      `parquet:"name=bool, type=BOOLEAN"`
	SString  []string  `parquet:"name=sstring, type=BYTE_ARRAY, repetitiontype=REPEATED, encoding=PLAIN_DICTIONARY"`
	SInt     []int     `parquet:"name=sint, type=INT32, repetitiontype=REPEATED"`
	SInt8    []int8    `parquet:"name=sint8, type=INT32, repetitiontype=REPEATED"`
	SInt16   []int16   `parquet:"name=sint16, type=INT32, repetitiontype=REPEATED"`
	SInt32   []int32   `parquet:"name=sint32, type=INT32, repetitiontype=REPEATED"`
	SInt64   []int64   `parquet:"name=sint64, type=INT64, repetitiontype=REPEATED"`
	SFloat32 []float32 `parquet:"name=sfloat32, type=FLOAT, repetitiontype=REPEATED"`
	SFloat64 []float64 `parquet:"name=sfloat64, type=DOUBLE, repetitiontype=REPEATED"`
	SBool    []bool    `parquet:"name=sbool, type=BOOLEAN, repetitiontype=REPEATED"`
	Embedded Embedded  `parquet:"name=embedded"`
	Created  int64     `parquet:"name=created, type=INT64, convertedtype=TIMESTAMP_MILLIS"`
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
		Embedded: l.Embedded,
		Created:  utils.TimeToMillis(l.Created),
	}
}
