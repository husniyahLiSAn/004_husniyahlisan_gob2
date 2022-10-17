# Summary

Golang telah menyediakan sebuah _package_ __testing__ untuk digunakan sebagai __unit testing__. Namun agar dapat menggunakannya terdapat aturan yang harus diikuti.
1. File testing harus diakhiri dengan penamaan _test.go
    > Misalkan kita mempunyai sebuah file bernama calculation.go, dan ketika kita ingin membuat file testing untuk file calculation.go, maka perlu membuat file dengan nama calculation_test.go. 
2. Function untuk testing harus diwali dengan penamaan Test.
    > Misalkan kita mempunyai suatu function bernama sum dan kita ingin membuat unit test dari function tersebut, maka kita dapat memberikan nama untuk function unit test kita seperti TestSum.
3. Function testing harus menerima satu parameter dengan tipe data __*testing.T__ dan tidak mengembalikan suatu return value apapun. Tipe data __*testing.T__ merupakan sebuah struct dari package testing untuk membuat suatu unit test.
