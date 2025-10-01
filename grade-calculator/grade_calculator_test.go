package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeC(t *testing.T) {
	expected_value := "C"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 35, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 52, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeD(t *testing.T) {
	expected_value := "D"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 82, Assignment)
	gradeCalculator.AddGrade("exam 1", 42, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 67, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGradeTypeToString(t *testing.T) {
	expected_value1 := "assignment"

	actual_value1 := Assignment.String()

	if expected_value1 != actual_value1 {
		t.Errorf("Expected GradeType.String() to return '%s'; got '%s' instead", expected_value1, actual_value1)
	}

	expected_value2 := "exam"

	actual_value2 := Exam.String()

	if expected_value2 != actual_value2 {
		t.Errorf("Expected GradeType.String() to return '%s'; got '%s' instead", expected_value2, actual_value2)
	}

	expected_value3 := "essay"

	actual_value3 := Essay.String()

	if expected_value3 != actual_value3 {
		t.Errorf("Expected GradeType.String() to return '%s'; got '%s' instead", expected_value3, actual_value3)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 67, Assignment)
	gradeCalculator.AddGrade("exam 1", 35, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 91, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetPassOrFailPass(t *testing.T) {
	expected_value := "Pass"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 35, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 91, Essay)

	actual_value := gradeCalculator.getPassOrFail()

	if expected_value != actual_value {
		t.Errorf("Expected getPassOrFail to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetPassOrFailFail(t *testing.T) {
	expected_value := "Fail"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 68, Assignment)
	gradeCalculator.AddGrade("exam 1", 35, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 91, Essay)

	actual_value := gradeCalculator.getPassOrFail()

	if expected_value != actual_value {
		t.Errorf("Expected getPassOrFail to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}
