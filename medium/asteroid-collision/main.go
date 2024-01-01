package main

import (
	"fmt"
)

func main() {
	fmt.Println(asteroidCollision([]int{5, 10, -5}), "[5, 10]")
	fmt.Println(asteroidCollision([]int{8, -8}), "[]")
	fmt.Println(asteroidCollision([]int{10, 2, -5}), "[10]")
	fmt.Println(asteroidCollision([]int{10, 2, -50}), "[-50]")
	fmt.Println(asteroidCollision([]int{-2, -1, 1, 2}), "[-2, -1, 1, 2]")
	fmt.Println(asteroidCollision([]int{-2, -2, 1, -2}), "[-2, -2, -2]")
	fmt.Println(asteroidCollision([]int{1, -1, -2, -2}), "[-2,-2]")
}

func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0, len(asteroids))
	stack = append(stack, asteroids[0])

	for _, asteroid := range asteroids[1:] {
		if len(stack) == 0 {
			stack = append(stack, asteroid)
			continue
		}
		for len(stack) > 0 {
			curAsteroid := stack[len(stack)-1]
			if curAsteroid > 0 && asteroid > 0 || curAsteroid < 0 && asteroid < 0 {
				stack = append(stack, asteroid)
				break
			}

			if curAsteroid < 0 && asteroid > 0 {
				stack = append(stack, asteroid)
				break
			}

			if curAsteroid > asteroid*-1 {
				break
			}
			if curAsteroid == asteroid*-1 {
				stack = stack[:len(stack)-1]
				break
			}

			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, asteroid)
				break
			}
		}
	}

	return stack
}
