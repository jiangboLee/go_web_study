package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {
	db, err := sql.Open("mysql", "root:147369@tcp(140.143.239.161:3306)/chitchat?parseTime=true")
	checkErr(err)
	//插入数据
	stmt, err := db.Prepare("insert users set uuid=?, name=?, email=?, password=?, created_at=?")
	checkErr(err)

	res, err := stmt.Exec(222555, "lee", "程s序猿", "5s555", time.Now())
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println("LastInsertId :", id)

	// //更新数据
	// stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	// checkErr(err)

	// res, err = stmt.Exec("jiang", id)
	// checkErr(err)

	// affect, err := res.RowsAffected()
	// checkErr(err)

	// fmt.Println("RowsAffected :", affect)

	// //查询数据
	// rows, err := db.Query("SELECT * FROM userinfo")
	// checkErr(err)

	// for rows.Next() {
	// 	var (
	// 		uid        int
	// 		username   string
	// 		department string
	// 		created    string
	// 	)
	// 	err = rows.Scan(&uid, &username, &department, &created)
	// 	checkErr(err)
	// 	fmt.Println(uid)
	// 	fmt.Println(username)
	// 	fmt.Println(department)
	// 	fmt.Println(created)
	// }

	// //删除数据
	// stmt, err = db.Prepare("delete from userinfo where uid=?")
	// checkErr(err)

	// res, err = stmt.Exec(id)
	// checkErr(err)

	// affect, err = res.RowsAffected()
	// checkErr(err)

	// fmt.Println(affect)

	db.Close()

}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
