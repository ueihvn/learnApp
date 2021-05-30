package main

import (
	"testing"
)

func TestCreate(t *testing.T) {
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
	if st.Id == 0 {
		t.Error("wrong id, was expecting 0")
	}
	if st.Mail != "mail@gmail.com" {
		t.Error("wrong mail,was expecting 'mail@gmail.com' but got", st.Mail)
	}
	if st.FullName != "hei lt" {
		t.Error("wrong fullname,was expecting 'hei lt' but got", st.FullName)
	}
	if st.UserName != "htl" {
		t.Error("wrong username,was expecting 'htl' but got", st.UserName)
	}
	if st.Password != "password" {
		t.Error("wrong password,was expecting 'password' but got", st.Password)
	}
}
