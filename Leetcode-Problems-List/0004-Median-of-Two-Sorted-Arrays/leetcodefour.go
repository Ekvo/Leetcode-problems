// Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.
//
// The overall run time complexity should be O(log (m+n)).
package leetcodefour

import "sort"

// Two Pointers
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var (
		lenght = len(nums1) + len(nums2)
		m1     = 0
		m2     = 0
	)
	for i1, i2 := 0, 0; i1+i2 <= lenght/2; {
		m2 = m1
		if i1 < len(nums1) && i2 < len(nums2) {
			if nums1[i1] < nums2[i2] {
				m1 = nums1[i1]
				i1++
			} else {
				m1 = nums2[i2]
				i2++
			}
		} else if i1 < len(nums1) {
			m1 = nums1[i1]
			i1++
		} else if i2 < len(nums2) {
			m1 = nums2[i2]
			i2++
		}
	}
	if lenght%2 != 0 {
		return float64(m1)
	}
	return float64(m1+m2) / 2.0
}

// Brute force
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var slice = append(nums1, nums2...)

	sort.Ints(slice)

	if len(slice) == 0 {
		return 0.0
	} else if len(slice)%2 != 0 {
		return float64(slice[len(slice)/2])
	}
	return float64(slice[len(slice)/2]+slice[len(slice)/2-1]) / 2.0
}
