package img

import (
	"io"
	"net/http"
	"os"
)

func UrlToFile(url string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	return createImg(res, "sample")
}

func createImg(res *http.Response, str string) error {

	var file *os.File
	var err error

	file, err = os.Create("./img/" + str)
	defer file.Close()
	if err != nil {
		return err

	}
	_, err = io.Copy(file, res.Body)
	if err != nil {
		return err
	}

	return nil
}
