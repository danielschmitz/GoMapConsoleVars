/*
Lê os argumentos da inicialização do programa (console),
criando um map chave/valor

Exemplo:

Ao chamar o programa: `go run *.go chave1=valor1 chave2=valor2`

Quando precisar do valor de chave1, use `GetArgsValue("chave1")`

	valor1, found := mapconsolevars.GetArgsValue("chave1")
	if found {
		fmt.Println("O valor de chave1 é", valor1)
	}

*/
package mapconsolevars
