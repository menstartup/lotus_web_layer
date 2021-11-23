package main

import (
	"fmt"
	"gitlab/lotus_chat_web_layer/server/schemadb/connection"
	"gitlab/lotus_chat_web_layer/server/schemadb/dao"
	"gitlab/lotus_chat_web_layer/server/schemadb/dataobject"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB = connection.NewConnectionToDB()
	// SchemaDAO   *mysql_dao.SchemaDAO   = mysql_dao.NewSchemaDAO(db)
	typeDAO   *dao.TypeDAO   = dao.NewTypeDAO(db)
	methodDAO *dao.MethodDAO = dao.NewMethodDAO(db)
)

type Data struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func main() {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		if c.Request.Method == http.MethodOptions {
			c.Status(http.StatusOK)
			return
		}

		c.Next()
	})
	router.Use(gin.Recovery())
	defer connection.CloseConnection(db)

	// typeData := &dataobject.Type{
	// 	Name:         "InputFile1",
	// 	SWLayer:      4,
	// 	Description:  "Defines a file uploaded by the client.",
	// 	DescWLink:    "https://core.telegram.org/method/upload.saveFilePart",
	// 	RelatedPages: "https://core.telegram.org/constructor/messages.exportedChatInvites",
	// }

	// id := TypeDAO.InsertOrUpdate(typeData)
	// fmt.Println(id)

	// router.POST("/createSchema", func(c *gin.Context) {
	// 	numLayer, err := strconv.Atoi(c.PostForm("schema_layer"))
	// 	if err != nil {
	// 		fmt.Printf("cannot convert string to int with err: %v \n", err)
	// 	}
	// 	schemaData := &layerdataobject.SchemaDO{
	// 		SchemaName:  c.PostForm("schema_name"),
	// 		Constructor: c.PostForm("constructor"),
	// 		SchemaDesc:  c.PostForm("schema_desc"),
	// 		SchemaLayer: int32(numLayer),
	// 		Type:        c.PostForm("type"),
	// 		RelatedPage: c.PostForm("related_page"),
	// 	}
	// 	if c.PostForm("schema_w_flag") == "false" {
	// 		schemaData.SchemaWFlag = 0
	// 	} else {
	// 		schemaData.SchemaWFlag = 1
	// 	}
	// 	if c.PostForm("flags") != "[]" {
	// 		schemaData.Flags = c.PostForm("flags")
	// 	}
	// 	if c.PostForm("input_types") != "[]" {
	// 		schemaData.InputTypes = c.PostForm("input_types")
	// 	}
	// 	id := SchemaDAO.InsertOrUpdate(schemaData)
	// 	fmt.Println("schemaData: ", schemaData)
	// 	fmt.Println("id inserted: ", id)

	// 	c.String(http.StatusOK, "Hello %s", "name")
	// })

	// router.POST("/createFunction", func(c *gin.Context) {
	// 	numLayer, err := strconv.Atoi(c.PostForm("layer"))
	// 	if err != nil {
	// 		fmt.Printf("cannot convert string to int with err: %v \n", err)
	// 	}
	// 	dataFunc := &layerdataobject.FunctionDO{
	// 		FunctionName:        c.PostForm("function_name"),
	// 		FunctionConstructor: c.PostForm("function_constructor"),
	// 		Layer:               int8(numLayer),
	// 		Type:                c.PostForm("output_type"),
	// 	}
	// 	if c.PostForm("function_flag") == "false" {
	// 		dataFunc.Flag = 0
	// 	} else {
	// 		dataFunc.Flag = 1
	// 	}
	// 	if c.PostForm("function_flags") != "[]" {
	// 		dataFunc.Flags = c.PostForm("function_flags")
	// 	}
	// 	if c.PostForm("input_types") != "[]" {
	// 		dataFunc.InputTypes = c.PostForm("input_types")
	// 	}

	// router.GET("/getAll", func(c *gin.Context) {
	// 	values := SchemaDAO.SelectAllSchema()
	// 	fmt.Println(values)
	// })
	// t := router.Group("/type")
	// {
	router.POST("/typecreate", func(c *gin.Context) {
		sWLayer, _ := strconv.Atoi(c.PostForm("sWLayer"))
		typeDO := &dataobject.Type{
			Name:         c.PostForm("name"),
			SWLayer:      sWLayer,
			Description:  c.PostForm("description"),
			DescWLink:    c.PostForm("descWLink"),
			RelatedPages: c.PostForm("relatedPages"),
		}
		fmt.Println("type_do: ", typeDO)
		id := typeDAO.InsertOrUpdate(typeDO)
		fmt.Println(id)
	})

	router.GET("/typesingle", func(c *gin.Context) {
		name := c.Query("name")
		do := typeDAO.SelectSingleType(name)
		fmt.Println("select type by name: ", do)
		if do == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "bad request",
			})

			return
		}
		c.JSON(http.StatusOK, do)
	})

	router.DELETE("/typedelete/:name", func(c *gin.Context) {
		name := c.Param("name")
		typeDO := typeDAO.SelectSingleType(name)
		if typeDO == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"msg":    "Invalid name type!",
			})

			return
		}
		id := typeDAO.DeleeteType(name)
		fmt.Println("id delete type: ", id)
		if id == 0 {
			c.JSON(http.StatusForbidden, gin.H{
				"status": http.StatusForbidden,
				"msg":    "Invalid Error!",
			})

			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"msg":    fmt.Sprintf("Deleted success type: %s", name),
		})
	})

	// formData.append("name", values.name)
	// formData.append("constructor", values.constructorVal)
	// formData.append("wFlag", values.wFlag ? values.WFlag : false)
	// formData.append("layer", values.layer)
	// formData.append("description", values.desc ? values.desc : "")
	// formData.append("desc_w_link", values.descWLink ? values.descWLink : "")
	// formData.append("additional", values.additional ? values.additional : "")
	// formData.append("relatedLink", values.linkRPages ? values.linkRPages : "")
	// formData.append("textRPages", values.textRPages ? values.textRPages : "")

	// formData.append("flags", values.flagsMethod ? JSON.stringify(values.flagsMethod) : "")
	// formData.append("entities", values.entitiesMethod ? JSON.stringify(values.entitiesMethod) : "")
	// formData.append("entities", values.errorsMethod ? JSON.stringify(values.errorsMethod) : "")
	// formData.append("result", values.result ? values.result : "")
	router.POST("/methodCreate", func(c *gin.Context) {
		fmt.Println("wFlag: , type: ", c.PostForm("wFlag"), reflect.TypeOf(c.PostForm("wFlag")))
		c.MultipartForm()
		for key, value := range c.Request.PostForm {
			fmt.Printf("%v = %v \n", key, value)
		}
		// check alreadt exist method
		
		methodLayer, _ := strconv.Atoi(c.PostForm("layer"))
		dataInsert := &dataobject.MethodDO{
			Name:        c.PostForm("name"),
			Constructor: c.PostForm("constructor"),
			Layer:       int8(methodLayer),
			Description: c.PostForm("description"),
			DescWLink:   c.PostForm("desc_w_link"),
			Additional:  c.PostForm("additional"),
			RelatedLink: c.PostForm("relatedLink"),
			LinkText:    c.PostForm("textRPages"),
			Flags:       c.PostForm("flags"),
			Entities:    c.PostForm("entities"),
			Errors:      c.PostForm("errors"),
		}
		if c.PostForm("wFlag") == "false" {
			dataInsert.WFlag = 0
		} else {
			dataInsert.WFlag = 1
		}
		methodDAO.InsertOrUpdate(dataInsert)
		fmt.Println("data_insert: ", dataInsert)
	})
	router.Run(":8079")
}
