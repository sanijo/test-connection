package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func getAllRows(db *sql.DB) error {
    // get all rows from the database
    rows, err := db.Query("SELECT * FROM users")
    if err != nil {
        return err
    }
    defer rows.Close()

    var firstName, lastName string
    var id int

    for rows.Next() {
        err := rows.Scan(&id, &firstName, &lastName)
        if err != nil {
            return err
        }
        fmt.Println(id, firstName, lastName)
    }

    if err = rows.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Println("---------------------------")
    
    return nil
}

func main() {
    // connect to database
    db, err := sql.Open("pgx", "host=localhost port=5432 user=postgres dbname=test-connect sslmode=disable")
    if err != nil {
        log.Fatal("error connecting to the database: ", err)
    }
    defer db.Close()
    fmt.Println("Successfully connected to database!")

    // test database connection
    err = db.Ping()
    if err != nil {
        log.Fatal("error pinging database: ", err)
    }
    fmt.Println("Successfully pinged database!")

    // get content from database
    err = getAllRows(db)
    if err != nil {
        log.Fatal("error getting content from database: ", err)
    }
    fmt.Println("Successfully retrieved all rows from database!")

    // insert a row into database
    query := `insert into users (first_name, last_name) values ($1, $2)`
    _, err = db.Exec(query, "John", "Doe")
    if err != nil {
        log.Fatal("error inserting row into database: ", err)
    }
    fmt.Println("Successfully inserted row into database!")

    // get content from database
    err = getAllRows(db)
    if err != nil {
        log.Fatal("error getting content from database: ", err)
    }

    // update a row in database
    query = `update users set first_name = $1 where id = $2`
    _, err = db.Exec(query, "Perica", 1)
    if err != nil {
        log.Fatal("error updating row in database: ", err)
    }
    fmt.Println("Successfully updated row in database!")

    // get content from database
    err = getAllRows(db)
    if err != nil {
        log.Fatal("error getting content from database: ", err)
    }

    // get row by id from database
    query = `select * from users where id = $1`
    row := db.QueryRow(query, 1)
    var firstName, lastName string
    var id int
    err = row.Scan(&id, &firstName, &lastName)
    if err != nil {
        log.Fatal("error getting row by id from database: ", err)
    }
    fmt.Println(id, firstName, lastName)
    fmt.Println("Successfully retrieved row by id from database!")

    // delete row from database
    query = `delete from users where id = $1`
    _, err = db.Exec(query, 1)
    if err != nil {
        log.Fatal("error deleting row from database: ", err)
    }
    fmt.Println("Successfully deleted row from database!")

    // get content from database
    err = getAllRows(db)
    if err != nil {
        log.Fatal("error getting content from database: ", err)
    }
}
