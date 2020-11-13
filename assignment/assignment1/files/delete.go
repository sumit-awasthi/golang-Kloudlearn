//deleting from a file
package main

import (
    "fmt"
    "io"
    "os"
)

var path = "test.txt"

func main() {
    deleteFile()
}
func deleteFile() {
    // delete file
    var err = os.Remove(path)
    if isError(err) {
        return
    }

    fmt.Println("File Deleted")
}

func isError(err error) bool {
    if err != nil {
        fmt.Println(err.Error())
    }

    return (err != nil)
}
