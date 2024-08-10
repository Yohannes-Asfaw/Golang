package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	
	var numSubjects int

	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name) 
	for {
		fmt.Print("Enter the number of subjects you have taken: ")
		_, err := fmt.Scanln(&numSubjects)

		if err != nil {
			fmt.Println("Invalid input. Please enter a positive integer.")
			reader.ReadString('\n') // Clear the buffer
		} else if numSubjects > 0 {
			break
		} else {
			fmt.Println("Invalid number of subjects. Please enter a positive integer.")
		}
	}

	subjects := make(map[string]float64)

	for i := 0; i < numSubjects; i++ {
		var subjectName string
		var grade float64

		fmt.Printf("Enter the name of subject %d: ", i+1)
		subjectName, _ = reader.ReadString('\n')
		subjectName = strings.TrimSpace(subjectName)

		for {
			fmt.Printf("Enter the grade for %s (0-100): ", subjectName)
			_, err := fmt.Scanln(&grade)

			if err != nil {
				fmt.Println("Invalid input. Please enter a number between 0 and 100.")
				reader.ReadString('\n') // Clear the buffer
			} else if grade >= 0 && grade <= 100 {
				break
			} else {
				fmt.Println("Invalid grade. Grades must be between 0 and 100. Please enter again.")
			}
		}

		subjects[subjectName] = grade
	}

	average := calculateAverage(subjects)

	fmt.Println(strings.Repeat("-", 40))
	fmt.Printf("Student Name: %s\n", name)
	fmt.Println("Subject Grades:")

	for subject, grade := range subjects {
		fmt.Printf("%-20s: %.2f\n", subject, grade)
	}

	fmt.Printf("Average Grade: %.2f\n", average)
	fmt.Println(strings.Repeat("-", 40))
}

func calculateAverage(subjects map[string]float64) float64 {
	var total float64
	for _, grade := range subjects {
		total += grade
	}
	return total / float64(len(subjects))
}
