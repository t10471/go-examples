package server

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"code.cloudfoundry.org/bytefmt"
)


func HttpMain(host, csvPath string, addContentLength bool) error {
	r := resCSV{csvPath: csvPath, addContentLength: addContentLength}
	http.HandleFunc("/csv", r.responseCSV)
	log.Println("starting server")
	return http.ListenAndServe(host, nil)
}

type resCSV struct {
	csvPath string
	addContentLength bool
}

func (c resCSV) responseCSV(w http.ResponseWriter, r *http.Request)  {
	f, _ := os.Open(c.csvPath)
	defer f.Close()

	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment; filename=res.csv")
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	file, err := os.Open(c.csvPath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("file open error: %v", err)
	}
	defer file.Close()
	cLength := 0
	if c.addContentLength {
		fileScanner := bufio.NewScanner(file)
		for fileScanner.Scan() {
			cLength += len(fileScanner.Bytes())
		}
		w.Header().Set("Content-length", strconv.Itoa(cLength))
		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			log.Printf("filed file.Seek: %v", err)
			return
		}
	}
	fileScanner := bufio.NewScanner(file)

	total := 0
	for fileScanner.Scan() {
		i, err := w.Write(append(fileScanner.Bytes(), []byte("\n")...))
		if err != nil {
			log.Printf("file fileScanner error: %v", err)
		}
		log.Printf("write bytes %d", i)
		total += i
	}
	if err := fileScanner.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("fileScanner.Err error: %v", err)
	}
	log.Printf("content-length %s sent size %s",
		bytefmt.ByteSize(uint64(cLength)),
		bytefmt.ByteSize(uint64(total)))
}