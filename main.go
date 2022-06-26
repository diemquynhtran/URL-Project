package main
 
import (
    "fmt"
    //"bufio"
    //"os"
    //"strconv"

)

func main() {

    /*input
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
    fmt.Print(2022- input)
    */

    /* MAP
    var listMap map[string]int = map[string]int{
        "apple": 5,
        "bear": 6,
        "orange": 10,
    }

    val, ok := listMap["bear"]

    fmt.Println(val, ok)

    mapTwo := make(map[string]int)
    */
    var x [2]int = [2]int{1,2}
    y := x
    y[0] = 100
    fmt.Println(x,y)

}

func absolute(x int) (result int) {
    if x < 0 {
        result = -x
        return
    }
    result = x
    return 
}

