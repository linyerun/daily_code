package main

// BubbleSort 做一个冒泡排序
func BubbleSort[E int | uint](a []E) {
	//第一次，去到最后一位
	//最后一次：去到1
	for i := len(a) - 1; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			if a[j-1] > a[j] {
				swap(j-1, j, a)
			}
		}
	}
}

func swap[E int | uint](i int, j int, a []E) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	a := []int{3, 2, 2, 1, 10, 1, -10}
	BubbleSort(a)
	for i := 0; i < len(a); i++ {
		println(a[i])
	}
}
