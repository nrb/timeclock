package main

import "log"

func main() {
	r := CreateRoutes()

	db, err := NewDB(DB_MODE_MEM)
	if err != nil {
		log.Fatalf("Could not create database: %s", err)
	}

	s := NewServer("", "", r, db)

	s.Run()

}
