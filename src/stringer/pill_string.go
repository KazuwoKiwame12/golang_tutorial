// Code generated by "stringer -type=Pill"; DO NOT EDIT.

package main

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[A-0]
	_ = x[B-1]
	_ = x[C-2]
}

const _Pill_name = "ABC"

var _Pill_index = [...]uint8{0, 1, 2, 3}

func (i Pill) String() string {
	if i < 0 || i >= Pill(len(_Pill_index)-1) {
		return "Pill(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Pill_name[_Pill_index[i]:_Pill_index[i+1]]
}
