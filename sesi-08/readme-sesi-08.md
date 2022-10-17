# Summary
1. __Decode__ data _JSON_ kepada sebuah tipe data lain adalah dengan menggunakan _function_ __json.Unmarshal__ dengan parameter pertama sebuah tipe data __slice of byte__ dan parameter kedua variable dari tipe data yang ingin disimpan.
    * Decode JSON To Struct
    * Decode JSON To Map
    * Decode JSON To Empty Interface
    * Decode JSON Array To Slice Of Struct
2. Function __url.Parse__ digunakan untuk parsing string ke bentuk url. Mengembalikan 2 data, variabel objek bertipe url.URL dan error(jika ada). Melalui variabel objek tersebut bisa didapatkan data yang dimiliki oleh URL yaitu jenis protokol yang digunakan, path yang diakses, query, dan lainnya.
3. Untuk mendokumentasikan __REST API__ dapat menggunakan _library_ __Swagger__. Swagger UI memberikan cara yang nyaman bagi consumer untuk menjelajahi API. __Swaggo__ dan __go-swagger__ adalah dua framework paling populer yang tersedia untuk menghasilkan dokumen.

---
## Swagger

**Step Swaggo:**

> mkdir swaggo
> cd swaggo
> 
> go mod init swaggo

**Library:**
> go get -u github.com/swaggo/swag/cmd/swag
> 
> go get -u github.com/swaggo/http-swagger
> 
> go get -u github.com/alecthomas/template
> 
> go get -u github.com/gorilla/mux

**Generate swaggo/docs:**
> swag init -g main.go

**Global:**
> go install github.com/swaggo/swag/cmd/swag

**MacOS: Add Path**
> export PATH=$(go env GOPATH)/bin:$PATH
