package pipe

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"syscall"

	utils "github.com/lucastavarex/Distributed-Systems-Course/Utils"
)

// O consumidor recebe um io.Reader e não tem retorno.
// Ele instanciará um novo scanner com o reader.
// Espera por cada elemento escaneado no buffer.
// Converte o conteúdo do scanner para inteiro.
// Verifica se o conteúdo recebido é 0, se sim, sai da função.
// Verifica se o conteúdo é um número primo.
// Imprime a mensagem e o valor.
// Não tem retorno.

func consumer(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())

		fmt.Printf("[CONSUMER] Message Received: %v\n", n)

		if n == 0 {
			fmt.Println("[CONSUMER] Process finished.")
			return
		}

		message := utils.IsPrime(n)

		fmt.Printf("[CONSUMER] Is the value %v prime? %s \n", n, message)
	}
}

// O produtor recebe um io.WriteCloser e não tem retorno.
// Ele instanciará um novo scanner com o os.Stdin.
// Espera pela entrada do usuário.
// Lê e converte a entrada do usuário para inteiro.
// Itera de 0 até o valor da entrada do usuário menos 1.
// Gera um número aleatório a partir do valor x anterior.
// Escreve no escritor w.
// Após o fim do loop, escreve 0 no escritor w e fecha o escritor.
// Não tem retorno.

func producer(w io.WriteCloser) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("\n\n\n[PRODUCER] Write the amount of numbers to be generated \n")

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	x := 1
	for i := 0; i < n; i++ {
		x = utils.GenerateRandomNumbers(x)
		fmt.Fprintf(w, "%v\n", x)
	}

	fmt.Fprint(w, "0\n")
	w.Close()
}

func Pipe() {
	pipe := make([]int, 2)
	syscall.Pipe(pipe)

	r := os.NewFile(uintptr(pipe[0]), "consumer")
	w := os.NewFile(uintptr(pipe[1]), "producer")

	id, _, _ := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)

	if id == 0 {
		producer(w)
	} else if id > 0 {
		consumer(r)
	} else {
		fmt.Println("[ERROR] Forked Failed.")
		return
	}
}
