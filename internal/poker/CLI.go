package poker

import (
	"bufio"
	"io"
	"strings"
	"time"
)

type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

type CLI struct {
	playerStore PlayerStore
	scanner     *bufio.Scanner
	alerter     BlindAlerter
}

func NewCLI(store PlayerStore, in io.Reader, alerter BlindAlerter) *CLI {
	return &CLI{playerStore: store, scanner: bufio.NewScanner(in), alerter: alerter}
}

func (cli *CLI) PlayPoker() {
	cli.alerter.ScheduleAlertAt(5*time.Second, 100)
	cli.playerStore.RecordWin(extractWinner(cli.readLine()))
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (cli *CLI) readLine() string {
	cli.scanner.Scan()
	return cli.scanner.Text()
}
