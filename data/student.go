package main

type Student struct {
	Id       int
	Mail     string
	FullName string
	UserName string
	Password string
}

func (st *Student) Create() (err error) {
	statement := "insert into students(mail, full_name, user_name, password) values ($1, $2, $3, $4) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(st.Mail, st.FullName, st.UserName, st.Password).Scan(&st.Id)
	return
}

func GetStudent(id int) (st Student, err error) {
	st = Student{}
	err = Db.QueryRow("select id, mail, full_name, user_name, password from students where id = $1", id).Scan(&st.Id, &st.Mail, &st.FullName, &st.UserName, &st.Password)
	return
}

func GetStudentsWithNumber(limit int) (sts []Student, err error) {
	rows, err := Db.Query("select id, mail, full_name, user_name, password from students limit $1", limit)
	if err != nil {
		return
	}
	for rows.Next() {
		st := Student{}
		err = rows.Scan(&st.Id, &st.Mail, &st.FullName, &st.UserName, &st.Password)
		if err != nil {
			return
		}
		sts = append(sts, st)
	}
	rows.Close()
	return
}

func (st *Student) Update() (err error) {
	_, err = Db.Exec("update students set mail = $2, full_name = $3, user_name = $4, password = $5 where id = $1", st.Id, st.Mail, st.FullName, st.UserName, st.Password)
	return
}

func (st *Student) Delete() (err error) {
	_, err = Db.Exec("delete from posts where id = $1", st.Id)
	return
}
