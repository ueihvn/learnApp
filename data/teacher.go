package data

type Teacher struct {
	Id        int
	Mail      string
	FullName  string
	UserName  string
	Password  string
	SubjectId int
}

func (teacher *Teacher) CreateTeacher() (err error) {
	_, err = Db.NamedExec(`insert into teachers (mail, full_name,  user_name, password, subject_id) values (:mail, :full_name, :user_name, :password, :subject_id)`, teacher)
	if err != nil {
		return
	}
	return
}
