// Solution to problem 1352 from leetcode
package leetcodeonethreefivetwo

type ProductOfNumbers struct {
	product []int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{
		product: []int{},
	}
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.product = []int{}
		return
	} else if num != 1 {
		this.multiply(num)
	}

	this.product = append(this.product, num)
}

func (this *ProductOfNumbers) multiply(num int) {
	alias := this.product
	for i := len(alias) - 1; i > -1; i-- {
		alias[i] *= num
	}
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	if len(this.product) < k {
		return 0
	}
	return this.product[len(this.product)-k]
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
