/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package balance

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"k8s.io/autoscaler/cluster-autoscaler/expander"
)

func TestBalanceExpander(t *testing.T) {
	e := NewStrategy()

	eo1a := expander.Option{
		Debug:     "EO1a",
		NodeCount: 5,
	}
	eo1b := expander.Option{
		Debug:     "EO1b",
		NodeCount: 2,
	}
	eo1c := expander.Option{
		Debug:     "EO1c",
		NodeCount: 5,
	}

	// when only one select it
	ret := e.BestOption([]expander.Option{eo1a}, nil)
	assert.Equal(t, *ret, eo1a)

	// test smaller should be picked
	ret = e.BestOption([]expander.Option{eo1a, eo1b}, nil)
	assert.True(t, assert.ObjectsAreEqual(*ret, eo1b))

	// equal keep first
	ret = e.BestOption([]expander.Option{eo1a, eo1c}, nil)
	assert.True(t, assert.ObjectsAreEqual(*ret, eo1a))

	// none
	ret = e.BestOption([]expander.Option{}, nil)
	assert.Nil(t, ret)
}
