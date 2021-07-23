package data

const createResource = `
INSERT INTO resources (
	type
)
VALUES (
	$1
)
RETURNING id, type, created_at
`

const getResourceById = `
SELECT
	id,
	type,
	created_at
FROM resources
WHERE id = $1
`

const getAllResource = `
SELECT
	id,
	type,
	created_at
FROM resources
`
const deleteResource = `
DELETE FROM resources
WHERE id = $1
`
