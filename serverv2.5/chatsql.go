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
			return []byte("!Ghostisok")
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
					return []byte("!" + name + "isok")
				}
			}
			return []byte("!Ghostisok")
		}
	case m[0] == '%':
		a := strings.Replace(m, "%", "", 1)
		c := strings.Split(a, "%")
		db, _ := sql.Open("mysql", "inmyroom:hyhy@tcp(127.0.0.1:3306)/mysql?")
		defer db.Close()
		rows, _ := db.Query("SELECT*FROM chatting.clientslist")
		var name string
		for rows.Next() {
			var id int
			var password string
			rows.Scan(&id, &name, &password)
			if c[0] == name {
				if c[1] == password {
					return []byte("!" + name + "isok")
				} else {
					return []byte("!" + name + "isnotok")
				}
			}
		}
		return []byte("!endisnotok")
	case m[0] == '?':
		a := strings.Replace(m, "?", "", 1)
		c := strings.Split(a, "@")
		if c[0] != "" && c[1] != "" {
			db, _ := sql.Open("mysql", "inmyroom:hyhy@tcp(127.0.0.1:3306)/mysql?")
			defer db.Close()
			rows, _ := db.Query("SELECT*FROM chatting.clientslist")
			var name string
			var id int
			for rows.Next() {
				var password string
				rows.Scan(&id, &name, &password)
				if c[0] == name {
					return []byte(name + " exists")
				}
			}
			stmt, _ := db.Prepare("INSERT INTO chatting.clientslist(id,name,password)VALUES(?,?,?)")
			defer stmt.Close()
			stmt.Exec(id+1, c[0], c[1])
			return []byte(c[0] + " joins us!\\(^v^)/")
		}
		return []byte("Attention!!No blank password!@(#*w*)@")
	default:
		return []byte("!Ghostisok")
	}
}
