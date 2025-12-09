# การรัน API (Gin_framework)

ก่อนอื่น ให้เข้าไปที่โฟลเดอร์ api และติดตั้ง dependencies:

```bash
cd Gin_framework/api
go mod tidy
```

#### a. รัน RESTful API
```bash
go run RESTful-API/main.go
```
- **ทดสอบ:** เปิด Browser หรือ Postman ไปที่ `http://localhost:8080/albums`

#### b. รัน GraphQL API
```bash
go run GraphQL/main.go
```
- **ทดสอบ:** เปิด Browser ไปที่ `http://localhost:8081/graphql` เพื่อใช้งาน GraphiQL interface

#### c. รัน Microservices (gRPC)
ต้องเปิด Terminal 2 หน้าต่าง:

**Terminal 1 (รัน Server):**
```bash
go run Microservices-gRPC/server/main.go
```

**Terminal 2 (รัน Client):**
```bash
go run Microservices-gRPC/client/main.go
```