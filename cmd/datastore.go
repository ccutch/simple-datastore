package main

import (
	"fmt"

	simpledb "../"
)

type User struct {
	Id   int    `simpledb:"id"`
	Name string `simpledb:"name"`
}

type Post struct {
	Title   string `simpledb:"title"`
	Content string `simpledb:"content"`
	Owner   User   `simpledb:"owner" relatesTo:"users"`
}

func main() {

	db := simpledb.Open("test.simple")
	defer db.MustDump()

	userSchema := simpledb.Schemafy(User{})
	postSchema := simpledb.Schemafy(Post{})

	usersTable := db.GetTable("users")
	postsTable := db.GetTable("posts")

	if usersTable == nil {
		usersTable, _ = db.DefineTable("users", userSchema)
	}

	if postsTable == nil {
		postsTable, _ = db.DefineTable("posts", postSchema)
	}

	user := User{
		Id:   1,
		Name: "Connor",
	}

	post := Post{
		Title:   "First Post",
		Content: "This is the body of the post",
		Owner:   user,
	}

	userNode := simpledb.Nodify(user)
	postNode := simpledb.Nodify(post)

	usersTable.Insert(userNode)
	postsTable.Insert(postNode)

	fmt.Println(db.Tables())

	// todo:
	// userMirror := User{}
	// postMirror := Post{}
	//
	// usersTable.FindWhere(&userMirror, "id", "1")
	// postsTable.FindWhere(&postMirror, "title", "First Post")
	//
	// fmt.Println(userMirror) // should be the same as user
	// fmt.Println(postMirror) // should be the same as post
}
