package main

import (
   "fmt"
   _ "os"
   "flag"
   "strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
/*
 * 초기 값("")에 의존하는 변수 선언
   var s, sep string
   for i := 1; i < len(os.Args); i++ {
      s += sep + os.Args[i]
      sep = " "
   }


 * 짧은 변수 선언 : 함수 안에서만 가능 
   s, sep := "", ""
   for _, arg := range os.Args[1:] {
      s += sep + arg
      sep = " "
   }

 * 변수 s는 for loop 동안 이전 값이 가비지 컬렉션 되므로 변경이 필요 
   fmt.Println("name : " + os.Args[0] + " params : " + strings.Join(os.Args[1:], " "))

 * flag 이용 */
   flag.Parse() // 플래그 변수들의 기본값을 갱신
   fmt.Print(strings.Join(flag.Args(), *sep))
   if !*n {
      fmt.Println()
   }   
}
