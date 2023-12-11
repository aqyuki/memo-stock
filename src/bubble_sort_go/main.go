package main

import "fmt"

func main() {
	src := []int{6, 8, 7, 5, 3, 4, 2, 1}
	dst := bubbleSort(src)

	fmt.Printf("src\t: %v\n", src)
	fmt.Printf("dst\t: %v\n", dst)
}

func bubbleSort(src []int) []int {
	if len(src) < 1 {
		return nil
	}
	dst := make([]int, len(src))
	copy(dst, src)

	// 配列のすべての要素を比較するためのループ
	for i := 0; i < len(dst)-1; i++ {

		// 実際に値を比べ移動させるためのループ
		for n := 0; n < len(dst)-1; n++ {
			// ここの条件式で昇順・降順を切り替えられる。
			if dst[n] < dst[n+1] {
				// swap
				tmp := dst[n]
				dst[n] = dst[n+1]
				dst[n+1] = tmp
			}
		}
	}
	return dst
}
