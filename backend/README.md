# Backend

Backend desenvolvido em GO para disponibilizar uma API controle de pressão arterial e glicemia.

## Como rodar
- Execute: `go run main.go`

## Rotas
- Glicemia
    - GET: `/blood-glucose` - Retorna todos os registros de glicemia
    - GET: `/blood-glucose/{id}` - Retorna um registro de glicemia
    - POST: `/blood-glucose` - Cria um registro de glicemia
    - DELETE: `/blood-glucose/{id}` - Deleta um registro de glicemia
---
- Pressão arterial
    - GET: `/blood-pressure` - Retorna todos os registros de pressão arterial
    - GET: `/blood-pressure/{id}` - Retorna um registro de pressão arterial
    - POST: `/blood-pressure` - Cria um registro de pressão arterial
    - DELETE: `/blood-pressure/{id}` - Deleta um registro de pressão arterial

## Estrutura de dados
- Glicemia
    ```json
    {
        "id": 1,
        "blood_glucose": 100,
        "created_at": "2021-01-01T00:00:00Z"
    }
    ```

- Pressão arterial
    ```json
    {
        "id": 1,
        "systolic": 100,
        "diastolic": 100,
        "pulse": 100,
        "created_at": "2021-01-01T00:00:00Z"
    }
    ```