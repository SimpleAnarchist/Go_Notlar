package dersler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)
// çalıştırmak için json-server --watch db.json
type Product struct{
	Id int `json:"id"` // altgr + virgül
	ProductName string `json:"productName"`
	CategoryId int `json:"categoryId"`
	UnitPrice float64 `json:"unitPrice"`
}

type Category struct {
	Id int `json:"id"`
	CategoryName string `json:"categoryName"`
}

func GetAllProducts() ([]Product,error){
	response,err := http.Get("http://localhost:3000/products")
	if err != nil{
	return nil, err
	}
	defer response.Body.Close()

	bodyByte,_ := ioutil.ReadAll(response.Body)
	var product []Product // product olan product array oluşturma

	json.Unmarshal(bodyByte,&product) //& aynı adresi kullanması için
	return product, nil
}

func AddProduct() (Product, error) {
	product := Product{ProductName: "Tablet",CategoryId: 1,UnitPrice: 3455.0}
	jsonProduct,err := json.Marshal(product)
	if err != nil{
		fmt.Println(err)
	}
	response,err :=	http.Post("http://localhost:3000/products",
		"application/json;charset=utf-8",
		bytes.NewBuffer(jsonProduct))
	if err != nil{
		return Product{}, err
	}
	defer response.Body.Close()

	bodyBytes,_ := ioutil.ReadAll(response.Body)

	var productResponse Product

	json.Unmarshal(bodyBytes,&productResponse)
	//fmt.Println("Kaydedildi: ", productResponse) // 0,0,0 döndürürse varolan bir datayı üzerine yazmadığı için
	return productResponse ,nil
}