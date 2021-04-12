/*
Le os argumentos da inicialização do programa, criando
um conjunto de chave/valor de acordo com SEPARATOR.

Exemplo:

Ao chamar o programa: go run *.go chave1=valor1 chave2=valor2

Passo 1: No main do programa, chamar ReadArgsConsole()
Passo 2: Quando precisar do valor de chave1, use GetArgsValue("chave1")

*/
package main

import (
	"os"
	"strings"
)

var mapArgs map[string]string
var arrayArgs []string

// O separador para delimitar o conjunto chave/valor na linha de comandos
// O padrão é =
const SEPARATOR string = "="

// Le os argumentos repassados pela linha de comando de execução
// do programa, tentando transformar os argumentos em um conjunto
// de chave/valor que pode ser acessado pelo método GetArgsValue
//
// Deve ser chamado na inicialização do programa
func ReadArgsConsole() {
	mapArgs = make(map[string]string)
	arrayArgs = os.Args
	for _, arg := range arrayArgs {
		splitArg := strings.Split(arg, SEPARATOR)
		if len(splitArg) == 2 {
			mapArgs[splitArg[0]] = splitArg[1]
		}
	}
}

// Retorna o valor especificado pela chave e
//  se a chave existe ou não
func GetArgsValue(key string) (interface{}, bool) {
	value := mapArgs[key]
	if value == "" {
		return "", false
	}
	return value, true
}
