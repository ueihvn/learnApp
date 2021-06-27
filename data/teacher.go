package data

type Teacher struct {
	Id        int    `db:"id"`
	Mail      string `db:"mail"`
	FullName  string `db:"full_name"`
	UserName  string `db:"user_name"`
	Password  string `db:"password"`
	SubjectId int    `db:"subject_id"`
}

func (teacher *Teacher) CreateTeacher() (err error) {
	err = Db.QueryRow("insert into teachers (mail, full_name, user_name, password, subject_id) values ($1, $2, $3, $4, $5) returning id",
		teacher.Mail,
		teacher.FullName,
		teacher.UserName,
		teacher.Password,
		teacher.SubjectId).Scan(&teacher.Id)
	if err != nil {
		return
	}
	return
}

func GetTeachersBySubjectId(subjectId int) (teachers []Teacher, err error) {
	rows, err := Db.Queryx("select id, mail, full_name, user_name, password, subject_id from teachers where subject_id = $1", subjectId)
	if err != nil {
		return
	}
	for rows.Next() {
		teacher := Teacher{}
		err = rows.StructScan(&teacher)
		if err != nil {
			return
		}
		teachers = append(teachers, teacher)
	}
	rows.Close()
	return
}

func GetTeacherById(teacherId int) (teacher Teacher, err error) {
	err = Db.QueryRowx("select id, mail, full_name, user_name, password, subject_id from teachers where id = $1", teacherId).StructScan(&teacher)
	return
}

func GetAllTeachers() (teachers []Teacher, err error) {
	rows, err := Db.Queryx("select id, mail, full_name, user_name, password, subject_id from teachers")
	if err != nil {
		return
	}

	for rows.Next() {
		teacher := Teacher{}
		err = rows.StructScan(teacher)
		if err != nil {
			return
		}
		teachers = append(teachers, teacher)
	}
	rows.Close()
	return
}

func (teacher *Teacher) Update() (err error) {
	_, err = Db.NamedExec(`update teachers set mail =:mail, full_name =:full_name, user_name =:user_name, password =:password,subject_id =:subject_id where id =:id`, teacher)
	return
}

func (teacher *Teacher) Delete() (err error) {
	_, err = Db.Exec("delete from students where id = $1", teacher.Id)
	return
}
