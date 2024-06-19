package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

type Grade string

const (
	A Grade = "A"
	B Grade = "B"
	C Grade = "C"
	F Grade = "F"
)

type student struct {
	firstName, lastName, university                string
	test1Score, test2Score, test3Score, test4Score int
}

type studentStat struct {
	student
	finalScore float32
	grade      Grade
}

func parseCSV(filePath string) []student {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	lineNumber := 0

	var students []student

	// Read and parse each line
	for {

		// Read each record from CSV
		record, err := reader.Read()
		
		if err != nil {
			break // Break if end of file
		}

		lineNumber++

		if lineNumber == 1 {
			continue // Skip the first line
		}

		currentStudent := student{}
		currentStudent.firstName = record[0]
		currentStudent.lastName = record[1]
		currentStudent.university = record[2]

		score1, _ := strconv.Atoi(record[3])
		score2, _ := strconv.Atoi(record[4])
		score3, _ := strconv.Atoi(record[5])
		score4, _ := strconv.Atoi(record[6])

		currentStudent.test1Score = score1
		currentStudent.test2Score = score2
		currentStudent.test3Score = score3
		currentStudent.test4Score = score4

		students = append(students, currentStudent)

		// fmt.Println("Current student is", currentStudent)
		// fmt.Println("-----------")
	}
	return students
}

func calculateGrade(students []student) []studentStat {
	return nil
}

func findOverallTopper(gradedStudents []studentStat) studentStat {
	return studentStat{}
}

func findTopperPerUniversity(gs []studentStat) map[string]studentStat {
	return nil
}
