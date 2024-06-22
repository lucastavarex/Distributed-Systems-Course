package signals

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

// BusyWait recebe o canal de sinais e o canal de saída e imprime mensagens quando recebe alguns sinais.
// Ele usa uma goroutine, uma forma de usar programação concorrente em Go.
// Para sincronizar a concorrência, é necessário usar um canal (neste caso, um Canal de Sinais).
// O loop while true possui um select-case que verifica se o sinal foi recebido e espera nessa condição.
// No entanto, por ter uma sentença default, ele entrará nessa seção toda vez que o loop for executado.
// Os sinais tratados são SIGHUP (1), SIGTERM (15), SIGQUIT (3).
// O exitChan recebe uma mensagem 0, encerrando a goroutine.
// Não tem retorno.

func busyWait(signalChan chan os.Signal, exitChan chan int) {
	fmt.Println("[RECEIVER] Running on Busy Wait mode. Please send a signal.")
	go func() {
		for {
			select {
			case signal := <-signalChan:

				if signal == syscall.SIGHUP {
					fmt.Println("[RECEIVER] SIGHUP received.")
				}
				if signal == syscall.SIGTERM {
					fmt.Println("[RECEIVER] SIGTERM received.")
				}
				if signal == syscall.SIGQUIT {
					fmt.Println("[RECEIVER] SIGQUIT received.")
					exitChan <- 0
				}
			default:
			}
		}
	}()
}

// BlockingWait recebe o canal de sinais e o canal de saída e imprime mensagens quando recebe alguns sinais.
// Ele usa uma goroutine, uma forma de usar programação concorrente em Go.
// Para sincronizar a concorrência, é necessário usar um canal (neste caso, um Canal de Sinais).
// A variável signal bloqueia o processo até que a variável signalChan receba um sinal.
// O loop while true possui um switch-case que verifica cada sinal recebido.
// Os sinais tratados são SIGHUP (1), SIGTERM (15), SIGQUIT (3).
// O exitChan recebe uma mensagem 0, encerrando a goroutine.
// Não tem retorno.

func blockingWait(signalChan chan os.Signal, exitChan chan int) {
	fmt.Println("[RECEIVER] Running on Blocking Wait mode. Please send a signal.")
	go func() {
		for {
			signal := <-signalChan
			switch signal {
			case syscall.SIGHUP:
				fmt.Println("[RECEIVER] SIGHUP received.")

			case syscall.SIGTERM:
				fmt.Println("[RECEIVER] SIGTERM received.")

			case syscall.SIGQUIT:
				fmt.Println("[RECEIVER] SIGQUIT received.")
				exitChan <- 0
			}
		}
	}()
}

// InstantiateChannels não recebe parâmetros e retorna uma tupla contendo 2 canais.
// Ele instancia o signalChannel, que é um canal que recebe valores de os.Signal.
// Ele instancia o exitChannel, que é um canal que recebe valores inteiros.
// Ele recebe a notificação de cada sinal enviado, evitando qualquer utilidade de sinal a menos que definido pelas funções anteriores.
// Não tem retorno.

func instantiateChannels() (chan os.Signal, chan int) {
	signalChannel := make(chan os.Signal, 1)

	signal.Notify(signalChannel)

	exitChannel := make(chan int)

	return signalChannel, exitChannel
}

func SignalReceiver() {
	fmt.Printf("\n\n\n[RECEIVER] Process pid : %d\n", os.Getpid())

	signalChan, exitChan := instantiateChannels()

	fmt.Printf("[RECEIVER] Write the mode you want to run the signal receiver. 0 to blocking wait or 1 to busy wait. \n")

	//Collect the input passed by the user
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	mode, _ := strconv.Atoi(scanner.Text())
	switch mode {
	case 0:
		blockingWait(signalChan, exitChan)
	case 1:
		busyWait(signalChan, exitChan)
	default:
		fmt.Println("[RECEIVER] Unknown parameter. Please, choose 0 for blocking wait or 1 for busy wait.")
		return
	}

	code := <-exitChan

	fmt.Println("[RECEIVER] Process finished.")

	os.Exit(code)
}
