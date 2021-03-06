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
 **描述：活动类型
 ***********************************************************
 */

 type Activity struct {
	 /*继承beego*/
	 beego.Controller
	 Serv services.ArticleService `inject:""`
 }

/**
 ************************************************************
 **方法::Activity::MemberShow
 **----------------------------------------------------------
 **描述::活动圈友秀
 **----------------------------------------------------------
 **参数:
 **param1:in-- Int : user_id    :: 用户ID
 **param2:in-- Int : page_num   :: 分页数量
 **param3:in-- Int : page_size  :: 分页大小
 **param4:in-- Int : type       :: 请求类型
 **----------------------------------------------------------
 **返回：
 **return:out--  JSON ：data  :: JSON对象
 **----------------------------------------------------------
 **日期:2018.08.15  Add by zwx
 ************************************************************
 */
 func (a *Activity) MemberShow() {

 }

/**
 ************************************************************
 **方法::Activity::GetLists
 **----------------------------------------------------------
 **描述::活动列表
 **----------------------------------------------------------
 **参数:
 **param1:in-- Int : user_id    :: 用户ID
 **param2:in-- Int : page_num   :: 分页数量
 **param3:in-- Int : page_size  :: 分页大小
 **param4:in-- Int : type       :: 请求类型
 **----------------------------------------------------------
 **返回：
 **return:out--  JSON ：data  :: JSON对象
 **----------------------------------------------------------
 **日期:2018.08.15  Add by zwx
 ************************************************************
 */
 func (a *Activity) GetLists() {

 }
/**
 ************************************************************
 **方法::Activity::GetVideos
 **----------------------------------------------------------
 **描述::获取视频列表
 **----------------------------------------------------------
 **参数:
 **param1:in-- Int : activity_id  :: 活动ID
 **param2:in-- Int : user_id      :: 用户ID
 **param3:in-- Int : page_num	  :: 分页数量
 **param4:in-- Int : page_size    :: 分页大小
 **param5:in-- Int : type		  :: 请求类型
 **----------------------------------------------------------
 **返回：
 **return:out--  JSON ：data    :: JSON对象
 **----------------------------------------------------------
 **日期:2018.08.15  Add by zwx
 ************************************************************
 */
 func (a *Activity) GetVideos() {

 }

/**
 ************************************************************
 **方法::Activity::VideoRank
 **----------------------------------------------------------
 **描述::视频排行榜
 **----------------------------------------------------------
 **参数:
 **param1:in-- Int : user_id    :: 用户ID
 **param2:in-- Int : page_num   :: 分页数量
 **param3:in-- Int : page_size  :: 分页大小
 **param4:in-- Int : type       :: 请求类型
 **----------------------------------------------------------
 **返回：
 **return:out--  JSON ：data    :: JSON对象
 **----------------------------------------------------------
 **日期:2018.08.15  Add by zwx
 ************************************************************
 */
 func (a *Activity) VideoRank() {

 }
/**
 ************************************************************
 **方法::Activity::CheckPartakeIn
 **----------------------------------------------------------
 **描述::视频排行榜
 **----------------------------------------------------------
 **参数:
 **param1:in-- Int : activity_id  :: 用户ID
 **param2:in-- Int : user_id      :: 用户ID
 **----------------------------------------------------------
 **返回：
 **return:out--  JSON ：data    :: JSON对象
 **----------------------------------------------------------
 **日期:2018.08.15  Add by zwx
 ************************************************************
 */
 func (a *Activity) CheckPartakeIn() {

 }