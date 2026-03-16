// This file is licensed under the terms of the MIT License (see LICENSE file)
// Copyright (c) 2026 Pavel Tsayukov p.tsayukov@gmail.com

package xerrors

// Tail ignores the first argument and returns the second argument, which is an error.
// It can wrap a call to a function that returns (T, error).
func Tail[T any](_ T, err error) error { return err }

// Tail2 ignores the first and second arguments and returns the third argument,
// which is an error. It can wrap a call to a function that returns (T1, T2, error).
func Tail2[T1, T2 any](_ T1, _ T2, err error) error { return err }
