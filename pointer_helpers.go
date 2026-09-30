// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.
package firecracker

// BoolValue will return a boolean value. If the pointer is nil, then false
// will be returned.
func BoolValue(b *bool) bool {
	if b == nil {
		return false
	}

	return *b
}

// Bool will return a pointer value of the given parameter.
//
//go:fix inline
func Bool(b bool) *bool {
	return new(b)
}

// StringValue will return a string value. If the pointer is nil, then an empty
// string will be returned.
func StringValue(str *string) string {
	if str == nil {
		return ""
	}

	return *str
}

// String will return a pointer value of the given parameter.
//
//go:fix inline
func String(str string) *string {
	return new(str)
}

// Int64 will return a pointer value of the given parameter.
//
//go:fix inline
func Int64(v int64) *int64 {
	return new(v)
}

// Int64Value will return an int64 value. If the pointer is nil, then zero will
// be returned.
func Int64Value(v *int64) int64 {
	if v == nil {
		return 0
	}

	return *v
}

// IntValue will return an int value. If the pointer is nil, zero will be
// returned.
func IntValue(v *int) int {
	if v == nil {
		return 0
	}

	return *v
}

// Int will return a pointer value of the given parameters.
//
//go:fix inline
func Int(v int) *int {
	return new(v)
}
