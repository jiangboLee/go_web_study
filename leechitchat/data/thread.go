package data

import (
	"time"
)

type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    string
	CreatedAt time.Time
}

//创建一个新的帖子
func (user *User) CreateThread(topic string) (conv Thread, err error) {
	statement := "INSERT threads SET uuid=?, topic=?, user_id=?, created_at=?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	res, err := stmt.Exec(createUUID(), topic, user.Id, time.Now())
	id, err := res.LastInsertId()
	err = Db.QueryRow("SELECT * FROM threads WHERE id=?", id).Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt)
	return
}

//得到所有的帖子
func Threads() (threads []Thread, err error) {
	rows, err := Db.Query("SELECT * FROM threads ORDER BY created_at DESC")
	if err != nil {
		return
	}
	for rows.Next() {
		conv := Thread{}
		if err = rows.Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt); err != nil {
			return
		}
		threads = append(threads, conv)
	}
	rows.Close()
	return
}

//得到用户谁开始这个帖子
func (thread *Thread) User() (user User) {
	user = User{}
	Db.QueryRow("SELECT * FROM users WHERE id=?", thread.UserId).Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

func (thread *Thread) CreatedAtDate() string {
	// loc, _ := time.LoadLocation("Local")
	// AcceptTime:="2015-01-12 16:44:33"
	// t, _ := time.ParseInLocation("2006-01-02 15:04:05", AcceptTime, loc)
	return thread.CreatedAt.Local().Format("Jan 2, 2006 at 3:11pm")
}

func (thread *Thread) NumReplies() (count int) {
	return 2
}
