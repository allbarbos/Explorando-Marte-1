package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func main() {
	var limiteInfX, limiteInfY int = 0, 0
	var limiteSupX, limiteSupY int
	var sondaX, sondaY int
	var sondaDir string
	var inputAction string
	var textFile string

	fmt.Print("\nInsira o nome do arquivo (com extensão): ")
	fmt.Scanln(&textFile)
	data, err := ioutil.ReadFile(textFile)    	
    for err != nil {
    	fmt.Print(err)
    	fmt.Print("\nInsira o nome do arquivo (com extensão): ")
		fmt.Scanln(&textFile)
		data, err = ioutil.ReadFile(textFile)    	
    }
    entrada := strings.Fields(string(data))

    //Limite direito superior
    limiteSupX, err = strconv.Atoi(entrada[0])
    limiteSupY, err = strconv.Atoi(entrada[1])

    for i := 2; i < len(entrada); i+=4 {
    	//Posição X e Y da sonda
    	sondaX, err = strconv.Atoi(entrada[i])
    	sondaY, err = strconv.Atoi(entrada[i+1])
    	//Direção da sonda
    	sondaDir = entrada[i+2]
    	//Sequência de ações
    	inputAction = entrada[i+3]

		for _, r := range inputAction {
			caracter := string(r)

			switch caracter {
			case "M":
				{
					if sondaDir == "N" && sondaY+1 <= limiteSupY {
						sondaY = sondaY + 1
						break
					}

					if sondaDir == "E" && sondaX+1 <= limiteSupX {
						sondaX = sondaX + 1
						break
					}

					if sondaDir == "W" && sondaX-1 >= limiteInfX {
						sondaX = sondaX - 1
						break
					}

					if sondaDir == "S" && sondaY-1 >= limiteInfY {
						sondaY = sondaY - 1
						break
					}
				}
			case "R":
				{
					if sondaDir == "N" {
						sondaDir = "E"
						break
					}

					if sondaDir == "E" {
						sondaDir = "S"
						break
					}

					if sondaDir == "W" {
						sondaDir = "N"
						break
					}

					if sondaDir == "S" {
						sondaDir = "W"
						break
					}
				}
			case "L":
				{
					if sondaDir == "N" {
						sondaDir = "W"
						break
					}

					if sondaDir == "E"	{
						sondaDir = "N"
						break
					}

					if sondaDir == "W"	{
						sondaDir = "S"
						break
					}

					if sondaDir == "S"	{
						sondaDir = "E"
						break
					}
				}
			default:
				fmt.Println("Entrada inválida. As entradas possíveis são: L, R ou M")
			}
		}

		fmt.Println(sondaX, sondaY, sondaDir)
	}
}