package queries

import (
	"database/connection"
)

func GetImages() {

}

// Implement batch insert for images
func InsertImages(images []ParsedInfo) {
	conn := connection.GetDB()
	defer conn.Close()
}
