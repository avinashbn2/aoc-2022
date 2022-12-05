package main

import (
  "os"
  "bufio"
  "log"
  "strconv"
  "fmt"

)
func main() {
  f, err := os.Open("dat.txt")
  if err!=nil {
    log.Fatal(err)
  }
  var sum , maxSum int64

  defer f.Close()
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    str:= scanner.Text()
    if str!="" {

    val, err:=strconv.ParseInt(str, 10, 64)
    if err!=nil {
      log.Fatal(err)
    }
      sum+=val
    }else{
      if maxSum<sum {
        maxSum = sum
      }
      sum=0
    }

    
  }

      if maxSum < sum {
        maxSum = sum
      }
      fmt.Println("sum", sum, maxSum)
  if err:= scanner.Err(); err!=nil {
    
    log.Fatal(err)
  }
  
}
