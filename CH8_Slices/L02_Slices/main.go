package main

import "errors"

// const (
// 	planFree = "free"
// 	planPro  = "pro"
// )

type Plan string

const (
	PlanFree Plan = "free"
	PlanPro  Plan = "pro"
)

// Changed below functions to use Plan type and slices
/*
func getMessageWithRetriesForPlan(plan Plan, messages []string) ([]string, error) {
	var s []string
	if plan == PlanFree {
		return messages[:2], nil
	} else if plan == PlanPro {
		return messages[:], nil
	} else {
		return s, errors.New("unsupported plan")
	}

}

func getMessageWithRetriesForPlan(plan Plan, messages []string) ([]string, error) {
	switch plan {
	case PlanFree:
		if len(messages) < 2 {
			return nil, errors.New("not enough messages for free plan")
		}
		return messages[:2], nil
	case PlanPro:
		return messages[:], nil
	default:
		return nil, errors.New("unsupported plan")
	}
}

*/

func getMessageWithRetriesForPlan(plan Plan, messages []string) ([]string, error) {
	var s []string
	switch plan {
	case PlanFree:
		if len(messages) < 2 {
			return nil, errors.New("not enough messages for free plan")
		}
		return messages[:2], nil
	case PlanPro:
		return messages, nil
	default:
		return s, errors.New("unsupported plan")
	}
}
