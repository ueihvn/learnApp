package data

const createCourse = `
INSET INTO courses (
	resource_id,
	name,
	description,
	subject_id,
) VALUES (
	$1, $2, $3, $4
) RETURNING id, resource_id, name, description, subject_id, is_delete, is_open
`
