# #VemPraGenesis! - Backend Challenge

Desafio para os futuros programadores back-end da Genesis.

## Introdução

Desenvolva uma **REST API** que faça conversão de moedas.

**Especifícações**:

* A URL da requisição deve seguir o seguinte formato:
    * http://localhost:8000/exchange/{amount}/{from}/{to}/{rate}
    * http://localhost:8000/exchange/10/BRL/USD/4.50
* A resposta deve seguir o seguinte formato:
   ```json
   {
     "valorConvertido": 45,
     "simboloMoeda": "$"
   }
   
   * Conversões:
    * De Real para Dólar;
    * De Dólar para Real;
    * De Real para Euro;
    * De Euro para Real;
    * De BTC para Dolar;
    * De BTC para Real;
   
   * Salvar os dados no banco de dados:
    * criar uma rotina para salvar o dados para consultas futuras

   ```

    

## Instruções

1. Efetue o fork deste repositório e crie um branch com o seu nome e sobrenome. (exemplo: fulano-dasilva)
2. Após finalizar o desafio, crie um Pull Request.
3. Aguarde algum contribuidor realizar o code review.
4. pode usar framework de sua preferência

## Pré-requisitos

* Linguagem Go
* Mysql

## Plus

* Testes unitários
* Clean code
* criar ambiente dentro do docker
* criar docker-compose.yml

## Dúvidas

Em caso de dúvidas, crie uma issue.
