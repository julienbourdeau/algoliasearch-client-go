// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
)

// ExtractAroundRadius returns the first found AroundRadiusOption from the
// given variadic arguments or nil otherwise.
func ExtractAroundRadius(opts ...interface{}) *opt.AroundRadiusOption {
	for _, o := range opts {
		if v, ok := o.(*opt.AroundRadiusOption); ok {
			return v
		}
	}
	return nil
}
