package esepunittests

type GradeCalculator struct {
	assignments_exams_essays []Grade
	//exams       []Grade
	//essays      []Grade
}

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{
		assignments_exams_essays: make([]Grade, 0),
		//exams:       make([]Grade, 0),
		//essays:      make([]Grade, 0),
	}
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()

	if numericalGrade >= 90 {
		return "A"
	} else if numericalGrade >= 80 {
		return "B"
	} else if numericalGrade >= 70 {
		return "C"
	} else if numericalGrade >= 60 {
		return "D"
	}

	return "F"
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	gc.assignments_exams_essays = append(gc.assignments_exams_essays, Grade{
			Name:  name,
			Grade: grade,
			Type:  gradeType,
		})
	/*switch gradeType {
	case Assignment:
		gc.assignments = append(gc.assignments, Grade{
			Name:  name,
			Grade: grade,
			Type:  Assignment,
		})
	case Exam:
		gc.exams = append(gc.exams, Grade{
			Name:  name,
			Grade: grade,
			Type:  Exam,
		})
	case Essay:
		gc.essays = append(gc.essays, Grade{
			Name:  name,
			Grade: grade,
			Type:  Essay,
		})
	}*/
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	assignment_average := computeAverage(gc.assignments_exams_essays,Assignment)
	exam_average := computeAverage(gc.assignments_exams_essays,Exam)
	essay_average := computeAverage(gc.assignments_exams_essays,Essay)

	weighted_grade := float64(assignment_average)*.5 + float64(exam_average)*.35 + float64(essay_average)*.15

	return int(weighted_grade)
}

func computeAverage(grades []Grade, gradetype GradeType) int {
	sum := 0
	num_of_valid_gradetype := 0

	for _, index := range grades {
		if (index.Type == gradetype) {
			sum += index.Grade
			num_of_valid_gradetype++
		}
	}


	return sum / num_of_valid_gradetype
}