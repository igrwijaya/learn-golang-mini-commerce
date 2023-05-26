package cli

import (
	"bufio"
	"fmt"
	"go-trial-class/config"
	"go-trial-class/entity"
	"os"
)

func ListProduct() {
	consoleReader := bufio.NewReader(os.Stdin)

	var products []entity.Product
	err := config.DB.Find(&products).Error

	if err != nil {
		ErrorHandler(err.Error())
		return
	}

	fmt.Println("*** List Products ***")

	for _, product := range products {
		product.PrintDetail()
	}

	var input string
	fmt.Println("Masukkan ID Product untuk melanjutkan order")
	fmt.Println("Tekan (m) untuk kembali ke menu utama")
	fmt.Println("Tekan (q) untuk keluar aplikasi")

	input, err = consoleReader.ReadString('\n')
	switch input {
	case "m\n":
		MainMenu()

	case "q\n":
		os.Exit(1)

	default:
		OrderProduct(input)
	}
}

func OrderProduct(id string) {
	consoleReader := bufio.NewReader(os.Stdin)

	var product entity.Product
	err := config.DB.Where("ID = ?", id).First(&product).Error

	if err != nil {
		return
	}

	product.PrintDetail()
}
