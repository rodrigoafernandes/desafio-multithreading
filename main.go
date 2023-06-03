package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func main() {
	Setup()
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(AppCFG.ApiTimeoutMS)*time.Second)
	defer cancel()
	cep := "09220-110"
	apiCepUrl := fmt.Sprintf("%s/%s.json", AppCFG.ApiCepUrl, cep)
	viaCepUrl := fmt.Sprintf("%s/%s/json", AppCFG.ViaCepUrl, cep)
	apiCepRequest, err := http.NewRequestWithContext(ctx, "GET", apiCepUrl, nil)
	if err != nil {
		panic(err)
	}
	viaCepRequest, err := http.NewRequestWithContext(ctx, "GET", viaCepUrl, nil)
	if err != nil {
		panic(err)
	}
	client := http.DefaultClient
	responseApiCep, err := client.Do(apiCepRequest)
	if err != nil {
		panic(err)
	}
	defer responseApiCep.Body.Close()
	if responseApiCep.StatusCode < 199 || responseApiCep.StatusCode > 299 {
		panic(errors.New(fmt.Sprintf("erro ao comunicar com a apicep. http status %d", responseApiCep.StatusCode)))
	}
	fmt.Println(responseApiCep.Body)
	responseViaCep, err := client.Do(viaCepRequest)
	if err != nil {
		panic(err)
	}
	if responseViaCep.StatusCode < 199 || responseViaCep.StatusCode > 299 {
		panic(errors.New(fmt.Sprintf("erro ao comunicar com a viacep. http status %d", responseApiCep.StatusCode)))
	}
}
