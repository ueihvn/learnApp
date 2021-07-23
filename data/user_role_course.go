package data

type UserRoleCourse struct {
	UserId   int `db:"user_id"`
	RoleId   int `db:"role_id"`
	CourseId int `db:"course_id"`
}

func (url *UserRoleCourse) Create() (err error) {
	stmtx, err := Db.Preparex(createUserRoleCourse)
	if err != nil {
		return
	}
	defer stmtx.Close()

	_, err = stmtx.Exec(
		url.UserId,
		url.RoleId,
		url.CourseId,
	)
	return
}

func GetUserRoleCourseByCourseId(courseId int) (urls []UserRoleCourse, err error) {
	stmtx, err := Db.Preparex(getUserRoleCourseByCourseId)
	if err != nil {
		return
	}
	defer stmtx.Close()

	rowxs, err := stmtx.Queryx(courseId)
	if err != nil {
		return
	}

	for rowxs.Next() {
		url := UserRoleCourse{}
		err = rowxs.StructScan(&url)
		if err != nil {
			return
		}
		urls = append(urls, url)
	}
	defer rowxs.Close()
	return
}
