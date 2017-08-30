// package main

// var a = "1321"
// var b = 123
// var c = true

// var (  // 这种因式分解关键字的写法一般用于声明全局变量
//     g int
//     h bool
// )

// const I int = 10

// func main() {
// 	d, e, f := a, b, c
// 	J := I
// 	J = 1
// 	println(d, e, f, J)
// 	println(&a,&d)
// }


////**************************************************************
////go const itoa
////**************************************************************
// package main

// import "fmt"
// const (
// 		a = iota   //0
// 		b          //1
// 		c          //2
// 		d = "ha"   //独立值，iota += 1
// 		e          //"ha"   iota += 1
// 		f = 100    //iota +=1
// 		g          //100  iota +=1
// 		h = iota   //7,恢复计数
// 		i          //8
// )
// func main() {
    
// 	fmt.Println(a,b,c,d,e,f,g,h,i)
	
// 	var dd  = 1
// 	var cc *int
// 	cc = &dd
// 	println(&dd,cc)

// 	if(a == 1){
	
// 	}else if( a == 2){

// 	}else{

// 	}
// }

////**************************************************************
////go const select
////**************************************************************
// package main

// import "fmt"

// func main() {
//    var c1, c2, c3 chan int
//    var i1, i2 int
//    select {
//       case i1 = <-c1:
//          fmt.Printf("received ", i1, " from c1\n")
//       case c2 <- i2:
//          fmt.Printf("sent ", i2, " to c2\n")
//       case i3, ok := (<-c3):  // same as: i3, ok := <-c3
//          if ok {
//             fmt.Printf("received ", i3, " from c3\n")
//          } else {
//             fmt.Printf("c3 is closed\n")
//          }
//       default:
//          fmt.Printf("no communication\n")
//    }    
// }

////**************************************************************
////函数闭包
////**************************************************************
// package main

// import "fmt"

// func getSequence() func() int {
//    i:=0
//    return func() int {
//       i+=1
//      return i  
//    }
// }

// func main(){
//    /* nextNumber 为一个函数，函数 i 为 0 */
//    nextNumber := getSequence()  

//    /* 调用 nextNumber 函数，i 变量自增 1 并返回 */
//    fmt.Println(nextNumber())
//    fmt.Println(nextNumber())
//    fmt.Println(nextNumber())
   
//    /* 创建新的函数 nextNumber1，并查看结果 */
//    nextNumber1 := getSequence()  
//    fmt.Println(nextNumber1())
//    fmt.Println(nextNumber1())
// }

// package main

// import (
//    "fmt"  
// )

// /* 定义函数 */
// type Circle struct {
//   radius float64
// }

// func main() {
//   var c1 Circle
//   c1.radius = 10.00
//   fmt.Println("Area of Circle(c1) = ", c1.getArea(2,"13123"))
// }

// //该 method 属于 Circle 类型对象中的方法
// func (cd Circle) getArea(t int, xx string) float64 {
//   //c.radius 即为 Circle 类型对象中的属性
//   return 3.14 * cd.radius * cd.radius * float64(t)
// }

package main

import (
  "fmt"
)

//init在main函数前被调用
func init(){
  fmt.Println("111")
}

func main(){
  fmt.Println("222")
}