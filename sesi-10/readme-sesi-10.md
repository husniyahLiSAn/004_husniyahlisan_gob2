# Summary
1. __Middleware__ merupakan sebuah fungsi yang akan terkesekusi sesudah maupun sebelum mencapai sebuah endpoint. _Middleware_ digunakan untuk logging atau untuk mengamankan sebuah endpoint seperti contohnya proses autentikasi dan autorisasi.
2. Untuk membuat middleware pada Golang, kita akan menggunakan package _net/http_ dengan menggunakan __multiplexer__ nya agar kita dapat melakukan kustomisasi.
3. __Json Web Token (JWT)__ adalah sebuah token berbentuk string panjang yang digunakan untuk pertukaran informasi antara 2 belah pihak atau client dan server yang biasa digunakan pada aplikasi web yang berbasis __Single Page Application (SPA)__. Untuk men-_decode_ JWT, dapat mengakses [website resmi JWT](https://jwt.io/).
