package dersler

import (
	"fmt"
	"os"
)

func edt(){
	f, err := os.Open("demo1.txt")
	if err != nil{
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("bulunamadı", pErr.Path)
			return
		}
	} else{
		fmt.Println(f.Name())
	}
}