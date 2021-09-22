package dersler


func Ankara(){
	println("a fonksiyonu başladı")
}

func Bursa()  {
	println("b fonksiyonu başladı")
	defer Ankara() // üstündeki kodlar çalıştıktan sonra çalıştır
}