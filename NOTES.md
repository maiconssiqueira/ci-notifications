--------
- [x] Gostei do README.md, ficou bem fácil de entender o que o projeto faz
https://github.com/maiconssiqueira/notifications-cli/blob/main/cmd/comments.go#L13 usar o RunE, para poder tratar melhor os erros da aplicação. Aqui tem um bom exemplo disso: https://github.com/PicPay/build-checker/blob/master/cmd/post_build.go#L26 e aqui e eu falei um pouco sobre esse assunto: https://eltonminetto.dev/post/2022-07-06-error-handling-cli-applications-golang 
<br>

- [x] Não vi problemas na criação do arquivo structs.go. Na comunidade Go não é uma prática muito comum nomear arquivos com o seu conteúdo (structs, interfaces, controllers, models, etc), mas não chega a ser um pecado mortal hehe. Uma sugestão seria mudar o nome do arquivo de structs.go para github.go.
<br>

- [ ] Você vai ter dificuldade em testar essa parte: https://github.com/maiconssiqueira/notifications-cli/blob/main/github/statuses.go#L54 Dá uma olhada no conceito de inversão de dependências no curso do Branas. Vou deixar vc chegar nessa parte dos testes para podermos discutir juntos isso.
<br>

- [ ] Se você der uma olhada neste post: https://medium.com/inside-picpay/organizando-um-projeto-e-convencionando-nomes-em-go-c18b3fa88ba0 vai ver um conceito legal que é o de pacotes de domínio. Aplicando isso no seu repo, o diretório github está no lugar certo, pq ele é um pacote do domínio da sua aplicação. Já o pacote http me parece conter apenas funcionalidades de apoio ao seu domínio. Minha sugestão é você criar um diretório chamado internal e colocar o pacote http dentro. Deixa mais organizado.
<br>

- [ ] Não é uma prática recomendada criar pacotes com nomes genéricos como utils. Sugiro você mudar o config para a raiz do projeto (transformando ele em um pacote de domínio) e remover o pacote utils
Quanto a dúvida da validação dos inputs, eu acho válido fazer isso na camada da CLI mesmo, e não deixar isso para o pacote github (camada de domínio).

####Resumindo
O código está ficando bem maneiro cara, parabéns. Fiz apenas algumas sugestões baseado no código atual. Tem alguns problemas de design mas você vai perceber eles quando começar a escrever os testes :sorriso_pequeno: Não vou apontar eles agora porque você perceber isso é um aprendizado muito bacana.
<br>

- [ ] Se você achar válido podemos fazer um pair quando você chegar a estas dificuldades na hora de escrever os testes e debatemos as soluções juntos. Acho que esse post pode te ajudar bastante quanto aos testes: https://medium.com/inside-picpay/testes-automatizados-em-go-aa5cf9ed672e

