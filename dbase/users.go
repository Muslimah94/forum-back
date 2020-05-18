package dbase

import (
	"database/sql"
	"fmt"

	models "../models"
)

// SelectUsers ...
func (db *DataBase) SelectUsers() ([]models.User, error) {
	rows, err := db.DB.Query(`SELECT ID, Nickname, RoleID FROM Users`)
	if err != nil {
		fmt.Println("SelectUsers db.Query ERROR:", err)
		return nil, err
	}
	defer rows.Close()
	var AllUsers []models.User
	for rows.Next() {
		var u models.User
		err = rows.Scan(&u.ID, &u.Nickname, &u.RoleID)
		if err != nil {
			fmt.Println("SelectUsers rows.Scan ERROR:", err)
			continue
		}
		AllUsers = append(AllUsers, u)
	}
	if err = rows.Err(); err != nil {
		fmt.Println("SelectUsers rows ERROR:", err)
		return nil, err
	}
	return AllUsers, nil
}

// SelectUserByID ...
func (db *DataBase) SelectUserByID(userID int) (models.User, error) {
	var u models.User
	rows, err := db.DB.Query(`SELECT ID, Nickname, RoleID FROM Users WHERE ID = ?`, userID)
	if err != nil {
		fmt.Println("SelectUserByID db.Query ERROR:", err)
		return u, err
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&u.ID, &u.Nickname, &u.RoleID)
		if err != nil {
			fmt.Println("SelectUserByID rows.Scan ERROR:", err)
			return u, err
		}
	}
	if err = rows.Err(); err != nil {
		fmt.Println("SelectUserByID rows ERROR:", err)
		return u, err
	}
	return u, nil
}

func (db *DataBase) SelectUser() {}

func (db *DataBase) CreateUser(new models.User, tx *sql.Tx) (int, error) {
	fmt.Println("CreateUSER")
	n := 0
	st, err := tx.Prepare(`INSERT INTO Users (Nickname, RoleID) VALUES (?,?)`)
	defer st.Close()
	if err != nil {
		fmt.Println("CreateUser Prepare", err)
		tx.Rollback()
		return n, err
	}
	_, err = st.Exec(new.Nickname, new.RoleID)
	if err != nil {
		fmt.Println("CreateUser Exec", err)
		tx.Rollback()
		return n, err
	}
	n, err = db.ReturnLastUserID()
	if err != nil {
		fmt.Println("CreateUser Exec", err)
		tx.Rollback()
		return n, err
	}
	return n, nil
}

// ReturnLastUserID ...
func (db *DataBase) ReturnLastUserID() (int, error) {
	n := 0
	rows, err := db.DB.Query(`SELECT ID FROM Users ORDER BY ID DESC LIMIT 1`)
	defer rows.Close()
	if err != nil {
		fmt.Println("ReturnLastUserID Query:", err)
		return n, err
	}
	for rows.Next() {
		err = rows.Scan(&n)
		if err != nil {
			fmt.Println("ReturnLastUserID rows.Scan:", err)
			continue
		}
	}
	if err = rows.Err(); err != nil {
		fmt.Println("ReturnLastUserID rows:", err)
		return n, err
	}
	return n, nil
}

// // AddNewUser ...
// func AddNewUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {

// 	//check session (true false) { est' ili net, timestamp < time.Now {
// 	// 	delete session
// 	// } else {
// 	// update current+3600
// 	// set cookie MaxAge 3600
// 	// *******
// 	// 	}
// 	// est' li net , redirect

// 	var user *models.Users
// 	handlers.ReceiveJSON(r, &user)
// 	st, err2 := db.Prepare(`INSERT INTO Users (Email, Nickname, Password, RoleID) VALUES (?,?,?,?)`)
// 	if err2 != nil {
// 		fmt.Println("AddNewUser db.Prepare", err2)
// 		http.Error(w, err2.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	_, err3 := st.Exec(user.Email, user.Nickname, user.Password, user.RoleID)
// 	if err3 != nil {
// 		fmt.Println("AddNewUser st.Exec", err3)
// 		http.Error(w, err3.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }

// // GetUserByID ...
// func GetUserByID(db *sql.DB, w http.ResponseWriter, r *http.Request, userID int) {
// 	var user models.Users
// 	rows := db.QueryRow(`SELECT * FROM Users WHERE ID = $1`, userID)
// 	err := rows.Scan(&user.ID, &user.Email, &user.Nickname, &user.Password, &user.RoleID)
// 	if err != nil {
// 		fmt.Println("GetUserByI rows.Scan ERROR:", err)
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	handlers.SendJSON(w, user)
// }

// // EditUserByID ...
// func EditUserByID(db *sql.DB, w http.ResponseWriter, r *http.Request, userID int) {
// 	var new *models.Users
// 	handlers.ReceiveJSON(r, &new)
// 	st, err2 := db.Prepare(`UPDATE Users SET Email = ?, Nickname = ?, Password = ?, RoleID = ? where ID = ?`)
// 	if err2 != nil {
// 		fmt.Println("EditUserByID db.Prepare:", err2)
// 		http.Error(w, err2.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	_, err3 := st.Exec(new.Email, new.Nickname, new.Password, new.RoleID, userID)
// 	if err3 != nil {
// 		fmt.Println("EditUserByID st.Exec:", err3)
// 		http.Error(w, err3.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }

// // DeleteUserByID ...
// func DeleteUserByID(db *sql.DB, w http.ResponseWriter, r *http.Request, userID int) {

// 	st, err1 := db.Prepare(`DELETE FROM Users WHERE ID = ?`)
// 	if err1 != nil {
// 		fmt.Println("DeleteUserByID db.Prepare:", err1)
// 		http.Error(w, err1.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	_, err2 := st.Exec(userID)
// 	if err2 != nil {
// 		fmt.Println("DeleteUserByID st.Exec:", err2)
// 		http.Error(w, err2.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }

// // GetUsersByRoleID ...
// func GetUsersByRoleID(db *sql.DB, w http.ResponseWriter, r *http.Request, roleID int) {
// 	rows, err1 := db.Query(`SELECT * FROM Users WHERE RoleID = ?`, roleID)
// 	if err1 != nil {
// 		fmt.Println("GetUsersByRoleID db.Query ERROR:", err1)
// 		http.Error(w, err1.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	defer rows.Close()
// 	var AllUsers []models.Users
// 	for rows.Next() {
// 		var u models.Users
// 		err2 := rows.Scan(&u.ID, &u.Email, &u.Nickname, &u.Password, &u.RoleID)
// 		if err2 != nil {
// 			fmt.Println("GetUsersByRoleID rows.Scan ERROR:", err2)
// 			continue
// 		}
// 		AllUsers = append(AllUsers, u)
// 	}
// 	if err3 := rows.Err(); err3 != nil {
// 		fmt.Println("GetUsersByRoleID rows ERROR:", err3)
// 		http.Error(w, err3.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	handlers.SendJSON(w, AllUsers)
// }
