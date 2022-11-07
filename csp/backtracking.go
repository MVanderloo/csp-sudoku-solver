package csp

import (
	"fmt"
	"time"
)

type Assignment map[Variable]int

type Inference struct {
	variable     Variable
	domain_value int
}

func (csp CSP) BacktrackingSearch(ac3 bool, forwardChecking bool, mrv bool, lcv bool) Assignment {
	start := time.Now()
	if ac3 {
		csp.AC3()
	}
	var res Assignment = csp.Backtrack(make(Assignment), forwardChecking, mrv, lcv)
	duration := time.Since(start)
	fmt.Println("Backtracking search:", duration.Milliseconds(), "(ms)")
	return res
}

func (csp CSP) Backtrack(assignment Assignment, forwardChecking bool, mrv bool, lcv bool) Assignment {
	if csp.isComplete(assignment) {
		return assignment
	}

	var variable = csp.selectUnassignedVariable(assignment, mrv)
	for _, value := range csp.orderDomainValues(assignment, variable, lcv) {
		assignment[variable] = value
		if csp.isConsistent(variable, assignment) {
			if forwardChecking {
				var inferences []Inference = csp.forwardCheck(variable, assignment)
				csp.addInferences(inferences)
				result := csp.Backtrack(assignment, forwardChecking, mrv, lcv)
				if result != nil {
					return result
				}
				csp.removeInferences(inferences)
			} else {
				result := csp.Backtrack(assignment, forwardChecking, mrv, lcv)
				if result != nil {
					return result
				}
			}
		}
		delete(assignment, variable)
	}

	return nil
}

func (csp CSP) isComplete(assignment Assignment) bool {
	// all variables have a value that is in their domain
	for var_id, variable := range csp.domains {
		if value, ok := assignment[var_id]; ok {
			if false && !variable.Contains(value) {
				return false
			}
		} else {
			return false
		}
	}

	//all constraints are satisifed
	// for _, constraint := range csp.constraints {
	// 	if !constraint.isSatisfied(assignment) {
	// 		return false
	// 	}
	// }

	return true
}

func (csp CSP) selectUnassignedVariable(assignment Assignment, mrv bool) Variable {
	if mrv {
		var min_values int
		var min_var Variable = -1
		for _, variable := range csp.variables {
			if _, ok := assignment[variable]; !ok {
				if min_var == -1 {
					min_values = len(csp.domains[variable])
					min_var = variable
				}
				if trial := len(csp.domains[variable]); trial < min_values {
					min_values = trial
					min_var = variable
				}

				if min_values == 1 {
					return min_var
				}
			}
		}
		return min_var
	} else {
		for variable := range csp.domains {
			if _, ok := assignment[variable]; !ok {
				return variable
			}
		}
		panic("All variables have been assigned")
	}
}

func (csp CSP) orderDomainValues(assignment Assignment, variable Variable, lcv bool) Domain {
	if !lcv {
		return csp.domains[variable]
	}
	var domain Domain = csp.domains[variable]
	var num_constraints = make([]int, len(domain))
	for i, val := range domain {
		for _, neighbor := range csp.getNeighbors(variable) {
			if _, ok := assignment[neighbor]; !ok && csp.domains[neighbor].Contains(val) {
				num_constraints[i] = num_constraints[i] + 1
			}
		}
	}

	for i := 0; i < len(domain); i++ {
		for j := i + 1; j < len(domain); j++ {
			if num_constraints[j] < num_constraints[i] {
				var temp = domain[i]
				domain[i] = domain[j]
				domain[j] = temp
			}
		}
	}

	return domain
}

/**
 * Checks that none of the constraints have been violated by the assignment of a variable
 * Only checks constraints relating to that variable
 **/
func (csp CSP) isConsistent(variable Variable, assignment Assignment) bool {
	for _, c := range csp.constraints {
		if c.constrains(variable) {
			switch c.getType() {
			case NOT_EQUALS:
				for _, cons_var := range c.constrained {
					assigned_value, isAssigned := assignment[cons_var]
					if cons_var != variable && isAssigned && assigned_value == assignment[variable] {
						return false
					}
				}

			case SUM:
				var sum int = 0
				var num_unassigned int = 0

				for _, cons_var := range c.constrained {
					assigned_value, isAssigned := assignment[cons_var]

					if isAssigned {
						sum += assigned_value
						if sum > c.sum {
							return false
						}
					} else {
						num_unassigned++
					}
				}

				if (num_unassigned == 0 && sum != c.sum) || sum > c.sum-min_sum(num_unassigned) {
					return false
				}
			}
		}
	}

	return true
}

/**
 * Evaluate domain values that could be removed and return them
 **/
func (csp CSP) forwardCheck(variable Variable, assignment Assignment) []Inference {
	var inferences = []Inference{}

	for _, neighbor := range csp.getNeighbors(variable) {
		if csp.domains[neighbor].Contains(assignment[variable]) {
			inferences = append(inferences, Inference{
				variable:     neighbor,
				domain_value: assignment[variable],
			})
		}
	}

	return inferences[:]
}

/**
 * Take the list of Inferences and apply to the CSP
 */
func (csp CSP) addInferences(inferences []Inference) {
	for _, inference := range inferences {
		csp.removeFromDomain(inference.variable, inference.domain_value)
	}
}

/**
 * Take the list of Inferences and remove from the CSP
 */
func (csp CSP) removeInferences(inferences []Inference) {
	for _, inference := range inferences {
		csp.addToDomain(inference.variable, inference.domain_value)
	}
}

/**
 * Returns the minimum sum of n non-repeating positive integers
 * Returns value of n(n+1)/2 for n 0-9 and otherwise calculates n(n+1)/2
 * used to prove the inconsistency of sum constraints which have unassigned variables
 */
func min_sum(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 3
	case 3:
		return 6
	case 4:
		return 10
	case 5:
		return 15
	case 6:
		return 21
	case 7:
		return 28
	case 8:
		return 36
	case 9:
		return 25
	default: // won't be called but exists for completeness
		return n * (n + 1) / 2
	}
}
