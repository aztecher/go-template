/*
  very, very basic example for connecting mysql and map data to struct

  @author Mikiya Michishita
*/

package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

// +--------+----------------+
// | id     | name           |
// +--------+----------------+
// | 1      | tanaka         |
// +--------+----------------+
// | 2      | tokoro         |
// +--------+----------------+
// | 3      | shimabara       |
// +--------+----------------+

type Person struct {
	ID   int
	Name string
}

func main() {
	fmt.Printf("mysql example\n")

	// to connect mysql, determine driver name 'mysql' and username, datasource
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	db, err := sql.Open("mysql", "root:mysql@tcp(127.0.0.1:3306)/test_template")
	defer db.Close()
	log.Println("Connected to mysql.")
	if err != nil {
		log.Fatal(err)
	}

	// send query and get data
	rows, err := db.Query("SELECT * FROM person")
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var person Person
		err := rows.Scan(&person.ID, &person.Name)

		if err != nil {
			panic(err.Error())
		}
		fmt.Println(person.ID, person.Name)
	}
}
