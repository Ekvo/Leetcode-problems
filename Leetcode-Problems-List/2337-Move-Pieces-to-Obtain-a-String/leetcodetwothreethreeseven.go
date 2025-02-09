// Solution to problem 2337 from leetcode
package leetcodetwothreethreeseven

// Queue
func canChange(start string, target string) bool {
	var (
		queueS = make([]*characterIndex, 0, len(start))
		queueT = make([]*characterIndex, 0, len(target))
	)
	for i := 0; i < len(start); i++ {
		if start[i] != '_' {
			queueS = append(queueS, &characterIndex{start[i], i})
		}
		if target[i] != '_' {
			queueT = append(queueT, &characterIndex{target[i], i})
		}
	}
	if len(queueS) != len(queueT) {
		return false
	}
	for len(queueS) > 0 {
		ptrS := queueS[0]
		ptrT := queueT[0]
		if ptrS.character() != ptrT.character() ||
			(ptrS.character() == 'L' && ptrS.index() < ptrT.index() ||
				ptrS.character() == 'R' && ptrS.index() > ptrT.index()) {
			return false
		}
		queueS = queueS[1:]
		queueT = queueT[1:]
	}
	return true
}

type characterIndex struct {
	ch byte
	id int
}

func (cI *characterIndex) index() int {
	return cI.id
}

func (cI *characterIndex) character() byte {
	return cI.ch
}

// Two Pointers
func canChange(start string, target string) bool {
	var (
		lenght = len(start)
		indexS = 0
		indexT = 0
	)
	for indexS < lenght || indexT < lenght {

		for indexS < lenght && start[indexS] == '_' {
			indexS++
		}
		for indexT < lenght && target[indexT] == '_' {
			indexT++
		}
		if indexS == lenght || indexT == lenght {
			break
		}
		if start[indexS] != target[indexT] ||
			(start[indexS] == 'L' && indexS < indexT ||
				start[indexS] == 'R' && indexS > indexT) {
			return false
		}
		indexS++
		indexT++
	}
	return indexS == lenght && indexT == lenght
}
