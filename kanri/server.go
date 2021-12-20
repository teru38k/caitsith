package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"kanri/gazo"
	"kanri/orm"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gocli/exitcode"
)

/*
func hontonoMain() {
	e := echo.New()
	e.POST("/gazo", createGazo)

}

func dropTable() {
	posql.DropTable()
}

func createTable() {
	posql.CreateTable()
}


func createGazo(c echo.Context) error {

	//prod_cd := c.FormValue("prod_cd")
	//customer_cd := c.FormValue("customer_cd")
	//customer_name := c.FormValue("customer_name")
	//memo := c.FormValue("memo")
	//gazo_path := c.FormValue("gazo_path")

	gazo := Gazo{
		Prod_cd:       "1",
		Customer_cd:   "00001",
		Customer_name: "test",
		Memo:          "memo",
		Gazo_path:     "/gazo_path",
	}

	return c.JSON(http.StatusOK, gazo)
}
*/

func createGazo(c echo.Context) error {
	prod_cd := c.FormValue("prod_cd")
	customer_cd := c.FormValue("customer_cd")
	customer_name := c.FormValue("customer_name")
	comment := c.FormValue("comment")
	img_name := c.FormValue("img_name")

	gazo.CreateGazo(prod_cd, customer_cd, customer_name, comment, img_name)

	return c.HTML(200, "created record successfully")
}

/*
func readGazo() {
	data, _ := gazo.ReadGazo()
	//bytes, _ := json.Marshal(&data)
	//ttt := string(bytes)
	fmt.Println(data[0]["prod_cd"])
	j_data, _ := json.Marshal(data)

	fmt.Println(string(j_data))

	//fmt.Println("data:", data)

}

*/
func findAllGazo(c echo.Context) error {
	id := ""
	prod_cd := ""
	customer_cd := ""
	customer_name := ""
	data, _ := gazo.FindGazo(id, prod_cd, customer_cd, customer_name)
	if err := c.Bind(data); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, data)
}

func findGazo(c echo.Context) error {
	id := c.FormValue("id")
	prod_cd := c.FormValue("prod_cd")
	customer_cd := c.FormValue("customer_cd")
	customer_name := c.FormValue("customer_name")

	fmt.Println("id: ", id, "prod_cd: ", prod_cd, "/ customer_cd: ", customer_cd, "/ customer_name: ", customer_name)

	var data []map[string]interface{}
	data, _ = gazo.FindGazo(id, prod_cd, customer_cd, customer_name)

	return c.JSON(http.StatusOK, data)
}

func updateGazo(c echo.Context) error {
	id := c.FormValue("id")
	prod_cd := c.FormValue("prod_cd")
	customer_cd := c.FormValue("customer_cd")
	customer_name := c.FormValue("customer_name")
	comment := c.FormValue("comment")
	img_name := c.FormValue("img_name")

	gazo.UpdateGazo(id, prod_cd, customer_cd, customer_name, comment, img_name)

	return c.HTML(200, "updated record successfully")

}

func upload(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	fileModel := strings.Split(file.Filename, ".")
	fileName := fileModel[0]
	extension := fileModel[1]
	dst, err := os.Create(fmt.Sprintf("/Users/teru/Docker-Compose/vue/frontend/src/assets/image/%s.%s", fileName, extension))
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(200, "Image has been saved.")
}

func deleteGazo(c echo.Context) error {
	id := c.FormValue("id")

	gazo.DeleteGazo(id)

	return c.HTML(200, "deleted:"+id)
}

func getimages(c echo.Context) error {
	var ls []string
	files, err := ioutil.ReadDir("./image")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		ls = append(ls, file.Name())
	}

	return c.JSON(http.StatusOK, ls)
}

func Run() exitcode.ExitCode {
	// create gorm.DB instance for PostgreSQL service
	gormCtx, err := orm.NewGORM()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitcode.Abnormal
	}
	defer gormCtx.Close()

	// query
	var results []map[string]interface{}
	tx := gormCtx.GetDb().Table("tablename").Find(&results) // "tablename" is not exist
	if tx.Error != nil {
		gormCtx.GetLogger().Error().Interface("error", errs.Wrap(tx.Error)).Send()
		return exitcode.Abnormal
	}

	return exitcode.Normal
}

func main() {
	//Run().Exit()
	//createGazo()
	//readGazo()

	e := echo.New()
	e.Use((middleware.CORS()))

	e.POST("/gazo", findGazo)
	e.POST("/upload", upload)
	e.POST("/create_gazo", createGazo)
	e.POST("/update_gazo", updateGazo)
	e.POST("/delete_gazo", deleteGazo)
	e.POST("/getimages", getimages)

	e.Logger.Fatal(e.Start(":1323"))
}
