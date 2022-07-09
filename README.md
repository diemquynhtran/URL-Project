# PACKET
## fmt
# STRING
 https://techmaster.vn/posts/34999/series-golang-co-ban-phan-14-kieu-du-lieu-string
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
Slices không có bất kì dữ liệu nào. Chúng là các tham chiếu đến mảng hiện có.
- Kiểu tham chiếu
```Golang
    var x []int = []int{1,2}
    y := x
    y[0] = 100
    fmt.Println(x,y)
    //Output: [100 2] [100 2]
    //Tạo slides sử dụng make
    func make([]T, len, cap)
```
- len: chiều dài slice (số phần tử hiện có trong slice). Cap: dung lượng của slide, khi dung lượng của slice = len, thì khi thêm bất kì phần tử nào vào slice thì cap += phần tử x 2
# MAP
- Syntax
```Golang
    var listMap map[string]int = map[string]int{...}
    //Or:
    listMap := make(map[string]int)
    //String/int are type of key/value

```
- Check element:
```Golang
    val, ok := listMap[key]
    //if key isn't exist ok=false
```
- Kiểu tham chiếu
```Golang
    var x []int = []int{1,2}
    y := x
    y[0] = 100
    fmt.Println(x,y)
    //Output: [100 2] [100 2]
```
# FUNCTION
- defer dùng để trì hoãn câu lệnh đến cuối cùng mới thực hiện (trước return)
- Variadic Functions: https://golangbot.com/variadic-functions/ 
- Closure: truy cập một biến bên ngoài phạm vi của nó
# POINTER
Biến a chứa địa chỉ của biến b thì a được gọi là con trỏ của b
```Golang
    var a = "hello world"
    var p *string //<nil>
    p = &a //address of a
    *p //value of a
    &p //address of p

```
# STRUCT
- Method on types
```Golang
    func (r receiver) func_name(param) return_type {}
    //func này hoạt động trên đối tượng nhận vào trong receiver 
```
non-struct:
```golang
type bien string
```

# INTERFACE
‘‘Interface xác định hành vi của một đối tượng’’. Nó chỉ xác định những gì đối tượng phải làm. Cách để thực hiện các hành vi này tùy thuộc vào đối tượng đó.
- Generics func
Định nghĩa được những hàm (phương thức hay method/function) chấp nhận các tham số chung chung, không cần chỉ định rõ nó thuộc kiểu dữ liệu gì. Tới khi hàm này được sử dụng, người gọi sẽ quyết định việc này.

# GOROUTINE

