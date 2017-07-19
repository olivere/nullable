// Copyright 2017 Oliver Eilhard. All rights reserved.
// Use of this source code is governed by a MIT-license.
// See http://olivere.mit-license.org/license.txt for details.

package nullable

import "time"

// -- Int --

// Int returns *v if v is not nil. Otherwise it returns 0.
func Int(v *int) int {
	return IntWithDefault(v, 0)
}

// IntWithDefault returns *v if v is not nil. Otherwise it returns d.
func IntWithDefault(v *int, d int) int {
	if v == nil {
		return d
	}
	return *v
}

// IntPtr returns a pointer to v.
func IntPtr(v int) *int {
	return &v
}

// IntSlice converts a slice of int pointers to a slice of
// int values. Elements that are nil are converted to its zero value.
func IntSlice(src []*int) []int {
	dst := make([]int, len(src))
	for i := 0; i < len(src); i++ {
		if s := src[i]; s != nil {
			dst[i] = *s
		} else {
			dst[i] = 0
		}
	}
	return dst
}

// IntPtrSlice converts a slice of int values to a slice of
// int pointers.
func IntPtrSlice(src []int) []*int {
	dst := make([]*int, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// -- Int32 --

// Int32 returns *v if v is not nil. Otherwise it returns 0.
func Int32(v *int32) int32 {
	return Int32WithDefault(v, 0)
}

// Int32WithDefault returns *v if v is not nil. Otherwise it returns d.
func Int32WithDefault(v *int32, d int32) int32 {
	if v == nil {
		return d
	}
	return *v
}

// Int32Ptr returns a pointer to v.
func Int32Ptr(v int32) *int32 {
	return &v
}

// Int32Slice converts a slice of int32 pointers to a slice of
// int32 values. Elements that are nil are converted to its zero value.
func Int32Slice(src []*int32) []int32 {
	dst := make([]int32, len(src))
	for i := 0; i < len(src); i++ {
		if v := src[i]; v != nil {
			dst[i] = *v
		}
	}
	return dst
}

// Int32PtrSlice converts a slice of int32 values to a slice of
// int32 pointers.
func Int32PtrSlice(src []int32) []*int32 {
	dst := make([]*int32, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// -- Int64 --

// Int64 returns *v if v is not nil. Otherwise it returns 0.
func Int64(v *int64) int64 {
	return Int64WithDefault(v, 0)
}

// Int64WithDefault returns *v if v is not nil. Otherwise it returns d.
func Int64WithDefault(v *int64, d int64) int64 {
	if v == nil {
		return d
	}
	return *v
}

// Int64Ptr returns a pointer to v.
func Int64Ptr(v int64) *int64 {
	return &v
}

// Int64Slice converts a slice of int64 pointers to a slice of
// int64 values. Elements that are nil are converted to its zero value.
func Int64Slice(src []*int64) []int64 {
	dst := make([]int64, len(src))
	for i := 0; i < len(src); i++ {
		if v := src[i]; v != nil {
			dst[i] = *v
		}
	}
	return dst
}

// Int64PtrSlice converts a slice of int64 values to a slice of
// int64 pointers.
func Int64PtrSlice(src []int64) []*int64 {
	dst := make([]*int64, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// -- Float32 --

// Float32 returns *v if v is not nil. Otherwise it returns 0.
func Float32(v *float32) float32 {
	return Float32WithDefault(v, 0)
}

// Float32WithDefault returns *v if v is not nil. Otherwise it returns d.
func Float32WithDefault(v *float32, d float32) float32 {
	if v == nil {
		return d
	}
	return *v
}

// Float32Ptr returns a pointer to v.
func Float32Ptr(v float32) *float32 {
	return &v
}

// Float32Slice converts a slice of float32 pointers to a slice of
// float32 values. Elements that are nil are converted to its zero value.
func Float32Slice(src []*float32) []float32 {
	dst := make([]float32, len(src))
	for i := 0; i < len(src); i++ {
		if v := src[i]; v != nil {
			dst[i] = *v
		}
	}
	return dst
}

// Float32PtrSlice converts a slice of float32 values to a slice of
// float32 pointers.
func Float32PtrSlice(src []float32) []*float32 {
	dst := make([]*float32, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// -- Float64 --

// Float64 returns *v if v is not nil. Otherwise it returns 0.
func Float64(v *float64) float64 {
	return Float64WithDefault(v, 0)
}

// Float64WithDefault returns *v if v is not nil. Otherwise it returns d.
func Float64WithDefault(v *float64, d float64) float64 {
	if v == nil {
		return d
	}
	return *v
}

// Float64Ptr returns a pointer to v.
func Float64Ptr(v float64) *float64 {
	return &v
}

// Float64Slice converts a slice of float64 pointers to a slice of
// float64 values. Elements that are nil are converted to its zero value.
func Float64Slice(src []*float64) []float64 {
	dst := make([]float64, len(src))
	for i := 0; i < len(src); i++ {
		if v := src[i]; v != nil {
			dst[i] = *v
		}
	}
	return dst
}

// Float64PtrSlice converts a slice of float64 values to a slice of
// float64 pointers.
func Float64PtrSlice(src []float64) []*float64 {
	dst := make([]*float64, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// -- String --

// String returns *v if v is not nil. Otherwise it returns "".
func String(v *string) string {
	return StringWithDefault(v, "")
}

// StringWithDefault returns *v if v is not nil. Otherwise it returns d.
func StringWithDefault(v *string, d string) string {
	if v == nil {
		return d
	}
	return *v
}

// StringPtr returns a pointer to v.
func StringPtr(v string) *string {
	return &v
}

// StringSlice converts a slice of string pointers to a slice of
// string values. Elements that are nil are converted to empty strings.
func StringSlice(src []*string) []string {
	dst := make([]string, len(src))
	for i := 0; i < len(src); i++ {
		if s := src[i]; s != nil {
			dst[i] = *s
		} else {
			dst[i] = ""
		}
	}
	return dst
}

// StringPtrSlice converts a slice of string values to a slice of
// string pointers.
func StringPtrSlice(src []string) []*string {
	dst := make([]*string, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// -- Bool --

// Bool returns *v if v is not nil. Otherwise it returns fase.
func Bool(v *bool) bool {
	return BoolWithDefault(v, false)
}

// BoolWithDefault returns *v if v is not nil. Otherwise it returns d.
func BoolWithDefault(v *bool, d bool) bool {
	if v == nil {
		return d
	}
	return *v
}

// BoolPtr returns a pointer to v.
func BoolPtr(v bool) *bool {
	return &v
}

// -- Time --

// Time returns *v if v is not nil. Otherwise it returns an empty date time.
func Time(v *time.Time) time.Time {
	var dt time.Time
	if v != nil {
		return *v
	}
	return dt
}

// TimeWithDefault returns *v if v is not nil. Otherwise it returns d.
func TimeWithDefault(v *time.Time, d time.Time) time.Time {
	if v == nil {
		return d
	}
	return *v
}

// TimePtr returns a pointer to v.
func TimePtr(v time.Time) *time.Time {
	return &v
}

// -- Duration --

// Duration returns *v if v is not nil. Otherwise it returns an empty duration.
func Duration(v *time.Duration) time.Duration {
	var d time.Duration
	if v != nil {
		return *v
	}
	return d
}

// DurationWithDefault returns *v if v is not nil. Otherwise it returns d.
func DurationWithDefault(v *time.Duration, d time.Duration) time.Duration {
	if v == nil {
		return d
	}
	return *v
}

// DurationPtr returns a pointer to v.
func DurationPtr(v time.Duration) *time.Duration {
	return &v
}
