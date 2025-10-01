package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != string(actual_value[0]) {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, string(actual_value[0]))
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != string(actual_value[0]) {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, string(actual_value[0]))
	}
}

func TestGetGradeC(t *testing.T) {
	expected_value := "C"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 70, Assignment)
	gradeCalculator.AddGrade("exam 1", 79, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 72, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != string(actual_value[0]) {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, string(actual_value[0]))
	}
}

func TestGetGradeD(t *testing.T) {
	expected_value := "D"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 60, Assignment)
	gradeCalculator.AddGrade("exam 1", 67, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 63, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != string(actual_value[0]) {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, string(actual_value[0]))
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 40, Assignment)
	gradeCalculator.AddGrade("exam 1", 50, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 55, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != string(actual_value[0]) {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, string(actual_value[0]))
	}
}

func TestString(t *testing.T) {
	expected_value := "assignment"
	if Assignment.String() != expected_value {
		t.Errorf("Expected string 'assignment' ; got '%s' instead", Assignment.String())
	}

	expected_value = "exam"
	if Exam.String() != expected_value {
		t.Errorf("Expected string 'exam' ; got '%s' instead", Exam.String())
	}

	expected_value = "essay"
	if Essay.String() != expected_value {
		t.Errorf("Expected string 'essay' ; got '%s' instead", Essay.String())
	}
}

func TestPass(t *testing.T) {

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 95, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	//planned method is to print final grades as 'A - Pass, D - Fail, ...' then checking string slices for comparison
	if actual_value[4:] != "Pass" {
		t.Errorf("Expected GetGrade to return 'Pass'; got 'Fail' instead")
	}
}

func TestFail(t *testing.T) {
	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 10, Assignment)
	gradeCalculator.AddGrade("exam 1", 30, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 65, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	//planned method is to print final grades as 'A - Pass, D - Fail, ...' then checking string slices for comparison
	if actual_value[4:] != "Fail" {
		t.Errorf("Expected GetGrade to return 'Fail'; got 'Pass' instead")
	}
}