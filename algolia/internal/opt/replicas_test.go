// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestReplicas(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.ReplicasOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.Replicas([]string{}...),
		},
		{
			opts:     []interface{}{opt.Replicas("value1")},
			expected: opt.Replicas("value1"),
		},
		{
			opts:     []interface{}{opt.Replicas("value1", "value2", "value3")},
			expected: opt.Replicas("value1", "value2", "value3"),
		},
	} {
		var (
			in  = ExtractReplicas(c.opts...)
			out opt.ReplicasOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.ElementsMatch(t, c.expected.Get(), out.Get())
	}
}

func TestReplicas_CommaSeparatedString(t *testing.T) {
	for _, c := range []struct {
		payload  string
		expected *opt.ReplicasOption
	}{
		{
			payload:  `""`,
			expected: opt.Replicas([]string{}...),
		},
		{
			payload:  `"value1"`,
			expected: opt.Replicas("value1"),
		},
		{
			payload:  `"value1,value2,value3"`,
			expected: opt.Replicas("value1", "value2", "value3"),
		},
	} {
		var got opt.ReplicasOption
		err := json.Unmarshal([]byte(c.payload), &got)
		require.NoError(t, err)
		require.ElementsMatch(t, c.expected.Get(), got.Get())
	}
}
