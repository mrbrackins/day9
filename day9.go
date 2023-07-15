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
	//trim the last piece off of userResponce making it not a number
	userResponse = userResponse[:len(userResponse)-1]

	userResponseInt, _ := strconv.Atoi(userResponse)

	if(isCurrentYear(currentYear, userResponseInt)){
		fmt.Printf("Correct, %d is the current year", currentYear)
	} else {
		differenceOfYears := currentYear - userResponseInt

		fmt.Printf("Unfortunately not. %d is the current year. You are %d year(s) off... ", currentYear, differenceOfYears)
	}

	currentMonth, _ := strconv.Atoi(getCurrentDate()[6:7])
	reader = bufio.NewReader(os.Stdin)
	
	
	fmt.Println("Current Month.")
	fmt.Println(currentMonth)
	fmt.Println("Please enter the current month...")
	
	userResponseMonth, _ := reader.ReadString('\n')
	userResponseMonth = strings.Replace(userResponseMonth, "\n", "", -1)
	//trim the last piece off of userResponce making it not a number
	userResponseMonth = userResponseMonth[:len(userResponseMonth)-1]

	userResponseMonthInt, _ := strconv.Atoi(userResponseMonth)

	if(isCurrentYear(currentMonth, userResponseMonthInt)){
		fmt.Printf("Correct, %d is the current month", currentMonth)
	} else {
		differenceOfMonths := currentMonth - userResponseMonthInt

		fmt.Printf("Unfortunately not. %d is the current month. You are %d month(s) off... ", currentMonth, differenceOfMonths)
	}


	currentDay, _ := strconv.Atoi(getCurrentDate()[8:10])
	reader = bufio.NewReader(os.Stdin)
	
	
	fmt.Println("Current Day.")
	fmt.Println(currentDay)
	fmt.Println("Please enter the current day...")
	
	usrResponseDay, _ := reader.ReadString('\n')
	usrResponseDay = strings.Replace(usrResponseDay, "\n", "", -1)
	//trim the last piece off of userResponce making it not a number
	usrResponseDay = usrResponseDay[:len(usrResponseDay)-1]

	usrResponseDayInt, _ := strconv.Atoi(usrResponseDay)

	if(isCurrentYear(currentDay, usrResponseDayInt)){
		fmt.Printf("Correct, %d is the current day", currentDay)
	} else {
		differenceOfDays := currentDay - usrResponseDayInt

		fmt.Printf("Unfortunately not. %d is the current day. You are %d day(s) off... ", currentDay, differenceOfDays)
	}
	

}


func getCurrentDate() string {
	return time.Now().String()
}

func isCurrentYear(curr int, uResponse int) bool{
if curr == uResponse {
		return true
	} else {
		return false
	}
}