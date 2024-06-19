package main

import (
	"encoding/csv"
	"fmt"
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

// Stringer function for student struct
func (s student) String() string {
	studentTemplate := "{ \n Student's name: %v \n Student's University: %v \n Student's score 1: %v \n Student's score 2: %v \n Student's score 3: %v \n Student's score 4: %v \n }\n"
	return fmt.Sprintf(studentTemplate, (s.firstName + " " + s.lastName), s.university, s.test1Score, s.test2Score, s.test3Score, s.test4Score)
}

type studentStat struct {
	student
	finalScore float32
	grade      Grade
}

func findFinalScore(scores ...int) float32 {
	total := 0
	for _, score := range scores {
		total += score
	}
	return float32(total) / float32(len(scores))
}

func findFinalGrade(finalScore float32) Grade {
	if finalScore < 35 {
		return F
	} else if finalScore >= 35 && finalScore < 50 {
		return C
	} else if finalScore >= 50 && finalScore < 70 {
		return B
	} else {
		return A
	}
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

		/*
			fmt.Println("Current student is", currentStudent)
			fmt.Println("-----------")
		*/

	}
	return students
}

func calculateGrade(students []student) []studentStat {

	var studentStats []studentStat

	for _, student := range students {
		currentStudentStat := studentStat{}
		currentStudentStat.student = student
		currentStudentFinalScore := findFinalScore(student.test1Score, student.test2Score, student.test3Score, student.test4Score)
		currentStudentStat.finalScore = currentStudentFinalScore
		currentStudentStat.grade = findFinalGrade(currentStudentFinalScore)

		studentStats = append(studentStats, currentStudentStat)
	}
	return studentStats
}

func findOverallTopper(gradedStudents []studentStat) studentStat {
	var overAllTopper studentStat
	overAllTopper.finalScore = -1.0

	for _, stat := range gradedStudents {
		if stat.finalScore > overAllTopper.finalScore {
			overAllTopper = stat
		}
	}
	return overAllTopper
}

func findTopperPerUniversity(gs []studentStat) map[string]studentStat {
	universityWiseStudentStat := make(map[string][]studentStat)
	topperPerUniversity := make(map[string]studentStat)

	for _, studentStat := range gs {
		universityWiseStudentStat[studentStat.university] = append(universityWiseStudentStat[studentStat.university], studentStat)
	}

	for k, v := range universityWiseStudentStat {
		currentUniversity := k
		currentUniversityTopper := findOverallTopper(v)
		topperPerUniversity[currentUniversity] = currentUniversityTopper
	}

	return topperPerUniversity
}
