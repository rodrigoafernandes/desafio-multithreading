package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func main() {
	Setup()
	fmt.Print("Informe o CEP a pesquisar(Utilize o formato 00000-000. Exemplo: 03977-250): ")
	var cep string
	fmt.Scanln(&cep)
	cep = strings.TrimSpace(cep)
	if len(strings.TrimSpace(cep)) != 9 || strings.Compare("00000-000", cep) == 0 {
		panic(fmt.Sprintf("O cep informado é inválido: %s", cep))
	}
	apiCepUrl := fmt.Sprintf("%s/%s.json", AppCFG.ApiCepUrl, cep)
	viaCepUrl := fmt.Sprintf("%s/%s/json", AppCFG.ViaCepUrl, cep)
	apiCepRequest, err := http.NewRequest("GET", apiCepUrl, nil)
	if err != nil {
		panic(err)
	}
	viaCepRequest, err := http.NewRequest("GET", viaCepUrl, nil)
	if err != nil {
		panic(err)
	}

	apiCepChannel := make(chan *http.Response)
	viaCepChannel := make(chan *http.Response)

	go buscaCep(apiCepRequest, apiCepChannel)
	go buscaCep(viaCepRequest, viaCepChannel)

	select {
	case apiCepChannelResponse := <-apiCepChannel:
		defer apiCepChannelResponse.Body.Close()
		type apiCepResponse struct {
			Code       string `json:"code"`
			State      string `json:"state"`
			City       string `json:"city"`
			District   string `json:"district"`
			Address    string `json:"address"`
			Status     int    `json:"status"`
			Ok         bool   `json:"ok"`
			StatusText string `json:"statusText"`
		}
		var apiCepRes apiCepResponse
		if err = json.NewDecoder(apiCepChannelResponse.Body).Decode(&apiCepRes); err != nil {
			panic(err)
		}
		if !apiCepRes.Ok {
			panic(fmt.Sprintf("erro ao consultar o cep %s", cep))
		}
		b, err := json.Marshal(apiCepRes)
		if err != nil {
			panic(err)
		}
		message := fmt.Sprintf("resultado: %s\nApi com a resposta mais rápida: %s\n", string(b), apiCepChannelResponse.Request.Host)
		fmt.Println(message)

	case viaCepChannelResponse := <-viaCepChannel:
		defer viaCepChannelResponse.Body.Close()
		type viaCepResponse struct {
			Cep         string `json:"cep"`
			Logradouro  string `json:"logradouro"`
			Complemento string `json:"complemento"`
			Bairro      string `json:"bairro"`
			Localidade  string `json:"localidade"`
			Uf          string `json:"uf"`
			Ibge        string `json:"ibge"`
			Gia         string `json:"gia"`
			Ddd         string `json:"ddd"`
			Siafi       string `json:"siafi"`
			Erro        bool   `json:"erro"`
		}
		var viaCepRes viaCepResponse
		if err = json.NewDecoder(viaCepChannelResponse.Body).Decode(&viaCepRes); err != nil {
			panic(err)
		}
		if viaCepRes.Erro {
			panic(fmt.Sprintf("erro ao consultar o cep %s", cep))
		}
		b, err := json.Marshal(viaCepRes)
		if err != nil {
			panic(err)
		}
		message := fmt.Sprintf("resultado: %s\nApi com a resposta mais rápida: %s\n", string(b), viaCepChannelResponse.Request.Host)
		fmt.Println(message)

	case <-time.After(time.Duration(AppCFG.ApiTimeoutMS) * time.Second):
		fmt.Println("timeout ao consultar as apis")
	}
}

func buscaCep(request *http.Request, canalResponse chan *http.Response) {
	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	if response.StatusCode < 199 || response.StatusCode > 299 {
		panic(errors.New(fmt.Sprintf("erro ao comunicar com a api. http status %d", response.StatusCode)))
	}
	canalResponse <- response
}
