package data

import (
	"fmt"
	"testing"
)

//test students_courses
func TestGetAllStudentOfCourse(t *testing.T) {
	students, err := GetAllStudentsOfCourse(1)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(students)
	}
}

func TestGetAllCourseOfStudent(t *testing.T) {
	students, err := GetAllCoursesOfStudent(1)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(students)
	}
}

//test setupDb
func TestInitDataDb(t *testing.T) {
	err := InitDataDb()
	if err != nil {
		t.Error(err)
	}
}

//test for course
func TestInitStudentsCourses(t *testing.T) {
	err := InitStudentsCourses()
	if err != nil {
		t.Error(err)
	}
}

// test for subject_module
func TestInitSubjectModules(t *testing.T) {
	err := InitSubjectModules()
	if err != nil {
		t.Error(err)
	}
}

//test for Subject
func TestInitSubject(t *testing.T) {
	err := InitSubject()
	if err != nil {
		t.Error(err)
	}
}

func TestGetAllSubject(t *testing.T) {
	subjects, err := GetAllSubject()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(subjects)
}

// test for Student
func TestCreateStudent(t *testing.T) {
	st := Student{
		Mail:     "mail@gmail.com",
		FullName: "hei lt",
		UserName: "htl",
		Password: "password",
	}

	err := st.Create()
	if err != nil {
		t.Error(err)
	}

}

func TestGetStudent(t *testing.T) {
	st, err := GetStudentById(1)

	if err != nil {
		t.Error(err)
	}
	fmt.Println(st)
}

func TestDeleteStudent(t *testing.T) {
	st := Student{Id: 1}
	err := st.Delete()
	if err != nil {
		t.Error(err)
	}
}
