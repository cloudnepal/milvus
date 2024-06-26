// Code generated by go generate; DO NOT EDIT
// This file is generated by go generate

package column

import (
	"fmt"

	"github.com/cockroachdb/errors"
	"github.com/milvus-io/milvus-proto/go-api/v2/schemapb"
	"github.com/milvus-io/milvus/client/v2/entity"
)

// ColumnBoolArray generated columns type for Bool
type ColumnBoolArray struct {
	ColumnBase
	name   string
	values [][]bool
}

// Name returns column name
func (c *ColumnBoolArray) Name() string {
	return c.name
}

// Type returns column entity.FieldType
func (c *ColumnBoolArray) Type() entity.FieldType {
	return entity.FieldTypeArray
}

// Len returns column values length
func (c *ColumnBoolArray) Len() int {
	return len(c.values)
}

func (c *ColumnBoolArray) Slice(start, end int) Column {
	l := c.Len()
	if start > l {
		start = l
	}
	if end == -1 || end > l {
		end = l
	}
	if end == -1 || end > l {
		end = l
	}
	return &ColumnBoolArray{
		ColumnBase: c.ColumnBase,
		name:       c.name,
		values:     c.values[start:end],
	}
}

// Get returns value at index as interface{}.
func (c *ColumnBoolArray) Get(idx int) (interface{}, error) {
	var r []bool // use default value
	if idx < 0 || idx >= c.Len() {
		return r, errors.New("index out of range")
	}
	return c.values[idx], nil
}

// FieldData return column data mapped to schemapb.FieldData
func (c *ColumnBoolArray) FieldData() *schemapb.FieldData {
	fd := &schemapb.FieldData{
		Type:      schemapb.DataType_Array,
		FieldName: c.name,
	}

	data := make([]*schemapb.ScalarField, 0, c.Len())
	for _, arr := range c.values {
		converted := make([]bool, 0, c.Len())
		for i := 0; i < len(arr); i++ {
			converted = append(converted, bool(arr[i]))
		}
		data = append(data, &schemapb.ScalarField{
			Data: &schemapb.ScalarField_BoolData{
				BoolData: &schemapb.BoolArray{
					Data: converted,
				},
			},
		})
	}
	fd.Field = &schemapb.FieldData_Scalars{
		Scalars: &schemapb.ScalarField{
			Data: &schemapb.ScalarField_ArrayData{
				ArrayData: &schemapb.ArrayArray{
					Data:        data,
					ElementType: schemapb.DataType_Bool,
				},
			},
		},
	}
	return fd
}

// ValueByIdx returns value of the provided index
// error occurs when index out of range
func (c *ColumnBoolArray) ValueByIdx(idx int) ([]bool, error) {
	var r []bool // use default value
	if idx < 0 || idx >= c.Len() {
		return r, errors.New("index out of range")
	}
	return c.values[idx], nil
}

// AppendValue append value into column
func (c *ColumnBoolArray) AppendValue(i interface{}) error {
	v, ok := i.([]bool)
	if !ok {
		return fmt.Errorf("invalid type, expected []bool, got %T", i)
	}
	c.values = append(c.values, v)

	return nil
}

// Data returns column data
func (c *ColumnBoolArray) Data() [][]bool {
	return c.values
}

// NewColumnBool auto generated constructor
func NewColumnBoolArray(name string, values [][]bool) *ColumnBoolArray {
	return &ColumnBoolArray{
		name:   name,
		values: values,
	}
}

// ColumnInt8Array generated columns type for Int8
type ColumnInt8Array struct {
	ColumnBase
	name   string
	values [][]int8
}

// Name returns column name
func (c *ColumnInt8Array) Name() string {
	return c.name
}

// Type returns column entity.FieldType
func (c *ColumnInt8Array) Type() entity.FieldType {
	return entity.FieldTypeArray
}

// Len returns column values length
func (c *ColumnInt8Array) Len() int {
	return len(c.values)
}

func (c *ColumnInt8Array) Slice(start, end int) Column {
	l := c.Len()
	if start > l {
		start = l
	}
	if end == -1 || end > l {
		end = l
	}
	return &ColumnInt8Array{
		ColumnBase: c.ColumnBase,
		name:       c.name,
		values:     c.values[start:end],
	}
}

// Get returns value at index as interface{}.
func (c *ColumnInt8Array) Get(idx int) (interface{}, error) {
	var r []int8 // use default value
	if idx < 0 || idx >= c.Len() {
		return r, errors.New("index out of range")
	}
	return c.values[idx], nil
}

// FieldData return column data mapped to schemapb.FieldData
func (c *ColumnInt8Array) FieldData() *schemapb.FieldData {
	fd := &schemapb.FieldData{
		Type:      schemapb.DataType_Array,
		FieldName: c.name,
	}

	data := make([]*schemapb.ScalarField, 0, c.Len())
	for _, arr := range c.values {
		converted := make([]int32, 0, c.Len())
		for i := 0; i < len(arr); i++ {
			converted = append(converted, int32(arr[i]))
		}
		data = append(data, &schemapb.ScalarField{
			Data: &schemapb.ScalarField_IntData{
				IntData: &schemapb.IntArray{
					Data: converted,
				},
			},
		})
	}
	fd.Field = &schemapb.FieldData_Scalars{
		Scalars: &schemapb.ScalarField{
			Data: &schemapb.ScalarField_ArrayData{
				ArrayData: &schemapb.ArrayArray{
					Data:        data,
					ElementType: schemapb.DataType_Int8,
				},
			},
		},
	}
	return fd
}

// ValueByIdx returns value of the provided index
// error occurs when index out of range
func (c *ColumnInt8Array) ValueByIdx(idx int) ([]int8, error) {
	var r []int8 // use default value
	if idx < 0 || idx >= c.Len() {
		return r, errors.New("index out of range")
	}
	return c.values[idx], nil
}

// AppendValue append value into column
func (c *ColumnInt8Array) AppendValue(i interface{}) error {
	v, ok := i.([]int8)
	if !ok {
		return fmt.Errorf("invalid type, expected []int8, got %T", i)
	}
	c.values = append(c.values, v)

	return nil
}

// Data returns column data
func (c *ColumnInt8Array) Data() [][]int8 {
	return c.values
}

// NewColumnInt8 auto generated constructor
func NewColumnInt8Array(name string, values [][]int8) *ColumnInt8Array {
	return &ColumnInt8Array{
		name:   name,
		values: values,
	}
}

// ColumnInt16Array generated columns type for Int16
type ColumnInt16Array struct {
	ColumnBase
	name   string
	values [][]int16
}

// Name returns column name
func (c *ColumnInt16Array) Name() string {
	return c.name
}

// Type returns column entity.FieldType
func (c *ColumnInt16Array) Type() entity.FieldType {
	return entity.FieldTypeArray
}

// Len returns column values length
func (c *ColumnInt16Array) Len() int {
	return len(c.values)
}

func (c *ColumnInt16Array) Slice(start, end int) Column {
	l := c.Len()
	if start > l {
		start = l
	}
	if end == -1 || end > l {
		end = l
	}
	return &ColumnInt16Array{
		ColumnBase: c.ColumnBase,
		name:       c.name,
		values:     c.values[start:end],
	}
}

// Get returns value at index as interface{}.
func (c *ColumnInt16Array) Get(idx int) (interface{}, error) {
	var r []int16 // use default value
	if idx < 0 || idx >= c.Len() {
		return r, errors.New("index out of range")
	}
	return c.values[idx], nil
}

// FieldData return column data mapped to schemapb.FieldData
func (c *ColumnInt16Array) FieldData() *schemapb.FieldData {
	fd := &schemapb.FieldData{
		Type:      schemapb.DataType_Array,
		FieldName: c.name,
	}

	data := make([]*schemapb.ScalarField, 0, c.Len())
	for _, arr := range c.values {
		converted := make([]int32, 0, c.Len())
		for i := 0; i < len(arr); i++ {
			converted = append(converted, int32(arr[i]))
		}
		data = append(data, &schemapb.ScalarField{
			Data: &schemapb.ScalarField_IntData{
				IntData: &schemapb.IntArray{
					Data: converted,
				},
			},
		})
	}
	fd.Field = &schemapb.FieldData_Scalars{
		Scalars: &schemapb.ScalarField{
			Data: &schemapb.ScalarField_ArrayData{
				ArrayData: &schemapb.ArrayArray{
					Data:        data,
					ElementType: schemapb.DataType_Int16,
				},
			},
		},
	}
	return fd
}

// ValueByIdx returns value of the provided index
// error occurs when index out of range
func (c *ColumnInt16Array) ValueByIdx(idx int) ([]int16, error) {
	var r []int16 // use default value
	if idx < 0 || idx >= c.Len() {
		return r, errors.New("index out of range")
	}
	return c.values[idx], nil
}

// AppendValue append value into column
func (c *ColumnInt16Array) AppendValue(i interface{}) error {
	v, ok := i.([]int16)
	if !ok {
		return fmt.Errorf("invalid type, expected []int16, got %T", i)
	}
	c.values = append(c.values, v)

	return nil
}

// Data returns column data
func (c *ColumnInt16Array) Data() [][]int16 {
	return c.values
}

// NewColumnInt16 auto generated constructor
func NewColumnInt16Array(name string, values [][]int16) *ColumnInt16Array {
	return &ColumnInt16Array{
		name:   name,
		values: values,
	}
}

// ColumnInt32Array generated columns type for Int32
type ColumnInt32Array struct {
	ColumnBase
	name   string
	values [][]int32
}

// Name returns column name
func (c *ColumnInt32Array) Name() string {
	return c.name
}

// Type returns column entity.FieldType
func (c *ColumnInt32Array) Type() entity.FieldType {
	return entity.FieldTypeArray
}

// Len returns column values length
func (c *ColumnInt32Array) Len() int {
	return len(c.values)
}

func (c *ColumnInt32Array) Slice(start, end int) Column {
	l := c.Len()
	if start > l {
		start = l
	}
	if end == -1 || end > l {
		end = l
	}
	return &ColumnInt32Array{
		ColumnBase: c.ColumnBase,
		name:       c.name,
		values:     c.values[start:end],
	}
}

// Get returns value at index as interface{}.
func (c *ColumnInt32Array) Get(idx int) (interface{}, error) {
	var r []int32 // use default value
	if idx < 0 || idx >= c.Len() {
		return r, errors.New("index out of range")
	}
	return c.values[idx], nil
}

// FieldData return column data mapped to schemapb.FieldData
func (c *ColumnInt32Array) FieldData() *schemapb.FieldData {
	fd := &schemapb.FieldData{
		Type:      schemapb.DataType_Array,
		FieldName: c.name,
	}

	data := make([]*schemapb.ScalarField, 0, c.Len())
	for _, arr := range c.values {
		converted := make([]int32, 0, c.Len())
		for i := 0; i < len(arr); i++ {
			converted = append(converted, int32(arr[i]))
		}
		data = append(data, &schemapb.ScalarField{
			Data: &schemapb.ScalarField_IntData{
				IntData: &schemapb.IntArray{
					Data: converted,
				},
			},
		})
	}
	fd.Field = &schemapb.FieldData_Scalars{
		Scalars: &schemapb.ScalarField{
			Data: &schemapb.ScalarField_ArrayData{
				ArrayData: &schemapb.ArrayArray{
					Data:        data,
					ElementType: schemapb.DataType_Int32,
				},
			},
		},
	}
	return fd
}

// ValueByIdx returns value of the provided index
// error occurs when index out of range
func (c *ColumnInt32Array) ValueByIdx(idx int) ([]int32, error) {
	var r []int32 // use default value
	if idx < 0 || idx >= c.Len() {
		return r, errors.New("index out of range")
	}
	return c.values[idx], nil
}

// AppendValue append value into column
func (c *ColumnInt32Array) AppendValue(i interface{}) error {
	v, ok := i.([]int32)
	if !ok {
		return fmt.Errorf("invalid type, expected []int32, got %T", i)
	}
	c.values = append(c.values, v)

	return nil
}

// Data returns column data
func (c *ColumnInt32Array) Data() [][]int32 {
	return c.values
}

// NewColumnInt32 auto generated constructor
func NewColumnInt32Array(name string, values [][]int32) *ColumnInt32Array {
	return &ColumnInt32Array{
		name:   name,
		values: values,
	}
}

// ColumnInt64Array generated columns type for Int64
type ColumnInt64Array struct {
	ColumnBase
	name   string
	values [][]int64
}

// Name returns column name
func (c *ColumnInt64Array) Name() string {
	return c.name
}

// Type returns column entity.FieldType
func (c *ColumnInt64Array) Type() entity.FieldType {
	return entity.FieldTypeArray
}

// Len returns column values length
func (c *ColumnInt64Array) Len() int {
	return len(c.values)
}

func (c *ColumnInt64Array) Slice(start, end int) Column {
	l := c.Len()
	if start > l {
		start = l
	}
	if end == -1 || end > l {
		end = l
	}
	return &ColumnInt64Array{
		ColumnBase: c.ColumnBase,
		name:       c.name,
		values:     c.values[start:end],
	}
}

// Get returns value at index as interface{}.
func (c *ColumnInt64Array) Get(idx int) (interface{}, error) {
	var r []int64 // use default value
	if idx < 0 || idx >= c.Len() {
		return r, errors.New("index out of range")
	}
	return c.values[idx], nil
}

// FieldData return column data mapped to schemapb.FieldData
func (c *ColumnInt64Array) FieldData() *schemapb.FieldData {
	fd := &schemapb.FieldData{
		Type:      schemapb.DataType_Array,
		FieldName: c.name,
	}

	data := make([]*schemapb.ScalarField, 0, c.Len())
	for _, arr := range c.values {
		converted := make([]int64, 0, c.Len())
		for i := 0; i < len(arr); i++ {
			converted = append(converted, int64(arr[i]))
		}
		data = append(data, &schemapb.ScalarField{
			Data: &schemapb.ScalarField_LongData{
				LongData: &schemapb.LongArray{
					Data: converted,
				},
			},
		})
	}
	fd.Field = &schemapb.FieldData_Scalars{
		Scalars: &schemapb.ScalarField{
			Data: &schemapb.ScalarField_ArrayData{
				ArrayData: &schemapb.ArrayArray{
					Data:        data,
					ElementType: schemapb.DataType_Int64,
				},
			},
		},
	}
	return fd
}

// ValueByIdx returns value of the provided index
// error occurs when index out of range
func (c *ColumnInt64Array) ValueByIdx(idx int) ([]int64, error) {
	var r []int64 // use default value
	if idx < 0 || idx >= c.Len() {
		return r, errors.New("index out of range")
	}
	return c.values[idx], nil
}

// AppendValue append value into column
func (c *ColumnInt64Array) AppendValue(i interface{}) error {
	v, ok := i.([]int64)
	if !ok {
		return fmt.Errorf("invalid type, expected []int64, got %T", i)
	}
	c.values = append(c.values, v)

	return nil
}

// Data returns column data
func (c *ColumnInt64Array) Data() [][]int64 {
	return c.values
}

// NewColumnInt64 auto generated constructor
func NewColumnInt64Array(name string, values [][]int64) *ColumnInt64Array {
	return &ColumnInt64Array{
		name:   name,
		values: values,
	}
}

// ColumnFloatArray generated columns type for Float
type ColumnFloatArray struct {
	ColumnBase
	name   string
	values [][]float32
}

// Name returns column name
func (c *ColumnFloatArray) Name() string {
	return c.name
}

// Type returns column entity.FieldType
func (c *ColumnFloatArray) Type() entity.FieldType {
	return entity.FieldTypeArray
}

// Len returns column values length
func (c *ColumnFloatArray) Len() int {
	return len(c.values)
}

func (c *ColumnFloatArray) Slice(start, end int) Column {
	l := c.Len()
	if start > l {
		start = l
	}
	if end == -1 || end > l {
		end = l
	}
	return &ColumnFloatArray{
		ColumnBase: c.ColumnBase,
		name:       c.name,
		values:     c.values[start:end],
	}
}

// Get returns value at index as interface{}.
func (c *ColumnFloatArray) Get(idx int) (interface{}, error) {
	var r []float32 // use default value
	if idx < 0 || idx >= c.Len() {
		return r, errors.New("index out of range")
	}
	return c.values[idx], nil
}

// FieldData return column data mapped to schemapb.FieldData
func (c *ColumnFloatArray) FieldData() *schemapb.FieldData {
	fd := &schemapb.FieldData{
		Type:      schemapb.DataType_Array,
		FieldName: c.name,
	}

	data := make([]*schemapb.ScalarField, 0, c.Len())
	for _, arr := range c.values {
		converted := make([]float32, 0, c.Len())
		for i := 0; i < len(arr); i++ {
			converted = append(converted, float32(arr[i]))
		}
		data = append(data, &schemapb.ScalarField{
			Data: &schemapb.ScalarField_FloatData{
				FloatData: &schemapb.FloatArray{
					Data: converted,
				},
			},
		})
	}
	fd.Field = &schemapb.FieldData_Scalars{
		Scalars: &schemapb.ScalarField{
			Data: &schemapb.ScalarField_ArrayData{
				ArrayData: &schemapb.ArrayArray{
					Data:        data,
					ElementType: schemapb.DataType_Float,
				},
			},
		},
	}
	return fd
}

// ValueByIdx returns value of the provided index
// error occurs when index out of range
func (c *ColumnFloatArray) ValueByIdx(idx int) ([]float32, error) {
	var r []float32 // use default value
	if idx < 0 || idx >= c.Len() {
		return r, errors.New("index out of range")
	}
	return c.values[idx], nil
}

// AppendValue append value into column
func (c *ColumnFloatArray) AppendValue(i interface{}) error {
	v, ok := i.([]float32)
	if !ok {
		return fmt.Errorf("invalid type, expected []float32, got %T", i)
	}
	c.values = append(c.values, v)

	return nil
}

// Data returns column data
func (c *ColumnFloatArray) Data() [][]float32 {
	return c.values
}

// NewColumnFloat auto generated constructor
func NewColumnFloatArray(name string, values [][]float32) *ColumnFloatArray {
	return &ColumnFloatArray{
		name:   name,
		values: values,
	}
}

// ColumnDoubleArray generated columns type for Double
type ColumnDoubleArray struct {
	ColumnBase
	name   string
	values [][]float64
}

// Name returns column name
func (c *ColumnDoubleArray) Name() string {
	return c.name
}

// Type returns column entity.FieldType
func (c *ColumnDoubleArray) Type() entity.FieldType {
	return entity.FieldTypeArray
}

// Len returns column values length
func (c *ColumnDoubleArray) Len() int {
	return len(c.values)
}

func (c *ColumnDoubleArray) Slice(start, end int) Column {
	l := c.Len()
	if start > l {
		start = l
	}
	if end == -1 || end > l {
		end = l
	}
	return &ColumnDoubleArray{
		ColumnBase: c.ColumnBase,
		name:       c.name,
		values:     c.values[start:end],
	}
}

// Get returns value at index as interface{}.
func (c *ColumnDoubleArray) Get(idx int) (interface{}, error) {
	var r []float64 // use default value
	if idx < 0 || idx >= c.Len() {
		return r, errors.New("index out of range")
	}
	return c.values[idx], nil
}

// FieldData return column data mapped to schemapb.FieldData
func (c *ColumnDoubleArray) FieldData() *schemapb.FieldData {
	fd := &schemapb.FieldData{
		Type:      schemapb.DataType_Array,
		FieldName: c.name,
	}

	data := make([]*schemapb.ScalarField, 0, c.Len())
	for _, arr := range c.values {
		converted := make([]float64, 0, c.Len())
		for i := 0; i < len(arr); i++ {
			converted = append(converted, float64(arr[i]))
		}
		data = append(data, &schemapb.ScalarField{
			Data: &schemapb.ScalarField_DoubleData{
				DoubleData: &schemapb.DoubleArray{
					Data: converted,
				},
			},
		})
	}
	fd.Field = &schemapb.FieldData_Scalars{
		Scalars: &schemapb.ScalarField{
			Data: &schemapb.ScalarField_ArrayData{
				ArrayData: &schemapb.ArrayArray{
					Data:        data,
					ElementType: schemapb.DataType_Double,
				},
			},
		},
	}
	return fd
}

// ValueByIdx returns value of the provided index
// error occurs when index out of range
func (c *ColumnDoubleArray) ValueByIdx(idx int) ([]float64, error) {
	var r []float64 // use default value
	if idx < 0 || idx >= c.Len() {
		return r, errors.New("index out of range")
	}
	return c.values[idx], nil
}

// AppendValue append value into column
func (c *ColumnDoubleArray) AppendValue(i interface{}) error {
	v, ok := i.([]float64)
	if !ok {
		return fmt.Errorf("invalid type, expected []float64, got %T", i)
	}
	c.values = append(c.values, v)

	return nil
}

// Data returns column data
func (c *ColumnDoubleArray) Data() [][]float64 {
	return c.values
}

// NewColumnDouble auto generated constructor
func NewColumnDoubleArray(name string, values [][]float64) *ColumnDoubleArray {
	return &ColumnDoubleArray{
		name:   name,
		values: values,
	}
}
