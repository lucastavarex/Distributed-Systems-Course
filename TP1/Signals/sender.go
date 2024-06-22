package signals

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"
)

// ProcessExists recebe um PID inteiro e retorna true se esse processo existir.
// Faz uma chamada de sistema Kill para o PID com o sinal 0 (verifica o acesso ao PID).
// Se for nil, isso significa que o processo existe.
// Retorna true se o processo existir e false se não existir.

func processExists(pid int) bool {
	killErr := syscall.Kill(pid, syscall.Signal(0))
	return killErr == nil
}

// InputParsing recebe uma string de entrada e retorna a entrada analisada como uma tupla.
// Divide a entrada contendo a vírgula e atribui a uma variável inputs.
// Remove os espaços em branco de cada elemento de inputs e converte para inteiro.
// Retorna uma tupla contendo o PID e o sinal inseridos.

func inputParsing(input string) (int, int) {
	inputs := strings.Split(input, ",")
	pid, _ := strconv.Atoi(strings.Trim(inputs[0], " "))
	signal, _ := strconv.Atoi(strings.Trim(inputs[1], " "))

	return pid, signal
}

func SignalSender() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("[SENDER] Write the pid of the process and the signal that you want to send. Ex: 3100, 2 \n")
	for scanner.Scan() {
		pid, signal := inputParsing(scanner.Text())
		if processExists(pid) {
			syscall.Kill(pid, syscall.Signal(signal))
			fmt.Println("[SENDER] Signal Sended.")
		} else {
			fmt.Println("[SENDER] Couldn't find the pid", pid)
		}

	}

}
