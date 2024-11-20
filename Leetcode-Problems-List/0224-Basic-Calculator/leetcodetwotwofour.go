// Given a string s representing a valid expression,
// implement a basic calculator to evaluate it, and return the result of the evaluation.
//
// Note: You are not allowed to use any built-in function
// which evaluates strings as mathematical expressions, such as eval().
package leetcodetwotwofour

import (
	"strconv"
	"unicode"
)

const (
	newExpression = '('
	endExpression = ')' //for verification
	unaryMinus    = '-'
)

type compute struct {
	data  string
	index int
}

// calculate - main function
func calculate(s string) int {
	var compute = newCalculate(s)
	return compute.expression()
}

func newCalculate(s string) *compute {
	return &compute{
		data:  s,
		index: 0,
	}
}

func (c *compute) work() bool {
	return -1 < c.index && c.index < len(c.data)
}

func (c *compute) increment() {
	c.index++
}

func (c *compute) decrement() {
	c.index--
}

func (c *compute) newSimbol() byte {
	var ch uint8 = ' '

	for c.work() && ch == ' ' {
		ch = c.data[c.index]
		c.increment()
	}

	return ch
}

func (c *compute) newValue() int {
	var (
		start  = c.index
		number = 0
	)

	for c.work() && unicode.IsDigit(rune(c.data[c.index])) {
		c.increment()
	}

	number, _ = strconv.Atoi(c.data[start:c.index])

	return number
}

func (c *compute) expression() int {
	var val int = c.prime()

	for c.work() {

		switch c.newSimbol() {
		case '+':
			val += c.prime()
		case '-':
			val -= c.prime()
		default:
			c.decrement()
			return val
		}

	}

	return val
}

func (c *compute) prime() int {
	var val = 0

	switch c.newSimbol() {
	case newExpression:
		val = c.expression()
		c.newSimbol() // find ')'
	case unaryMinus:
		val = -c.prime()
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		c.decrement() // unget
		val = c.newValue()
	}

	return val
}
