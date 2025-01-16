// Solution to problem 2429 from leetcode
package leetcodetwofourtwonine

func minimizeXor(num1 int, num2 int) int {
	num1Bits := countOfBits(num1)
	num2Bits := countOfBits(num2)

	bits := abs(num1Bits - num2Bits)

	if num1Bits < num2Bits {
		num1 = incrementOfBits(num1, bits)
	} else if num1Bits > num2Bits {
		num1 = decrementOfBits(num1, bits)
	}

	return num1
}

func countOfBits(num int) int {
	count := 0

	for num != 0 {
		count += num & 1
		num >>= 1
	}

	return count
}

func decrementOfBits(num, bits int) int {
	for bits > 0 {
		num &= num - 1
		bits--
	}

	return num
}

func incrementOfBits(num, bits int) int {
	position := 0

	for bits > 0 {

		for (num>>position)&1 == 1 {
			position++
		}

		num |= 1 << position
		bits--
	}

	return num
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
