# Cara ou Coroa?

## Contexto
Para dar continuidade ao nosso processo, temos um desafio! Gostamos muito de jogar e queremos que você crie um joguinho simples de [Cara ou Coroa](https://en.wikipedia.org/wiki/Coin_flipping).

## Requisitos

O programa deve ser escrito em **Golang** e deve ser possível executar sem erros.

### Regras de negócio
Através do programa, o comportamento deve:
- Ser capaz de perguntar ao usuário (você quer jogar?) 
- Ser capaz de perguntar ao usuário (cara ou coroa?)
- Aleatoriamente retornar um valor de cara ou coroa
- Identificar o resultado e informar se o jogador acertou ou não
- Ser capaz de perguntar ao usuário (quer jogar novamente?)
- Poder armazenar quantas vezes o jogador ganhou e perdeu ao longo das rodadas
- Mostrar o resultado de rodadas ganhas e perdidas quando o jogador desistir de jogar

### Específicos
- Deve seguir a arquitetura [REST](https://restfulapi.net/)
- Deve seguir os principios do [12 factor app](https://12factor.net/pt_br/)
- As respostas do jogador devem ser enviados via console (prompt) ou em qualquer método melhor
- O sorteio da moeda deve ser aleatório, com 50% de chance para cada lado

## Avaliação
A ideia aqui é entender como você toma suas decisões e como você desenvolve através de multiplas funcionalidades.

Pontos que vamos avaliar:
- Commits
    - como você evoluiu seu pensamento durante o projeto, pontualidade e clareza.
- Testes
    - É ideal a criação de um arquivo de teste. É nesse arquivo que será verificado se o que você escreveu faz sentido. Vide https://code.tutsplus.com/pt/tutorials/lets-go-testing-golang-programs--cms-26499 .
- Complexidade
    - Código bom é código legivel e simples (https://medium.com/trainingcenter/golang-d94e16d4b383).
- Dependências
    - O ecosistema (https://github.com/avelino/awesome-go) da linguagem possui diversas ferramentas para serem usadas, use-as bem!
- Documentação
    - Qual versão de Go você usou?
    - Quais bibliotecas e ferramentas usou?
    - Como se utiliza a sua aplicação?
    - Como executamos os testes e a aplicação?
- Considerações
    - Golang é uma linguagem principalmente funcional. Seguro não criar toda a aplicação dentro de uma única função. Separe bem o que cada coisa faz, isso deixa a leitura de código mais clara e facilita o desenvolvimento e execução dos testes.
