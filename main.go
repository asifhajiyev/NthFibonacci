package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	var option int
menuLoop:
	for ok := true; ok; ok = option <= 3 {
		showOptionMenu()
		scanner.Scan()
		option, _ := strconv.Atoi(scanner.Text())
	inputLoop:
		for true {
			if option > 4 || option < 1 {
				break inputLoop
			} else if option == 4 {
				break menuLoop
			}
			var input, result int
			input = getInputN()
			if input == 0 {
				break inputLoop
			} else if input == -1 {
				continue
			}
			switch option {
			case 1:
				result = fibonacciUsingRecursion(input)
			case 2:
				result = fibonacciUsingDynamicProgramming(input)
			case 3:
				result = fibonacciUsingSpaceOptimization(input)
			}
			fmt.Println(input, "number fib:", result)
		}
	}
}

func showOptionMenu() {
	fmt.Println("Choose method number to find")
	fmt.Println(`
1. Recursion 
2. Using Dynamic Programming 
3. Space optimized way
4. Exit`)
}

func getInputN() int {
	fmt.Print("Enter n to find nth fibonacci number: ")
	scanner.Scan()
	s := scanner.Text()

	if s == "exit" {
		return 0
	}
	input, err := validateInput(s)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return input
}

func validateInput(inputAsString string) (int, error) {
	inputAsNumber, err := strconv.Atoi(inputAsString)
	if err != nil {
		return -1, errors.New("non numeric character")
	}
	if inputAsNumber <= 0 {
		return -1, errors.New("not defined")
	}
	return inputAsNumber, nil
}

func fibonacciUsingRecursion(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacciUsingRecursion(n-1) + fibonacciUsingRecursion(n-2)
}

func fibonacciUsingDynamicProgramming(n int) int {
	arr := make([]int, n+2)
	arr[0] = 0
	arr[1] = 1
	for i := 2; i < n+2; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[n]
}

func fibonacciUsingSpaceOptimization(n int) int {
	a, b := 0, 1
	var c int
	for i := 2; i <= n; i++ {
		c = a + b
		a = b
		b = c
	}
	return b
}
