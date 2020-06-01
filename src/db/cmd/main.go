package main

import (
	"fmt"

	"gitlab.com/yoursown/gocb/src/yoblog"
	yoblogmodel "gitlab.com/yoursown/gocb/src/yoblog/model"
)

func main() {
	yoblog.InitStore()
	yoblog.InsertBlog(&yoblogmodel.Blog{})
	rows, _ := yoblog.FetchBlogs()
	fmt.Println(rows)

	yoblog.DeleteBlog("dc5ed9ca-41f7-11ea-8d0b-122d51db1d75")
	rows, _ = yoblog.FetchBlogs()
	fmt.Println(rows)
	yoblog.CloseStore()
}
