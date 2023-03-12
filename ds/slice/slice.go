package slice

// ISlice is an interface of Slice for iterator
type ISlice interface {
	Len() int
	At(position int) interface{}
	Set(position int, val interface{})
}

// Slice definition
type Slice []interface{}

// IntSlice is a redefinition []int
type IntSlice []int

// UIntSlice is a definition []uint
type UIntSlice []uint

// Int8Slice is redefinition []int8
type Int8Slice []int8

// UInt8Slice is a redefinition []uint8
type UInt8Slice []uint8

// Int16Slice is redefinition []int16
type Int16Slice []int16

// UInt16Slice is a redefinition []uint16
type UInt16Slice []uint16

// Int32Slice is redefinition []int32
type Int32Slice []int32

// UInt32Slice is a redefinition []uint32
type UInt32Slice []uint32

// Int64Slice is a redefinition []int64
type Int64Slice []int64

// UInt64Slice is a redefinition []uint64
type UInt64Slice []uint64

// Float32Slice is a redefinition []float32
type Float32Slice []float32

// Float64Slice is a redefinition []float64
type Float64Slice []float64

// StringSlice is a redefinition []string
type StringSlice []string

// Tells the compiler XXSlice implements ISlice
var _ ISlice = Slice(nil)
var _ ISlice = IntSlice(nil)
var _ ISlice = UIntSlice(nil)
var _ ISlice = Int8Slice(nil)
var _ ISlice = UInt8Slice(nil)
var _ ISlice = Int16Slice(nil)
var _ ISlice = UInt16Slice(nil)
var _ ISlice = Int32Slice(nil)
var _ ISlice = Int32Slice(nil)
var _ ISlice = Int64Slice(nil)
var _ ISlice = Float32Slice(nil)
var _ ISlice = Float64Slice(nil)
var _ ISlice = StringSlice(nil)

// Len returns the length of s
func (s Slice) Len() int {
	 //TODO: Complete me!
}

// At returns the value at position
func (s Slice) At(position int) interface{} {
	 //TODO: Complete me!
}

// Set sets value at position
func (s Slice) Set(position int, val interface{}) {
	 //TODO: Complete me!
}

// Begin returns the first iterator of s
func (s Slice) Begin() *SliceIterator {
	 //TODO: Complete me!
}

// End returns the end iterator of s
func (s Slice) End() *SliceIterator {
	 //TODO: Complete me!
}

// First returns the first iterator of s
func (s Slice) First() *SliceIterator {
	 //TODO: Complete me!
}

// Last returns the last iterator of s
func (s Slice) Last() *SliceIterator {
	 //TODO: Complete me!
}

// Len returns the length of s
func (s IntSlice) Len() int {
	 //TODO: Complete me!
}

// At returns the value at position
func (s IntSlice) At(position int) interface{} {
	 //TODO: Complete me!
}

// Set sets value at position
func (s IntSlice) Set(position int, val interface{}) {
	 //TODO: Complete me!
}

// Begin returns the first iterator of s
func (s IntSlice) Begin() *SliceIterator {
	 //TODO: Complete me!
}

// End returns the end iterator of s
func (s IntSlice) End() *SliceIterator {
	 //TODO: Complete me!
}

// First returns the first iterator of s
func (s IntSlice) First() *SliceIterator {
	 //TODO: Complete me!
}

// Last returns the last iterator of s
func (s IntSlice) Last() *SliceIterator {
	 //TODO: Complete me!
}

// Len returns the length of s
func (s UIntSlice) Len() int {
	 //TODO: Complete me!
}

// At returns the value at position
func (s UIntSlice) At(position int) interface{} {
	 //TODO: Complete me!
}

// Set sets value at position
func (s UIntSlice) Set(position int, val interface{}) {
	 //TODO: Complete me!
}

// Begin returns the first iterator of s
func (s UIntSlice) Begin() *SliceIterator {
	 //TODO: Complete me!
}

// End returns the end iterator of s
func (s UIntSlice) End() *SliceIterator {
	 //TODO: Complete me!
}

// First returns the first iterator of s
func (s UIntSlice) First() *SliceIterator {
	 //TODO: Complete me!
}

// Last returns the last iterator of s
func (s UIntSlice) Last() *SliceIterator {
	 //TODO: Complete me!
}

// Len returns the length of s
func (s Int8Slice) Len() int {
	 //TODO: Complete me!
}

// At returns the value at position
func (s Int8Slice) At(position int) interface{} {
	 //TODO: Complete me!
}

// Set sets value at position
func (s Int8Slice) Set(position int, val interface{}) {
	 //TODO: Complete me!
}

// Begin returns the first iterator of s
func (s Int8Slice) Begin() *SliceIterator {
	 //TODO: Complete me!
}

// End returns the end iterator of s
func (s Int8Slice) End() *SliceIterator {
	 //TODO: Complete me!
}

// First returns the first iterator of s
func (s Int8Slice) First() *SliceIterator {
	 //TODO: Complete me!
}

// Last returns the last iterator of s
func (s Int8Slice) Last() *SliceIterator {
	 //TODO: Complete me!
}

// Len returns the length of s
func (s UInt8Slice) Len() int {
	 //TODO: Complete me!
}

// At returns the value at position
func (s UInt8Slice) At(position int) interface{} {
	 //TODO: Complete me!
}

// Set sets value at position
func (s UInt8Slice) Set(position int, val interface{}) {
	 //TODO: Complete me!
}

// Begin returns the first iterator of s
func (s UInt8Slice) Begin() *SliceIterator {
	 //TODO: Complete me!
}

// End returns the end iterator of s
func (s UInt8Slice) End() *SliceIterator {
	 //TODO: Complete me!
}

// First returns the first iterator of s
func (s UInt8Slice) First() *SliceIterator {
	 //TODO: Complete me!
}

// Last returns the last iterator of s
func (s UInt8Slice) Last() *SliceIterator {
	 //TODO: Complete me!
}

// Len returns the length of s
func (s Int16Slice) Len() int {
	 //TODO: Complete me!
}

// At returns the value at position
func (s Int16Slice) At(position int) interface{} {
	 //TODO: Complete me!
}

// Set sets value at position
func (s Int16Slice) Set(position int, val interface{}) {
	 //TODO: Complete me!
}

// Begin returns the first iterator of s
func (s Int16Slice) Begin() *SliceIterator {
	 //TODO: Complete me!
}

// End returns the end iterator of s
func (s Int16Slice) End() *SliceIterator {
	 //TODO: Complete me!
}

// First returns the first iterator of s
func (s Int16Slice) First() *SliceIterator {
	 //TODO: Complete me!
}

// Last returns the last iterator of s
func (s Int16Slice) Last() *SliceIterator {
	 //TODO: Complete me!
}

// Len returns the length of s
func (s UInt16Slice) Len() int {
	 //TODO: Complete me!
}

// At returns the value at position
func (s UInt16Slice) At(position int) interface{} {
	 //TODO: Complete me!
}

// Set sets value at position
func (s UInt16Slice) Set(position int, val interface{}) {
	 //TODO: Complete me!
}

// Begin returns the first iterator of s
func (s UInt16Slice) Begin() *SliceIterator {
	 //TODO: Complete me!
}

// End returns the end iterator of s
func (s UInt16Slice) End() *SliceIterator {
	 //TODO: Complete me!
}

// First returns the first iterator of s
func (s UInt16Slice) First() *SliceIterator {
	 //TODO: Complete me!
}

// Last returns the last iterator of s
func (s UInt16Slice) Last() *SliceIterator {
	 //TODO: Complete me!
}

// Len returns the length of s
func (s Int32Slice) Len() int {
	 //TODO: Complete me!
}

// At returns the value at position
func (s Int32Slice) At(position int) interface{} {
	 //TODO: Complete me!
}

// Set sets value at position
func (s Int32Slice) Set(position int, val interface{}) {
	 //TODO: Complete me!
}

// Begin returns the first iterator of s
func (s Int32Slice) Begin() *SliceIterator {
	 //TODO: Complete me!
}

// End returns the end iterator of s
func (s Int32Slice) End() *SliceIterator {
	 //TODO: Complete me!
}

// First returns the first iterator of s
func (s Int32Slice) First() *SliceIterator {
	 //TODO: Complete me!
}

// Last returns the last iterator of s
func (s Int32Slice) Last() *SliceIterator {
	 //TODO: Complete me!
}

// Len returns the length of s
func (s UInt32Slice) Len() int {
	 //TODO: Complete me!
}

// At returns the value at position
func (s UInt32Slice) At(position int) interface{} {
	 //TODO: Complete me!
}

// Set sets value at position
func (s UInt32Slice) Set(position int, val interface{}) {
	 //TODO: Complete me!
}

// Begin returns the first iterator of s
func (s UInt32Slice) Begin() *SliceIterator {
	 //TODO: Complete me!
}

// End returns the end iterator of s
func (s UInt32Slice) End() *SliceIterator {
	 //TODO: Complete me!
}

// First returns the first iterator of s
func (s UInt32Slice) First() *SliceIterator {
	 //TODO: Complete me!
}

// Last returns the last iterator of s
func (s UInt32Slice) Last() *SliceIterator {
	 //TODO: Complete me!
}

// Len returns the length of s
func (s Int64Slice) Len() int {
	 //TODO: Complete me!
}

// At returns the value at position
func (s Int64Slice) At(position int) interface{} {
	 //TODO: Complete me!
}

// Set sets value at position
func (s Int64Slice) Set(position int, val interface{}) {
	 //TODO: Complete me!
}

// Begin returns the first iterator of s
func (s Int64Slice) Begin() *SliceIterator {
	 //TODO: Complete me!
}

// End returns the end iterator of s
func (s Int64Slice) End() *SliceIterator {
	 //TODO: Complete me!
}

// First returns the first iterator of s
func (s Int64Slice) First() *SliceIterator {
	 //TODO: Complete me!
}

// Last returns the last iterator of s
func (s Int64Slice) Last() *SliceIterator {
	 //TODO: Complete me!
}

// Len returns the length of s
func (s UInt64Slice) Len() int {
	 //TODO: Complete me!
}

// At returns the value at position
func (s UInt64Slice) At(position int) interface{} {
	 //TODO: Complete me!
}

// Set sets value at position
func (s UInt64Slice) Set(position int, val interface{}) {
	 //TODO: Complete me!
}

// Begin returns the first iterator of s
func (s UInt64Slice) Begin() *SliceIterator {
	 //TODO: Complete me!
}

// End returns the end iterator of s
func (s UInt64Slice) End() *SliceIterator {
	 //TODO: Complete me!
}

// First returns the first iterator of s
func (s UInt64Slice) First() *SliceIterator {
	 //TODO: Complete me!
}

// Last returns the last iterator of s
func (s UInt64Slice) Last() *SliceIterator {
	 //TODO: Complete me!
}

// Len returns the length of s
func (s Float32Slice) Len() int {
	 //TODO: Complete me!
}

// At returns the value at position
func (s Float32Slice) At(position int) interface{} {
	 //TODO: Complete me!
}

// Set sets value at position
func (s Float32Slice) Set(position int, val interface{}) {
	 //TODO: Complete me!
}

// Begin returns the first iterator of s
func (s Float32Slice) Begin() *SliceIterator {
	 //TODO: Complete me!
}

// End returns the end iterator of s
func (s Float32Slice) End() *SliceIterator {
	 //TODO: Complete me!
}

// First returns the first iterator of s
func (s Float32Slice) First() *SliceIterator {
	 //TODO: Complete me!
}

// Last returns the last iterator of s
func (s Float32Slice) Last() *SliceIterator {
	 //TODO: Complete me!
}

// Len returns the length of s
func (s Float64Slice) Len() int {
	 //TODO: Complete me!
}

// At returns the value at position
func (s Float64Slice) At(position int) interface{} {
	 //TODO: Complete me!
}

// Set sets value at position
func (s Float64Slice) Set(position int, val interface{}) {
	 //TODO: Complete me!
}

// Begin returns the first iterator of s
func (s Float64Slice) Begin() *SliceIterator {
	 //TODO: Complete me!
}

// End returns the end iterator of s
func (s Float64Slice) End() *SliceIterator {
	 //TODO: Complete me!
}

// First returns the first iterator of s
func (s Float64Slice) First() *SliceIterator {
	 //TODO: Complete me!
}

// Last returns the last iterator of s
func (s Float64Slice) Last() *SliceIterator {
	 //TODO: Complete me!
}

// Len returns the length of s
func (s StringSlice) Len() int {
	 //TODO: Complete me!
}

// At returns the value at position
func (s StringSlice) At(position int) interface{} {
	 //TODO: Complete me!
}

// Set sets value at position
func (s StringSlice) Set(position int, val interface{}) {
	 //TODO: Complete me!
}

// Begin returns the first iterator of s
func (s StringSlice) Begin() *SliceIterator {
	 //TODO: Complete me!
}

// End returns the end iterator of s
func (s StringSlice) End() *SliceIterator {
	 //TODO: Complete me!
}

// First returns the first iterator of s
func (s StringSlice) First() *SliceIterator {
	 //TODO: Complete me!
}

// Last returns the last iterator of s
func (s StringSlice) Last() *SliceIterator {
	 //TODO: Complete me!
}
