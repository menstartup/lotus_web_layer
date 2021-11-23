package dao

import (
	"context"
	"fmt"
	"gitlab/lotus_chat_web_layer/server/layerenv"
	"gitlab/lotus_chat_web_layer/server/schemadb/dataobject"

	"github.com/jmoiron/sqlx"
)

type LayerDAO struct {
	db *sqlx.DB
}

func NewLayerDAO(db *sqlx.DB) *LayerDAO {
	return &LayerDAO{
		db,
	}
}

// NumberLayer          int8   `db:"number_layer"`
// ConstructorName      string `db:"constructor_name"`
// FunctionName         string `db:"function_name"`
// LinkLayerDescription string `db:"link_layer_description"`

// func InsertOrUpdate()
func (dao *LayerDAO) InsertOrUpdate(do *dataobject.LayerDO) int64 {
	myCtx, cancelFunc := context.WithTimeout(context.Background(), layerenv.SQLTimeOut)
	defer cancelFunc()

	var query = "insert into layers (number_layer, constructor_name, function_name, link_layer_description) values (:number_layer, :constructor_name, :function_name, :link_layer_description) on duplicate key update number_layer = values(number_layer)"
	r, err := dao.db.NamedExecContext(myCtx, query, do)
	if err != nil {
		errDesc := fmt.Errorf("NamedExec in InsertOrUpdate(%v), error: %v", do, err)
		fmt.Println(errDesc)
		// sendLogToTelegram
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

// func SelectByLayer()
func (dao *LayerDAO) SelectByLayer(numLayer int8) *dataobject.LayerDO {
	myCtx, cancelFunc := context.WithTimeout(context.Background(), layerenv.SQLTimeOut)
	defer cancelFunc()

	var query = "select * from layers where number_layer = ?"
	rows, err := dao.db.QueryxContext(myCtx, query, numLayer)
	if err != nil {
		errDesc := fmt.Errorf("ueryx in SelectByLayer(_), error: %v", err)
		fmt.Println(errDesc)

		return nil
	}

	defer rows.Close()
	do := &dataobject.LayerDO{}
	if rows.Next() {
		err := rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Errorf("StructScan in SelectByLayer(_), error: %v", err)
			fmt.Println(errDesc)
			return nil
		}
	} else {
		return nil
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Errorf("rows in SelectByLayer(_), error: %v", err)
		fmt.Println(errDesc)

		return nil
	}

	return do
}
