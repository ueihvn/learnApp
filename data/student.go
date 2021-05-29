package data

type Student struct {
	Id       int
	Mail     int
	FullName string
	UserName string
	Password string
}

func getStudent(id int) (st Student, err error) {
	st = Student{}
	err = Db.QueryRow("select id, mail, full_name, user_name, password from students where id = $1", id).Scan(&st.Id, &st.Mail, &st.FullName, &st.UserName, &st.Password)
	return
}

func (st *Student) create() (err error) {
	statement := "insert into students(mail, full_name, user_name, password) values ($1, $2, $3, $4) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(st.Mail, st.FullName, st.UserName, st.Password).Scan(&st.Id)
	return
}
