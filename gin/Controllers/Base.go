package Controllers
import (
	"database/sql"
	_ "database/sql"
)
import _ "github.com/go-sql-driver/mysql"
type Json struct {
	// contains filtered or unexported fields
	Code int
	Message string
	Data interface{}
}
type Paging struct {
	Page      int64   //当前页
	Pagesize  int64   //每页条数
	Total     int64   //总条数
	PageCount int64   //总页数
	Nums      []int64 //分页序数
	NumsCount int64   //总页序数
}
type SetData struct {
	Item interface{}
	Paging Paging
}
func initSql() (*sql.DB ,error){
	host := "qa_cloudsports_wangliangliang:5qsv38hA3gyS0i3R@tcp(rm-2zeuj8ep5xv20r667mo.mysql.rds.aliyuncs.com:3306)/qa_maibu"
	db, err := sql.Open("mysql", host)
	if err!= nil{
		return db,err
	}
	return db ,nil
}