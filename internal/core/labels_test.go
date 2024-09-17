package core

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMergeCustomLabels(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		dst := map[string]string{"A": "1", "B": "2"}
		src := map[string]string{"B": "X", "C": "3"}

		err := MergeCustomLabels(dst, src)
		require.NoError(t, err)
		require.Equal(t, map[string]string{"A": "1", "B": "X", "C": "3"}, dst)
	})

	t.Run("invalid-prefix", func(t *testing.T) {
		dst := map[string]string{"A": "1", "B": "2"}
		src := map[string]string{"B": "X", LabelLang: "go"}

		err := MergeCustomLabels(dst, src)

		require.EqualError(t, err, `cannot use prefix "org.testcontainers" for custom labels`)
		require.Equal(t, map[string]string{"A": "1", "B": "X"}, dst)
	})

	t.Run("nil destination and empty source", func(t *testing.T) {
		src := map[string]string{}

		err := MergeCustomLabels(nil, src)

		require.NoError(t, err)
	})

	t.Run("nil-destination", func(t *testing.T) {
		src := map[string]string{"A": "1"}
		err := MergeCustomLabels(nil, src)
		require.Error(t, err)
	})
}
