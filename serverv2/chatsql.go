package main

import (
	"database/sql"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func Rfromsql(ms []byte) []byte {
	m := string(ms)
	switch {
	case m[0] == '@':
		a := strings.Replace(m, "@", "", 1)
		return []byte(a)
	case m[0] == '#':
		a := strings.Replace(m, "#", "", 1)
		if a == "" {
			return []byte("Ghost")
		} else {
			db, _ := sql.Open("mysql", "inmyroom:hyhy@tcp(127.0.0.1:3306)/mysql?")
			defer db.Close()
			rows, _ := db.Query("SELECT*FROM chatting.clientslist")
			for rows.Next() {
				var id int
				var name string
				var password string
				rows.Scan(&id, &name, &password)
				if a == name {
					return []byte(name + "isok")
				}
			}
			return []byte("Ghostisok")
		}
	case m[0] == '%':
		a := strings.Replace(m, "%", "", 1)
		db, _ := sql.Open("mysql", "inmyroom:hyhy@tcp(127.0.0.1:3306)/mysql?")
		defer db.Close()
		rows, _ := db.Query("SELECT*FROM chatting.clientslist")
		var name string
		for rows.Next() {
			var id int
			var password string
			rows.Scan(&id, &name, &password)
			if a == password {
				return []byte(name + "isok")
			}
		}
		return []byte(name + "isnotok")
	default:
		return []byte("Ghostisok")
	}
}
