package main

import (
	"bufio"
	"errors"
	"fmt"
	"go-gorm/database"
	"go-gorm/models"
	"os"

	"gorm.io/gorm"
)

func main() {
	database.StartDB()

	Route()
}

func createUser(email string) {
	db := database.GetDB()

	User := models.User{
		Email: email,
	}

	err := db.Create(&User).Error

	if err != nil {
		fmt.Println("Error creating user data:", err)
		return
	}

	fmt.Println("New User Data")
}

func getAllUser() {
	db := database.GetDB()

	user := models.User{}

	err := db.First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User data not found")
			return
		}
		print("Error finding user:", err)
	}

	fmt.Printf("User Data: %+v\n\n", user)
}

func getUserById(id uint) {
	db := database.GetDB()

	user := models.User{}

	err := db.First(&user, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User data not found")
			return
		}
		print("Error finding user:", err)
	}

	fmt.Printf("User Data: %+v\n\n", user)
}

func updateUserById(id uint, email string) {
	db := database.GetDB()

	user := models.User{}

	err := db.Model(&user).Where("id = ?", id).Updates(models.User{Email: email}).Error

	if err != nil {
		fmt.Println("Error updating user data:", err)
		return
	}

	fmt.Printf("Update user's email: %+v \n\n", user.Email)
}

func deleteUserById(id uint) {
	db := database.GetDB()

	user := models.User{}

	err := db.Where("id = ?", id).Delete(&user).Error

	if err != nil {
		fmt.Println("Error Deleting Product:", err.Error())
		return
	}

	fmt.Printf("Product with id %d has been successfully deleted\n\n", id)
}

func createProduct(userId uint, brand string, name string) {
	db := database.GetDB()

	Product := models.Product{
		UserID: userId,
		Brand:  brand,
		Name:   name,
	}

	err := db.Create(&Product).Error

	if err != nil {
		fmt.Println("Error creating product data:", err)
		return
	}

	fmt.Println("New Product Data:", Product)
}

func getUsersWithProduct() {
	db := database.GetDB()

	users := models.User{}
	err := db.Preload("Products").Find(&users).Error

	if err != nil {
		fmt.Println("Error getting user datas with product", err)
		return
	}

	fmt.Println("User Datas With Products")
	fmt.Printf("%+v\n\n", users)
}

func deleteProductById(id uint) {
	db := database.GetDB()

	product := models.Product{}

	err := db.Where("id = ?", id).Delete(&product).Error

	if err != nil {
		fmt.Println("Error Deleting Product:", err.Error())
		return
	}

	fmt.Printf("Product with id %d has been successfully deleted\n\n", id)
}

func Route() {
	var point int
	var loop = 1
	for i := 0; i <= loop; i++ {
		if loop == 1 {
			fmt.Println("==============================")
			fmt.Println("1. Create User")
			fmt.Println("2. Get All User")
			fmt.Println("3. Get User By Id")
			fmt.Println("4. Update User By Id")
			fmt.Println("5. Delete User By Id")
			fmt.Println("6. Create Product")
			fmt.Println("7. Get Users With Product")
			fmt.Println("8. Delete Product By Id")
			fmt.Println("9. Exit")
			fmt.Println("==============================")
			fmt.Print("Masukkan Route : ")
			fmt.Scanln(&point)
			if point == 1 || point == 2 || point == 3 || point == 4 || point == 5 || point == 6 || point == 7 || point == 8 || point == 9 {
				switch point {
				case 1:
					var email string
					fmt.Println("==============================")
					fmt.Print("Masukkan Email user yang akan dibuat: ")
					fmt.Scanln(&email)

					createUser(email)
					i = 0
					loop = 1
				case 2:
					getAllUser()
					i = 0
					loop = 1
				case 3:
					var id uint
					fmt.Println("==============================")
					fmt.Print("Masukkan Id user yang akan di lihat: ")
					fmt.Scanln(&id)
					getUserById(id)
					i = 0
					loop = 1
				case 4:
					var email string
					var id uint
					fmt.Println("==============================")
					fmt.Print("Masukkan Id user yang akan di ubah: ")
					fmt.Scanln(&id)
					fmt.Println("==============================")
					fmt.Print("Masukkan email user yang baru: ")
					fmt.Scanln(&email)
					updateUserById(id, email)
					i = 0
					loop = 1
				case 5:
					var id uint
					fmt.Println("==============================")
					fmt.Print("Masukkan Id user: ")
					fmt.Scanln(&id)
					deleteUserById(id)
					i = 0
					loop = 1
				case 6:
					var id uint
					consoleReader := bufio.NewReader(os.Stdin)
					fmt.Println("==============================")
					fmt.Print("Masukkan Id user: ")
					fmt.Scanln(&id)
					fmt.Println("==============================")
					fmt.Print("Masukkan brand produk: ")
					brand, _ := consoleReader.ReadString('\n')
					fmt.Println("==============================")
					fmt.Print("Masukkan nama produk: ")
					nama, _ := consoleReader.ReadString('\n')
					createProduct(id, brand, nama)
					i = 0
					loop = 1
				case 7:
					fmt.Println("==============================")
					getUsersWithProduct()
					i = 0
					loop = 1
				case 8:
					var id uint
					fmt.Println("==============================")
					fmt.Print("Masukkan Id user: ")
					fmt.Scanln(&id)
					deleteProductById(id)
					i = 0
					loop = 1
				case 9:
					i = 1
					loop = 0
					fmt.Println("==============================")
					fmt.Println("Terima Kasih")
					fmt.Println("==============================")
				}
			} else {
				i = 0
				loop = 1
				fmt.Println("Route tidak ada menu hanya 1 hingga 9")
			}
		}
	}
}
