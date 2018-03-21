package data

import (
	"fmt"
	"time"
)

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
}

//新建sessionCreateSession
func (user *User) CreateSession() (session Session, err error) {
	statement := "INSERT sessions SET uuid=?, email=?, user_id=?, created_at=?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	res, err := stmt.Exec(createUUID(), user.Email, user.Id, time.Now())
	id, err := res.LastInsertId()
	err = Db.QueryRow("SELECT * FROM sessions WHERE id=?", id).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return
}

//取出session
func (user *User) Session() (session Session, err error) {
	session = Session{}
	fmt.Println("quchusession.........")
	err = Db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE user_id=?", user.Id).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return
}

// Check if session is valid in the database
func (session *Session) Check() (valid bool, err error) {
	fmt.Println(session.Uuid)
	err = Db.QueryRow("SELECT * FROM sessions WHERE uuid=?", session.Uuid).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	if err != nil {
		fmt.Println(err)
		valid = false
		return
	}
	if session.Id != 0 {
		valid = true
	}
	return
}

//在数据库删除session
func (session *Session) DeleteByUUID() (err error) {
	stmt, err := Db.Prepare("DELETE FROM sessions WHERE uuid=?")
	if err != nil {
		fmt.Println("删除失败")
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(session.Uuid)
	fmt.Println("删除成功")
	return
}

//通过session得到user
func (session *Session) User() (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id=?", session.UserId).Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}

//新建用户
func (user *User) Create() (err error) {
	statement := "insert users set uuid=?, name=?, email=?, password=?, created_at=?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(createUUID(), user.Name, user.Email, Encrypt(user.Password), time.Now())
	if err != nil {

		return err
	} else {
		return nil
	}
}

//根据用户得到邮箱
func UserByEmail(email string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT * FROM users WHERE email=?", email).Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}
