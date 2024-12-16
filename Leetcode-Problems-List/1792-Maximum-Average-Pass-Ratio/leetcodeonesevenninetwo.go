// There is a school that has classes of students and each class will be having a final exam.
// You are given a 2D integer array classes, where classes[i] = [passi, totali].
// You know beforehand that in the ith class, there are totali total students,
// but only passi number of students will pass the exam.
//
// You are also given an integer extraStudents.
// There are another extraStudents brilliant students that are guaranteed to
// pass the exam of any class they are assigned to.
// You want to assign each of the extraStudents students to a class in a way
// that maximizes the average pass ratio across all the classes.
//
// The pass ratio of a class is equal to the number of students of the class that will
// pass the exam divided by the total number of students of the class.
// The average pass ratio is the sum of pass ratios of all the classes divided by the number of the classes.
//
// Return the maximum possible average pass ratio after assigning the extraStudents students.
// Answers within 10-5 of the actual answer will be accepted.
package leetcodeonesevenninetwo

import "container/heap"

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	var (
		arrSchool = make(School, 0, len(classes))

		//looking for the maximum score increase
		findAvrDif = func(i int) float64 {
			avrD := float64(classes[i][0]+1)/float64(classes[i][1]+1) -
				float64(classes[i][0])/float64(classes[i][1])
			return avrD
		}
	)
	for i := 0; i < len(classes); i++ {
		arrSchool = append(arrSchool, &IndexClass{i, findAvrDif(i)})
	}
	heap.Init(&arrSchool)

	for i := 0; i < extraStudents; i++ {
		classes[arrSchool[0].Index][0]++
		classes[arrSchool[0].Index][1]++
		arrSchool[0].AverageDifference = findAvrDif(arrSchool[0].Index)
		heap.Fix(&arrSchool, 0)
	}
	averagePass := 0.0
	for i := 0; i < len(classes); i++ {
		averagePass += float64(classes[i][0]) / float64(classes[i][1])
	}
	return averagePass / float64(len(classes))
}

func (s School) Len() int      { return len(s) }
func (s School) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s School) Less(i, j int) bool {
	return s[i].AverageDifference > s[j].AverageDifference
}
func (s *School) Push(v any) { *s = append(*s, v.(*IndexClass)) }
func (s *School) Pop() any {
	old := *s
	lenght := len(old)
	elem := old[lenght-1]
	*s = old[:lenght-1]
	return elem
}

type School []*IndexClass

type IndexClass struct {
	Index             int
	AverageDifference float64
}
