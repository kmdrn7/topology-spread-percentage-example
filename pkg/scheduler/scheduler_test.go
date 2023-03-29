package scheduler

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScheduler_CalculateTopologySpreadPercentage(t *testing.T) {
	type field struct {
		name              string
		replicacount      int32
		schedulingrules   []SchedulingRule
		schedulingresults []SchedulingResult
	}
	tests := []field{
		{
			name:         "0 replicas",
			replicacount: 0,
			schedulingrules: []SchedulingRule{
				{
					Key:                     "role",
					Value:                   "gke-app-regular-pool-123500",
					Weight:                  10,
					TopologySpreadPercetage: 25,
				},
				{
					Key:                     "role",
					Value:                   "gke-app-spot-pool-123500",
					Weight:                  1,
					TopologySpreadPercetage: 75,
				},
			},
			schedulingresults: []SchedulingResult{
				{
					Node:         "gke-app-regular-pool-123500",
					ReplicaCount: 0,
				},
			},
		},
	}

	nodes := []Node{
		{
			Name: "gke-app-regular-pool-123500",
			Labels: map[string]string{
				"role":        "gke-app-regular-pool-123500",
				"preemptible": "false",
			},
		},
		{
			Name: "gke-app-spot-pool-123500",
			Labels: map[string]string{
				"role":        "gke-app-spot-pool-123500",
				"preemptible": "true",
			},
		},
	}

	s := NewScheduler(nodes)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			schedulingresults, err := s.CalculateTopologySpreadPercentage(tt.replicacount, tt.schedulingrules)
			if err != nil {
				t.Errorf(err.Error())
			}
			if eq := assert.ObjectsAreEqualValues(schedulingresults, tt.schedulingresults); !eq {
				log.Println("expected and returned results are not equal, expected vs returned:")
				log.Printf("expected: %v \n", tt.schedulingresults)
				log.Printf("returned: %v \n", schedulingresults)
				t.Fail()
			}
		})
	}
}
