package thirdpartypackage

import (
	"database/sql"
	"fmt"
	"log"
)

var DbConnection *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	DbConnection, err := sql.Open(
		"mysql",
		// fmt.Sprintf("practice:practice@tcp(127.0.0.1:%d)/practice?parseTime=true", 33306),
		fmt.Sprintf("practice:practice@tcp(practice-db:%d)/practice?parseTime=true", 3306),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer DbConnection.Close()

	cmd := `CREATE TABLE IF NOT EXISTS person(
		name VARCHAR(50),
		age INTEGER
		)`
	_, err = DbConnection.Exec(cmd)
	if err != nil {
		log.Fatal(err)
	}

	// cmd = `INSERT INTO person(name, age) VALUES (?, ?)`
	// _, err = DbConnection.Exec(cmd, "Nancy", 20)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// cmd = `UPDATE person SET  age = ? WHERE name  = ?`
	// _, err = DbConnection.Exec(cmd, 21, "Nancy")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// cmd = `SELECT  * FROM  person`
	// rows, _ := DbConnection.Query(cmd)
	// defer rows.Close()
	// var pp []Person
	// for rows.Next() {
	// 	var p Person
	// 	err := rows.Scan(&p.Name, &p.Age)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	pp = append(pp, p)
	// }
	// err = rows.Err()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, p := range pp {
	// 	fmt.Println(p.Name, p.Age)
	// }
	// log.Println("#####################")

	// cmd = `SELECT * FROM person where age = ?`
	// rows := DbConnection.QueryRow(cmd, 21)
	// var p Person
	// err = rows.Scan(&p.Name, &p.Age)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		log.Println("no row")
	// 	} else {
	// 		log.Println(err)
	// 	}
	// }
	// log.Println(p.Name, p.Age)

	cmd = `DELETE FROM person WHERE name  = ?`
	_, err = DbConnection.Exec(cmd, "Nancy")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("#####################")
}
