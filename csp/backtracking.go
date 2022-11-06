package csp

type Assignment map[int]int

func (a Assignment) contains(id int) bool {
	_, ok := a[id]
	return ok
}

func (csp CSP) BacktrackingSearch() Assignment {
	return csp.Backtrack(make(Assignment))
}

func (csp CSP) Backtrack(assignment Assignment) Assignment {
	if csp.isCompleteAssignment(assignment) {
		return assignment
	}

	var variable = csp.selectUnassignedVariable(assignment)
	for _, value := range csp.orderDomainValues(variable) {
		if csp.isConsistent(variable, assignment) {
			assignment[variable] = value
			inferences = csp.inference(variable, assignment)

			if inferences != nil {
				csp.addInferences(inferences)
				result = csp.Backtrack(assignment)

				if result != nil {
					return result
				}
			}

			delete(assignment, variable)
		}
	}
	return nil
}

func (csp CSP) isCompleteAssignment(assignment Assignment) bool {
	for id, variable := range csp.variables {
		if value, ok := assignment[id]; ok {
			if !variable.Contains(value) {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func (csp CSP) selectUnassignedVariable(assignment Assignment) int {
	for id, _ := range csp.variables {
		if _, ok := assignment[id]; !ok {
			return id
		}
	}
	panic("All variables have been assigned")
}

// the book says this function takes assignment so do this later
func (csp CSP) orderDomainValues(variable int) []int {
	return csp.variables[variable].domain[:]
}

func (csp CSP) isConsistent(variable int, assignment )

// function BACKTRACK(sp, assignment) returns a solution or failure
// 	if assignment is complete then return assignment
// 	var = SELECT-UNASSIGNED-VARIABLE(csp, assignment)
// 	for each value in ORDER-DOMAIN-VALUES(csp, var, assignment) do
// 		if value is consistent with assignment then
// 			add {var = value} to assignment
// 			inferences = INFERENCE(csp, var, assignment)
// 			if inferences != failure then
// 				add inferences to csp
// 				result = BACKTRACK(sp, assignment)
// 				if result != failure then return result
// 				remove inferences from csp
// 			remove {var=value} from assignment
// 	return failure
