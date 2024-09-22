package bot

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
)

func getPhoto(path string) string {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("Error occurred while getting photo by path")
		return ""
	}
	var base64Encoding string
	mimeType := http.DetectContentType(bytes)
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}
	base64Encoding += base64.StdEncoding.EncodeToString(bytes)
	log.Println("Photo converted to base64Encoding", base64Encoding)
	return base64Encoding
}
