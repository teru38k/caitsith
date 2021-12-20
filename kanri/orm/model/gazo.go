package model

import "gorm.io/gorm"

type Gazo struct {
	gorm.Model
	Prod_cd       string `gorm:"size:30" param:"prod_cd" query:"prod_cd" header:"prod_cd" form:"prod_cd" json:"prod_cd" xml:"prod_cd"`
	Customer_cd   string `gorm:"size:30" param:"customer_cd" query:"customer_cd" header:"customer_cd" form:"customer_cd" json:"customer_cd" xml:"customer_cd"`
	Customer_name string `gorm:"size:255" param:"customer_name" query:"customer_name" header:"customer_name" form:"customer_name" json:"customer_name" xml:"customer_name"`
	Comment       string `gorm:"size:500" param:"comment" query:"comment" header:"comment" form:"comment" json:"comment" xml:"comment"`
	Img_name      string `gorm:"size:255" param:"img_name" query:"img_name" header:"img_name" form:"img_name" json:"img_name" xml:"img_name"`
}
