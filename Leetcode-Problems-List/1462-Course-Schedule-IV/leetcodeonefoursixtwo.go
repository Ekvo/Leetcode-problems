// Solution to problem 1462 from leetcode
package leetcodeonefoursixtwo

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	directions := make(map[int][]int)

	for _, p := range prerequisites {
		directions[p[0]] = append(directions[p[0]], p[1])
	}

	endCourse := -1

	var findCourse func(visited []bool, current, countCourses int) bool

	findCourse = func(visited []bool, current, countCourses int) bool {
		if countCourses == 0 {
			return false
		}

		if current == endCourse {
			return true
		}
		visited[current] = true

		for _, direct := range directions[current] {

			if !visited[direct] && findCourse(visited, direct, countCourses-1) {
				return true
			}
		}

		return false
	}

	result := make([]bool, len(queries))

	for i, q := range queries {
		endCourse = q[1]
		visited := make([]bool, numCourses)
		result[i] = findCourse(visited, q[0], numCourses)
	}

	return result
}
