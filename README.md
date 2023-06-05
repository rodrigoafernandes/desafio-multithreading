# desafio-multithreading

Este desafio consiste em buscar o resultado mais rápido entre as Apis: [ApiCep](https://cdn.apicep.com/file/apicep) e [ViaCep](http://viacep.com.br/ws).

## Requisitos

Os requisitos para este desafio são:

- Acatar a API que entregar a resposta mais rápida e descartar a resposta mais lenta.

- O resultado da request deverá ser exibido no command line, bem como qual API a enviou.

- Limitar o tempo de resposta em 1 segundo. Caso contrário, o erro de timeout deve ser exibido.

## Usage

```bash
go run main.go
```
Será solicitado o CEP a ser consultado via commandline<br/>
A aplicação validará se o CEP não foi informado e se o cep é igual a 00000-000, caso alguma dessas validações seja verdadeira, irá retornar um erro.<br/>
Caso o CEP informado não seja localizado nas Apis, um erro será retornado.<br/>