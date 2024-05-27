package main

import (
	"fmt"
	"math"
)

func main() {
	var count int
	result := 0
	fmt.Scanln(&count)
	building := make([]int, count)
	visible := make([]int, count)

	for i := 0; i < count; i++ {
		fmt.Scan(&building[i])
	}

	for i := 0; i < count; i++ {
		// 왼쪽 방향으로 보이는 빌딩, 기울기 작아야 보임
		lastInclination := math.Inf(1)
		for left := i - 1; left >= 0; left-- {
			inclination := float64(building[i]-building[left]) / float64(i-left)
			if inclination < lastInclination {
				visible[i]++
				lastInclination = inclination
			}
		}

		// 오른쪽 방향으로 보이는 빌딩
		lastInclination = math.Inf(-1)
		for right := i + 1; right < count; right++ {
			inclination := float64(building[right]-building[i]) / float64(right-i)
			if inclination > lastInclination {
				visible[i]++
				lastInclination = inclination
			}
		}

		if result < visible[i] {
			result = visible[i]
		}
	}
	fmt.Println(result)
}
