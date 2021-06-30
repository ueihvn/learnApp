package data

type Course struct {
	Id        int
	ClassId   int
	StartDate string //yyyy-mm-dd
	EndDate   string //yyyy-mm-dd
	Grade     int
}

func (course *Course) Create() (err error) {
	err = Db.QueryRow("insert into courses (class_id, start_date, end_date, grade) values ($1, $2, $3, $4) returning id",
		course.ClassId,
		course.StartDate,
		course.EndDate,
		course.Grade).Scan(&course.Id)
	return
}

func GetCoursesByGrade(grade int) (courses []Course, err error) {
	rows, err := Db.Queryx("select id, class_id, start_date, end_date, grade from courses where grade = $1", grade)
	if err != nil {
		return
	}

	for rows.Next() {
		course := Course{}
		err = rows.StructScan(&course)
		if err != nil {
			return
		}
		courses = append(courses, course)
	}
	defer rows.Close()
	return
}

func GetCoursesByClassId(classId int) (courses []Course, err error) {
	rows, err := Db.Queryx("select id, class_id, start_date, end_date, grade from courses where grade = $1", classId)
	if err != nil {
		return
	}

	for rows.Next() {
		course := Course{}
		err = rows.StructScan(&course)
		if err != nil {
			return
		}
		courses = append(courses, course)
	}
	defer rows.Close()
	return
}
