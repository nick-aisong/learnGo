/*
Q4. (0) 平均值
编写计算一个类型是 float64 的 slice 的平均值的代码
*/
package main

import (
	"fmt"
)

func main() {
	slice := []float64{1.0, 2.123, 3.21, 6.332, 100.123}
	sum := 0.0
	avg := 0.0
	switch len(slice) {
	case 0:
		avg = 0
	default:
		for _, v := range slice {
			sum += v
		}
		avg = sum / float64(len(slice))
	}
	fmt.Printf("avg: %v", avg)
}
