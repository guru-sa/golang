package main

/*
 * Stream mode
 *
import (
   "bufio"
   "fmt"
   "os"
)

func main() {
   counts := make(map[string]int) // NOTE: 키가 string이고, 값이 int인 map 생성
   files := os.Args[1:]
   if len(files) == 0 {
      countLines(os.Stdin, counts)
   } else {
      for _, arg := range files {
         f, err := os.Open(arg)
	 if err != nil {
            fmt.Fprintf(os.Stderr, "dup: %v\n", err)
	    continue
	 }
	 countLines(f, counts)
	 f.Close()
      }
   }
   for line, n := range counts {
      if n > 1 {
         fmt.Println("%d\t%s\n", n, line)
      }
   }
}

func countLines(f *os.File, counts map[string]int) {
   input := bufio.NewScanner(f)
   for input.Scan() {
      fmt.Println(input.Text())
      counts[input.Text()]++
   }
   // NOTE: input.Err()에서의 잠재적 오류는 무시
}

 * Buffer Mode 
 */

import (
   "fmt"
   "io/ioutil"
   "os"
   "strings"
)

func main() {
   counts := make(map[string]int)
   for _, filename := range os.Args[1:] {
      data, err := ioutil.ReadFile(filename)
      if err != nil {
         fmt.Fprintf(os.Stderr, "dup: %v\n", err)
	 continue
      }
      for _, line := range strings.Split(string(data), "\n") {
         counts[line]++
      }
   }
   for line, n := range counts {
      if n > 1 {
         fmt.Printf("%d\t%s\n", n, line)
      }
   }
}

