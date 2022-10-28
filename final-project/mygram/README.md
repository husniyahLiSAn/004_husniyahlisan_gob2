## Biodata
- Nama Lengkap : Husniyah Lisan
- Kode Peserta : 149368582100-4
---

# PANDUAN

1. Buat database terlebih dahulu pada SQL Workbrench
2. Buat file **.env** untuk setting konfigurasi database. Formatnya dapat dilihat pada file **.env.example**
3. Buat *struct model* untuk masing-masing tabel **User, Photo, Comment, Social Media** dalam folder **models**
4. Buat file **db** untuk konfigurasi database dalam folder **config**
5. Buat folder **helpers** yang akan digunakan untuk menyimpan file yang mana setiap fungsinya dapat melakukan satu pekerjaan yang spesifik tanpa ada ketergantungan terhadap fungsi yang lain.
6. Buat folder **middlewares** yang akan digunakan untuk menyimpan file yang mana fungsinya dapat mengakses permintaan object (req), respons object (res), dan setiap siklus permintaan dan respon tersebut (next)
7. Buat folder **controllers** yang akan digunakan untuk menyimpan file *controller* yang digunakan untuk menentukan flow dari masing-masing *table* dalam artian **CRUD (Create, Read, Update, Delete)**
8. Buat file **router** dalam folder **router** untuk konfigurasi rute yang menentukan controller dan method/fungsi yang akan dieksekusi

---

## HOW TO RUN
1. Inisialisasi aplikasi dengan perintah berikut pada terminal
    ```
    go mod init mygram
    ```
2. Instal *library* yang akan diperlukan dengan perintah sebagai berikut pada terminal
    ```
    go get -u golang.org/x/crypto/bcrypt
    go get -u github.com/lib/pq
    go get -u gorm.io/gorm
    go get -u gorm.io/driver/postgres
    go get -u ithub.com/gin-gonic/gin
    go get -u github.com/asaskevich/govalidator
    go get -u github.com/dgrijalva/jwt-go
    go get -u github.com/joho/godotenv
    ```
3. Namun jika sudah terdapat file **go.mod**, cukup jalankan perintah berikut untuk meng-*update library* pada terminal
    ```
    go mod tidy
    ```
4. Sebelum menjalankan aplikasi, setting konfigurasi environment terlebih dahulu. Pastikan terdapat file **.env** yang berisikan konfigurasi yang dibutuhkan aplikasi
5. Jalankan aplikasi mygram dengan perintah berikut
    ```
    go run main.go
    ```
6. Untuk terminate, tekan **Ctrl+C**

---

## Postman Collection
https://www.getpostman.com/collections/2531f5e6efe20c077c1f