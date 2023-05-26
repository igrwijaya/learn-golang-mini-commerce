package cli

import (
	"fmt"
	"os"
)

func MainMenu() {
	fmt.Println("Selamat Data di Mini Ecommerce")
	fmt.Println("------------------------------")

	var input string
	fmt.Println("Tekan (1) untuk melanjutkan ke list Produk")
	fmt.Println("Tekan (2) untuk melanjutkan ke list Order")
	fmt.Println("Tekan (q) untuk keluar dari aplikasi")

	_, err := fmt.Scanln(&input)

	if err != nil {
		fmt.Println(err.Error())
	}

	switch input {
	case "1":
		ListProduct()
	case "2":
	//
	case "q":
		os.Exit(1)
	default:
		MainMenu()
	}
}
