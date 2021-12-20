package posql

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Gazo struct {
	gorm.Model
	Prod_cd       string `gorm:"size:30"`
	Customer_cd   string `gorm:"size:30"`
	Customer_name string `gorm:"size:255"`
	Memo          string `gorm:"size:500"`
	Gazo_path     string `gorm:"size:255"`
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func DBConn() *gorm.DB {
	dsn := "host=localhost user=root password=root dbname=db1 port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	checkError(err)
	return db
}

func DropTable() {
	// pgadmin user=admin@admin.jp pass=admin
	//db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=root password=root dbname=db1 sslmode=disable")
	db := DBConn()
	defer db.Close()

	_, err = db.Exec("DROP TABLE gazo")
	checkError(err)
	fmt.Println("Finished droped gazo table")

}

func CreateTable() {
	// pgadmin user=admin@admin.jp pass=admin
	
	db := DBConn()
	defer db.Close()

	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.AutoMigrate(&Gazo{})
	checkError(err)
	fmt.Println("Finished creating gazo table")

}

func InsertGazoTable(prod_cd string, customer_cd string, customer_name string, memo string, gazo_path string) {
	dsn := "host=localhost user=root password=root dbname=db1 port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	defer db.Close()
	checkError(err)

	_, err = db. exec
}
