package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// 获取随机数
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	//fmt.Println("The secret number is ", secretNumber)

	// 读取用户输入
	fmt.Println("Please input your guess")
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
			//return
			continue
		}
		input = strings.TrimSuffix(input, "\n")

		// 转换为数字
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer value")
			//return
			continue
		}
		//fmt.Println("You guess is ", guess)

		if guess > secretNumber {
			fmt.Println("Your guess is bigger than the secret number. PLease try again")
		} else if guess < secretNumber {
			fmt.Println("Your guess is small than the secret number. Please try again")
		} else {
			fmt.Println("Correct,you legend!")
			break
		}
	}
}
