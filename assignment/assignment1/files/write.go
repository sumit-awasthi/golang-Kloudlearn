//writing in file

package main

import (
    "fmt"
    "io"
    "os"
)

var path = "test.txt"

func main() {
    writeFile()
}
func writeFile() {
    // Open file using READ & WRITE permission.
    var file, err = os.OpenFile(path, os.O_RDWR, 0644)
    if isError(err) {
        return
    }
    defer file.Close()

    // Write some text line-by-line to file.
    _, err = file.WriteString("Hello \n")
    if isError(err) {
        return
    }
    _, err = file.WriteString("World \n")
    if isError(err) {
        return
    }

    // Save file changes.
    err = file.Sync()
    if isError(err) {
        return
    }

    fmt.Println("File Updated Successfully.")
}
