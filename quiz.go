// quiz.go
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	filename = "./problems.csv"
)

var (
	// to track the correct answers
	correct = 1
)

func main() {

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal("os failed to open csv file", err)

	}

	r := csv.NewReader(f)
	records, err := r.ReadAll()

	if err != nil {
		log.Fatal("Error opening using csvReader:", err)
	}

	fmt.Println("Enter any key to start")
	var answer string
	fmt.Scanf("%s \n", &answer)

	rand.Seed(time.Now().Unix())
	for i := 1; i < 11; i++ {
		randomnu := rand.Intn(len(records))
		fmt.Println("Question:", i, records[randomnu][0])
		fmt.Scanf("%s\n", &answer)
		if strings.TrimSpace(answer) == strings.TrimSpace(records[randomnu][1]) {
			correct++
		}

	}

	fmt.Println("Overall score :", correct, " out of ", 11)

}
