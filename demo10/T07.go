package main

//func numberOfPairs(a []int, nums2 []int, diff int) (ans int64) {
//	for i, x := range nums2 {
//		a[i] -= x
//	}
//	b := append(sort.IntSlice{}, a...)
//	b.Sort()
//	var bt BIT = make([]int, len(a)+1)
//	for _, x := range a {
//		idx := b.Search(x + diff + 1) // 找到大于(x+diff)的左边界,之所以加1是避免了下面对+1越界处理
//		ans += int64(bt.query(idx))   // 看[1,idx]有多少个, 本来应该是[0,idx-1]的, dddd
//		bt.add(b.Search(x) + 1)       //和x一样映射同一个下标不影响结果的
//	}
//	return
//}
//
//type BIT []int
//
//func (b BIT) add(x int) {
//	for x < len(b) {
//		b[x]++
//		x += x & -x // 从右到左数到1, 把这部分加进来
//	}
//}
//
//func (b BIT) query(x int) (res int) {
//	for x > 0 {
//		res += b[x]
//		x &= x - 1 // 从右到左把最开始的1变0
//	}
//	return
//}
