package dersler

import (
	"fmt"
)

type borderException struct{
	parameter int
	message string
}
//pointer vermek performansı artırıyor
func (b *borderException) Error() string{
	return fmt.Sprintf("%d---%s ", b.parameter,b.message)
}

func TahminEt2(tahmin int) (string,error) {
	if tahmin <1 || tahmin >100{
		return "", &borderException{tahmin,"sınırların dışında kaldı"}
	}
	return "Bildiniz", nil //boş değer gibi yollamak
}

