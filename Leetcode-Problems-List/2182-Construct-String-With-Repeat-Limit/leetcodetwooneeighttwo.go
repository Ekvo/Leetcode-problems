// You are given a string s and an integer repeatLimit.
// Construct a new string repeatLimitedString using the characters of s such that
// no letter appears more than repeatLimit times in a row.
// You do not have to use all characters from s.
//
// Return the lexicographically largest repeatLimitedString possible.
//
// A string a is lexicographically larger than a string b if in the first position
// where a and b differ, string a has a letter that appears later in the alphabet
// than the corresponding letter in b.
// If the first min(a.length, b.length) characters do not differ,
// then the longer string is the lexicographically larger one.
package leetcodetwooneeighttwo

import (
	"bytes"
	"container/heap"
)

const unlimit = '~'

var curSimbol byte = unlimit

func repeatLimitedString(s string, repeatLimit int) string {
	var simbolsCount = make([]int, 26)
	for i := 0; i < len(s); i++ {
		simbolsCount[s[i]-'a']++
	}
	var lC = make(letters, 0, 26)
	for i, count := range simbolsCount {
		if count != 0 {
			lC = append(lC, &simbolCount{byte(i + 'a'), count})
		}
	}
	heap.Init(&lC)
	var buff bytes.Buffer
	buff.Grow(len(s))

	for lC.Len() > 0 {
		sC := heap.Pop(&lC).(*simbolCount)
		loop := repeatLimit
		if curSimbol != unlimit && curSimbol > sC.ch {
			loop = 1
		}
		for j := 0; j < loop && sC.count > 0; j++ {
			buff.WriteByte(sC.ch)
			sC.count--
		}
		if sC.count > 0 && lC.Len() != 0 {
			curSimbol = sC.ch
			heap.Push(&lC, sC)
		} else {
			curSimbol = unlimit
		}
	}
	return buff.String()
}

type letters []*simbolCount

type simbolCount struct {
	ch    byte
	count int
}

func (l letters) Len() int      { return len(l) }
func (l letters) Swap(i, j int) { l[i], l[j] = l[j], l[i] }
func (l letters) Less(i, j int) bool {
	if l[i].ch == curSimbol && l[i].ch > l[j].ch {
		return !(l[i].ch > l[j].ch)
	} else if l[j].ch == curSimbol && l[i].ch < l[j].ch {
		return !(l[i].ch < l[j].ch)
	}
	return l[i].ch > l[j].ch
}
func (l *letters) Push(obj any) { *l = append(*l, obj.(*simbolCount)) }
func (l *letters) Pop() any {
	old := *l
	lenght := len(old)
	elem := old[lenght-1]
	*l = old[:lenght-1]
	return elem
}
