package csp

// func backtrack(csp CSP, assignment Value) {
// 	e
// }

// func AC3[V Value, X Variable[V], C Constraint_t](csp CSP[V, X, C]) bool {
// 	var queue []Arc[V] = getArcs(csp.getVariables(), csp.getConstraints())

// 	for len(queue) > 0 {
// 		var x, y X
// 		constrained, queue = removeFirst(queue)
// 		if revise(csp, constrained) {
// 		}

// 	}

// 	return true
// }

// func revise[V Value, X Variable[V], C Constraint_t](csp CSP[V, X, C], xi X, xj X) bool {
// 	var revised = false

// 	for _, val := range xi.getDomain() {
// 		if !satisfiesConstraints(val, csp.getConstraintsOf(xi, xj)) {
// 			xi.removeVal(val)
// 			revised = true
// 		}
// 	}

// 	return revised
// }

// func satisfiesConstraints[Value V](val V, constraints []Constraint[C, V, X]) {

// }

// function AC-3(csp) returns false if an inconsistency is found and true otherwise
//  inputs: csp, a binary CSP with components (X, D, C)
//  local variables: queue, a queue of arcs, initially all the arcs in csp

//  while queue is not empty do
//    (Xi, Xj) ← REMOVE-FIRST(queue)
//    if REVISE(csp, Xi, Xj) then
//      if size of Di = 0 then return false
//      for each Xk in Xi.NEIGHBORS − {Xj} do
//       add(Xk, Xi) to queue
//  return true

// function REVISE(csp, Xi, Xj) returns true iff we revise the domain of Xi
//  revised ← false
//  for each x in Di do
//    if no value y in Dj allows (x, y) to satisfy the constraint between Xi and Xj then
//     delete x from Di
//     revised ← true
//  return revised
