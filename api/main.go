package main

import (
	"fmt"
	"strings"

	"github.com/fvbommel/sortorder"
	"github.com/hookenz/mytest/pkg/libhello"
)

type (
	Pair struct {
		Name  string `json:"name" example:"name"`
		Value string `json:"value" example:"value"`
	}
	InstanceType struct {
		Pair
	}
)

type InstanceTypeByName []InstanceType

func (t InstanceTypeByName) Len() int {
	return len(t)
}

func (t InstanceTypeByName) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t InstanceTypeByName) Less(i, j int) bool {

	// Put General Purpose instance types first
	if strings.HasPrefix(t[i].Name, "General Purpose") {
		if !strings.HasPrefix(t[j].Name, "General Purpose") {
			return true
		}
	} else if strings.HasPrefix(t[j].Name, "General Purpose") {
		return false
	}

	return sortorder.NaturalLess(t[i].Name, t[j].Name)
}

func main() {
	types := []InstanceType{
		InstanceType{
			Pair: Pair{
				Name:  "General Purpose",
				Value: "t3.medium",
			},
		},
		InstanceType{
			Pair: Pair{
				Name:  "Compute Optimized",
				Value: "c5.large",
			},
		},
		InstanceType{
			Pair: Pair{
				Name:  "Memory Optimized",
				Value: "r5.large",
			},
		},
	}

	fmt.Printf("Types: %v\n", types)
	libhello.Hello()
}
