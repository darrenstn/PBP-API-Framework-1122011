package controllers

import (
	"strconv"

	m "eksplorasi2/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
)

// GetAllUsers..
// func GetAllUsers(w http.ResponseWriter, r *http.Request) {
// 	db := connect()
// 	defer db.Close()

// 	query := "SELECT id, name, age, address, email FROM users"

// 	user_id := r.URL.Query()["user_id"]

// 	if user_id != nil {
// 		query += " WHERE id='" + user_id[0] + "'"
// 	}

// 	rows, err := db.Query(query)
// 	if err != nil {
// 		log.Println(err)
// 		var response m.Response
// 		response.Status = 400
// 		response.Message = "Error"
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	var user m.User
// 	var users []m.User

// 	for rows.Next() {
// 		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address, &user.Email); err != nil {
// 			log.Println(err)
// 			return
// 		} else {
// 			users = append(users, user)
// 		}
// 	}

// 	w.Header().Set("Content-Type", "application/json")

// 	var response m.UsersResponse
// 	response.Status = 200
// 	response.Message = "Success"
// 	response.Data = users
// 	json.NewEncoder(w).Encode(response)
// }

func GetAllUsers(c buffalo.Context) error {
	// Connect to the database
	db := connect()
	defer db.Close()

	// Initial query to fetch all users
	query := "SELECT id, name, age, address, email FROM users"

	// Check if user_id parameter is provided in the request
	userID := c.Param("user_id")
	if userID != "" {
		query += " WHERE id = '" + userID + "'"
	}

	// Execute the query
	rows, err := db.Query(query)
	if err != nil {
		// Handle database query error
		return c.Render(400, render.String("Error"))
	}

	// Define slices to hold user data
	var user m.User
	var users []m.User

	// Loop through the rows and scan into user struct
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address, &user.Email); err != nil {
			// Handle scanning error
			return c.Render(400, render.String("Error"))
		} else {
			users = append(users, user)
		}
	}

	// Set response header
	c.Response().Header().Set("Content-Type", "application/json")

	// Prepare response data
	response := m.UsersResponse{
		Status:  200,
		Message: "Success",
		Data:    users,
	}

	// Render response as JSON and send
	return c.Render(200, render.JSON(response))
}

// // InsertUser..
// func InsertUser(w http.ResponseWriter, r *http.Request) {
// 	db := connect()
// 	defer db.Close()

// 	err := r.ParseForm()
// 	if err != nil {
// 		return
// 	}

// 	name := r.Form.Get("name")
// 	age, _ := strconv.Atoi(r.Form.Get("age"))
// 	address := r.Form.Get("address")
// 	email := r.Form.Get("email")
// 	password := r.Form.Get("password")

// 	_, errQuery := db.Exec("INSERT INTO users(name, age, address, email, password) values(?, ?, ?, ?, ?)",
// 		name,
// 		age,
// 		address,
// 		email,
// 		password,
// 	)

// 	var response m.Response
// 	if errQuery == nil {
// 		response.Status = 200
// 		response.Message = "Success"
// 	} else {
// 		response.Status = 400
// 		response.Message = "Insert Failed!"
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(response)
// }

func InsertUser(c buffalo.Context) error {
	// Connect to the database
	db := connect()
	defer db.Close()

	// Parse form data
	err := c.Request().ParseForm()
	if err != nil {
		return err
	}

	// Extract form values
	name := c.Request().Form.Get("name")
	ageStr := c.Request().Form.Get("age")
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		return err
	}
	address := c.Request().Form.Get("address")
	email := c.Request().Form.Get("email")
	password := c.Request().Form.Get("password")

	// Execute the query
	_, err = db.Exec("INSERT INTO users(name, age, address, email, password) VALUES(?, ?, ?, ?, ?)",
		name, age, address, email, password)
	if err != nil {
		return err
	}

	// Prepare response
	response := m.Response{}
	response.Status = 200
	response.Message = "Success"

	// Render response as JSON and send
	return c.Render(200, render.JSON(response))
}

// // UpdateUser..
// func UpdateUser(w http.ResponseWriter, r *http.Request) {
// 	db := connect()
// 	defer db.Close()

// 	err := r.ParseForm()
// 	if err != nil {
// 		return
// 	}

// 	id, _ := strconv.Atoi(r.Form.Get("id"))
// 	name := r.Form.Get("name")
// 	age, _ := strconv.Atoi(r.Form.Get("age"))
// 	address := r.Form.Get("address")
// 	email := r.Form.Get("email")
// 	password := r.Form.Get("password")

// 	query, _ := db.Prepare("UPDATE users SET name=?, age=?, address=?, email=?, password=? WHERE id=?")
// 	_, errQuery := query.Exec(name, age, address, email, password, id)

// 	var response m.Response
// 	if errQuery == nil {
// 		response.Status = 200
// 		response.Message = "Success"
// 	} else {
// 		response.Status = 400
// 		response.Message = "Update user Failed!"
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(response)
// }

func UpdateUser(c buffalo.Context) error {
	// Connect to the database
	db := connect()
	defer db.Close()

	// Parse form data
	err := c.Request().ParseForm()
	if err != nil {
		return err
	}

	// Extract form values
	idStr := c.Request().Form.Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}
	name := c.Request().Form.Get("name")
	ageStr := c.Request().Form.Get("age")
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		return err
	}
	address := c.Request().Form.Get("address")
	email := c.Request().Form.Get("email")
	password := c.Request().Form.Get("password")

	// Execute the query
	query, err := db.Prepare("UPDATE users SET name=?, age=?, address=?, email=?, password=? WHERE id=?")
	if err != nil {
		return err
	}
	_, err = query.Exec(name, age, address, email, password, id)
	if err != nil {
		return err
	}

	// Prepare response
	response := m.Response{}
	response.Status = 200
	response.Message = "Success"

	// Render response as JSON and send
	return c.Render(200, render.JSON(response))
}

// // DeleteUser..
// func DeleteUser(w http.ResponseWriter, r *http.Request) {
// 	db := connect()
// 	defer db.Close()

// 	err := r.ParseForm()
// 	if err != nil {
// 		return
// 	}

// 	vars := mux.Vars(r)
// 	userId := vars["user_id"]

// 	_, errQuery := db.Exec("DELETE FROM users WHERE id=?",
// 		userId,
// 	)

// 	if errQuery == nil {
// 		sendSuccessResponse(w)
// 	} else {
// 		sendErrorResponse(w)
// 	}
// }

func DeleteUser(c buffalo.Context) error {
	// Connect to the database
	db := connect()
	defer db.Close()

	// Extract user_id from URL parameters
	userID := c.Param("user_id")

	// Execute the delete query
	_, err := db.Exec("DELETE FROM users WHERE id=?", userID)
	if err != nil {
		// Handle delete error
		return c.Error(400, err)
	}

	// Prepare response
	response := m.Response{
		Status:  200,
		Message: "Success",
	}

	// Render response as JSON and send
	return c.Render(200, render.JSON(response))
}
