package gobe

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Turn any model into a JSON
func ToJSON(model any) string {
	res, err := json.MarshalIndent(model, "", " ")
	if err != nil {
		return ""
	}
	return string(res)
}

// Generate random numbers in any length. Usually used for generating OTP
func GenerateRandomNumber(length int) int {
	var numbers []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		num := rand.Intn(10)
		numbers = append(numbers, num)
	}
	var numStr []string
	for _, n := range numbers {
		res := strconv.Itoa(n)
		numStr = append(numStr, res)
	}
	numStrJoined := strings.Join(numStr, "")
	result, err := strconv.Atoi(numStrJoined)
	if err != nil {
		fmt.Printf("failed to generate random number with error: %s", err.Error())
	}
	return result
}
