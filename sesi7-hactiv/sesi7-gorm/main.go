package main

import (
	"fmt"
	"sesi7-gorm/database"
	"sesi7-gorm/repository"
)

func main() {
	//db conn
	db := database.StartDB()
	userRepo := repository.NewUserRepo(db)
	productRepo := repository.NewProductRepo(db)

	// ============== USERS ===================

	// user := models.User{
	// 	Email: "user2@hactiv.com",
	// }

	// //create user
	// err := userRepo.CreateUser(&user)

	// if err != nil {
	// 	fmt.Println("Error :", err)
	// 	return
	// }

	fmt.Println("Created Success")
	fmt.Println("")

	//GET ALL USERS
	fmt.Println("=== GET ALL USER ===")
	employees, err := userRepo.GetAllUsers()
	if err != nil {
		fmt.Println("Error when getAllUsers :", err.Error())
		return
	}

	//print result
	for k, emp := range *employees {
		fmt.Println("User:", k+1)
		emp.Print()
		fmt.Println("")
	}

	//get user by id
	fmt.Println("=== GET USER BY ID ===")
	usr, err := userRepo.GetUserById(3)
	if err != nil {
		fmt.Println("Error when getUserById:", err.Error())
		return
	}
	usr.Print()

	//DELETE USER BY ID
	fmt.Println("=== DELETE USER BY ID ===")
	err = userRepo.DeleteUserById(1)
	if err != nil {
		fmt.Println("Error when Delete user by id : ", err.Error())
		return
	}
	fmt.Println("DELETE SUCCESS")
	fmt.Println("")

	// UPDATE DATA BY ID
	err = userRepo.UpdateUserById(3, "hello@gmail.com")
	if err != nil {
		fmt.Println("Error update user by id :", err.Error())
		return
	}

	// ================== PRODUCTS ====================

	// product := models.Product{
	// 	Name:   "Topi",
	// 	Brand:  "NY",
	// 	UserID: 3,
	// }

	// //create product
	// err = productRepo.CreateProduct(&product)

	// if err != nil {
	// 	fmt.Println("Error :", err)
	// 	return
	// }

	//get all product
	fmt.Println("")
	fmt.Println("=== GET ALL PRODUCT ===")
	products, err := productRepo.GetAllProduct()
	if err != nil {
		fmt.Println("Error when get all Product :", err.Error())
		return
	}

	for _, v := range *products {
		v.Print()
	}

	// get product by id
	fmt.Println("")
	fmt.Println("=== GET PRODUCT BY ID ===")
	productById, err := productRepo.GetProductById(2)
	if err != nil {
		fmt.Println("Error when get product by id :", err.Error())
		return
	}

	productById.Print()

	//update product
	fmt.Println("")
	fmt.Println("=== UPDATE PRODUCT BY ID ===")
	err = productRepo.UpdateProductById(3, "bandana", "ABC")
	if err != nil {
		fmt.Println("Error when update product :", err.Error())
		return
	}

	//delete product
	fmt.Println("")
	fmt.Println("=== DELETE PRODUCT BY ID ===")
	err = productRepo.DeleteProductById(3)
	if err != nil {
		fmt.Println("Error when delete product :", err.Error())
		return
	}
	fmt.Println("success delete product")
	fmt.Println("")
}
