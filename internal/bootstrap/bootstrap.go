package bootstrap

import (
	"github.com/nicolasvasquez/matchProducts/internal/models"
	"fmt"
)

func Init() {
	uno := models.Products{Title: "Cubre Asiento 3d Azul Sparco Funda 30161"}

	dos := models.Products{Title: "Cubre Asiento 3d Azul Sparco Funda 30161 / Fernapet"}
	tres := models.Products{Title: "Cubre Asiento Auto 3d Azul Sparco Funda 30161 / Fernapet"}
	cuatro := models.Products{Title: "Cubre Asiento Auto 3d Gris Sparco Funda 30162 / Fernapet"}
	cinco := models.Products{Title: "Kit Sparco Funda Individual 3d + Cubre Volante Azul / Zf"}

	fmt.Println(uno.TitleCosine(dos))
	fmt.Println(uno.TitleCosine(tres))
	fmt.Println(uno.TitleCosine(cuatro))
	fmt.Println(uno.TitleCosine(cinco))

	fmt.Printf("\n")

	fmt.Printf("TEST\n")

	fmt.Println(uno.TitleProcessedCosine(dos))
	fmt.Println(uno.TitleProcessedCosine(tres))
	fmt.Println(uno.TitleProcessedCosine(cuatro))
	fmt.Println(uno.TitleProcessedCosine(cinco))
}