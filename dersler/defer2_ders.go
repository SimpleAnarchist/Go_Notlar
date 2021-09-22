package dersler

func test(number int32) string  {
	defer prnt() //çalışır çünkü daha bir şey return edilmedi
	if number%2==0{
		return "Cift Sayi"
	}
	if number%2 != 0 {
		return "Tek Sayi"
	}
	prnt()
	defer prnt() //çalışmaz bunlar çünkü fonksiyonun içinde bir şey return edilirse işi bitter
	return "" //iki if çalışmazsa diye bir şeyi return etmek için go zorluyor
}

func dene(){
	sayi := test(10)
	println(sayi)
}

func prnt(){
	println("bitti")
}