package main  /* package 必须在 import 之前
               不是 main 的包都叫做库。
               类似 if __name__ == '__main__' 的作用
               说明这个程序可以独立执行
              */

import "fmt"  // 导入标准输出的库

func main() {
    fmt.Printf("Hello, world; 世界，你好")
}
