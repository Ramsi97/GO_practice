package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculateAverage(grades []float64) float64 {
	if len(grades) == 0 {
		return 0
	}

	sum := 0.0

	for _, grade := range grades {
		sum += grade
	}

	return sum / float64(len(grades))

}

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter number of subjects: ")
	numInput, _ := reader.ReadString('\n')

	numInput = strings.TrimSpace(numInput)
	numSubjects, err := strconv.Atoi(numInput)

	if err != nil || numSubjects <= 0 {
		fmt.Println("Invalid number of subjects. Please enter a positive integer.")
		return
	}

	subjects := make(map[string]float64)
	var grades []float64

	for i := 1; i <= numSubjects; i++ {

		fmt.Printf("\nEnter name of subject %d: ", i)
		subjectName, _ := reader.ReadString('\n')
		subjectName = strings.TrimSpace(subjectName)

		fmt.Printf("Enter grade for %s (0–100): ", subjectName)
		gradeInput, _ := reader.ReadString('\n')
		gradeInput = strings.TrimSpace(gradeInput)
		grade, err := strconv.ParseFloat(gradeInput, 64)
		if err != nil || grade < 0 || grade > 100 {
			fmt.Println("Invalid grade. Please enter a number between 0 and 100.")
			i--
			continue
		}

		subjects[subjectName] = grade
		grades = append(grades, grade)
	}

	average := calculateAverage(grades)

	fmt.Printf("\n--- Grade Report for %s ---\n", name)
	for subject, grade := range subjects {
		fmt.Printf("%-15s : %.2f\n", subject, grade)
	}
	fmt.Printf("\nAverage Grade: %.2f\n", average)

	switch {
	case average >= 90:
		fmt.Println("Excellent performance! 🎉")
	case average >= 75:
		fmt.Println("Good job! 👍")
	case average >= 50:
		fmt.Println("You passed, keep improving!")
	default:
		fmt.Println("Needs improvement. 💪")
	}

}
