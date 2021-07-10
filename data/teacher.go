package data

import "log"

type Teacher struct {
	Id        int    `db:"id"`
	Mail      string `db:"mail"`
	FullName  string `db:"full_name"`
	UserName  string `db:"user_name"`
	Password  string `db:"password"`
	SubjectId int    `db:"subject_id"`
}

func (teacher *Teacher) Create() (err error) {
	err = Db.QueryRow("insert into teachers (mail, full_name, user_name, password, subject_id) values ($1, $2, $3, $4, $5) returning id",
		teacher.Mail,
		teacher.FullName,
		teacher.UserName,
		teacher.Password,
		teacher.SubjectId).Scan(&teacher.Id)
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

func GetTeacherByUserName(userName string) (teacher Teacher, err error) {
	err = Db.QueryRowx("select id, mail, full_name, user_name, password, subject_id from teachers where user_name = $1", userName).StructScan(&teacher)
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
	_, err = Db.Exec("delete from teachers where id = $1", teacher.Id)
	return
}

func InitTeacher() (err error) {
	subjectInit, err := GetAllSubjectInit()
	if err != nil {
		return err
	}

	teachers := []Teacher{
		{
			Mail:      "emailThay1@gmail.com",
			FullName:  "Full name thay1",
			UserName:  "thay1",
			Password:  "passwordthay1",
			SubjectId: subjectInit["Toán Học"].Id,
		},
		{
			Mail:      "emailThay2@gmail.com",
			FullName:  "Full name thay2",
			UserName:  "thay2",
			Password:  "passwordthay2",
			SubjectId: subjectInit["Vật Lý"].Id,
		},
		{
			Mail:      "emailThay3@gmail.com",
			FullName:  "Full name thay3",
			UserName:  "thay3",
			Password:  "passwordthay3",
			SubjectId: subjectInit["Hóa Học"].Id,
		},
		{
			Mail:      "emailThay4@gmail.com",
			FullName:  "Full name thay4",
			UserName:  "thay4",
			Password:  "passwordthay4",
			SubjectId: subjectInit["Sinh Học"].Id,
		},
		{
			Mail:      "emailThay5@gmail.com",
			FullName:  "Full name thay5",
			UserName:  "thay5",
			Password:  "passwordthay5",
			SubjectId: subjectInit["Tiếng Anh"].Id,
		},
	}

	for _, teacher := range teachers {
		err := teacher.Create()
		if err != nil {
			return err
		}
	}
	return
}

func GetAllTeacherInit() (mapTeacher map[string]Teacher, err error) {
	mapTeacher = map[string]Teacher{}
	teacherInit := []string{"thay1", "thay2", "thay3", "thay4", "thay5"}

	for _, teacherUserName := range teacherInit {
		teacher, err := GetTeacherByUserName(teacherUserName)
		if err != nil {
			log.Fatal(err)
		}
		mapTeacher[teacherUserName] = teacher
	}

	return
}
