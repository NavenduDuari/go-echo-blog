package main

import (
	"github.com/NavenduDuari/go-echo-blog/src/db/blog"
)

func main() {
	blog.InitStore()
	// blog.InsertBlog(&model.Blog{})
	// rows, _ := blog.FetchBlogs()
	// fmt.Println(rows)

	// blog.DeleteBlog("dc5ed9ca-41f7-11ea-8d0b-122d51db1d75")
	// rows, _ = blog.FetchBlogs()
	// fmt.Println(rows)
	blog.CloseStore()
}
