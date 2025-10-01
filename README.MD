# ESEP Grade Calculator

This library is a simple utility to calculate a weighted average of your grades. 

The GradeCalculator supports 3 types of grades: 
 - 1. Assignments (50% of total grade)
 - 2. Exams (35% of total grade)
 - 3. Essays (15% of total grade)

 Instructions for using the grade calculator:

 - Instantiate a GradeCalculator. Ex: `gradeCalculator := NewGradeCalculator()`

 - Add assignment grades by providing a name, grade, and grade type. Ex: `gradeCalculator.AddGrade("exam 1", 85, Exam)`

 - To compute a final letter grade, use the GetFinalGrade() method

 - To compute a numeric weighted average, use the calculateNumericalGrade() method

Runtime requirements:
Golang 1.24+

Run unit tests:
```
cd grade_calculator/
go test ./...
```

Run unit tests with code coverage:
```
cd grade_calculator/
go test -coverprofile=coverage.txt ./...
```