# ESEP Grade Calculator

This library serves as a utility to assist with grade calculations for assignments, exams and essays
in classrooms.  

## How it works
Each grade is represented by a 'Grade' struct, storing:
- A string (name of the assignment, exam or essay)
- An integer (the grade itself)
- Grade type ('Assignment', 'Exam' or 'Essay')

The calculator will save each grade struct in separate slices, which is important for distinguishing 
them during calculations.

Each grade type is weighted differently, which impacts the final calculation:

| Category    | Weight |
|-------------|--------|
| Assignments | 50%    |
| Exams       | 35%    |
| Essays      | 15%    |

The final average grade will be calculated using all grade structs, with final weighting being:

| Range  | Grade |
|--------|-------|
| 90–100 | A     |
| 80–89  | B     |
| 70–79  | C     |
| 60–69  | D     |
| <60    | F     |


## Requirements:
Golang 1.24+ is required for runtime.


## Testing: 
To run unit tests, type the following in command prompt. Typing "./..." will automatically find
the "grade_calculator_test.go" file in the grade calculator directory:
```
cd grade_calculator/
go test ./...
```

To run unit tests with code coverage, type the following:
```
cd grade_calculator/
go test -coverprofile=coverage.txt ./...
```