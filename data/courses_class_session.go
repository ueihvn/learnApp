package data

type CoursesClassSession struct {
	CourseId        int    `db:"course_id"`
	ClassSessionId  int    `db:"class_session_id`
	SubjectModuleId int    `db: "subject_module_id`
	At_date         string `db: "subject_module_id`
}

func (cLS *CoursesClassSession) Create() (err error) {
	query := "insert into courses_class_session (course_id, class_session_id, subject_module_id) values ($1, $2, $3)"
	_, err = Db.Query(query,
		cLS.CourseId,
		cLS.ClassSessionId,
		cLS.SubjectModuleId)
	return
}

/*
func (cLS *CoursesClassSession) Update() (err error) {
	query := `
		update
			courses_class_session
		set
			course_id= :course_id, class_session_id= :class_session_id, subject_module_id= :subject_module_id
		where `
}
*/
