package dbase

import (
	"fmt"

	models "../models"
)

// CountReactionsToPost ...
func (db *DataBase) CountReactionsToPost(t int, postID int) (int, error) {
	num := 0
	rows, err := db.DB.Query(`SELECT COUNT(*) FROM Reactions WHERE Type = ? AND PostID = ?`, t, postID)
	defer rows.Close()
	if err != nil {
		fmt.Println("CountReactionsToPost Query:", err)
		return 0, err
	}
	if rows.Next() {
		err = rows.Scan(&num)
		if err != nil {
			fmt.Println("CountReactionsToPost rows.Scan:", err)
		}
	}
	if err = rows.Err(); err != nil {
		fmt.Println("CountReactionsToPost rows:", err)
		return 0, err
	}
	return num, nil
}

// CountReactionsToComment ...
func (db *DataBase) CountReactionsToComment(t int, commentID int) (int, error) {
	num := 0
	rows, err := db.DB.Query(`SELECT COUNT(*) FROM Reactions WHERE Type = ? AND CommentID = ?`, t, commentID)
	defer rows.Close()
	if err != nil {
		fmt.Println("CountReactionsToComment Query:", err)
		return 0, err
	}
	if rows.Next() {
		err = rows.Scan(&num)
		if err != nil {
			fmt.Println("CountReactionsToComment rows.Scan:", err)
		}
	}
	if err = rows.Err(); err != nil {
		fmt.Println("CountReactionsToComment rows:", err)
		return 0, err
	}
	return num, nil
}

// CreateReaction ...
func (db *DataBase) CreateReaction(new models.Reaction) error {

	if new.PostID == 0 {
		st, err := db.DB.Prepare(`INSERT INTO Reactions (AuthorID, Type, CommentID) VALUES (?,?,?)`)
		defer st.Close()
		if err != nil {
			fmt.Println("CreateReaction Prepare", err)
			return err
		}
		_, err = st.Exec(new.AuthorID, new.Type, new.CommentID)
		if err != nil {
			fmt.Println("CreateReaction Exec", err)
			return err
		}
	} else {
		st, err := db.DB.Prepare(`INSERT INTO Reactions (AuthorID, Type, PostID) VALUES (?,?,?)`)
		defer st.Close()
		if err != nil {
			fmt.Println("CreateReaction Prepare", err)
			return err
		}
		_, err = st.Exec(new.AuthorID, new.Type, new.PostID)
		if err != nil {
			fmt.Println("CreateReaction Exec", err)
			return err
		}
	}
	return nil
}

// SelectReaction ...
func (db *DataBase) SelectReaction(new models.Reaction) (int, error) {
	num := 0
	if new.PostID == 0 {
		rows, err := db.DB.Query(`SELECT COUNT(*) FROM Reactions WHERE Type = ? AND AuthorID = ? AND CommentID = ?`, new.Type, new.AuthorID, new.CommentID)
		defer rows.Close()
		if err != nil {
			fmt.Println("SelectReaction Query[comment]:", err)
			return 0, err
		}
		if rows.Next() {
			err = rows.Scan(&num)
			if err != nil {
				fmt.Println("SelectReaction[comment] rows.Scan:", err)
			}
		}
		if err = rows.Err(); err != nil {
			fmt.Println("SelectReaction[comment] rows:", err)
			return 0, err
		}
	} else {
		rows, err := db.DB.Query(`SELECT COUNT(*) FROM Reactions WHERE Type = ? AND AuthorID = ? AND PostID = ?`, new.Type, new.AuthorID, new.PostID)
		defer rows.Close()
		if err != nil {
			fmt.Println("SelectReaction Query[post]:", err)
			return 0, err
		}
		if rows.Next() {
			err = rows.Scan(&num)
			if err != nil {
				fmt.Println("SelectReaction[post] rows.Scan:", err)
			}
		}
		if err = rows.Err(); err != nil {
			fmt.Println("SelectReaction[post] rows:", err)
			return 0, err
		}
	}
	return num, nil
}

// UpdateReaction ...
func (db *DataBase) UpdateReaction(new models.Reaction) error {

	if new.PostID == 0 {
		stmt, err := db.DB.Prepare(`UPDATE Reactions SET type = ? WHERE AuthorID = ? AND CommentID = ?`)
		defer stmt.Close()
		if err != nil {
			fmt.Println("UpdateReaction Prepare[comment]", err)
			return err
		}
		_, err = stmt.Exec(new.Type, new.AuthorID, new.CommentID)
		if err != nil {
			fmt.Println("UpdateReaction Exec[comment]", err)
			return err
		}
	} else {
		stmt, err := db.DB.Prepare(`UPDATE Reactions SET type = ? WHERE AuthorID = ? AND PostID = ?`)
		defer stmt.Close()
		if err != nil {
			fmt.Println("UpdateReaction Prepare[post]", err)
			return err
		}
		_, err = stmt.Exec(new.Type, new.AuthorID, new.PostID)
		if err != nil {
			fmt.Println("UpdateReaction Exec[post]", err)
			return err
		}
	}
	return nil
}
