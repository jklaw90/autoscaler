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
	"math"

	"k8s.io/autoscaler/cluster-autoscaler/expander"
	schedulerframework "k8s.io/kubernetes/pkg/scheduler/framework"
)

type balance struct {
}

// NewStrategy returns an expansion strategy that picks the smallest between node groups
func NewStrategy() expander.Strategy {
	return &balance{}
}

// BestOptions selects from the expansion options at balance
func (r *balance) BestOptions(expansionOptions []expander.Option, nodeInfo map[string]*schedulerframework.NodeInfo) []expander.Option {
	best := r.BestOption(expansionOptions, nodeInfo)
	if best == nil {
		return nil
	}
	return []expander.Option{*best}
}

// BestOption selects from the expansion options at balance
func (r *balance) BestOption(expansionOptions []expander.Option, nodeInfo map[string]*schedulerframework.NodeInfo) *expander.Option {
	if len(expansionOptions) <= 0 {
		return nil
	}

	smallestNodeCount := math.MaxInt
	var choice *expander.Option
	for i := 0; i < len(expansionOptions); i++ {
		opt := expansionOptions[i]

		// skip 0 count node groups
		if opt.NodeCount == 0 {
			continue
		}

		if opt.NodeCount < smallestNodeCount {
			smallestNodeCount = opt.NodeCount
			choice = &opt
		} else if opt.NodeCount == smallestNodeCount {
			// if equal node count pick the group with more pods
			// TODO could improve with some kind of calculations with aws or check cpu/mem
			if len(opt.Pods) > len(choice.Pods) {
				choice = &opt
			}
		}
	}

	return choice
}
