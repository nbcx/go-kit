package to

import (
	"fmt"
	"strconv"
	"strings"
)

type Value string

func ValueF(v any) Value {
	return Value(fmt.Sprint(v))
}

// Int returns input as an int
func (v Value) Int() int {
	i, _ := v.IntE()
	return i
}

func (v Value) IntE() (int, error) {
	return strconv.Atoi(string(v))
}

func (v Value) Int8() int8 {
	i, _ := v.Int8E()
	return i
}

// Int8E return input as an int8
func (v Value) Int8E() (int8, error) {
	i64, err := strconv.ParseInt(string(v), 10, 8)
	return int8(i64), err
}

// GetUint8 return input as an uint8
func (v Value) Uint8() (i uint8) {
	i, _ = v.Uint8E()
	return
}

// GetUint8 return input as an uint8
func (v Value) Uint8E() (uint8, error) {
	u64, err := strconv.ParseUint(string(v), 10, 8)
	return uint8(u64), err
}

func (v Value) Int16() (i int16) {
	i, _ = v.Int16E()
	return
}

// GetInt16 returns input as an int16
func (v Value) Int16E() (int16, error) {
	i64, err := strconv.ParseInt(string(v), 10, 16)
	return int16(i64), err
}

func (v Value) Uint16() (i uint16) {
	i, _ = v.Uint16E()
	return
}

// Uint16E returns input as an uint16 and return error if has
func (v Value) Uint16E() (uint16, error) {
	u64, err := strconv.ParseUint(string(v), 10, 16)
	return uint16(u64), err
}

func (v Value) Int32() (i int32) {
	i, _ = v.Int32E()
	return
}

// Int32E returns input as an int32
func (v Value) Int32E() (int32, error) {
	i64, err := strconv.ParseInt(string(v), 10, 32)
	return int32(i64), err
}

func (v Value) Uint32() (i uint32) {
	i, _ = v.Uint32E()
	return
}

// Uint32E returns input as an uint32
func (v Value) Uint32E() (uint32, error) {
	u64, err := strconv.ParseUint(string(v), 10, 32)
	return uint32(u64), err
}

func (v Value) Int64() (i int64) {
	i, _ = v.Int64E()
	return
}

// Int64E returns input value as int64
func (v Value) Int64E() (int64, error) {
	return strconv.ParseInt(string(v), 10, 64)
}

func (v Value) Uint64() (i uint64) {
	i, _ = v.Uint64E()
	return
}

// Uint64E returns input value as uint64
func (v Value) Uint64E() (uint64, error) {
	return strconv.ParseUint(string(v), 10, 64)
}

func (v Value) Bool() (i bool) {
	i, _ = v.BoolE()
	return
}

// GetBool returns input value as bool
func (v Value) BoolE() (bool, error) {
	return strconv.ParseBool(string(v))
}

func (v Value) Float() (i float64) {
	i, _ = v.FloatE()
	return
}

// FloatE returns input value as float64
func (v Value) FloatE() (float64, error) {
	return strconv.ParseFloat(string(v), 64)
}

// String returns input value as float64
func (v Value) String() string {
	return string(v)
}

func (v Value) Split(sep string) Values[string] {
	sp := strings.Split(string(v), sep)
	return sp
}

type Values[T any] []T

func ValuesF[T any](v []T) Values[T] {
	return Values[T](v)
}

func (v Values[T]) Int32() (i []int32) {
	for _, vv := range v {
		i = append(i, ValueF(vv).Int32())
	}
	return
}

func (v Values[T]) Int() (i []int) {
	for _, vv := range v {
		i = append(i, ValueF(vv).Int())
	}
	return
}

func (v Values[T]) String() (i []string) {
	for _, vv := range v {
		i = append(i, ValueF(vv).String())
	}
	return
}

func (v Values[T]) First() (val Value) {
	for _, vv := range v {
		return ValueF(vv)
	}
	return
}

type ValueM[K comparable, V any] map[K]V

func ValueMF[K comparable, V any](v map[K]V) ValueM[K, V] {
	return ValueM[K, V](v)
}

func (v ValueM[K, V]) String() (i map[K]string) {
	i = make(map[K]string)
	for k, vv := range v {
		i[k] = ValueF(vv).String()
	}
	return
}

func (v ValueM[K, V]) Int() (i map[K]int) {
	i = make(map[K]int)
	for k, vv := range v {
		i[k] = ValueF(vv).Int()
	}
	return
}

func (v ValueM[K, V]) IntE() (i map[K]int) {
	i = make(map[K]int)
	for k, vv := range v {
		i[k] = ValueF(vv).Int()
	}
	return
}

func (v ValueM[K, V]) KInt() (i map[int]V) {
	i = make(map[int]V)
	for k, vv := range v {
		i[ValueF(k).Int()] = vv
	}
	return
}

func (v ValueM[K, V]) KVInt() (i map[int]int) {
	i = make(map[int]int)
	for k, vv := range v {
		i[ValueF(k).Int()] = ValueF(vv).Int()
	}
	return
}
