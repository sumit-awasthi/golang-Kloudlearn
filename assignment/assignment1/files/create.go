//creating a file 

package main

import (
    "fmt"
    "io"
    "os"
)

var path = "test.txt"

func main() {
    createFile()
}
func createFile() {
    // check if file exists
    var _, err = os.Stat(path)

    // create file if not exists
    if os.IsNotExist(err) {
        var file, err = os.Create(path)
        if isError(err) {
            return
        }
        defer file.Close()
    }

    fmt.Println("File Created Successfully", path)
}