package client

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"

	"code.cloudfoundry.org/bytefmt"
)
var mem runtime.MemStats

func PrintMemory() {
	runtime.ReadMemStats(&mem)
	log.Printf("Alloc %s, TotalAlloc %s, HeapAlloc %s, HeapSys %s",
		bytefmt.ByteSize(mem.Alloc),
		bytefmt.ByteSize(mem.TotalAlloc),
		bytefmt.ByteSize(mem.HeapAlloc),
		bytefmt.ByteSize(mem.HeapSys))
}
func HttpMain(host, csv string, useReadAll bool) error {
	PrintMemory()
	err := realMain(host, csv, useReadAll)
	PrintMemory()
	return err
}

func realMain(host, csv string, useReadAll bool) error {
	client := http.Client{}

	resp, err := client.Get(fmt.Sprintf("http://%s/csv", host))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	log.Println(resp.Header)


	file, err := os.Create(csv)
	if err != nil {
		return err
	}
	defer file.Close()
	if useReadAll {
		return readAll(resp.Body, file)
	}
	read, err := io.Copy(file, resp.Body)
	log.Printf("csv size %s", bytefmt.ByteSize(uint64(read)))
	if err != nil && err != io.EOF {
		log.Printf("failed to io.Copy: %v", err)
		return err
	}
	return nil
}

func readAll(body io.Reader, file io.Writer) error {
	b, err := io.ReadAll(body)
	log.Printf("csv size %s", bytefmt.ByteSize(uint64(len(b))))
	if err != nil {
		return err
	}
	_, err = file.Write(b)
	if err != nil {
		log.Printf("failed to Write: %v", err)
		return err
	}
	return nil
}