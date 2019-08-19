package main

import (
	"crypto/md5"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)
import _ "database/sql"
import _ "github.com/go-sql-driver/mysql"
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
type  lists struct {
	Blocks []*cloud_slide_show
}
type Result struct {
	Code    int    `json:"code"`
	Data    []*cloud_slide_show `json:"data"`
	Message string `json:"message"`
}
func initSql() *sql.DB  {
	host := "qa_cloudsports_wangliangliang:5qsv38hA3gyS0i3R@tcp(rm-2zeuj8ep5xv20r667mo.mysql.rds.aliyuncs.com:3306)/qa_maibu"
	db, err := sql.Open("mysql", host)
	if err!= nil{
		fmt.Println("数据库初始化失败！")
	}
	return db
}
func main() {
	//实现文件读取handler
	fileHanddler := http.FileServer(http.Dir("./video"))
	//注册文件服务
	http.Handle("/video/", http.StripPrefix("/video/", fileHanddler))
	//注册上传文件
	http.HandleFunc("/api/upload",uploadHander)
	//注册进servermux 将不同的url请求交给对应的handler
	http.HandleFunc("/say", sayHello)
	http.HandleFunc("/api/list", getList)
	http.HandleFunc("/v1/api-member/carousel-map", getImgList)
	//启动web服务
	http.ListenAndServe(":8080", nil)
}
func getImgList(w http.ResponseWriter, r *http.Request){
	setAllowOrigin(w,r)
	db := initSql()
	defer db.Close()
	row, err := db.Query("select * from cloud_slide_show order by place_number")
	if err != nil {
		fmt.Println("sql异常！")
	}
	nums :=&Result{}
	aaa  := []*cloud_slide_show{}
	for row.Next()  {
		data := &cloud_slide_show{}
		if err := row.Scan(&data.Id, &data.Name, &data.Pic, &data.create_time, &data.update_time, &data.create_id, &data.Place_number,&data.Url,&data.Status,&data.Type); err != nil {
			log.Println("Scan failed:", err)
		}
		fmt.Println(*nums)
		nums.Data = append(nums.Data,data)
		aaa = append(aaa,data)
	}
	retJson,_ := json.Marshal(aaa)
	w.Write(retJson)
	return
}


func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hellow word"))
}
func uploadHander(w http.ResponseWriter, r *http.Request) {
	//判断文件大小
	r.Body = http.MaxBytesReader(w, r.Body, 15*1024*1024)
	err := r.ParseMultipartForm(15 * 1024 * 1024)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//获取上传的文件
	file, fileHeander, err := r.FormFile("uploadFile")
	//检查文件类型
	ret := strings.HasSuffix(fileHeander.Filename, ".mp4")
	fmt.Print(ret)
	if ret == false {
		http.Error(w, "not mp4", http.StatusInternalServerError)
		return
	}
	//获取文件名称
	md5Byte := md5.Sum([]byte(fileHeander.Filename + time.Now().String() ))
	md5Str := fmt.Sprintf("%x", md5Byte)
	newFileName := md5Str + ".mp4"
	//写入文件
	dst, err := os.Create("./video/" + newFileName)
	defer dst.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()
	if _, err := io.Copy(dst, file); err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}
//获取视频列表
func getList(w http.ResponseWriter, r *http.Request)  {
	setAllowOrigin(w,r)
	files ,_ := filepath.Glob("video/*")
	var  ret [] string
	for _,file := range files{
		ret = append(ret,"http://"+r.Host+"/video/"+filepath.Base(file))
	}
	retJson,_ := json.Marshal(ret)
	w.Write(retJson)
	return
}
func setAllowOrigin(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	return
}


