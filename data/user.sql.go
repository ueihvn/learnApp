package data

const createUser = `
INSET INTO users (
	email,
	full_name,
	user_name,
	password
) VALUES (
	$1, $2, $3, $4
) RETURNING id, email, full_name, user_name, password, created_at, updated_at
`
const getUserById = `
SELECT
	id,
	email,
	full_name,
	user_name,
	password,
	created_at,
	updated_at
FROM users
WHERE id = $1
`
const getAllUser = `
SELECT
	id,
	email,
	full_name,
	user_name,
	password,
	created_at,
	updated_at
FROM users
`

const updateUser = `
UPDATE users
SET
`

const deleteUser = `
DELETE FROM users
WHERE id = $1
`
