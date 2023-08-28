package main
import "fmt"
func main() {
    learnGoTo()
}
func learnGoTo() {
    i := 0
    fmt.Println("here")
START:
    for i < 10 {
        if i%2 == 0 {
            i++
            goto START
        }  
        fmt.Println(i)
        i++
    }
}