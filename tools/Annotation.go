package tools

import (
	"github.com/astaxie/beego"
)

/**
 ***********************************************************
 **作者：zwx<lockfeel1215@gmail.com>
 **---------------------------------------------------------
 **时间：2017-09-09
 **---------------------------------------------------------
 **描述：GO语言注释样例
 ***********************************************************
 */
type ObjectDemo struct {
	beego.Controller
}

/**
 ************************************************************
 **方法::ObjectDemo::Post
 **----------------------------------------------------------
 **描述::无参数无返回值方法注释
 **----------------------------------------------------------
 **参数:
 **param:in--    无
 **----------------------------------------------------------
 **返回：
 **return:out--  无
 **----------------------------------------------------------
 **日期:2017.09.09  Add by zwx
 ************************************************************
 */
func (objectDemo *ObjectDemo) Post() {
}

/**
 ************************************************************
 **方法::ObjectDemo::maxNum
 **----------------------------------------------------------
 **描述::私有方法注释
 **----------------------------------------------------------
 **参数:
 **param1:in-- Int : num1 :: 数字
 **param2:in-- Int : num2 :: 数字
 **----------------------------------------------------------
 **返回：
 **return:out-- Int : result :: 最终结果
 **----------------------------------------------------------
 **日期:2017.09.09  Add by zwx
 ************************************************************
 */
func (objectDemo *ObjectDemo) maxNum (num1,num2 int) int {
	return num2 - num1
}

/**
 ************************************************************
 **方法::ObjectDemo::FindMuse
 **----------------------------------------------------------
 **描述::多值方法注释
 **----------------------------------------------------------
 **参数:
 **param1:in-- String : str :: 字符串
 **param2:in-- Int : age    :: 年龄
 **----------------------------------------------------------
 **返回：
 **return1:out-- Int : age    :: 最终结果
 **return2:out-- String : str :: 最终结果
 **----------------------------------------------------------
 **示例：
 **case1:use-- : FindMuse("Mr.z",1) :: 不同参数不同使用
 **case2:use-- : FindMuse("Mr.z",2) :: 不同参数不同使用
 **----------------------------------------------------------
 **日期:2017.09.09  Add by zwx
 ************************************************************
 */
func (objectDemo *ObjectDemo) FindMuse (str string,age int) (int,string){
	return age,str
}