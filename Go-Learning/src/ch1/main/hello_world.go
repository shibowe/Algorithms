package main //包，表明代码所在的模块（包）

import (
	"fmt"
	"os"
)

func main(){
	fmt.Println(os.Args)
	if len(os.Args)>1{
		fmt.Println("hello world,",os.Args[1])
	}
}