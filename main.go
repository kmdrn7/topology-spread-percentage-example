package main

import (
	"log"

	"github.com/kmdrn7/topology-spread-percentage-example/pkg/scheduler"
)

func main() {
	s := &scheduler.Scheduler{}
	var podreplicas, p1, p2, node1, node2 int32
	for i := 0; i < 20; i++ {
		podreplicas, p1, p2 = int32(i), 90, 10
		node1, node2 = s.CalculateTopologySpreadPercentage(podreplicas, p1, p2)
		log.Printf("podreplicas:%d, p1:%d, p2:%d, node1:%d, node2:%d \n", podreplicas, p1, p2, node1, node2)
	}
}
