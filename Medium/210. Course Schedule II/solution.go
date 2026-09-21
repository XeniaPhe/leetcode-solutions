package course_schedule_II

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph, state, order := make([][]int, numCourses), make([]byte, numCourses), make([]int, 0, numCourses)
    for p, i := prerequisites, 0; i < len(p); graph[p[i][0]], i = append(graph[p[i][0]], p[i][1]), i + 1 { }
    for stack, i := make([][2]int, 0), 0; i < numCourses; i += 1 {
        if state[i] == 0 {
        	dfs:
        	for stack, state[i] = append(stack, [2]int{i, 0}), 1; len(stack) > 0; {
				list, idx := graph[stack[len(stack) - 1][0]], len(stack) - 1
        	    for j := stack[idx][1]; j < len(list); j += 1 {
        	        if state[list[j]] == 0 {
        	            stack, stack[idx][1], state[list[j]] = append(stack, [2]int{list[j], 0}), j + 1, 1
        	            continue dfs
					} else if state[list[j]] == 1 {
        	            return nil
        	        }
        	    }

        	    order, stack, state[stack[idx][0]] = append(order, stack[idx][0]), stack[:len(stack) - 1], 2
        	}
		}
    }

    return order
}