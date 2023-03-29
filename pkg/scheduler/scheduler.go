package scheduler

import (
	"math"
)

type Scheduler struct{}

func (s *Scheduler) CalculateTopologySpreadPercentage(podreplicas int32, p1, p2 int32) (int32, int32) {
	if podreplicas == 0 {
		return 0, 0
	} else if podreplicas == 1 {
		return 1, 0
	}
	node1 := int32(math.Round(float64(podreplicas) * float64(p1) / float64(100)))
	node2 := podreplicas - node1
	return node1, node2
}
