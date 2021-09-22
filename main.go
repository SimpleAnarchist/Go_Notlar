package main

import (
	"Seyis/dersler"
	"fmt"
)

func main(){
	//dersler.Bursa() //derslerden bursayı çalıştır
	//fmt.Println(dersler.TahminEt2(44))
	//dersler.GonderPost()
	//dersler.AddProduct()
	products,_ := dersler.GetAllProducts()
	//fmt.Println(products)

	for i := 0; i < len(products); i++ {
		fmt.Println(products[i].ProductName)
	}
}