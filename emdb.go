package main

import (
	"emdb/db"
	"log"
)

func main() {

	// configure database
	// TODO the dsn should read from configuration file
	// TODO 设置尝试连接的超时时长，默认的超时时长过长
	log.Println("try to connect to the database.")
	err := db.SetDSN("root:123456@tcp(192.168.68.30:3306)/emdb?charset=utf8")

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("connect to database successfully.")

}
