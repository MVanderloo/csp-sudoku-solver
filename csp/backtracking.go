package csp

import "fmt"

type Assignment map[int]int

// func (a Assignment) contains(id int) bool {
// 	_, ok := a[id]
// 	return ok
// }

type Inference struct {
	variable     int
	domain_value int
}

func (csp CSP) BacktrackingSearch(forwardChecking bool, mrv bool, lcv bool) Assignment {
	return csp.Backtrack(make(Assignment), forwardChecking, mrv, lcv)
}

func (csp CSP) Backtrack(assignment Assignment, forwardChecking bool, mrv bool, lcv bool) Assignment {
	fmt.Println("assignment:", assignment)
	fmt.Println("isComplete(assignment) =", csp.isCompleteAssignment(assignment))
	if csp.isCompleteAssignment(assignment) {
		return assignment
	}

	var variable = csp.selectUnassignedVariable(assignment)
	fmt.Println("selecting variable", variable)
	for _, value := range csp.orderDomainValues(variable) {
		fmt.Printf("csp.isConsistent(value: %v, variable: %v, %v) = %v\n", value, variable, assignment, csp.isConsistent(value, variable, assignment))
		if csp.isConsistent(value, variable, assignment) {
			assignment[variable] = value
			if forwardChecking {
				var inferences []Inference = csp.forwardCheck(variable, assignment)
				if len(inferences) != 0 {
					csp.addInferences(inferences)
					result := csp.Backtrack(assignment, forwardChecking, mrv, lcv)
					if result != nil {
						return result
					}
					csp.removeInferences(inferences)
				}
			} else {
				result := csp.Backtrack(assignment, forwardChecking, mrv, lcv)
				if result != nil {
					return result
				}
			}
		}
	}
	return nil
}

// function BACKTRACK(sp, assignment) returns a solution or failure
// 	if assignment is complete then return assignment
// 	var = SELECT-UNASSIGNED-VARIABLE(csp, assignment)
// 	for each value in ORDER-DOMAIN-VALUES(csp, var, assignment) do
// 		if value is consistent with assignment then
// 			add {var = value} to assignment
// 			inferences = INFERENCE(csp, var, assignment)
// 			if inferences != failure then
// 				add inferences to csp
// 				result = BACKTRACK(csp, assignment)
// 				if result != failure then return result
// 				remove inferences from csp
// 			remove {var=value} from assignment
// 	return failure

func (csp CSP) isCompleteAssignment(assignment Assignment) bool {
	// all variables have a value that is in their domain
	for var_id, variable := range csp.variables {
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

func (csp CSP) selectUnassignedVariable(assignment Assignment) int {
	for variable := range csp.variables {
		if _, ok := assignment[variable]; !ok {
			return variable
		}
	}
	panic("All variables have been assigned")
}

// the book says this function takes assignment so do this later
func (csp CSP) orderDomainValues(variable int) []int {
	return csp.variables[variable].domain[:]
}

/**
 * Checks that none of the constraints have been violated by the assignment of a variable
 * Only checks constraints relating to that variable
 **/
func (csp CSP) isConsistent(value int, variable int, assignment Assignment) bool {

	for _, c := range csp.constraints {
		if c.constrains(variable) {
			switch c.getType() {
			case NOT_EQUALS:

				for _, cons_var := range c.constrained {
					assigned_value, isAssigned := assignment[cons_var]
					if cons_var != variable && isAssigned && assigned_value == value {
						// fmt.Println("IsConsistent check:", cons_var, variable, isAssigned, assigned_value, value)
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

				if (num_unassigned == 0 && sum == c.sum) || sum < c.sum {
					return false
				}
			}
		}
	}

	return true
}

func (csp CSP) forwardCheck(variable int, assignment Assignment) []Inference {
	return nil
}

func (csp CSP) addInferences(inferences []Inference) {
	for _, inference := range inferences {
		csp.RemoveFromDomain(inference.variable, inference.domain_value)
	}
}

func (csp CSP) removeInferences(inferences []Inference) {
	for _, inference := range inferences {
		csp.AddToDomain(inference.variable, inference.domain_value)
	}
}

/**
 * Returns the minimum sum of n non-repeating positive integers
 * Returns value of n(n+1)/2 for n 0-9 and otherwise calculates n(n+1)/2
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
	default:
		return n * (n + 1) / 2
	}
}
