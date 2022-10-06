# Summary
1. Untuk menggunakan package _database/sql_ kita perlu menginstal sebuah sql driver, misalnya untuk postgresql
   ```
   go get -u github.com/lib/pq
   ```
2. Create database dapat dijalankan pada __GUI pgAdmin__ atau GUI lainnya maupun langsung melalui __psql__ pada terminal
3. Untuk koneksi ke database, kita perlu meregistrasikan konfigurasi PostgreSQL yang telah diberikan saat pertama kali instal beserta nama database
4. __Gorm__ adalah ORM yang cukup populer untuk bahasa Go, yang dimana Gorm telah menyediakan berbagai macam fitur seperti auto migration, eager loading, association, query method, dan lain-lain.
5. Ketikkan perintah berikut untuk menginstal Gorm
   ```
   go get -u gorm.io/gorm
   ```
6. Cara menginstal driver PostgreSQL dari Gorm dengan menjalankan perintah berikut
   ```
   go get gorm.io/driver/postgres
   ```
7. Gorm juga telah menyediakan hooks yang dapat kita gunakan, salah satunya adalah method BeforeCreate. Method BeforeCreate akan tereksekusi sebelum melakukan create data.
8. Untuk melakukan _join statement_ dapat menggunakan eager loading dari Gorm. Caranya adalah dengan menggunakan method Preload dan kita perlu memberikan nama table untuk parameter method Preload.