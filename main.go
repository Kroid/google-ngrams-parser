package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
)

var (
	start_year   = 2000
	total_counts = make(map[int]int)

	files_path        = "/home/kroid/research/english/raw/extracted/"
	total_counts_path = "/home/kroid/research/english/raw/googlebooks-eng-all-totalcounts-20120701.txt"
)

func init() {
	b, err := ioutil.ReadFile(total_counts_path)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "	")
	for _, line := range lines {
		fields := strings.Split(line, ",")

		year, err := strconv.Atoi(fields[0])
		if err != nil {
			log.Fatal(err)
		}

		count, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatal(err)
		}

		total_counts[year] = count
	}
}

func main() {
	// return
	files_info, err := ioutil.ReadDir(files_path)
	if err != nil {
		panic(err)
	}

	for _, ele := range files_info {
		basename := ele.Name()
		parse_file(path.Join(files_path, basename))
	}
}

func parse_file(filepath string) (err error) {
	fmt.Println(filepath)
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parse_line(scanner.Text())
	}

	return nil
}

func parse_line(line string) {

}
