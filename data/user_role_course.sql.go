package data

const createUserRoleCourse = `
INSERT INTO user_role_course (
	user_id,
	role_id,
	course_id
) 
VALUES (
	$1, $2, $3
)
`

const getUserRoleCourseByCourseId = `
SELECT
	user_id,
	role_id,
	course_id
FROM user_role_course
WHERE course_id = $1
`
