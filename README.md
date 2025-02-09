# API de Diaristas

## Introdução
Esta API permite a gestão de um marketplace de diaristas, onde clientes podem contratar serviços de limpeza e diaristas podem se cadastrar para oferecer seus serviços.

## Tecnologias Utilizadas
- **Linguagem:** Go (Golang)
- **Framework:** Fiber
- **ORM:** GORM
- **Banco de Dados:** PostgreSQL
- **Autenticação:** JWT (em futuras implementações)

---

# Endpoints da API

## Usuários (`/api/users`)

### Criar um usuário
**POST** `/api/users`
```json
{
  "name": "João Silva",
  "email": "joao@email.com",
  "phone": "11999999999",
  "password": "senha123",
  "role": "cliente"
}
```
**Resposta:**
```json
{
  "id": 1,
  "name": "João Silva",
  "email": "joao@email.com",
  "phone": "11999999999",
  "role": "cliente",
  "created_at": "2025-02-08T12:00:00Z"
}
```

### Buscar todos os usuários
**GET** `/api/users`

### Buscar um usuário por ID
**GET** `/api/users/:id`

### Atualizar um usuário
**PUT** `/api/users/:id`

### Deletar um usuário
**DELETE** `/api/users/:id`

---

## Endereços (`/api/addresses`)

### Criar um endereço
**POST** `/api/addresses`
```json
{
  "user_id": 1,
  "street": "Rua Exemplo",
  "number": "100",
  "city": "São Paulo",
  "state": "SP",
  "zipcode": "01000-000",
  "latitude": -23.55052,
  "longitude": -46.63331
}
```

### Buscar todos os endereços
**GET** `/api/addresses`

### Buscar um endereço por ID
**GET** `/api/addresses/:id`

### Atualizar um endereço
**PUT** `/api/addresses/:id`

### Deletar um endereço
**DELETE** `/api/addresses/:id`

---

## Diaristas (`/api/diarists`)

### Criar um diarista
**POST** `/api/diarists`
```json
{
  "user_id": 2,
  "bio": "Tenho 5 anos de experiência em limpeza residencial.",
  "experience_years": 5,
  "price_per_hour": 50.00,
  "available": true
}
```

### Buscar todos os diaristas
**GET** `/api/diarists`

### Buscar um diarista por ID
**GET** `/api/diarists/:id`

### Atualizar um diarista
**PUT** `/api/diarists/:id`

### Deletar um diarista
**DELETE** `/api/diarists/:id`

---

## Serviços (`/api/services`)

### Criar um serviço
**POST** `/api/services`
```json
{
  "client_id": 1,
  "diarist_id": 2,
  "address_id": 3,
  "status": "pendente",
  "total_price": 150.00,
  "duration_hours": 3.00,
  "scheduled_at": "2025-02-10T10:00:00Z"
}
```

### Buscar todos os serviços
**GET** `/api/services`

### Buscar um serviço por ID
**GET** `/api/services/:id`

### Atualizar um serviço
**PUT** `/api/services/:id`

### Deletar um serviço
**DELETE** `/api/services/:id`

---

## Pagamentos (`/api/payments`)

### Criar um pagamento
**POST** `/api/payments`
```json
{
  "service_id": 1,
  "client_id": 1,
  "diarist_id": 2,
  "amount": 150.00,
  "status": "pendente",
  "method": "pix"
}
```

### Buscar todos os pagamentos
**GET** `/api/payments`

### Buscar um pagamento por ID
**GET** `/api/payments/:id`

### Atualizar um pagamento
**PUT** `/api/payments/:id`

### Deletar um pagamento
**DELETE** `/api/payments/:id`

---

## Avaliações (`/api/reviews`)

### Criar uma avaliação
**POST** `/api/reviews`
```json
{
  "service_id": 1,
  "reviewer_id": 1,
  "reviewed_id": 2,
  "rating": 5,
  "comment": "Serviço excelente!"
}
```

### Buscar todas as avaliações
**GET** `/api/reviews`

### Buscar uma avaliação por ID
**GET** `/api/reviews/:id`

### Atualizar uma avaliação
**PUT** `/api/reviews/:id`

### Deletar uma avaliação
**DELETE** `/api/reviews/:id`

---

# Autoria e Contato
- **Desenvolvedor:** Daniel Rocha Tavares da Silva
- **GitHub:** [github.com/de.maricaense](https://github.com/danielrocha1)
- **Email:** daniel.rochats@gmail.com
- **Instagram:** [@de.maricaense](https://instagram.com/dev.maricaense)

