package dataobject

type LayerDO struct {
	NumberLayer          int8   `db:"number_layer"`
	ConstructorName      string `db:"constructor_name"`
	FunctionName         string `db:"function_name"`
	LinkLayerDescription string `db:"link_layer_description"`
}
