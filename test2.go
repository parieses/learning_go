package main

import (
	"fmt"
)

type  Phone interface {
	Call()
}
type IPhone struct {
}
func (iPhone IPhone) Call() {
	fmt.Println("I am iPhone, I can call you!")
}

func main()  {
	var phone Phone
	phone = new(IPhone)
	phone.Call()
	//数组 固定长度数组无法append
	var balances [10] float32
	var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	balances[0] =255.2
	// append works on nil slices.
	//balances = append(balances,20.1)
	fmt.Println(balances)
	fmt.Println(balance)
	var a int= 20   /* 声明实际变量 */
	var ip *int        /* 声明指针变量 */

	ip = &a  /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a  )

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip )

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip )
	/**
	a 变量的地址是: 20818a220
	ip 变量储存的指针地址: 20818a220
	ip 变量的值: 20
	 */
	//切片
	//var yinzhengjie []types //类型
	var yinzhengjie []IPhone
	var lenS int =2
	var slice1 []int = make([]int, lenS)
	slice1 = append(slice1,1)
	fmt.Println(yinzhengjie)
	fmt.Println(slice1)
	//MAP
	var maps map[string]string
	println(maps)
	var countryCapitalMap  = make(map[string]string )/*创建集合 */
	countryCapitalMap["dawd "] ="dwd"
	fmt.Println(countryCapitalMap)
	countryCapitalMaps := make(map[int]string)
	i :=0
	var aa = []string{"a","b","c","d"}
	for _,item := range aa {
		countryCapitalMaps[i] = item
		i++
	}

	fmt.Println(countryCapitalMap)
	fmt.Println(countryCapitalMaps)

}