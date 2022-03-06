package interface_learnning

import "fmt"

//
//  DataWriter
//  @Description: 定义一个数据写入器
//
type DataWriter interface {
	// 方法签名
	WriteData(data interface{}) error
}

//  定义文件结构
type file struct {
}

//
//  WriteData
//  @Description: 实现DataWriter接口的WriteData方法
//  @receiver d
//  @param data
//  @return error
//
func (d *file) WriteData(data interface{}) error {
	fmt.Println("WriteData:", data)
	return nil
}

//
//  Code1
//  @Description: 主方法
//
func Code1() {
	f := new(file)

	var writer DataWriter

	writer = f

	err := writer.WriteData("data")
	if err != nil {
		fmt.Printf("writeData error: %v\n", err)
	}
}
