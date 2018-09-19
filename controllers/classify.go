package controllers

import (
	"github.com/astaxie/beego"
	"xiaomatv.cn/services"
)

/**
 ***********************************************************
 **作者：zwx<lockfeel1215@gmail.com>
 **---------------------------------------------------------
 **时间：2018-08-15
 **---------------------------------------------------------
 **描述：分类类型
 ***********************************************************
 */
 type Classify struct {
	 /*继承beego*/
	 beego.Controller
	 Serv services.ArticleService `inject:""`
 }

/**
 ************************************************************
 **方法::Classify::GetLists
 **----------------------------------------------------------
 **描述::获取分类数据
 **----------------------------------------------------------
 **参数:
 **param1:in-- Int : user_id    :: 用户ID
 **param2:in-- Int : page_num   :: 分页数量
 **param3:in-- Int : page_size  :: 分页大小
 **param4:in-- Int : type		:: 请求类型
 **----------------------------------------------------------
 **返回：
 **return:out--  JSON ：data  :: JSON对象
 **----------------------------------------------------------
 **日期:2018.08.15  Add by zwx
 ************************************************************
 */
 func (a *Classify) GetLists (){

 }

/**
 ************************************************************
 **方法::Classify::GetVideoLists
 **----------------------------------------------------------
 **描述::获取视频分类
 **----------------------------------------------------------
 **参数:
 **param1:in-- Int : user_id    :: 用户ID
 **param2:in-- Int : page_num   :: 分页数量
 **param3:in-- Int : page_size  :: 分页大小
 **param4:in-- Int : type		:: 视频分类
 **----------------------------------------------------------
 **返回：
 **return:out--  JSON ：data  :: JSON对象
 **----------------------------------------------------------
 **日期:2018.08.15  Add by zwx
 ************************************************************
 */
 func (a *Classify) GetVideoLists() {

 }

/**
 ************************************************************
 **方法::Classify::GetUserClassify
 **----------------------------------------------------------
 **描述::获取用户分类
 **----------------------------------------------------------
 **参数:
 **param1:in-- Int : user_id    :: 用户ID
 **----------------------------------------------------------
 **返回：
 **return:out--  JSON ：data  :: JSON对象
 **----------------------------------------------------------
 **日期:2018.08.15  Add by zwx
 ************************************************************
 */
 func (a *Classify) GetUserClassify() {

 }

/**
 ************************************************************
 **方法::ClassifyController::GetVideoClassify
 **----------------------------------------------------------
 **描述::获取视频
 **----------------------------------------------------------
 **参数:
 **param1:in-- Int : user_id    :: 用户ID
 **----------------------------------------------------------
 **返回：
 **return:out--  JSON ：data  :: JSON对象
 **----------------------------------------------------------
 **日期:2018.08.15  Add by zwx
 ************************************************************
 */
 func (a *Classify) GetVideoClassify() {

 }

/**
 ************************************************************
 **方法::Classify::GetInterest
 **----------------------------------------------------------
 **描述::获取视频兴趣
 **----------------------------------------------------------
 **参数:
 **param1:in-- Int : user_id    :: 用户ID
 **----------------------------------------------------------
 **返回：
 **return:out--  JSON ：data  :: JSON对象
 **----------------------------------------------------------
 **日期:2018.08.15  Add by zwx
 ************************************************************
 */
 func (a *Classify) GetInterest() {

 }

/**
 ************************************************************
 **方法::Classify::GetNewsList
 **----------------------------------------------------------
 **描述::获取文章详情
 **----------------------------------------------------------
 **参数:
 **param1:in-- Int : user_id    :: 用户ID
 **param2:in-- Int : page_num   :: 分页数量
 **param3:in-- Int : page_size  :: 分页大小
 **param4:in-- Int : type		:: 请求类型
 **----------------------------------------------------------
 **返回：
 **return:out--  JSON ：data  :: JSON对象
 **----------------------------------------------------------
 **日期:2018.08.15  Add by zwx
 ************************************************************
 */
 func (a *Classify) GetNewsList() {

 }