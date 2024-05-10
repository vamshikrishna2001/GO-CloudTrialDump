package Utils

import (
	"log"
    "bufio"
    "os"
    "encoding/json"
)

func AWS_resource_reader_from_txt(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        log.Fatal("error reading the file")
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}

func Json_writer(path string,dict interface{}){
    file, err := os.Create(path)
    if err != nil {
        log.Fatal("Error creating file")
        
    }
    defer file.Close()
    encoder := json.NewEncoder(file)
    err = encoder.Encode(dict)

    if err != nil{
        log.Fatal("Error creating file")
    }
}
