package resource

import (
	"database/sql"
	"log"

	"gorm.io/gorm"
)

type Repository struct {
	sql  *sql.DB
	gorm *gorm.DB
}

func (r *Repository) CreateWithMySLQ(comment Comment) {
	q := "INSERT IGNORE INTO `comments` (id, post_id, name, email, body) VALUES (?, ?, ?, ?, ?);"
	insert, err := r.sql.Prepare(q)
	if err != nil {
		log.Println(err)
	}

	insert.Exec(comment.Id, comment.PostId, comment.Name, comment.Email, comment.Body)
	insert.Close()
}

func (r *Repository) CreateWithGORM(comment Comment) {
	result := r.gorm.Create(&comment)
	if result.Error != nil {
		log.Println(result.Error)
	}
}

func NewRepository(sql *sql.DB, gorm *gorm.DB) *Repository {
	return &Repository{
		sql:  sql,
		gorm: gorm,
	}
}
