package csp

import (
	"Sudoku-CSP/util"
)

/**
 * Binary constraint holds ids of 2 variables where v1 is constrained by v2
 * All arcs represent a NOT_EQUALS constraint
 **/
type Arc struct {
	x Variable
	y Variable
}

/**
 * Converts a constraint into an equivalent set of arcs for AC-3
 * The sum constraint will return an empty set as arc conversion is difficult
 **/
func (c Constraint) toArcs() []Arc {
	var arcs = []Arc{}
	switch c.getType() {
	case SUM:
		return arcs // not generating arcs for sum constraints

	case NOT_EQUALS:
		for i, x1 := range c.constrained {
			for j, x2 := range c.constrained {
				if i == j {
					continue
				}

				arcs = append(arcs, Arc{x1, x2})
			}
		}
		return arcs

	default:
		return arcs
	}
}

func Arc_reduce(x Domain, y Domain) (bool, Domain) {
	var change bool = false

	for _, x_value := range x {
		if !y.ContainsOtherThan(x_value) {
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

		if wasRevised, revised_var := Arc_reduce(csp.domains[arc.x], csp.domains[arc.y]); wasRevised {
			csp.domains[arc.x] = revised_var
			if len(csp.domains[arc.x]) == 0 {
				return false
			}

			var neighbors []Variable = csp.getNeighbors(arc.x)
			for _, neighbor := range neighbors {
				var new_arc = Arc{neighbor, arc.x}
				if !util.Contains(queue, new_arc) {
					queue = append(queue, new_arc)
				}
			}
		}
	}

	return true
}
