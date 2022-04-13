// Utilizamos a palavra var, o nome da variável que queremos referenciar e seu tipo.

package main

import (
	"fmt"
)

//Utilizamos := (dois pontos, igual), usamos apenas quando vamos criar e atribuir um valor
//dentro do escopo de uma função.

func main() {

	idade := 15
	altura := 1.78
	nome := "Guilherme"
	fmt.Println("Nome:", nome, "-", "Idade:", idade, "anos", "-", "Altuda:", altura)
}

//Quando a variável receber um novo valor, devemos utilizar apenas o sinal de = (igual).
// Caso contrário, receberemos a seguinte mensagem dizendo que não há novas variáveis seguido do
