package main
import(
 "github.com/golang-collections/collections/stack"
 "fmt")
func main() {
    s := stack.New()
    s.Push(5)
    s.Push(3)
    s.Push(4)
for s.Len() != 0 {
    val := s.Pop()
fmt.Print(val, " ")
}
}