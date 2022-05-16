package main

import "fmt"

func main() {
	// Kiritilgan 3ta sonnning yigindisi kopaytmasi, o'rtalamasi,Katta va kichigini topish.
	var son1, son2, son3, kopaytma, yigindi, ortalama, katta, kichik float32
	fmt.Println("Birinchi sondi kiriting:")
	fmt. Scanf("%v", &son1)
	fmt.Println("Ikkinchi sondi kiriting:")
	fmt. Scanf("%v", &son2)
	fmt.Println("Uchinchi sondi kiriting:")
	fmt. Scanf("%v", &son3)


	yigindi = son1 + son2 + son3
	kopaytma = son1 * son2 * son3
	ortalama = yigindi / 3
	kichik = son1
	if son2 < kichik {
		kichik = son2
	}
	if son3 < kichik {
		kichik = son3
	}

	katta = son1
	if son2 > katta {
		katta = son2
	}
	if son3 > katta {
		katta = son3
	}

	fmt.Println("Katta son:", katta, "Kichik son:", kichik, "O'rta arifmetik:", ortalama, "Yig'indi:", yigindi, "Ko'paytma:", kopaytma)
}
