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

type Post struct {
	Id        int
	Uuid      string
	Body      string
	UserId    int
	ThreadId  int
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
	return thread.CreatedAt.Local().Format("Jan 2, 2006 at 3:11pm")
}

func (post *Post) User() (user User) {
	user = User{}
	Db.QueryRow("SELECT * FROM users WHERE id=?", post.UserId).Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

func (post *Post) CreatedAtDate() string {
	return post.CreatedAt.Local().Format("Jan 2, 2006 at 3:11pm")
}

func (user *User) CreatePost(conv Thread, body string) (post Post, err error) {
	statement := "INSERT posts SET uuid=?, body=?, user_id=?, thread_id=?, created_at=?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	res, err := stmt.Exec(createUUID(), body, user.Id, conv.Id, time.Now())
	id, err := res.LastInsertId()
	err = Db.QueryRow("SELECT * FROM posts WHERE id=?", id).Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt)
	return
}

func (thread *Thread) Posts() (posts []Post, err error) {
	rows, err := Db.Query("SELECT * FROM posts WHERE thread_id=?", thread.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post{}
		if err = rows.Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt); err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

func (thread *Thread) NumReplies() (count int) {
	return 2
}

//根据uuid得到帖子
func ThreadByUUID(uuid string) (conv Thread, err error) {
	conv = Thread{}
	err = Db.QueryRow("SELECT * FROM threads WHERE uuid=?", uuid).Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt)
	return
}
