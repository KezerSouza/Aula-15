package main
import "fmt"


func pow(y, z int) int {
    if(z == 0) {
        return 1
    }else {
        return y * pow(y, z-1)
    }
	
	
}

func main() {
    var x int
    var n int
    fmt.Scan(&x)
    fmt.Scan(&n)
	fmt.Println(pow(x, n))
}
