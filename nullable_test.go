// Copyright 2017 Oliver Eilhard. All rights reserved.
// Use of this source code is governed by a MIT-license.
// See http://olivere.mit-license.org/license.txt for details.

package nullable

import (
	"testing"
	"time"
)

// -- Int --

func TestInt(t *testing.T) {
	one := int(1)
	two := int(2)
	tests := []struct {
		Input  *int
		Output int
	}{
		{Input: nil, Output: 0},
		{Input: &one, Output: 1},
		{Input: &two, Output: 2},
	}

	for i, tt := range tests {
		if have, want := Int(tt.Input), tt.Output; have != want {
			t.Errorf("#%d: have Int(%v) = %v, want %v", i, tt.Input, have, want)
		}
	}
}

func TestIntWithDefault(t *testing.T) {
	one := int(1)
	two := int(2)
	tests := []struct {
		Input   *int
		Default int
		Output  int
	}{
		{Input: nil, Default: 0, Output: 0},
		{Input: nil, Default: 1, Output: 1},
		{Input: &one, Default: 2, Output: 1},
		{Input: &two, Default: 0, Output: 2},
	}

	for i, tt := range tests {
		if have, want := IntWithDefault(tt.Input, tt.Default), tt.Output; have != want {
			t.Errorf("#%d: have IntWithDefault(%v) = %v, want %v", i, tt.Input, have, want)
		}
	}
}

func TestIntPtr(t *testing.T) {
	one := int(1)
	tests := []struct {
		Input  int
		Output *int
	}{
		{Input: one, Output: &one},
	}

	for i, tt := range tests {
		have, want := IntPtr(tt.Input), tt.Output
		if have == nil || want == nil {
			t.Fatalf("#%d: have IntPtr(%v) = %v, want %v", i, tt.Input, have, want)
		}
		if *have != *want {
			t.Errorf("#%d: have IntPtr(%v) = %v, want %v", i, tt.Input, have, want)
		}
	}
}

func TestIntSlice(t *testing.T) {
	one := int(1)
	two := int(2)
	tests := []struct {
		Input  []*int
		Output []int
	}{
		{
			Input:  []*int{&one, &two},
			Output: []int{one, two},
		},
	}

	for i, tt := range tests {
		have, want := IntSlice(tt.Input), tt.Output
		if haveLen, wantLen := len(have), len(want); haveLen != wantLen {
			t.Fatalf("#%d: have len(IntSlice(%v)) = %d, want %d", i, tt.Input, haveLen, wantLen)
		}
		for j := 0; j < len(have); j++ {
			if x, y := have[j], want[j]; x != y {
				t.Errorf("#%d: have IntSlice(%v)[%d] = %v, want %v", i, tt.Input, j, x, y)
			}
		}
	}
}

func TestIntPtrSlice(t *testing.T) {
	one := int(1)
	two := int(2)
	tests := []struct {
		Input  []int
		Output []*int
	}{
		{
			Input:  []int{one, two},
			Output: []*int{&one, &two},
		},
	}

	for i, tt := range tests {
		have, want := IntPtrSlice(tt.Input), tt.Output
		if haveLen, wantLen := len(have), len(want); haveLen != wantLen {
			t.Fatalf("#%d: have len(IntPtrSlice(%v)) = %d, want %d", i, tt.Input, haveLen, wantLen)
		}
		for j := 0; j < len(have); j++ {
			if x, y := *have[j], *want[j]; x != y {
				t.Errorf("#%d: have IntPtrSlice(%v)[%d] = %v, want %v", i, tt.Input, j, x, y)
			}
		}
	}
}

// -- Int32 --

func TestInt32(t *testing.T) {
	one := int32(1)
	two := int32(2)
	tests := []struct {
		Input  *int32
		Output int32
	}{
		{Input: nil, Output: 0},
		{Input: &one, Output: 1},
		{Input: &two, Output: 2},
	}

	for i, tt := range tests {
		if have, want := Int32(tt.Input), tt.Output; have != want {
			t.Errorf("#%d: have Int32(%v) = %v, want %v", i, tt.Input, have, want)
		}
	}
}

func TestInt32WithDefault(t *testing.T) {
	one := int32(1)
	two := int32(2)
	tests := []struct {
		Input   *int32
		Default int32
		Output  int32
	}{
		{Input: nil, Default: 0, Output: 0},
		{Input: nil, Default: 1, Output: 1},
		{Input: &one, Default: 2, Output: 1},
		{Input: &two, Default: 0, Output: 2},
	}

	for i, tt := range tests {
		if have, want := Int32WithDefault(tt.Input, tt.Default), tt.Output; have != want {
			t.Errorf("#%d: have Int32WithDefault(%v) = %v, want %v", i, tt.Input, have, want)
		}
	}
}

func Test32IntSlice(t *testing.T) {
	one := int32(1)
	two := int32(2)
	tests := []struct {
		Input  []*int32
		Output []int32
	}{
		{
			Input:  []*int32{&one, &two},
			Output: []int32{one, two},
		},
	}

	for i, tt := range tests {
		have, want := Int32Slice(tt.Input), tt.Output
		if haveLen, wantLen := len(have), len(want); haveLen != wantLen {
			t.Fatalf("#%d: have len(Int32Slice(%v)) = %d, want %d", i, tt.Input, haveLen, wantLen)
		}
		for j := 0; j < len(have); j++ {
			if x, y := have[j], want[j]; x != y {
				t.Errorf("#%d: have Int32Slice(%v)[%d] = %v, want %v", i, tt.Input, j, x, y)
			}
		}
	}
}

func TestInt32PtrSlice(t *testing.T) {
	one := int32(1)
	two := int32(2)
	tests := []struct {
		Input  []int32
		Output []*int32
	}{
		{
			Input:  []int32{one, two},
			Output: []*int32{&one, &two},
		},
	}

	for i, tt := range tests {
		have, want := Int32PtrSlice(tt.Input), tt.Output
		if haveLen, wantLen := len(have), len(want); haveLen != wantLen {
			t.Fatalf("#%d: have len(Int32PtrSlice(%v)) = %d, want %d", i, tt.Input, haveLen, wantLen)
		}
		for j := 0; j < len(have); j++ {
			if x, y := *have[j], *want[j]; x != y {
				t.Errorf("#%d: have Int32PtrSlice(%v)[%d] = %v, want %v", i, tt.Input, j, x, y)
			}
		}
	}
}

// -- Int64 --

func TestInt64(t *testing.T) {
	one := int64(1)
	two := int64(2)
	tests := []struct {
		Input  *int64
		Output int64
	}{
		{Input: nil, Output: 0},
		{Input: &one, Output: 1},
		{Input: &two, Output: 2},
	}

	for i, tt := range tests {
		if have, want := Int64(tt.Input), tt.Output; have != want {
			t.Errorf("#%d: have Int64(%v) = %v, want %v", i, tt.Input, have, want)
		}
	}
}

func TestInt64WithDefault(t *testing.T) {
	one := int64(1)
	two := int64(2)
	tests := []struct {
		Input   *int64
		Default int64
		Output  int64
	}{
		{Input: nil, Default: 0, Output: 0},
		{Input: nil, Default: 1, Output: 1},
		{Input: &one, Default: 2, Output: 1},
		{Input: &two, Default: 0, Output: 2},
	}

	for i, tt := range tests {
		if have, want := Int64WithDefault(tt.Input, tt.Default), tt.Output; have != want {
			t.Errorf("#%d: have Int64WithDefault(%v) = %v, want %v", i, tt.Input, have, want)
		}
	}
}

func TestInt64Slice(t *testing.T) {
	one := int64(1)
	two := int64(2)
	tests := []struct {
		Input  []*int64
		Output []int64
	}{
		{
			Input:  []*int64{&one, &two},
			Output: []int64{one, two},
		},
	}

	for i, tt := range tests {
		have, want := Int64Slice(tt.Input), tt.Output
		if haveLen, wantLen := len(have), len(want); haveLen != wantLen {
			t.Fatalf("#%d: have len(Int64Slice(%v)) = %d, want %d", i, tt.Input, haveLen, wantLen)
		}
		for j := 0; j < len(have); j++ {
			if x, y := have[j], want[j]; x != y {
				t.Errorf("#%d: have Int64Slice(%v)[%d] = %v, want %v", i, tt.Input, j, x, y)
			}
		}
	}
}

func TestInt64PtrSlice(t *testing.T) {
	one := int64(1)
	two := int64(2)
	tests := []struct {
		Input  []int64
		Output []*int64
	}{
		{
			Input:  []int64{one, two},
			Output: []*int64{&one, &two},
		},
	}

	for i, tt := range tests {
		have, want := Int64PtrSlice(tt.Input), tt.Output
		if haveLen, wantLen := len(have), len(want); haveLen != wantLen {
			t.Fatalf("#%d: have len(Int64PtrSlice(%v)) = %d, want %d", i, tt.Input, haveLen, wantLen)
		}
		for j := 0; j < len(have); j++ {
			if x, y := *have[j], *want[j]; x != y {
				t.Errorf("#%d: have Int64PtrSlice(%v)[%d] = %v, want %v", i, tt.Input, j, x, y)
			}
		}
	}
}

// -- Float32 --

func TestFloat32(t *testing.T) {
	one := float32(1)
	two := float32(2)
	tests := []struct {
		Input  *float32
		Output float32
	}{
		{Input: nil, Output: 0},
		{Input: &one, Output: 1},
		{Input: &two, Output: 2},
	}

	for i, tt := range tests {
		if have, want := Float32(tt.Input), tt.Output; have != want {
			t.Errorf("#%d: have Float32(%v) = %v, want %v", i, tt.Input, have, want)
		}
	}
}

func TestFloat32WithDefault(t *testing.T) {
	one := float32(1)
	two := float32(2)
	tests := []struct {
		Input   *float32
		Default float32
		Output  float32
	}{
		{Input: nil, Default: 0, Output: 0},
		{Input: nil, Default: 1, Output: 1},
		{Input: &one, Default: 2, Output: 1},
		{Input: &two, Default: 0, Output: 2},
	}

	for i, tt := range tests {
		if have, want := Float32WithDefault(tt.Input, tt.Default), tt.Output; have != want {
			t.Errorf("#%d: have Float32WithDefault(%v) = %v, want %v", i, tt.Input, have, want)
		}
	}
}

func TestFloat32Slice(t *testing.T) {
	one := float32(1.5)
	two := float32(2.7)
	tests := []struct {
		Input  []*float32
		Output []float32
	}{
		{
			Input:  []*float32{&one, &two},
			Output: []float32{one, two},
		},
	}

	for i, tt := range tests {
		have, want := Float32Slice(tt.Input), tt.Output
		if haveLen, wantLen := len(have), len(want); haveLen != wantLen {
			t.Fatalf("#%d: have len(Float32Slice(%v)) = %d, want %d", i, tt.Input, haveLen, wantLen)
		}
		for j := 0; j < len(have); j++ {
			if x, y := have[j], want[j]; x != y {
				t.Errorf("#%d: have Float32Slice(%v)[%d] = %v, want %v", i, tt.Input, j, x, y)
			}
		}
	}
}

func TestFloat32PtrSlice(t *testing.T) {
	one := float32(1)
	two := float32(2)
	tests := []struct {
		Input  []float32
		Output []*float32
	}{
		{
			Input:  []float32{one, two},
			Output: []*float32{&one, &two},
		},
	}

	for i, tt := range tests {
		have, want := Float32PtrSlice(tt.Input), tt.Output
		if haveLen, wantLen := len(have), len(want); haveLen != wantLen {
			t.Fatalf("#%d: have len(Float32PtrSlice(%v)) = %d, want %d", i, tt.Input, haveLen, wantLen)
		}
		for j := 0; j < len(have); j++ {
			if x, y := *have[j], *want[j]; x != y {
				t.Errorf("#%d: have Float32PtrSlice(%v)[%d] = %v, want %v", i, tt.Input, j, x, y)
			}
		}
	}
}

// -- Float64 --

func TestFloat64(t *testing.T) {
	one := float64(1)
	two := float64(2)
	tests := []struct {
		Input  *float64
		Output float64
	}{
		{Input: nil, Output: 0},
		{Input: &one, Output: 1},
		{Input: &two, Output: 2},
	}

	for i, tt := range tests {
		if have, want := Float64(tt.Input), tt.Output; have != want {
			t.Errorf("#%d: have Float64(%v) = %v, want %v", i, tt.Input, have, want)
		}
	}
}

func TestFloat64WithDefault(t *testing.T) {
	one := float64(1)
	two := float64(2)
	tests := []struct {
		Input   *float64
		Default float64
		Output  float64
	}{
		{Input: nil, Default: 0, Output: 0},
		{Input: nil, Default: 1, Output: 1},
		{Input: &one, Default: 2, Output: 1},
		{Input: &two, Default: 0, Output: 2},
	}

	for i, tt := range tests {
		if have, want := Float64WithDefault(tt.Input, tt.Default), tt.Output; have != want {
			t.Errorf("#%d: have Float64WithDefault(%v) = %v, want %v", i, tt.Input, have, want)
		}
	}
}

func TestFloat64Slice(t *testing.T) {
	one := float64(1.5)
	two := float64(2.7)
	tests := []struct {
		Input  []*float64
		Output []float64
	}{
		{
			Input:  []*float64{&one, &two},
			Output: []float64{one, two},
		},
	}

	for i, tt := range tests {
		have, want := Float64Slice(tt.Input), tt.Output
		if haveLen, wantLen := len(have), len(want); haveLen != wantLen {
			t.Fatalf("#%d: have len(Float64Slice(%v)) = %d, want %d", i, tt.Input, haveLen, wantLen)
		}
		for j := 0; j < len(have); j++ {
			if x, y := have[j], want[j]; x != y {
				t.Errorf("#%d: have Float64Slice(%v)[%d] = %v, want %v", i, tt.Input, j, x, y)
			}
		}
	}
}

func TestFloat64PtrSlice(t *testing.T) {
	one := float64(1)
	two := float64(2)
	tests := []struct {
		Input  []float64
		Output []*float64
	}{
		{
			Input:  []float64{one, two},
			Output: []*float64{&one, &two},
		},
	}

	for i, tt := range tests {
		have, want := Float64PtrSlice(tt.Input), tt.Output
		if haveLen, wantLen := len(have), len(want); haveLen != wantLen {
			t.Fatalf("#%d: have len(Float64PtrSlice(%v)) = %d, want %d", i, tt.Input, haveLen, wantLen)
		}
		for j := 0; j < len(have); j++ {
			if x, y := *have[j], *want[j]; x != y {
				t.Errorf("#%d: have Float64PtrSlice(%v)[%d] = %v, want %v", i, tt.Input, j, x, y)
			}
		}
	}
}

// -- String --
func TestString(t *testing.T) {
	one := "one"
	two := "two"
	tests := []struct {
		Input  *string
		Output string
	}{
		{Input: nil, Output: ""},
		{Input: &one, Output: one},
		{Input: &two, Output: two},
	}

	for i, tt := range tests {
		if have, want := String(tt.Input), tt.Output; have != want {
			t.Errorf("#%d: have String(%v) = %v, want %v", i, tt.Input, have, want)
		}
	}
}

func TestStringWithDefault(t *testing.T) {
	one := "one"
	two := "two"
	tests := []struct {
		Input   *string
		Default string
		Output  string
	}{
		{Input: nil, Default: "", Output: ""},
		{Input: nil, Default: one, Output: one},
		{Input: &one, Default: two, Output: one},
		{Input: &two, Default: "", Output: two},
	}

	for i, tt := range tests {
		if have, want := StringWithDefault(tt.Input, tt.Default), tt.Output; have != want {
			t.Errorf("#%d: have StringWithDefault(%v) = %v, want %v", i, tt.Input, have, want)
		}
	}
}

func TestStringPtr(t *testing.T) {
	one := "one"
	tests := []struct {
		Input  string
		Output *string
	}{
		{Input: one, Output: &one},
	}

	for i, tt := range tests {
		have, want := StringPtr(tt.Input), tt.Output
		if have == nil || want == nil {
			t.Fatalf("#%d: have StringPtr(%v) = %v, want %v", i, tt.Input, have, want)
		}
		if *have != *want {
			t.Errorf("#%d: have StringPtr(%v) = %v, want %v", i, tt.Input, have, want)
		}
	}
}

func TestStringSlice(t *testing.T) {
	one := "one"
	two := "two"
	tests := []struct {
		Input  []*string
		Output []string
	}{
		{
			Input:  []*string{&one, &two},
			Output: []string{one, two},
		},
	}

	for i, tt := range tests {
		have, want := StringSlice(tt.Input), tt.Output
		if haveLen, wantLen := len(have), len(want); haveLen != wantLen {
			t.Fatalf("#%d: have len(StringSlice(%v)) = %d, want %d", i, tt.Input, haveLen, wantLen)
		}
		for j := 0; j < len(have); j++ {
			if x, y := have[j], want[j]; x != y {
				t.Errorf("#%d: have StringSlice(%v)[%d] = %v, want %v", i, tt.Input, j, x, y)
			}
		}
	}
}

func TestStringPtrSlice(t *testing.T) {
	one := "one"
	two := "two"
	tests := []struct {
		Input  []string
		Output []*string
	}{
		{
			Input:  []string{one, two},
			Output: []*string{&one, &two},
		},
	}

	for i, tt := range tests {
		have, want := StringPtrSlice(tt.Input), tt.Output
		if haveLen, wantLen := len(have), len(want); haveLen != wantLen {
			t.Fatalf("#%d: have len(StringPtrSlice(%v)) = %d, want %d", i, tt.Input, haveLen, wantLen)
		}
		for j := 0; j < len(have); j++ {
			if x, y := *have[j], *want[j]; x != y {
				t.Errorf("#%d: have StringPtrSlice(%v)[%d] = %v, want %v", i, tt.Input, j, x, y)
			}
		}
	}
}

// -- Bool --

func TestBool(t *testing.T) {
	tr := true
	f := false
	tests := []struct {
		Input  *bool
		Output bool
	}{
		{Input: nil, Output: false},
		{Input: &tr, Output: true},
		{Input: &f, Output: false},
	}

	for i, tt := range tests {
		if have, want := Bool(tt.Input), tt.Output; have != want {
			t.Errorf("#%d: have Bool(%v) = %v, want %v", i, tt.Input, have, want)
		}
	}
}

func TestBoolWithDefault(t *testing.T) {
	tr := true
	f := false
	tests := []struct {
		Input   *bool
		Default bool
		Output  bool
	}{
		{Input: nil, Default: false, Output: false},
		{Input: nil, Default: true, Output: true},
		{Input: &tr, Default: false, Output: true},
		{Input: &f, Default: true, Output: false},
	}

	for i, tt := range tests {
		if have, want := BoolWithDefault(tt.Input, tt.Default), tt.Output; have != want {
			t.Errorf("#%d: have BoolWithDefault(%v) = %v, want %v", i, tt.Input, have, want)
		}
	}
}

// -- Time --

func TestTime(t *testing.T) {
	var zero time.Time
	one := time.Date(2017, 1, 2, 12, 14, 59, 0, time.UTC)
	two := time.Date(1982, 11, 23, 23, 11, 9, 0, time.UTC)
	tests := []struct {
		Input  *time.Time
		Output time.Time
	}{
		{Input: nil, Output: zero},
		{Input: &one, Output: one},
		{Input: &two, Output: two},
	}

	for i, tt := range tests {
		if have, want := Time(tt.Input), tt.Output; have != want {
			t.Errorf("#%d: have Time(%v) = %v, want %v", i, tt.Input, have, want)
		}
	}
}

func TestTimeWithDefault(t *testing.T) {
	var zero time.Time
	one := time.Date(2017, 1, 2, 12, 14, 59, 0, time.UTC)
	two := time.Date(1982, 11, 23, 23, 11, 9, 0, time.UTC)
	tests := []struct {
		Input   *time.Time
		Default time.Time
		Output  time.Time
	}{
		{Input: nil, Default: zero, Output: zero},
		{Input: nil, Default: one, Output: one},
		{Input: &one, Default: zero, Output: one},
		{Input: &two, Default: one, Output: two},
	}

	for i, tt := range tests {
		if have, want := TimeWithDefault(tt.Input, tt.Default), tt.Output; have != want {
			t.Errorf("#%d: have TimeWithDefault(%v) = %v, want %v", i, tt.Input, have, want)
		}
	}
}

// -- Duration --

func TestDuration(t *testing.T) {
	var zero time.Duration
	one := time.Duration(63 * time.Second)
	two := time.Duration(3 * time.Minute)
	tests := []struct {
		Input  *time.Duration
		Output time.Duration
	}{
		{Input: nil, Output: zero},
		{Input: &one, Output: one},
		{Input: &two, Output: two},
	}

	for i, tt := range tests {
		if have, want := Duration(tt.Input), tt.Output; have != want {
			t.Errorf("#%d: have Duration(%v) = %v, want %v", i, tt.Input, have, want)
		}
	}
}

func TestDurationWithDefault(t *testing.T) {
	var zero time.Duration
	one := time.Duration(63 * time.Second)
	two := time.Duration(3 * time.Minute)
	tests := []struct {
		Input   *time.Duration
		Default time.Duration
		Output  time.Duration
	}{
		{Input: nil, Default: zero, Output: zero},
		{Input: nil, Default: one, Output: one},
		{Input: &one, Default: zero, Output: one},
		{Input: &two, Default: one, Output: two},
	}

	for i, tt := range tests {
		if have, want := DurationWithDefault(tt.Input, tt.Default), tt.Output; have != want {
			t.Errorf("#%d: have DurationWithDefault(%v) = %v, want %v", i, tt.Input, have, want)
		}
	}
}
