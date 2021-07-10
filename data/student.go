package data

import "log"

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
	err = Db.QueryRowx("select id, mail, full_name, user_name, password from students where id = $1", id).StructScan(&st)
	return
}

func GetStudentByUserName(userName string) (st Student, err error) {
	err = Db.QueryRowx("select id, mail, full_name, user_name, password from students where user_name = $1", userName).StructScan(&st)
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

func InitStudent() (err error) {
	students := []Student{
		{
			Mail:     "mailSt1@gmail.com",
			FullName: "full name St1",
			UserName: "St1",
			Password: "passwordSt1",
		},
		{
			Mail:     "mailSt2@gmail.com",
			FullName: "full name St2",
			UserName: "St2",
			Password: "passwordSt2",
		},
		{
			Mail:     "mailSt3@gmail.com",
			FullName: "full name St3",
			UserName: "St3",
			Password: "passwordSt3",
		},
	}

	for _, student := range students {
		err = student.Create()
		if err != nil {
			return
		}
	}
	return
}

func GetAllStudentInit() (mapStudent map[string]Student, err error) {
	mapStudent = map[string]Student{}
	initStudent := []string{"St1", "St2", "St3"}

	for _, studentUsername := range initStudent {
		student, err := GetStudentByUserName(studentUsername)
		if err != nil {
			log.Fatal(err)
		}
		mapStudent[studentUsername] = student
	}
	return
}
