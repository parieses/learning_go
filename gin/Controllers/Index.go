package Controllers

import (
	"fmt"
	"gin/Models"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type cloud_slide_show struct {
	Id int64
	Name string
	Pic string
	create_time int
	update_time int
	create_id int
	Place_number int
	Url string
	Status int
	Type int
}



func WebRoot(context *gin.Context) {
	var c = Models.Get()
	c = c.GetConf()
	fmt.Println(c.Host)
	//context.String(http.StatusOK, c)
}
func LoginEndpoint(context *gin.Context)  {
	db ,err:= initSql()
	if err!= nil{
		context.JSON(200,"初始化失败")
		return
	}
	page := context.Param("page")

	int,err:=strconv.Atoi(page)
	if err != nil {
		int =0
	}
	count :=1
	defer db.Close()

	var  str  = "select * from cloud_slide_show order by place_number LIMIT %d,%d"
	str = fmt.Sprintf(str,count*int,count)
	row, err := db.Query(str)


	if err != nil {
		fmt.Println("sql异常！")
	}
	aaa  := []*cloud_slide_show{}
	for row.Next()  {
		data := &cloud_slide_show{}
		if err := row.Scan(&data.Id, &data.Name, &data.Pic, &data.create_time, &data.update_time, &data.create_id, &data.Place_number,&data.Url,&data.Status,&data.Type); err != nil {
			log.Println("Scan failed:", err)
		}
		aaa = append(aaa,data)
	}

		dd := SetData{Item:aaa,}
	//var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	Json := Json{
		Code:200,
		Message:"获取成功",
		Data:dd,
	}
	context.JSON(200,Json)
}
