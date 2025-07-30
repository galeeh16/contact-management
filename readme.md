Ini adalah program golang untuk crud contact sederhana membuat contact.

### library:
- [fiber](https://gofiber.io/)
- [gorm](https://gorm.io/)
- [godotenv](https://github.com/joho/godotenv)
- [golang-jwt](https://github.com/golang-jwt/jwt)

### Untuk menjalankan aplikasi
- Copy file ```.env.example``` ke file ```.env``` (sesuaikan isinya)
- Buat databasenya, yang sekarang pake postgresq, lalu jalankan file ```script.sql```
- Jalankan perintah ```go mod tidy``` atau ```make tidy```
- Untuk menjalankan aplikasi, ketik ```go run main.go``` atau ```make run```
- Aplikasi running pada ```http://localhost:8000/```

### Endpoint Auth dan Contact
| Keterangan                     | URL                                          |
|--------------------------------|----------------------------------------------|
| Register                       | [POST] localhost:8000/api/v1/auth/register    |
| Login                          | [POST] localhost:8000/api/v1/auth/login       |
| Me/Profile                     | [POST] localhost:8000/api/v1/auth/me         |
| Get All Contact (belum diimplementasi)| [GET] localhost:8000/api/v1/contacts        |
| Find Contact by Phone          | [GET] localhost:8000/api/v1/contacts/:phone  |
| Create Contact                 | [POST] localhost:8000/api/v1/contacts        |
| Update Contact by ID (belum diimplementasi) | [PUT] localhost:8000/api/v1/contacts/:id     |
| Delete Contact by ID  (belum diimplementasi)| [DELETE] localhost:8000/api/v1/contacts/:id  |