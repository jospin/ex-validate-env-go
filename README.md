# LIB para validar env de projetos que estão no argocd

## Como usar
Para poder validar as variáveis de ambiente (ignorando as secrets) ele deverá ser adicionado no build-input do projeto, e configurar o bitbucket-pipeline para chamara a bin passando o nome do projeto e o ambiente de validação 
ex:

`./build-input/ex-validate-env-go/validate ms-platform-go homologation`

Perceba que no ms-platform-go ele foi criado dentro do build input em uma pasta para ele, passa como parâmetro o projeto e o ambiente