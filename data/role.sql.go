package data

const createRole = `
INSERT INTO roles (
	type
)
VALUES (
	$1
)
RETURNING id, type, created_at
`

const getRoleById = `
SELECT
	id,
	type,
	created_at
FROM roles
WHERE id = $1
`

const getAllRole = `
SELECT
	id,
	type,
	created_at
FROM resources
`
const deleteRole = `
DELETE FROM roles
WHERE id = $1
`
