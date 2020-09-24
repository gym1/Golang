
package main
import "fmt"
import "math"
import "strconv"


func fib_1(fib_n []int, N int) int {
   
    if N < 2 {
        return N
    }

    fib_n[0] = 0
    fib_n[1] = 1
    fib_n[2] = 1
   
    for i := 3; i <= N; i++ {

        fib_n[i] = fib_n[i-1] + fib_n[i-2]

    }

    return fib_n[N]

}

// A math approach - golden ratio
// F(n) = (golden ratio ^ n) / sqrt(5)
func fib_2(fib_m [] int, N int) int {
    var goldenRatio float64 = float64((1 + math.Sqrt(5)) / 2);
    
    fib_m[N] = int(math.Round(math.Pow(goldenRatio, float64(N)) / math.Sqrt(5)))
    
    return fib_m[N];
}

func main() {
    
    var fib_count int
    var result int
    fmt.Printf("Please Enter Fibonacci Sequence Number\n")
    fmt.Scan(&fib_count)
    fib_n := make([]int, fib_count+3)
    // fib_m := make([]int, fib_count+3)
    result = fib_1(fib_n, fib_count)
    // result = fib_2(fib_n, fib_count)
    fmt.Printf("\n")
    fmt.Print("Fibonacci Number\n", result)
    fmt.Printf("\n")
    fmt.Printf("Fibonacci Sequence\n")
    for i := 0; i < fib_count; i++ {
       fmt.Print(strconv.Itoa(fib_n[i]) + " ")
    }
}
