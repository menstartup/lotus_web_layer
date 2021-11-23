package dao

// import (
// 	"context"
// 	"fmt"
// 	"gitlab/lotus_chat_web_layer/server/layerenv"
// 	"gitlab/lotus_chat_web_layer/server/schemadb/do"

// 	"github.com/jmoiron/sqlx"
// )

// type SchemaDAO struct {
// 	db *sqlx.DB
// }

// func NewSchemaDAO(db *sqlx.DB) *SchemaDAO {
// 	return &SchemaDAO{
// 		db: db,
// 	}
// }

// // ConstructorName string `db:"constructor_name"`
// //     ConstructorNum  string `db:"constructor_num"`
// //     Flag            int    `db:"flag"`
// //     Entities        string `db:"entities"`
// //     Type            string `db:"type"`
// //     Additional      string `db:"additional"`
// // func InsertOrUpdate
// func (d *SchemaDAO) InsertOrUpdate(do *do.SchemaDO) int64 {
// 	fmt.Println("do: ", do)
// 	myCtx, cancelFunc := context.WithTimeout(context.Background(), layerenv.SQLTimeOut)
// 	defer cancelFunc()

// 	var query = "insert into schemas (schema_name, constructor, schema_desc, schema_layer, schema_w_flag, flags, input_types, type, additionnal, related_page) values (:schema_name, :constructor, :schema_desc, :schema_layer, :schema_w_flag, :flags, :input_types, :type, :additionnal, related_page) on duplicate key update constructor = value(constructor)"
// 	r, err := d.db.NamedExecContext(myCtx, query, do)
// 	if err != nil {
// 		errDesc := fmt.Errorf("NamedExec in InsertOrUpdate(%v), error: %v", do, err)
// 		fmt.Println(errDesc)

// 		return 0
// 	}

// 	id, err := r.LastInsertId()
// 	if err != nil {
// 		errDesc := fmt.Errorf("LastInsertId in InsertOrUpdate(_), error: %v", err)
// 		fmt.Println(errDesc)

// 		return 0
// 	}

// 	return id
// }

// // func SelectBySchemaName()
// func (d *SchemaDAO) SelectBySchemaName(schemaName string) *do.SchemaDO {
// 	myCtx, cancelFunc := context.WithTimeout(context.Background(), layerenv.SQLTimeOut)
// 	defer cancelFunc()

// 	var query = "select * from schemas where constructor = ?"
// 	rows, err := d.db.QueryxContext(myCtx, query, schemaName)
// 	if err != nil {
// 		errDesc := fmt.Errorf("QueryxContext in SelectBySchemaName(%s), error: %v", schemaName, err)
// 		fmt.Println(errDesc)

// 		return nil
// 	}

// 	defer rows.Close()
// 	do := &do.SchemaDO{}
// 	if rows.Next() {
// 		err := rows.StructScan(do)
// 		if err != nil {
// 			errDesc := fmt.Errorf("rows SelectBySchemaName(_), error: %v", err)
// 			fmt.Println(errDesc)

// 			return nil
// 		}
// 	} else {
// 		return nil
// 	}

// 	err = rows.Err()
// 	if err != nil {
// 		errDesc := fmt.Errorf("rows.Err in SelectBySchemaName(_), error: %v", err)
// 		fmt.Println(errDesc)

// 		return nil
// 	}

// 	return do
// }

// func (d *SchemaDAO) SelectByType(schemaType string) []do.SchemaDO {
// 	myCtx, cancelFunc := context.WithTimeout(context.Background(), layerenv.SQLTimeOut)
// 	defer cancelFunc()

// 	var query = "select * from schemas where type = ?"
// 	rows, err := d.db.QueryxContext(myCtx, query, schemaType)
// 	if err != nil {
// 		errDesc := fmt.Errorf("QueryxContext in SelectByType(%s), error: %v", schemaType, err)
// 		fmt.Println(errDesc)

// 		return []do.SchemaDO{}
// 	}
// 	defer rows.Close()
// 	values := []do.SchemaDO{}
// 	for rows.Next() {
// 		v := do.SchemaDO{}

// 		err := rows.StructScan(&v)
// 		if err != nil {
// 			errDesc := fmt.Errorf("StructScan in SelectByType(_), error: %v", err)
// 			fmt.Println(errDesc)

// 			continue
// 		}
// 		values = append(values, v)
// 	}

// 	err = rows.Err()
// 	if err != nil {
// 		errDesc := fmt.Errorf("rows in SelectByType(_), error: %v", err)
// 		fmt.Println(errDesc)

// 		return []do.SchemaDO{}
// 	}

// 	return values
// }

// func (d *SchemaDAO) SelectAllSchema() []do.SchemaDO {
// 	myCtx, cancelFunc := context.WithTimeout(context.Background(), layerenv.SQLTimeOut)
// 	defer cancelFunc()

// 	var query = "select * from schemas"
// 	rows, err := d.db.QueryxContext(myCtx, query)
// 	if err != nil {
// 		errDesc := fmt.Errorf("QueryxContext in SelectAllSchema, error: %v", err)
// 		fmt.Println(errDesc)

// 		return nil
// 	}

// 	defer rows.Close()
// 	values := []do.SchemaDO{}
// 	for rows.Next() {
// 		v := do.SchemaDO{}

// 		err := rows.StructScan(&v)
// 		if err != nil {
// 			errDesc := fmt.Errorf("StructScan in SelectAllSchema(_), error: %v", err)
// 			fmt.Println(errDesc)

// 			continue
// 		}
// 		values = append(values, v)
// 	}

// 	err = rows.Err()
// 	if err != nil {
// 		errDesc := fmt.Errorf("rows in SelectAllSchema(_), error: %v", err)
// 		fmt.Println(errDesc)

// 		return []do.SchemaDO{}
// 	}

// 	return values
// }
