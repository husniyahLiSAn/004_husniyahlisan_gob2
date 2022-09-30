# Summary
1. _Interface_ pada Golang adalah sebuah tipe data yang merupakan kumpulan dari definisi satu atau lebih method dan jika ingin menggunakannya maka perlu mengimplementasikan method-method yang telah didefinisikan sebelumnya
2. Method yang tidak didefinisikan pada interface sebelumnya dapat menggunakan _type assertion_ untuk menanggulangi error pada saat compile
3. _Empty interface_ merupakan suatu tipe data yang dapat menerima segala tipe data pada Golang
4. _Reflect_ digunakan untuk melakukan inspeksi variable, mengambil informasinya bahkan dapat memanipulasinya. Cakupan informasi yang bisa didapatkan sangat luas, seperti struktur variable, tipe, nilai pointer dan masih banyak lagi. Fungsi yang paling penting untuk diketahui adalah:
   ```
   reflect.ValueOf()
   reflect.TypeOf()
   ```
5. Jika untuk mengembalikan nilai interface kosong dapat menggunakan method berikut:
   ```
   .Interface()
   ```
6. Goroutine merupakan sebuah thread ringan pada Golang untuk melakukan concurrency. Arti dari concurrency sendiri adalah mengeksekusi sebuah proses secara independen atau berhadapan lebih dari satu tugas dalam waktu yang sama
7. Waitgroup merupakan sebuah struct dari package sync yang digunakan untuk melakukan sinkronisasi terhadap go routine
