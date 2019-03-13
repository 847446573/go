package main

import (
	//"net/http"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	r := gin.Default()
	// Ping test
	r.POST("/check/api/sdk",Adrequest)
	r.POST("/ad/imp/:pid/:space/:cid",Imp)
	r.POST("/ad/click/:pid/:space/:cid",Click)
	r.POST("/ad/video_play_monitor/:pid/:space/:cid/:flag",VideoPlayMonitor)
	r.POST("/ad/material_play_monitor/:pid/:space/:cid/:flag/:duration",MaterialPlayMonitor)
	r.POST("/ad/material_play_monitor/:pid/:space/:cid/:flag",MaterialPlayMonitor)
	r.POST("/adx/qry/:sspId",Qry)
	r.POST("/sms/click/:campaign/:identifier_type/:identifier/:create_time",Smsclick)
	r.POST("/content/:trigger_id",AdContent)
	r.POST("/list/:trigger_id",AdList)
	r.POST("/prize/picked/:trigger_id",Prizepicked)
	r.POST("/raffle/:pid/:trigger_id/:activity_id/:prize_id",Prizeraffle)
	r.POST("/prize/imp/:pid/:trigger_id/:activity_id/:prize_id",Prizeimp)
	r.POST("/prize/click/:pid/:trigger_id/:activity_id/:prize_id",Prizeclick)
	r.POST("/winner_list/:trigger_id",WinnerList)
	r.POST("/saveWcmnp/:trigger_id",Wcmnp)
	r.POST("/postback",Postback)
	r.POST("/commodity",Commodity)
	r.POST("/commit",AdCommit)
	r.POST("/orderList",OrderList)
	r.POST("/detail",Adetail)
	r.POST("/order/finish",Afinish)
	r.POST("/ad/trace/:type/:pid/:adunit/:cid/:content",SspTrace)
	r.POST("/fill/click",FillClick)
	r.POST("/fill/postback",Postback)
	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
