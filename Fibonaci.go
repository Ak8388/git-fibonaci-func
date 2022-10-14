package fibonaci

func Fibonaci(nilai int)[]int{
	var tamp = 1
	var tamp2 = 0
	var angka []int
	for i := 0; i<nilai; i++{

		if i % 2 == 0{
			angka = append(angka, tamp2)
			tamp2 += tamp
		} else{
			angka = append(angka, tamp)
			tamp += tamp2
		}
	}
	return angka
}