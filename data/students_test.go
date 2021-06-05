package data

import (
	"fmt"
	"testing"
)

func TestCreate(t *testing.T) {
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
	if st.Id == 0 {
		t.Error("wrong id, was not expecting 0")
	}

	err = st1.Create()

	if err != nil {
		t.Error(err)
	}
	if st.Id == 0 {
		t.Error("wrong id, was expecting 0")
	}
}

func TestGetStudent(t *testing.T) {
	st, err := GetStudent(4)

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
