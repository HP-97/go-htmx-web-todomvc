package libsql

import (
	"database/sql"

	"github.com/HP-97/go-htmx-web-todomvc/internal/core/ports"
)

func NewTodoDBRepository(connStr string) (ports.TodoRepository, error) {
	db, err := sql.Open("sqlite", connStr);
}
