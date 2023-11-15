package main

func main() {

}

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	inDegrees := make([]int, numCourses)

	for _, pr := range prerequisites {
		graph[pr[1]] = append(graph[pr[1]], pr[0])
		inDegrees[pr[0]]++
	}

	queue := []int{}
	for v, inDegree := range inDegrees {
		if inDegree == 0 {
			queue = append(queue, v)
		}
	}

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		numCourses--
		for _, w := range graph[v] {
			inDegrees[w]--
			if inDegrees[w] == 0 {
				queue = append(queue, w)
			}
		}
	}
	return numCourses == 0
}
