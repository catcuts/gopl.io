// 练习 1.3： 做实验测量潜在低效的版本和使用了strings.Join的版本的运行时间差异。
//（1.6节讲解了部分time包，11.4节展示了如何写标准测试程序，以得到系统性的性能评测。）

package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	s, sep := "", " "
	l := []string{"a", "b"}
	flag := 0
	timeSpend := float64(0)

	seq := []string{}

	for i := 0; i < 100000; i++ {
		timeStart := time.Now()
		s += sep + l[flag]
		timeSpend += time.Since(timeStart).Seconds()

		seq = append(seq, l[flag])
		if flag == 0 {
			flag = 1
		} else {
			flag = 0
		}
	}
	fmt.Printf("version '+='          costs: %.2fs\n", timeSpend)

	timeStart := time.Now()
	s = strings.Join(seq, sep)
	timeSpend = time.Since(timeStart).Seconds()
	fmt.Printf("version 'string.join' costs: %.2fs\n", timeSpend)
}
