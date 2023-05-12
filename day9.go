package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println(getCurrentDate())
	currentYear, _ := strconv.Atoi(getCurrentDate()[0:4])
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Current Year.")
	fmt.Println(currentYear)
	fmt.Println("Please enter the current year...")
	userResponse, _ := reader.ReadString('\n')
	userResponse = strings.Replace(userResponse, "\n", "", -1)

	userResponseInt, _ := strconv.Atoi(userResponse)
	fmt.Println(userResponseInt)

	if currentYear == userResponseInt {
		fmt.Println("issa match")
	}

}

func getCurrentDate() string {
	return time.Now().String()
}
