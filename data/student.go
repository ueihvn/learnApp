package data

type Student struct {
	Id       int    `db:"id"`
	Mail     string `db:"mail"`
	FullName string `db:"full_name"`
	UserName string `db:"user_name"`
	Password string `db:"password"`
}

func (st *Student) Create() (err error) {
	statementx := "insert into students(mail, full_name, user_name, password) values ($1, $2, $3, $4) returning id"
	stmt, err := Db.Preparex(statementx)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRowx(st.Mail, st.FullName, st.UserName, st.Password).Scan(&st.Id)
	return
}

func GetStudentById(id int) (st Student, err error) {
	st = Student{}
	err = Db.QueryRowx("select id, mail, full_name, user_name, password from students where id = $1", id).StructScan(&st)
	return
}

func GetStudentsWithNumber(limit int) (sts []Student, err error) {
	rows, err := Db.Queryx("select id, mail, full_name, user_name, password from students limit $1", limit)
	if err != nil {
		return
	}
	for rows.Next() {
		st := Student{}
		err = rows.StructScan(&st)
		if err != nil {
			return
		}
		sts = append(sts, st)
	}
	rows.Close()
	return
}

func (st *Student) Update() (err error) {
	_, err = Db.NamedExec(`update students set mail =:mail, full_name =:full_name, user_name =:user_name, password =:password where id =:id`, st)
	return
}

func (st *Student) Delete() (err error) {
	_, err = Db.Exec("delete from students where id = $1", st.Id)
	return
}
