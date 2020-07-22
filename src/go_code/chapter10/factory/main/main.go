package main

import(
	"fmt"
	"go_code/chapter10/factory/model"
)

func main (){
	var stu model.Student = model.Student{
		Name : "Tom",
		Score : 99.9,
	}
	fmt.Println(stu)
} 