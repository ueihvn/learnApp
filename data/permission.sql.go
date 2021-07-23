package data

const createPermission = `
INSERT INTO permissions (
	type
)
VALUES (
	$1
)
RETURNING id, type, created_at
`

const getPermissionById = `
SELECT
	id,
	type,
	created_at
FROM permissions
WHERE id = $1
`

const getAllPermission = `
SELECT
	id,
	type,
	created_at
FROM permissions
`
const deletePermission = `
DELETE FROM permissions
WHERE id = $1
`
