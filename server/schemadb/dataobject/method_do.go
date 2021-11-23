package dataobject

type MethodDO struct {
	Name        string `db:"name"`
	Constructor string `db:"constructor"`
	WFlag       int8   `db:"w_flag"`
	Layer       int8   `db:"layer"`
	Flags       string `db:"flags"`
	Entities    string `db:"entities"`
	Errors      string `db:"errors"`
	Description string `db:"description"`
	DescWLink   string `db:"desc_w_link"`
	RelatedLink string `db:"related_link"`
	LinkText    string `db:"link_text"`
	Additional  string `db:"additional"`
	Result      string `db:"result"`
}
