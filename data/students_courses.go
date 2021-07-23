package data

type StudentCourse struct {
	StudentId int    `db:"student_id"`
	CourseId  int    `db:"course_id"`
	JoinAt    string `db:"join_at"` //yyyy-mm-dd
}

func (studentCourse *StudentCourse) Create() (err error) {
	_, err = Db.Query("insert into students_courses (student_id, course_id) values($1, $2)", studentCourse.StudentId, studentCourse.CourseId)
	return
}

func GetAllStudentsOfCourse(courseId int) (students []Student, err error) {
	query := `
		select 
			students.id,
			students.mail,
			students.full_name,
			students.user_name,
			students.password
		from 
			students join students_courses on students.id = students_courses.student_id 
		where students_courses.course_id = $1`

	rows, err := Db.Queryx(query, courseId)
	if err != nil {
		return
	}

	for rows.Next() {
		student := Student{}
		err = rows.StructScan(&student)
		if err != nil {
			return
		}
		students = append(students, student)
	}
	defer rows.Close()
	return
}

func GetAllCoursesOfStudent(studentId int) (courses []Course, err error) {
	query := `
		select
		courses.id,
		courses.class_id,
		courses.start_date,
		courses.end_data,
		courses.grade
	from
		courses join students_courses on courses.id = students_courses.course_id
	where students_courses.student_id = $1
	`
	rows, err := Db.Queryx(query, studentId)
	for rows.Next() {
		course := Course{}
		err = rows.StructScan(&course)
		if err != nil {
			return
		}
		courses = append(courses, course)
	}
	return
}

/*
func InitStudentsCourses() (err error) {
	mapCourse, err := GetAllCourseInit()
	if err != nil {
		return
	}

	mapStudent, err := GetAllStudentInit()
	if err != nil {
		return
	}

	for _, course := range mapCourse {
		for _, student := range mapStudent {
			studentCourse := StudentCourse{
				StudentId: student.Id,
				CourseId:  course.Id,
			}
			err = studentCourse.Create()
			if err != nil {
				return
			}
		}
	}

	return
}
*/
