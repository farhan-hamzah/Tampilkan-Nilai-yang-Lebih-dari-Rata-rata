package main
import "fmt"

func main(){
	const NMAX int = 100
	var A[NMAX]int
	var B[NMAX]int
	var i, n, j, count int
	var rata2 float64
	fmt.Scan(&n)

	for i = 0; i < n; i++{
		fmt.Scan(&A[i])
	}
	for i = 0; i < n; i++{
		for j = i+1; j < n; j++{
			rata2 = float64(A[i] + A[j])/2
			if rata2 > 5{
				B[count] = A[i]
				count++
				B[count] = A[j]
				break
			}
		}
	}
	fmt.Print("[")
    for i = 0; i < count; i++ {
        if i > 0 {
            fmt.Print(" ")
        }
        fmt.Print(B[i])
    }
    fmt.Println("]")

}