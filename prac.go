package main
import "fmt"

const NMAX = 20
type arrSuara [NMAX]int


func main()  {
	var suara arrSuara
	var input, suaraMasuk, suaraSah int
	fmt.Scan(&input)

	suaraMasuk = 0

	for input!=0{
		suara[suaraMasuk] = input
		suaraMasuk++
		if input > 0 && input <= 20 {
			suaraSah++
		}
		fmt.Scan(&input)
	}
	fmt.Println("Suara masuk : ",suaraMasuk)
	fmt.Println("Suara sah : ",suaraSah)

	for i := 1; i <=20; i++ {
		var jumlahSuara = cariSuara(suara, suaraMasuk, i)
		if jumlahSuara >0 {
			fmt.Println(i," : ",jumlahSuara)
		}
	}


}
func cariSuara(suara arrSuara, suaraMasuk int, x int)  int{
	var jumlah int
	jumlah = 0
	for i := 0; i < suaraMasuk; i++ {
		if suara[i] == x{
			jumlah++
		}
	}
	return jumlah
}