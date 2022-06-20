package main
import (
 "fmt"
)

type MyQueue struct {
  queueArr []int
}

func Constructor() MyQueue {
  var queueObject  MyQueue
  return queueObject
}

func (this *MyQueue) Push(x int) {
  this.queueArr = append(this.queueArr, x)
}

func (this *MyQueue) Pop() int {
  firstElem := this.queueArr[0]
  if len(this.queueArr) > 1 {
    this.queueArr = this.queueArr[1:len(this.queueArr)]
  } else {
    this.queueArr = []int{}
  }
  return firstElem
}

func (this *MyQueue) Peek() int {
  return this.queueArr[0]
}

func (this *MyQueue) Empty() bool {
  return len(this.queueArr) == 0
}

func main() {
  obj := Constructor()
  obj.Push(1)
  fmt.Println(obj.Peek())
  obj.Pop()
  fmt.Println(obj.Empty())

}
