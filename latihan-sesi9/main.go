package main

import (
	"encoding/json"
	"fmt"
	"latihan-sesi9/middleware"
	"latihan-sesi9/models"
	"latihan-sesi9/params"
	"net/http"
)

func main() {
	/*
		Latihan :
		Buatlah sebuah REST API untuk :
			- Register
			  ini cukup dengan menginputkan username dan password.
			  data user bisa lebih dari 1, jadi silahkan gunakan slice.

			- GetAllUsers
			  untuk cek seluruh user. endpoint ini khusus untuk user yg sudah
			  didaftarkan.

			- AddNewProducts
			  yang bisa hit endpoint ini hanyalah user yang sudah di daftarkan.
			  data yang dibutuhkan adalah :
			  	- Nama
				- Brand
				- Stok
				- Price

			- GetProducts
			  tidak ada auth disini. jadi ini adalah API open. akan ngereturn :
			  payload : [
				  {
					  nama 	: "",
					  brand : "",
					  stok 	: 0,
					  price	: 0
				  }
			  ]

			- GetProductByBrand
			  tidak ada auth disini. jadi ini adalah API open.
			  Brand akan di dapat dari query. akan ngereturn :
			  payload : [
				  {
					  nama 	: "",
					  brand : "",
					  stok 	: 0,
					  price	: 0
				  }
			  ]
	*/

	http.HandleFunc("/users/register", RegisterUser) //POST
	http.HandleFunc("/users", GetAllUsers)           //GET
	http.HandleFunc("/product", AddNewProducts)      //POST
	http.HandleFunc("/products", GetAllProducts)     //GET

	server := new(http.Server)
	port := ":4444"

	fmt.Println("server running at port :", port)
	server.Addr = port
	server.ListenAndServe()

}

func RegisterUser(rw http.ResponseWriter, r *http.Request) {
	method := r.Method
	rw.Header().Add("Content-Type", "application/json")
	if method == http.MethodPost {
		var request params.CreateUser

		err := json.NewDecoder(r.Body).Decode(&request)

		if err != nil {
			json.NewEncoder(rw).Encode(params.Response{
				Status:         http.StatusBadRequest,
				Message:        "BAD REQUEST",
				AdditionalInfo: err.Error(),
			})
		}

		usr := models.User{
			Username: request.Username,
			Password: request.Password,
		}

		models.Users = append(models.Users, usr)

		json.NewEncoder(rw).Encode(params.Response{
			Status:         http.StatusAccepted,
			Message:        "ACCEPTED",
			AdditionalInfo: "Create User Request Accepted.",
		})
	} else {
		json.NewEncoder(rw).Encode(params.Response{
			Status:         http.StatusMethodNotAllowed,
			Message:        "METHOD NOT ALLOWED",
			AdditionalInfo: "METHOD NOT ALLOWED",
		})
	}
}

func GetAllUsers(rw http.ResponseWriter, r *http.Request) {

	if !middleware.Auth(rw, r) {
		return
	}

	method := r.Method
	if method == http.MethodGet {
		usersData := models.GetUsers()
		json.NewEncoder(rw).Encode(params.Response{
			Status:  http.StatusOK,
			Message: "OK",
			Payload: usersData,
		})
	} else {
		json.NewEncoder(rw).Encode(params.Response{
			Status:         http.StatusMethodNotAllowed,
			Message:        "METHOD NOT ALLOWED",
			AdditionalInfo: "METHOD NOT ALLOWED",
		})
	}
}

func GetAllProducts(rw http.ResponseWriter, r *http.Request) {

	method := r.Method
	if method == http.MethodGet {

		query := r.URL.Query()
		brand := query.Get("brand")

		if brand == "" {
			productData := models.GetProducts()
			json.NewEncoder(rw).Encode(params.Response{
				Status:  http.StatusOK,
				Message: "OK",
				Payload: productData,
			})
			return
		} else {
			productData, err := models.GetProductsByBrand(brand)

			if err != nil {
				json.NewEncoder(rw).Encode(params.Response{
					Status:         http.StatusMethodNotAllowed,
					Message:        "NOT FOUND",
					AdditionalInfo: "METHOD NOT ALLOWED",
				})
			}
			json.NewEncoder(rw).Encode(params.Response{
				Status:  http.StatusOK,
				Message: "OK",
				Payload: productData,
			})
		}

	} else {
		json.NewEncoder(rw).Encode(params.Response{
			Status:         http.StatusMethodNotAllowed,
			Message:        "METHOD NOT ALLOWED",
			AdditionalInfo: "METHOD NOT ALLOWED",
		})
	}
}

func AddNewProducts(rw http.ResponseWriter, r *http.Request) {
	if !middleware.Auth(rw, r) {
		return
	}

	method := r.Method
	rw.Header().Add("Content-Type", "application/json")
	if method == http.MethodPost {
		var request params.CreateProduct

		err := json.NewDecoder(r.Body).Decode(&request)

		if err != nil {
			json.NewEncoder(rw).Encode(params.Response{
				Status:         http.StatusBadRequest,
				Message:        "BAD REQUEST",
				AdditionalInfo: err.Error(),
			})
		}

		prdct := models.Product{
			Name:  request.Name,
			Brand: request.Brand,
			Stock: request.Stock,
			Price: request.Price,
		}

		models.Products = append(models.Products, prdct)

		json.NewEncoder(rw).Encode(params.Response{
			Status:         http.StatusAccepted,
			Message:        "ACCEPTED",
			AdditionalInfo: "Create User Request Accepted.",
		})
	} else {
		json.NewEncoder(rw).Encode(params.Response{
			Status:         http.StatusMethodNotAllowed,
			Message:        "METHOD NOT ALLOWED",
			AdditionalInfo: "METHOD NOT ALLOWED",
		})
	}
}
