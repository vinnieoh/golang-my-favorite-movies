# Golang My Favorite Movies - Projeto de Filmes e Séries Favoritos

## Descrição do Projeto

Este projeto é um sistema de gerenciamento de filmes e séries favoritos, onde os usuários podem adicionar seus filmes e séries preferidos, bem # Golang My Favorite Movies - Projeto de Filmes e Séries Favoritos

## Descrição do Projeto

Este projeto é um sistema de gerenciamento de filmes e séries favoritos, onde os usuários podem adicionar seus filmes e séries preferidos, bem como comentar sobre eles. O sistema consome a API do TMDB para buscar e salvar dados na plataforma. Ele é desenvolvido com Golang para o backend, utilizando PostgreSQL e Redis como bancos de dados.

## Funcionalidades

- Adicionar filmes e séries aos favoritos.
- Adicionar comentários aos filmes e séries.
- Buscar detalhes de filmes e séries na API do TMDB.
- Armazenamento de dados em PostgreSQL.
- Cache utilizando Redis para melhorar o desempenho.

## Tecnologias Utilizadas

- **Backend**: Golang, Gin (framework web)
- **Banco de Dados**: PostgreSQL, Redis
- **Autenticação**: JWT (JSON Web Tokens)
- **ORM**: GORM
- **API Externa**: The Movie Database (TMDB) API
- **Migrações de Banco de Dados**: Gormigrate

## Pré-requisitos

- Docker e Docker Compose instalados.
- Criar uma conta na [TMDB](https://www.themoviedb.org/) para obter a API Key.

## Configuração do Ambiente

1. Clone o repositório:
   ```sh
   git clone https://github.com/vinnieoh/golang-my-favorite-movies.git
   cd golang-my-favorite-movies
2. Crie o arquivo .env para o backend a partir do exemplo fornecido:
     ```sh
     cp ./api/dotenv_files/.env_exemplo_backend ./api/dotenv_files/.env
3. Edite o arquivo .env e adicione a sua TMDB API Key:
     ```sh
     API_MOVIE=<sua_api_key_da_tmdb>
4. Certifique-se de que o Docker e Docker Compose estão instalados em sua máquina.

5. Para iniciar o projeto, utilize o comando:
    ```sh
    docker-compose up --build
    ```
    Este comando irá:
    - Construir e iniciar o contêiner do PostgreSQL.
    - Construir e iniciar o contêiner do Redis.
    - Construir e iniciar o contêiner do backend Golang.

6. Certifique-se de que todos os serviços estão funcionando corretamente verificando os logs dos contêineres.

## Rotas da API

### Usuários
- GET /v1/users: Lista todos os usuários.
- GET /v1/users/:id: Obtém um usuário específico pelo ID.
- POST /v1/users: Cria um novo usuário.
- PUT /v1/users/:id: Atualiza um usuário específico pelo ID.
- DELETE /v1/users/:id: Deleta um usuário específico pelo ID.
- POST /v1/login: Realiza o login de um usuário e retorna um token JWT.

### The Movie Database (TMDB) API
- GET /v1/trending-all-week-br: Obtém os conteúdos em alta da semana em português brasileiro.
- GET /v1/trending-all-day-br: Obtém os conteúdos em alta do dia em português brasileiro.
- GET /v1/search-content/:content: Busca conteúdos pelo nome em português brasileiro.
- GET /v1/movie-id/:id: Obtém detalhes de um filme específico pelo ID da API do TMDB.
- GET /v1/tv-show-id/:id: Obtém detalhes de uma série específica pelo ID da API do TMDB.

### Filmes
- GET /v1/movies: Lista todos os filmes.
- GET /v1/movies/:id: Obtém um filme específico pelo ID.
- POST /v1/movies: Adiciona um novo filme aos favoritos.
- PUT /v1/movies/:id: Atualiza um filme específico pelo ID.
- DELETE /v1/movies/:id: Deleta um filme específico pelo ID.

### Séries
- GET /v1/tvshows: Lista todas as séries.
- GET /v1/tvshows/:id: Obtém uma série específica pelo ID.
- POST /v1/tvshows: Adiciona uma nova série aos favoritos.
- PUT /v1/tvshows/:id: Atualiza uma série específica pelo ID.
- DELETE /v1/tvshows/:id: Deleta uma série específica pelo ID.

## Observações
- Certifique-se de que você possui uma API Key válida da TMDB. Caso não tenha, registre-se no site da TMDB e crie sua key.

- O projeto é configurado para ser executado em contêineres Docker, facilitando a configuração do ambiente e a implantação.

- Para personalizar ou modificar o projeto, edite os arquivos de configuração e o código-fonte conforme necessário.

## Licença

Este projeto está licenciado sob a Licença MIT. Consulte o arquivo LICENSE para obter mais informações.