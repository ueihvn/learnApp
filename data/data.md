## package data explantation

### setupdb.go

github.com/lib/pq is Postgres driver

### student.go

#### Create

statement is SQL prepared statement: SQL statement template, can replace certain values during execution. Prepared statements are often used to execure statemenst repeatedly

use db.Prepare(statement) to create a prepared statement
execute prepared statement by using QueryRow method
When QueryRow method return only a single reference to an sql.Row struct, no errors.
Because QueryRow is often used with the method Scan on the Row struct, which copies the values in the row into its parameters

#### GetStudent

Using QueryRow() method on sql.Db struct not sql.Stmt struct don't have to need a prepared statement

#### Update, Delte

use Db.Exec() because don't have to update receiver, so don't need scan returned results -> using Db.Exec() return sql.Result and error, is much faster
