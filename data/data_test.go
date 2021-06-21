package data

import (
	"fmt"
	"testing"
)

//test for Subject
func TestInitSubject(t *testing.T) {
	err := InitSubject()
	if err != nil {
		t.Error(err)
	}
}

func TestGetSubjectByType(t *testing.T) {
	subject, err := GetSubjectByType("Toán Học")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(subject)
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

	st1 := Student{
		Mail:     "st1mail@gmail.com",
		FullName: "fullname st1 là 2",
		UserName: "st1username",
		Password: "st1password",
	}

	err := st.Create()
	if err != nil {
		t.Error(err)
	}

	err = st1.Create()
	if err != nil {
		t.Error(err)
	}

}

func TestGetStudent(t *testing.T) {
	st, err := GetStudentById(4)

	if err != nil {
		t.Error(err)
	}
	fmt.Println(st)
}

func TestDeleteStudent(t *testing.T) {
	st := Student{Id: 3}
	err := st.Delete()
	if err != nil {
		t.Error(err)
	}
}
