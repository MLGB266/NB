package models

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beedb"
	_ "github.com/go-sql-driver/mysql"
)

type Blog struct {
	Id      int `PK`
	Name    string
	Sex     string
	Address string
	Title   string
	Content string
}

func GetLink() beedb.Model {
	db, err := sql.Open("mysql", "root:/8520*963.@/blog?charset=utf8")
	checkErr(err)
	orm := beedb.New(db)
	return orm
}

///////////////////////////////////////////////////////////////展示/////////////////////////////////////////

func GetAll() []Blog {

	//db, err := sql.Open("mysql", "root:/8520*963.@/blog?charset=utf8")
	//checkErr(err)
	db := GetLink()
	var blogs Blog
	bloglist := make([]Blog, 0)
	rows, err := db.Db.Query("SELECT * FROM `blog`.`myblog`")
	checkErr(err)

	defer rows.Close()

	for rows.Next() {

		var Id int
		var Name string
		var Sex string
		var Address string
		var Title string
		var Content string

		err = rows.Scan(&Id, &Name, &Sex, &Address, &Title, &Content)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		blogs.Id, blogs.Name, blogs.Sex, blogs.Address = Id, Name, Sex, Address

		bloglist = append(bloglist, blogs)

	}

	err = rows.Err()
	checkErr(err)

	return bloglist
}

///////////////////////////////////////////////////////////////展示/////////////////////////////////////////////

////////////////////////////////View////////////////////////////////////////

func GetBlog(id int) (blogs Blog) {

	db := GetLink()
	rows, err := db.Db.Query("SELECT * FROM `blog`.`myblog` WHERE Id = ?", id)
	checkErr(err)

	defer rows.Close()

	for rows.Next() {
		var Id int
		var Name string
		var Sex string
		var Address string
		var Title string
		var Content string

		err = rows.Scan(&Id, &Name, &Sex, &Address, &Title, &Content)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		blogs.Id, blogs.Name, blogs.Sex, blogs.Address, blogs.Title, blogs.Content = Id, Name, Sex, Address, Title, Content

	}

	return

}

////////////////////////////////View////////////////////////////////////////

///////////////////////////////////////////////////////////////Save////////////////////////////////////////

func SaveBlog(blog Blog) (blogs Blog) {
	db := GetLink()

	//db, err := sql.Open("mysql", "root:/8520*963.@/blog?charset=utf8")
	//checkErr(err)

	stmt, err := db.Db.Prepare("UPDATE `blog`.`myblog` SET  `Name`=?, `Sex`=?,`Address`=?,`Title`=?,`Content`=? WHERE  `Id`=?;")
	checkErr(err)

	res, err := stmt.Exec(blog.Name, blog.Sex, blog.Address, blog.Title, blog.Content, blog.Id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("affect", affect)

	db.Db.Close()

	return

}

///////////////////////////////////////////////////////////////Save////////////////////////////////////////

//////////////////////////Delete////////////////////////////////////////

func DelBlog(id int) {
	db := GetLink()
	//db, err := sql.Open("mysql", "root:/8520*963.@/blog?charset=utf8")
	//checkErr(err)

	stmt, err := db.Db.Prepare("delete from myblog  Where Id=? ")
	checkErr(err)

	res, err := stmt.Exec(id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("affect", affect)

	db.Db.Close()

	return

}

///////////////////////////Delete////////////////////////////////////////

/////////////////////////////Add////////////////////////////////////////////////////////
func AddBlog(blog Blog, id int) (blogs Blog) {
	db := GetLink()
	//db, err := sql.Open("mysql", "root:/8520*963.@/blog?charset=utf8")
	//checkErr(err)

	stmt, err := db.Db.Prepare("INSERT INTO `blog`.`myblog` SET `Name`=?, `Sex`=?,`Address`=?,`Title`=?,`Content`=?, `Id`=?;")
	checkErr(err)

	res, err := stmt.Exec(blog.Name, blog.Sex, blog.Address, blog.Title, blog.Content, blog.Id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("affect", affect)

	db.Db.Close()

	return

}

/////////////////////////////Add////////////////////////////////////////////////////////

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
