# PACKET
## fmt
# ARRAY
- Kiểu tham trị
```Golang
    var x [2]int = [2]int{1,2}
    y := x
    y[0] = 100
    fmt.Println(x,y)
    //Output: [1 2] [100 2]
```
# SLICE
- Kiểu tham chiếu
```Golang
    var x []int = []int{1,2}
    y := x
    y[0] = 100
    fmt.Println(x,y)
    //Output: [100 2] [100 2]
```
# MAP
## Syntax
```Golang
    var listMap map[string]int = map[string]int{...}
Or:
    listMap := make(map[string]int)
String/int are type of key/value

```
- Check element:
```Golang
    val, ok := listMap[key]
```
if key isn't exist ok=false
# FUNCTION
- defer dùng để trì hoãn câu lệnh đến cuối cùng mới thực hiện (trước return)


