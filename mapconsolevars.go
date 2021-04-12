package mapconsolevars

import (
	"os"
	"strings"
)

var mapArgs map[string]string
var arrayArgs []string
var initializated bool = false

// O separador para delimitar o conjunto chave/valor na linha de comandos
// O padrão é =
const SEPARATOR string = "="

// Le os argumentos repassados pela linha de comando de execução
// do programa, tentando transformar os argumentos em um conjunto
// de chave/valor que pode ser acessado pelo método GetArgsValue
func readArgsConsole() {
	mapArgs = make(map[string]string)
	arrayArgs = os.Args
	for _, arg := range arrayArgs {
		splitArg := strings.Split(arg, SEPARATOR)
		if len(splitArg) == 2 {
			mapArgs[splitArg[0]] = splitArg[1]
		}
	}
}

// Retorna o valor especificado pela chave e se a chave existe ou não
//
// Por exemplo:
//
// 	valor1, found := mapconsolevars.GetArgsValue("chave1")
//	if found {
//		fmt.Println("O valor de chave1 é", valor1)
//	}
func GetArgsValue(key string) (interface{}, bool) {
	if !initializated {
		readArgsConsole()
		initializated = true
	}
	value := mapArgs[key]
	if value == "" {
		return "", false
	}
	return value, true
}
