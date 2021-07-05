package data

type CoursesClassSession struct {
	CourseId        int
	ClassSessionId  int
	SubjectModuleId int
	At_date         string
}

func (cLS *CoursesClassSession) Create() (err error) {
	_, err = Db.Query("insert into courses_class_session (course_id, class_session_id, subject_module_id) values ($1, $2, $3)",
		cLS.CourseId,
		cLS.ClassSessionId,
		cLS.SubjectModuleId)
	return
}
