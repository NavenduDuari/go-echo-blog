package store

import (
	"database/sql"
	"time"

	"github.com/NavenduDuari/go-echo-blog/src/common/constant"
	"github.com/NavenduDuari/go-echo-blog/src/db/blog/model"
	"github.com/lib/pq"
)

func InsertBlog(db *sql.DB, newBlog *model.Blog) error {
	// _, err := db.Exec("INSERT INTO \""+constant.PostgressTableBlog+
	// 	"\" ( title, body, author, timestamp, categories, tags) VALUES($1, ROW($2, $3, $4)::body_component, $5, $6, $7, $8)",
	// 	newBlog.Title,
	// 	newBlog.Body.Title,
	// 	newBlog.Body.Content,
	// 	pq.Array(newBlog.Body.Images),
	// 	newBlog.Author,
	// 	time.Now(),
	// 	pq.Array(newBlog.Categories),
	// 	pq.Array(newBlog.Tags))

	_, err := db.Exec("INSERT INTO \""+constant.PostgressTableBlog+
		"\" ( title, body, author, timestamp, categories, tags) VALUES($1, $2, $3, $4, $5, $6)",
		newBlog.Title,
		newBlog.Body,
		newBlog.Author,
		time.Now(),
		pq.Array(newBlog.Categories),
		pq.Array(newBlog.Tags))

	if err != nil {
		return err
	}
	return nil
}

func FetchBlogs(db *sql.DB) ([]model.Blog, error) {
	var blogs []model.Blog

	rows, err := db.Query("Select * from \"" + constant.PostgressTableBlog + "\"")
	if err != nil || rows == nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var blog model.Blog

		err := rows.Scan(&blog.Id,
			&blog.Title,
			// (&blog.Body.Title, &blog.Body.Content, pq.Array(&blog.Body.Images)::model.BodyComponent,
			&blog.Body,
			&blog.Author,
			&blog.Timestamp,
			pq.Array(&blog.Categories),
			pq.Array(&blog.Tags))
		if err != nil {
			return blogs, err
		}
		blogs = append(blogs, blog)
	}
	return blogs, nil
}

func DeleteBlog(db *sql.DB, id string) error {
	_, err := db.Exec("DELETE from \""+constant.PostgressTableBlog+"\" WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateBlog(db *sql.DB, id string, updatedBlog model.Blog) error {
	_, err := db.Exec("UPDATE \""+constant.PostgressTableBlog+
		"\" SET title = $1, body = $2, author = $3, timestamp = $4, categories = $5, tags = $6 WHERE id = $7",
		updatedBlog.Title,
		updatedBlog.Body,
		updatedBlog.Author,
		time.Now(),
		pq.Array(updatedBlog.Categories),
		pq.Array(updatedBlog.Tags),
		id)
	if err != nil {
		return err
	}
	return nil
}
