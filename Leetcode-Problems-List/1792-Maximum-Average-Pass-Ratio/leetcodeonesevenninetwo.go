// Solution to problem 1792 from leetcode
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
