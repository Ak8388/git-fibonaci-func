package fibonaci

func Fibonaci(nilai int)[]int{
	var tamp = 1
	var angka []int
	for i := 0; i<nilai; i++{

		if i % 2 == 0{
			angka = append(angka, tamp)
			tamp += i
		} else{
			angka = append(angka, i)
		}
	}
	return angka
}