// ZigZag in Golang
package main
  
import (
    "fmt"
)
  
func zigzag(n int) []int {
    z := make([]int, n*n)
    i := 1
    m := n * 2
    for p := 1; p <= m; p++ {
        x := p - n
        if x < 0 {
            x = 0
        }
        y := p - 1
        if y > n-1 {
            y = n - 1
        }
        j := m - p
        if j > p {
            j = p
        }
        for k := 0; k < j; k++ {
            if p&1 == 0 {
                z[(x+k)*n+y-k] = i
            } else {
                z[(y-k)*n+x+k] = i
            }
            i++
        }
    }
  
    return z
}
  
func main() {
    num := 4
    len := 6
    for i, draw := range zigzag(num) {
        fmt.Printf("%*d ", len, draw)
        if i%num == num-1 {
            fmt.Println("")
        }
    }
}