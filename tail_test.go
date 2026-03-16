// This file is licensed under the terms of the MIT License (see LICENSE file)
// Copyright (c) 2026 Pavel Tsayukov p.tsayukov@gmail.com

package xerrors_test

import (
	"errors"
	"math/rand/v2"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ergosit/xerrors"
)

func Test_Tail(t *testing.T) {
	require.EqualError(t, xerrors.Tail(rand.Int(), errors.New(errMsg)), errMsg)
	require.EqualError(t,
		xerrors.Tail(func() (string, error) {
			return "Hello, world!", errors.New(errMsg)
		}()),
		errMsg,
	)
}

func Test_Tail2(t *testing.T) {
	require.EqualError(t, xerrors.Tail2(rand.Int(), rand.Float64(), errors.New(errMsg)), errMsg)
	require.EqualError(t,
		xerrors.Tail2(func() (string, *int, error) {
			return "Hello, world!", nil, errors.New(errMsg)
		}()),
		errMsg,
	)
}
