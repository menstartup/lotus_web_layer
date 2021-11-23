package dao

import (
	"context"
	"fmt"
	"gitlab/lotus_chat_web_layer/server/layerenv"
	"gitlab/lotus_chat_web_layer/server/schemadb/dataobject"

	"github.com/jmoiron/sqlx"
)

type TypeDAO struct {
	db *sqlx.DB
}

func NewTypeDAO(db *sqlx.DB) *TypeDAO {
	return &TypeDAO{
		db,
	}
}

// func InsertOrUpdate
func (d *TypeDAO) InsertOrUpdate(do *dataobject.Type) int64 {
	myCtx, cancelFunc := context.WithTimeout(context.Background(), layerenv.SQLTimeOut)
	defer cancelFunc()

	var query = "insert into types (name, s_w_layer, description, desc_w_link, related_pages) values (:name, :s_w_layer, :description, :desc_w_link, :related_pages) on duplicate key update name = values(name), s_w_layer = values(s_w_layer), description = values(description), desc_w_link = values(desc_w_link), related_pages = values(related_pages)"
	r, err := d.db.NamedExecContext(myCtx, query, do)
	if err != nil {
		errDesc := fmt.Errorf("NamedExec in InsertOrUpdate(%v), error: %v", do, err)
		fmt.Println(errDesc)

		return 0
	}

	id, err := r.LastInsertId()
	if err != nil {
		errDesc := fmt.Errorf("LastInsertId in InsertFile(_), error: %v", err)
		fmt.Println(errDesc)

		return 0
	}

	return id
}

func (d *TypeDAO) SelectSingleType(name string) *dataobject.Type {
	myCtx, cancelFunc := context.WithTimeout(context.Background(), layerenv.SQLTimeOut)
	defer cancelFunc()

	var query = "select name, s_w_layer, description, desc_w_link, related_pages from types where name = ?"
	rows, err := d.db.QueryxContext(myCtx, query, name)
	if err != nil {
		errDesc := fmt.Errorf("Queryx in SelectSingleType(_), error: %v", err)
		fmt.Println(errDesc)

		return nil
	}

	defer rows.Close()
	do := &dataobject.Type{}
	if rows.Next() {
		err := rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Errorf("StructScan in SelectSingleType(_), error: %v", err)
			fmt.Println(errDesc)

			return nil
		}
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Errorf("rows in SelectSingleType(_), error: %v", err)
		fmt.Println(errDesc)

		return nil
	}

	return do
}

func (d *TypeDAO) DeleeteType(name string) int64 {
	myCtx, cancelFunc := context.WithTimeout(context.Background(), layerenv.SQLTimeOut)
	defer cancelFunc()

	var query = "delete from types where name = ?"
	r, err := d.db.ExecContext(myCtx, query, name)

	if err != nil {
		errDesc := fmt.Errorf("Exec in Delete(_), error: %v", err)
		fmt.Println(errDesc)

		return 0
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Errorf("RowsAffected in Delete(_), error: %v", err)
		fmt.Println(errDesc)

		return 0
	}

	return rows
}
