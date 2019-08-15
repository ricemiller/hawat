package main

import (
	"bytes"
	"fmt"
	"regexp"
	"runtime"

	//"time"
	"bufio"
	"hawat/board"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var ALLOWED_COMMANDS = [...]string{"uci", "debug", "isready", "setoption", "register", "ucinewgame", "position", "go", "stop", "ponderhit", "quit", "print", "perft", "divide", "diff"}

const ENGINE_NAME = "Hawat"
const ENGINE_AUTHOR = "ricemiller"

/*

func background(quit chan struct{}) {
    for {
        select{
        default:
            x = x+1
            time.Sleep(1000 * time.Millisecond)
        case <-quit:
            return
        }
    }
}
*/

func UCIcommand(r *bufio.Reader) []string {
	cmd, _ := r.ReadString('\n')
	cmd = strings.TrimSpace(cmd)
	return strings.Fields(cmd)
}

func IsGoodCMD(cmd string) bool {
	for _, allowedCMD := range ALLOWED_COMMANDS {
		if allowedCMD == cmd {
			return true
		}
	}
	return false
}

func ClearCMDGarbage(cmd []string) []string {
	if len(cmd) == 0 {
		return cmd
	}

	for !IsGoodCMD(cmd[0]) {
		cmd = cmd[1:]
		if len(cmd) == 0 {
			return cmd
		}
	}
	return cmd
}

func main() {
	var b board.Board
	r := bufio.NewReader(os.Stdin)

	//quit := make(chan struct{})

	// go background(quit) //start processing

	// defer close(quit)

	for {
		cmd := UCIcommand(r)
		cmd = ClearCMDGarbage(cmd)
		if len(cmd) > 0 {

			switch cmd[0] {
			case "uci":
				fmt.Printf("id name %s\n", ENGINE_NAME)
				fmt.Printf("id author %s\n", ENGINE_AUTHOR)
				// SEND OPTIONS
				fmt.Printf("uciok\n")

			case "debug":

			case "isready":
				fmt.Printf("readyok\n")

			case "setoption":
				// RECEIVE OPTIONS

			case "register":
				fmt.Printf("This is FOSS, baby\n")

			case "ucinewgame":
				b.Init()

			case "position":
				cmd = cmd[1:]
				if len(cmd) > 0 {
					switch cmd[0] {
					case "fen":
						cmd = cmd[1:]
						if len(cmd) > 5 {
							b.SetFEN(cmd[0:6])
							cmd = cmd[6:]
						}

					case "startpos":
						b.Init()
						cmd = cmd[1:]
					}

					if len(cmd) > 0 {
						if cmd[0] == "moves" {
							cmd = cmd[1:]
							b.Moves()
						}
					}
				}

			case "go":

			case "stop":

			case "ponderhit":

			case "print":
				b.Print()

			case "perft":
				cmd = cmd[1:]
				if len(cmd) > 0 {
					depth, _ := strconv.Atoi(cmd[0])
					fmt.Printf("PERFT TEST\n")
					fmt.Printf("############################################################\n\n")

					nodes := b.Perft(depth)
					fmt.Printf("DEPTH: %d, NODES: %d\n", depth, nodes)
				}

			case "divide":
				cmd = cmd[1:]
				if len(cmd) > 0 {
					depth, _ := strconv.Atoi(cmd[0])
					fmt.Printf("DIVIDE TEST\n")
					fmt.Printf("############################################################\n\n")

					movesMap, totalNodes := b.Divide(depth)
					for move, numNodes := range movesMap {
						fmt.Printf("%s %d\n", move, numNodes)
					}
					fmt.Printf("DEPTH: %d, NODES: %d\n", depth, totalNodes)
				}

			case "diff":
				cmd = cmd[1:]
				if len(cmd) > 0 {
					var stdout, buffer bytes.Buffer
					movesSF := make(map[string]int)
					depth, _ := strconv.Atoi(cmd[0])
					sf := exec.Command("C:\\Users\\JJ\\go\\src\\hawat\\stockfish.exe")

					movesHawat, _ := b.Divide(depth)

					sf.Stdout = &stdout
					sf.Stdin = &buffer
					sf.Stderr = os.Stderr

					buffer.Write([]byte("position fen " + b.FEN + "\n"))
					buffer.Write([]byte("go perft " + cmd[0] + "\n"))
					buffer.Write([]byte("quit\n"))

					var lineSep string
					if runtime.GOOS == "windows" {
						lineSep = "\r\n"
					} else {
						lineSep = "\n"
					}
					err := sf.Run()

					out := strings.Split(string(stdout.Bytes()), lineSep)

					for _, s := range out {
						isAMove, _ := regexp.MatchString("([a-h][1-8]){2}[qnbr]?: [0-9]+", s)
						if isAMove {
							s = strings.Replace(s, ":", "", -1)
							node := strings.Fields(s)
							numMoves, _ := strconv.Atoi(node[1])
							movesSF[node[0]] = numMoves
						}
					}

					mismatch := false
					fmt.Printf("\t\tHawat\t\tSF\n")
					fmt.Printf("\t\t-----\t\t-----\n")
					for node, numNodes := range movesHawat {
						warning := ""
						if movesSF[node] != numNodes {
							mismatch = true
							warning = "[*] Mismatch\n"
						}
						fmt.Printf("%s\t%d\t\t\t%d %s\n", node, numNodes, movesSF[node], warning)
					}

					if !mismatch {
						fmt.Printf("\n[+] No mismatch found!\n")
						if len(movesHawat) != len(movesSF){
							fmt.Printf("[!] Different number of results!")
						}
					}

					if err != nil {
						fmt.Fprintln(os.Stderr, err)
					}

				}

			case "quit":
				os.Exit(0)
			default:
			}
		}
	}
}
