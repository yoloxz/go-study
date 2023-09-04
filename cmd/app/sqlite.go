package app

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func NewSql() {
	// 打开数据库
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 创建表
	createTableSQL := `CREATE TABLE users (
                id INTEGER PRIMARY KEY,
                name TEXT,
                email TEXT,
                age INTEGER
            );`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		panic(err)
	}

	// 插入数据
	insertSQL := `INSERT INTO users (name, email, age) VALUES ("John Doe", "john@example.com", 25);`
	_, err = db.Exec(insertSQL)
	if err != nil {
		panic(err)
	}

	// 查询数据
	rows, err := db.Query("SELECT id, name, email, age FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, age int
		var name, email string
		err = rows.Scan(&id, &name, &email, &age)
		if err != nil {
			panic(err)
		}
		fmt.Printf("ID: %d, Name: %s, Email: %s, Age: %d\\n", id, name, email, age)
	}
}
