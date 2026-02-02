// This file is licensed under the terms of the MIT License (see LICENSE file)
// Copyright (c) 2026 Pavel Tsayukov p.tsayukov@gmail.com

package xerrors_test

import (
	"errors"
	"math/rand/v2"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ergosit/xerrors"
)

func Test_Must(t *testing.T) {
	t.Run("no panic", func(t *testing.T) {
		require.NotPanics(t, func() {
			t.Helper()
			val := rand.Int()
			assert.Equal(t, val, xerrors.Must(val, nil))
		})

		wrappedNoError := func(val int) (int, error) { return val, nil }
		require.NotPanics(t, func() {
			t.Helper()
			val := rand.Int()
			assert.Equal(t, val, xerrors.Must(wrappedNoError(val)))
		})
	})

	t.Run("panic", func(t *testing.T) {
		const errMsg = "some error"

		require.PanicsWithError(t, errMsg, func() {
			val := rand.Int()
			_ = xerrors.Must(val, errors.New(errMsg))
		})

		wrappedWithError := func(val int) (int, error) { return val, errors.New(errMsg) }
		require.PanicsWithError(t, errMsg, func() {
			val := rand.Int()
			_ = xerrors.Must(wrappedWithError(val))
		})
	})
}
