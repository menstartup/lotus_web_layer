package dataobject

type Type struct {
	Name         string `db:"name"`
	SWLayer      int    `db:"s_w_layer"`
	Description  string `db:"description"`
	DescWLink    string `db:"desc_w_link"`
	RelatedPages string `db:"related_pages"`
}
