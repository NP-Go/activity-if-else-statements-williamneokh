package main

import (
	"fmt"
	"strconv"
)

func evaluateValue(value int64) string {
	replyMessage := ""
	//insert your code here
	if value%5 == 0 && value%6 == 0 {
		replyMessage = "The value is divisible by 5 and 6"
	} else {
		replyMessage = "The value is NOT divisible by 5 and 6"
	}
	//Do not remove these lines
	fmt.Println(replyMessage)
	return replyMessage

}

func main() {
	var compareValue string
	fmt.Println("Enter an integer: ")
	_, _ = fmt.Scanln(&compareValue)

	//conversion of value
	valueInt, _ := strconv.ParseInt(compareValue, 10, 0)
	evaluateValue(valueInt)

}
