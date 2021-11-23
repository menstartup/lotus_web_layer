package dao

import (
	"context"
	"fmt"
	"gitlab/lotus_chat_web_layer/server/layerenv"
	"gitlab/lotus_chat_web_layer/server/schemadb/dataobject"

	"github.com/jmoiron/sqlx"
)

type MethodDAO struct {
	db *sqlx.DB
}

func NewMethodDAO(db *sqlx.DB) *MethodDAO {
	return &MethodDAO{
		db,
	}
}

func (d *MethodDAO) InsertOrUpdate(do *dataobject.MethodDO) int64 {
	myCtx, cancelFunc := context.WithTimeout(context.Background(), layerenv.SQLTimeOut)
	defer cancelFunc()

	var query = "insert into methods (name, constructor, w_flag, flags, entities, errors, description, desc_w_link, related_link, link_text, additional, result, layer) values (:name, :constructor, :w_flag, :flags, :entities, :errors, :description, :desc_w_link, :related_link, :link_text, :additional, :result, :layer)"
	r, err := d.db.NamedExecContext(myCtx, query, do)
	if err != nil {
		errDesc := fmt.Errorf("NamedExec in InsertFunction(%v), error: %v", do, err)
		fmt.Println(errDesc)

		return 0
	}
	id, err := r.LastInsertId()
	if err != nil {
		errDesc := fmt.Errorf("LastInsertId in InsertFunction(_), error: %v", err)
		fmt.Println(errDesc)

		return 0
	}

	return id
}

func (d *MethodDAO) SelectByName(name string) *dataobject.MethodDO {
	myCtx, cancelFunc := context.WithTimeout(context.Background(), layerenv.SQLTimeOut)
	defer cancelFunc()

	var query = "select * from methods where name = ?"
	rows, err := d.db.QueryxContext(myCtx, query, name)
	if err != nil {
		errDesc := fmt.Errorf("QueryxContext in SelectByName(%s), error: %v", name, err)
		fmt.Println(errDesc)

		return nil
	}

	defer rows.Close()
	do := &dataobject.MethodDO{}
	if rows.Next() {
		err := rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Errorf("rows SelectByName(_), error: %v", err)
			fmt.Println(errDesc)

			return nil
		}
	} else {
		return nil
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Errorf("rows.Err in SelectByName(_), error: %v", err)
		fmt.Println(errDesc)

		return nil
	}

	return do
}

func (d *MethodDAO) SelectAllMethod() []dataobject.MethodDO {
	myCtx, cancelFunc := context.WithTimeout(context.Background(), layerenv.SQLTimeOut)
	defer cancelFunc()

	var query = "select * from methods"
	rows, err := d.db.QueryxContext(myCtx, query)
	if err != nil {
		errDesc := fmt.Errorf("QueryxContext in SelectAllMethod, error: %v", err)
		fmt.Println(errDesc)

		return nil
	}

	defer rows.Close()
	values := []dataobject.MethodDO{}
	for rows.Next() {
		v := dataobject.MethodDO{}

		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Errorf("StructScan in SelectAllMethod(_), error: %v", err)
			fmt.Println(errDesc)

			continue
		}
		values = append(values, v)
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Errorf("rows in SelectAllMethod(_), error: %v", err)
		fmt.Println(errDesc)

		return []dataobject.MethodDO{}
	}

	return values
}

func (d *MethodDAO) DeleteMethod(name string) int64 {
	myCtx, cancelFunc := context.WithTimeout(context.Background(), layerenv.SQLTimeOut)
	defer cancelFunc()

	var query = "delete from methods where name = ?"
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
