package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type SingleTable struct {
	Id          int
	Key1        string
	Key2        int
	Key3        string
	KeyPart1    string
	KeyPart2    string
	KeyPart3    string
	CommonField string
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(192.168.1.135:3306)/test_mysql")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//createData(db)
	//showData(db)
	compare(db)
}

func compare(db *sql.DB) {
	prev := time.Now()

	//换个顺序就变了,笑死

	db.QueryRow("select * from single_table where `key1` = ?", "key1 0")
	fmt.Println("二级索引搜索花费的时间：", time.Since(prev))

	prev = time.Now()

	db.QueryRow("select * from single_table where `id` = ?", 1)
	fmt.Println("主键搜索花费的时间：", time.Since(prev))

}

func createData(db *sql.DB) {
	for i := 0; i < 10000; i++ {
		res, err := db.Exec("insert into single_table(key1,key2,key3,key_part1,key_part2,key_part3,common_field) values (?,?,?,?,?,?,?)",
			fmt.Sprintln("key1", i),
			i,
			fmt.Sprintln("key3", i),
			fmt.Sprintln("key_part1", i),
			fmt.Sprintln("key_part2", i),
			fmt.Sprintln("key_part3", i),
			fmt.Sprintln("common_field", i),
		)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println(i, res)
	}
}

func showData(db *sql.DB) {
	rows, err := db.Query("select * from single_table limit 10")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		body := new(SingleTable)
		if err := rows.Scan(&body.Id, &body.Key1, &body.Key2, &body.Key3, &body.KeyPart1, &body.KeyPart2, &body.KeyPart3, &body.CommonField); err != nil {
			log.Println(err)
		} else {
			fmt.Printf("%v\n", *body)
		}
	}
}
