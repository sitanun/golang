package main
import "fmt"

func main() {
   alphabets := [5]string{"A", "B", "C", "D", "E"}
   fmt.Println(alphabets)

   x := alphabets[0:2]
   fmt.Println(x) 

   y := x[2:4]
   fmt.Println(y)

   z := y[0:1]
   fmt.Println(z)
}