package csp

import (
	"Sudoku-CSP/util"
)

func Arc_reduce(x Variable, y Variable) (bool, Variable) {
	var change bool = false

	for _, x_value := range x.domain {
		if !y.DomainContainsOtherThan(x_value) {
			x = x.Remove(x_value)
			change = true
		}
	}

	return change, x
}

func (csp CSP) AC3() bool {
	var queue = []Arc{}
	for _, constraint := range csp.constraints {
		queue = append(queue, constraint.toArcs()...)
	}

	var arc Arc
	for len(queue) > 0 {
		arc, queue = util.RemoveLast(queue)

		if wasRevised, revised_var := Arc_reduce(csp.variables[arc.x1], csp.variables[arc.x2]); wasRevised {
			csp.variables[arc.x1] = revised_var
			if len(csp.variables[arc.x1].domain) == 0 {
				return false
			}

			var neighbors []int = csp.getNeighbors(arc.x1)
			for _, neighbor := range neighbors {
				var new_arc = Arc{neighbor, arc.x1}
				if !util.Contains(queue, new_arc) {
					queue = append(queue, new_arc)
				}
			}
		}
	}

	return true
}
