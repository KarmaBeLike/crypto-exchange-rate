package main

import "log"

func main() {
	initDB()
	defer db.Close()

	go connectToBinance()

	r := setupRouter()

	log.Println("Server is running on port 8080")
	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Server Error:", err)
	}
}
