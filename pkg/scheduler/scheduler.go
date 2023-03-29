package scheduler

import (
	"math"
)

type Node struct {
	Name   string
	Labels map[string]string
}

type Scheduler struct {
	Node []Node
}

type SchedulingRule struct {
	Key                     string
	Value                   string
	Weight                  int32
	TopologySpreadPercetage int32
}

type SchedulingResult struct {
	Node         string
	ReplicaCount int32
}

func NewScheduler(nodes []Node) *Scheduler {
	return &Scheduler{
		Node: nodes,
	}
}

// CalculateTopologySpreadPercentage return scheduling information for pods
func (s *Scheduler) CalculateTopologySpreadPercentage(podreplicas int32, schedulingrule []SchedulingRule) ([]SchedulingResult, error) {
	// return empty SchedulingResult
	return []SchedulingResult{
		{
			Node:         "gke-app-regular-pool-123500",
			ReplicaCount: 0,
		},
	}, nil
}

func (s *Scheduler) SimpleCalculateTopologySpreadPercentage(podreplicas int32, p1, p2 int32) (int32, int32) {
	if podreplicas == 0 {
		return 0, 0
	} else if podreplicas == 1 {
		return 1, 0
	}
	node1 := int32(math.Round(float64(podreplicas) * float64(p1) / float64(100)))
	node2 := podreplicas - node1
	return node1, node2
}
