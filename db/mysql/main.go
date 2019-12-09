/*
  very, very basic example for connecting mysql and map data to struct

  @author Mikiya Michishita
*/

package main

import (
	// "database/sql"
	"fmt"
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
}
