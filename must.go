// This file is licensed under the terms of the MIT License (see LICENSE file)
// Copyright (c) 2026 Pavel Tsayukov p.tsayukov@gmail.com

package xerrors

// Must returns its first passed argument or panics if the second one is non-nil
// error. It can wrap a call to a function that returns (T, error).
func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}
