package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
	"sync"
)

type Result struct {
	operation string
	value     string
}

func calculate(a, b *big.Int, operation string, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	var result *big.Int
	var resultStr string

	switch operation {
	case "add":
		result = new(big.Int).Add(a, b)
		resultStr = fmt.Sprintf("%s + %s = %s", a.String(), b.String(), result.String())
	case "sub":
		result = new(big.Int).Sub(a, b)
		resultStr = fmt.Sprintf("%s - %s = %s", a.String(), b.String(), result.String())
	case "mul":
		result = new(big.Int).Mul(a, b)
		resultStr = fmt.Sprintf("%s * %s = %s", a.String(), b.String(), result.String())
	case "div":
		if b.Cmp(big.NewInt(0)) != 0 {
			result = new(big.Int).Div(a, b)
			remainder := new(big.Int).Mod(a, b)
			if remainder.Cmp(big.NewInt(0)) == 0 {
				resultStr = fmt.Sprintf("%s / %s = %s", a.String(), b.String(), result.String())
			} else {
				resultStr = fmt.Sprintf("%s / %s = %s (remainder %s)", a.String(), b.String(), result.String(), remainder.String())
			}
		} else {
			resultStr = "division by zero"
		}
	}

	results <- Result{operation: operation, value: resultStr}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("input numbers greater than 1048576 (2^20)")

	fmt.Print("input number a: ")
	aStr, _ := reader.ReadString('\n')
	aStr = strings.TrimSpace(aStr)

	fmt.Print("input number b: ")
	bStr, _ := reader.ReadString('\n')
	bStr = strings.TrimSpace(bStr)

	a := new(big.Int)
	b := new(big.Int)

	_, okA := a.SetString(aStr, 10)
	_, okB := b.SetString(bStr, 10)

	if !okA || !okB {
		fmt.Println("error: invalid number")
		return
	}

	minValue := big.NewInt(1048576)
	if a.Cmp(minValue) <= 0 || b.Cmp(minValue) <= 0 {
		fmt.Printf("warning: numbers should be greater than %s\n", minValue.String())
	}

	fmt.Printf("\n\na = %s\nb = %s\n\n", a.String(), b.String())

	results := make(chan Result, 4)
	var wg sync.WaitGroup

	operations := []string{"add", "sub", "mul", "div"}

	fmt.Println("Processing...")

	for _, op := range operations {
		wg.Add(1)
		go calculate(a, b, op, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	fmt.Println("\nresults:")
	fmt.Println("===========")

	for result := range results {
		fmt.Printf("%-10s: %s\n", result.operation, result.value)
	}

	//fmt.Println("\nВсе операции выполнены!")
}
