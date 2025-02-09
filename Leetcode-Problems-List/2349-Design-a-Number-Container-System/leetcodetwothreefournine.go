// Solution to problem 2349 from leetcode
package leetcodetwothreefournine

import "sort"

/**
 * Your NumberContainers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Change(index,number);
 * param_2 := obj.Find(number);
 */

type store struct {
	addr     []int
	addrSort bool
}

type NumberContainers struct {
	indexNumber map[int]int
	numberIndex map[int]store
}

func Constructor() NumberContainers {
	return NumberContainers{
		indexNumber: make(map[int]int),
		numberIndex: make(map[int]store),
	}
}

func (this *NumberContainers) Change(index int, number int) {
	if num, ex := this.indexNumber[index]; ex {
		if num == number {
			return
		}
		this.remove(index, num)
	}
	this.add(index, number)
}

func (this *NumberContainers) add(index int, number int) {
	this.indexNumber[index] = number
	this.numberIndex[number] = store{
		addr:     append(this.numberIndex[number].addr, index),
		addrSort: false,
	}
}

func (this *NumberContainers) remove(index int, number int) {
	s := this.numberIndex[number]

	for i := 0; i < len(s.addr); i++ {
		if s.addr[i] == index {
			s.addr = append(s.addr[:i], s.addr[i+1:]...)

			if len(s.addr) < 1 {
				delete(this.numberIndex, number)
			} else {
				this.numberIndex[number] = s
			}
			break
		}
	}
	delete(this.indexNumber, index)
}

func (this *NumberContainers) Find(number int) int {
	if s, ex := this.numberIndex[number]; ex {
		if !s.addrSort {
			sort.Ints(s.addr)
			this.numberIndex[number] = store{
				addr:     s.addr,
				addrSort: true,
			}
		}

		return s.addr[0]
	}

	return -1
}
