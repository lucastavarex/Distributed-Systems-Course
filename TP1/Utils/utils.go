package utils

import (
	"math"
	"math/rand"
	"time"
)

// GetSquareRoot recebe um número inteiro e retorna sua raiz quadrada.
// É necessário converter o inteiro para float64 por causa da função sqrt.
// Arredonda para cima a raiz quadrada obtida porque a saída é float64.
// Converte novamente para inteiro por causa da saída desta função.
// Retorna a raiz quadrada de um número.

func getSquareRoot(number int) int {
	return int(math.Ceil(math.Sqrt(float64(number))))
}

// IsPrime recebe um número inteiro e retorna uma string.
// Irá iterar de 2 até a raiz quadrada do número menos 1.
// Verifica se o número é divisível por i.
// Retorna a string "false" ou "true".

func IsPrime(number int) string {
	for i := 2; i < getSquareRoot(number); i++ {
		if number%i == 0 {
			return "false"
		}
	}
	return "true"
}

var seed = rand.New(rand.NewSource(time.Now().UnixNano()))

// GenerateRandomNumbers recebe um inteiro x e retorna um número aleatório.
// Usará uma semente que gera um número de 0 a 99.
// Soma x + 1 com este número aleatório.
// Retorna um número inteiro aleatório.

func GenerateRandomNumbers(x int) int {
	return x + seed.Intn(100) + 1
}
