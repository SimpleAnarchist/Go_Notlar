package dersler

import (
	"errors"
	"fmt"
)

func tahminEt(tahmin int) (string, error){
	if tahmin <1 || tahmin >100{
		return "", errors.New("1-100 arası sayı girin")
	}
	return "bildiniz", nil
}

func Tahmin2(){
	fmt.Println(tahminEt(100))
	fmt.Println(tahminEt(50))
	fmt.Println(tahminEt(-50))
}