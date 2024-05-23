package consoleApp

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"wor/bst"
)

//------------------------------------------------------------------------------

var (
	red                string = "\033[31m"
	reset              string = "\033[0m"
	ErrConvertStrToInt error  = errors.New("bst: String to int conversion errord")
)

//------------------------------------------------------------------------------

func convertStringToInt(str string) (int, error) {
	data, err := strconv.Atoi(str)
	if err != nil {
		return 0, ErrConvertStrToInt
	}
	return data, nil
}

//------------------------------------------------------------------------------

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
	PrintCommands()
}

//------------------------------------------------------------------------------

func PrintCommands() {
	fmt.Println("Available commands:")
	fmt.Println("  insert <number1> <number2> ... <numberN> - Inserts numbers into the tree")
	fmt.Println("  remove <number1> <number2> ... <numberN> - Removes numbers from the tree")
	fmt.Println("  replace <old number> to <new number> - Replaces the old number to the new one in the tree")
	fmt.Println("  clear - Clears the tree")
	fmt.Println("  exit - Exits the program")
	fmt.Println()
}

//------------------------------------------------------------------------------

func Display(tree bst.BST) {
	fmt.Println("size", tree.Size())
	fmt.Println("preorder: ", tree.Preorder())
	fmt.Println("inorder: ", tree.Inorder())
	fmt.Println("postorder: ", tree.Postorder())
	fmt.Println(tree.TreeString())
}

//------------------------------------------------------------------------------

func WaitingInput(tree bst.BST) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter the command: ")
		text, _ := reader.ReadString('\n')
		textArr := strings.Split(strings.TrimSpace(text), " ")
		command := strings.ToLower(textArr[0])

		if command == "exit" {
			break
		}

		if command == "clear" {
			tree.Root(nil)
			tree.Size(0)
			clearScreen()
			continue
		}

		if command == "insert" || command == "remove" {
			numbers := make([]int, 0, len(textArr)-1)
			for _, dataStr := range textArr[1:] {
				data, err := convertStringToInt(dataStr)
				if err != nil {
					fmt.Printf("%s\rError: %s {%s}%s\n", red, err, dataStr, reset)
					continue
				}
				numbers = append(numbers, data)
			}

			fmt.Print(numbers)

			count := 0
			for _, data := range numbers {
				var err error
				if command == "insert" {
					err = tree.Insert(data)
				} else {
					err = tree.Remove(data)
				}

				if err != nil {
					fmt.Printf("%s\rError: %s {%d}%s\n", red, err, data, reset)
					continue
				}
				count++
			}
			fmt.Printf("Number of successful %s nodes: %d out of %d\n", command, count, len(textArr)-1)
			fmt.Println()
			Display(tree)
			fmt.Println("────────────────────────────────────────────────────────────────────────")
		}

		if command == "replace" {
			dataStrOld := textArr[1]
			dataStrNew := textArr[3]
			dataOld, errOld := convertStringToInt(dataStrOld)
			dataNew, errNew := convertStringToInt(dataStrNew)
			if errOld != nil || errNew != nil {
				if errOld != nil {
					fmt.Printf("%s\rError: %s {%s}%s\n", red, errOld, dataStrOld, reset)
				}

				if errNew != nil {
					fmt.Printf("%s\rError: %s {%s}%s\n", red, errNew, dataStrNew, reset)
				}

				continue
			}

			err := tree.Replace(dataOld, dataNew)
			if err != nil {
				fmt.Printf("%s\rError: %s %s\n", red, err, reset)
				continue
			}
			fmt.Println()
			Display(tree)
			fmt.Println("────────────────────────────────────────────────────────────────────────")
		}
	}
}
