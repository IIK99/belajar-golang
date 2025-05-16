package main

import (
	"fmt"
	"math"
)

func main() {
	// math round bisa membulatkan angka ke atas atau ke bawah, tergantung angka ke yang terdekat
	fmt.Println(math.Round(1.8), "round up")
	fmt.Println(math.Round(1.4), "round down")

	// math ceil membulatkan angka ke atas
	fmt.Println(math.Ceil(1.4), "ceil up")

	// math floor membulatkan angka ke bawah
	fmt.Println(math.Floor(1.8), "floor down")

	// math abs untuk mendapatkan nilai absolut dari suatu angka
	fmt.Println(math.Abs(-1.8), "abs")

	// math max untuk mendapatkan nilai maksimum dari dua angka
	fmt.Println(math.Max(1.8, 2.5), "max")

	// math min untuk mendapatkan nilai minimum dari dua angka
	fmt.Println(math.Min(1.8, 2.5), "min")

	// math pow untuk mendapatkan pangkat dari suatu angka
	fmt.Println(math.Pow(2, 3), "pow") // 2^3 = 8

	// math sqrt untuk mendapatkan akar dari suatu angka
	fmt.Println(math.Sqrt(9), "sqrt") // 3.1622776601683795

	// math sin untuk mendapatkan sinus dari suatu angka
	fmt.Println(math.Sin(math.Pi / 2), "sin") // 1

	// math cos untuk mendapatkan cosinus dari suatu angka
	fmt.Println(math.Cos(math.Pi / 2), "cos") 

	// math tan untuk mendapatkan tangen dari suatu angka
	fmt.Println(math.Tan(math.Pi / 4), "tan") // 1

	// math asin untuk mendapatkan arcsinus dari suatu angka
	fmt.Println(math.Asin(1), "asin") 

	// math acos untuk mendapatkan arcosinus dari suatu angka
	fmt.Println(math.Acos(1), "acos") // 0

	// math atan untuk mendapatkan arctangen dari suatu angka
	fmt.Println(math.Atan(1), "atan") // 0.7853981633974483

	// math atan2 untuk mendapatkan arctangen dari dua angka
	fmt.Println(math.Atan2(1, 1), "atan2") // 0.7853981633974483

	// math mod untuk mendapatkan sisa dari suatu angka
	fmt.Println(math.Mod(10, 3), "mod") // 1

	// math exp untuk mendapatkan eksponen dari suatu angka
	fmt.Println(math.Exp(1), "exp") // 2.718281828459045

	// math log untuk mendapatkan logaritma dari suatu angka
	fmt.Println(math.Log(math.E), "log") // 1

	// math log10 untuk mendapatkan logaritma dari suatu angka
	fmt.Println(math.Log10(100), "log10") // 2

	// math log2 untuk mendapatkan logaritma dari suatu angka
	fmt.Println(math.Log2(4), "log2") // 2

	// math log1p untuk mendapatkan logaritma dari suatu angka
	fmt.Println(math.Log1p(math.E), "log1p") 
	
}