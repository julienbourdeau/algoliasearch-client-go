// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestAnchoring(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.AnchoringOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.Anchoring(""),
		},
		{
			opts:     []interface{}{opt.Anchoring("")},
			expected: opt.Anchoring(""),
		},
		{
			opts:     []interface{}{opt.Anchoring("content of the string value")},
			expected: opt.Anchoring("content of the string value"),
		},
	} {
		var (
			in  = ExtractAnchoring(c.opts...)
			out opt.AnchoringOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, *c.expected, out)
	}
}
