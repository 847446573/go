package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "io/ioutil"
  "fmt"
  "encoding/json"
  "net"
  "net/url"
  "strings"
  "time"
  "strconv"
)

func Adrequest(c *gin.Context) {
	  data, _ := ioutil.ReadAll(c.Request.Body)
	  mainMap := make(map[string]interface{})
	  t  := make(map[string]interface{})
	  subMap  := make(map[string]string)
	  jsonData := []byte(data)
	  err := json.Unmarshal(jsonData,&t)
      if err != nil {
      	fmt.Printf("error")
      }
	  subMap["ip"]   = getIp()
	  mainMap["req"] = t
	  mainMap["ext"] = subMap
	  rdata,_ := json.Marshal(mainMap)
	  fmt.Printf(string(rdata))
	  result := "[]"
	  c.String(200,result)	  
}


func Imp(c *gin.Context) {
	pid   := c.Param("pid")
	space := c.Param("space")
	cid   := c.Param("cid")
	mainMap := make(map[string]map[string]string)
	extMap  := make(map[string]string)
	reqMap  := make(map[string]string)
	extMap["ip"]     = getIp()
	reqMap["pid"]    =  pid
	reqMap["space"]  =  space
	reqMap["cid"]    =  cid
	mainMap["req"]   =  reqMap
	mainMap["ext"]   =  extMap
	data,_ := json.Marshal(mainMap)
	fmt.Printf(string(data))
	result := "[]"
	c.String(200,result)
	
}

func Click(c *gin.Context) {
	fd    := c.Query("fd")
	pid   := c.Param("pid")
	space := c.Param("space")
	cid   := c.Param("cid")
	reqMap,extMap,data := make(map[string]string),make(map[string]string),make(map[string]map[string]string)
	extMap["ip"] = getIp()
	reqMap["pid"]   = pid
	reqMap["space"] = space
	reqMap["cid"]   = cid
	data["req"] = reqMap
	data["ext"] = extMap
	result,err := json.Marshal(data)
	if err != nil {
		c.String(500,"[]")
	}
	fmt.Printf(string(result))
	if fd != "" {
		decodeurl,err := url.QueryUnescape(fd)
		if err !=  nil {
			c.String(200,"[]")
		}
		locationUrl := marProReplace(decodeurl)
		c.Redirect(http.StatusMovedPermanently,locationUrl)
		return
	}
	c.String(200,"[]")
}


func VideoPlayMonitor(c *gin.Context) {
	pid   := c.Param("pid")
	space := c.Param("space")
	cid   := c.Param("cid")
	flag  := c.Param("flag")
	reqMap,extMap,dataMap := make(map[string]string),make(map[string]string),make(map[string]map[string]string)
	reqMap["pid"]   = pid
	reqMap["space"] = space
	reqMap["cid"]   = cid
	reqMap["flag"]  = flag
	extMap["ip"]    = getIp()
	dataMap["req"]  = reqMap
	dataMap["ext"]  = extMap
	data,err := json.Marshal(dataMap)
	if err != nil {
		c.String(500,"[]")
	}
	fmt.Printf(string(data))
	result := "[]"
	c.String(200,result)
}

/*type ReqData struct {
	Pid        string  `uri:"pid" binding:"required"`
	Space      string  `uri:"space" binding:"required"`
	Cid        string  `uri:"cid" binding:"required"`
	flag       string  `uri:"flag" binding:"required"`
	duration   string  `uri:"duration" binding:"required"`
}*/

func MaterialPlayMonitor(c *gin.Context) {
	pid   := c.Param("pid")
	space := c.Param("space")
	cid   := c.Param("cid")
	flag  := c.Param("flag")
	duration  := c.Param("duration")
	reqMap,extMap,dataMap := make(map[string]string),make(map[string]string),make(map[string]map[string]string)
	reqMap["pid"]   = pid
	reqMap["space"] = space
	reqMap["cid"]   = cid
	if flag ==  "inter" {
		reqMap["duration"]  = duration
	}
	reqMap["flag"]  = flag
	extMap["ip"]    = getIp()
	dataMap["req"]  = reqMap
	dataMap["ext"]  = extMap
	data,err := json.Marshal(dataMap)
	if err != nil {
		c.String(500,"[]")
	}
	fmt.Printf(string(data))
	result := "[]"
	c.String(200,result)

}


func Qry(c *gin.Context) {
	sspId := c.Param("sspId")
	data,_ := ioutil.ReadAll(c.Request.Body)
	reqMap,extMap,mainMap := make(map[string]interface{}),make(map[string]string),make(map[string]interface{})
	tempMap := make(map[string]interface{})
	err     := json.Unmarshal([]byte(data),&tempMap)
	if err != nil {
		c.String(200,"Param is not json")
	}
	reqMap["req"] = tempMap
	extMap["adunit"] = sspId
	extMap["ip"] = getIp()
	mainMap["req"] = reqMap
	mainMap["ext"] = extMap
    res,_ := json.Marshal(mainMap)
    fmt.Printf(string(res))
}

func Smsclick(c *gin.Context) {
	campaign := c.Param("campaign")
	identifierType := c.Param("identifier_type")
	identifier := c.Param("identifier")
	createTime := c.Param("create_time")
	lp := c.Query("lp")
	reqMap,extMap,mainMap := make(map[string]string),make(map[string]string),make(map[string]map[string]string)
	reqMap["campaign"] = campaign
	reqMap["identifier_type"] = identifierType
	reqMap["identifier"] = identifier
	reqMap["create_time"] = createTime
	extMap["ip"] = getIp()
	mainMap["req"] = reqMap
	mainMap["ext"] = extMap
	res,_ := json.Marshal(mainMap)
	fmt.Printf(string(res))
	if lp != "" {
		c.Redirect(http.StatusMovedPermanently,lp)
		return
	}
	c.String(200,"[]")
}

func AdContent(c *gin.Context) {
	jsonData := make(map[string]interface{})
	body,_ := ioutil.ReadAll(c.Request.Body)
	triggerId := c.Param("trigger_id")
	err := json.Unmarshal([]byte(body),&jsonData)
	if err != nil {
		c.String(200,"json decode error")
	}
	delete(jsonData,"_url")
	extMap,mainMap := make(map[string]interface{}),make(map[string]interface{})
    jsonData["trigger_id"] = triggerId
    extMap["ip"] = getIp()
    mainMap["req"] = jsonData
    mainMap["ext"] = extMap
    res,_ := json.Marshal(mainMap)
    fmt.Printf(string(res))
    c.String(200,"[]")
}

func AdList(c *gin.Context) {
	jsonData := make(map[string]interface{})
	body,_ := ioutil.ReadAll(c.Request.Body)
	triggerId := c.Param("trigger_id")
	err := json.Unmarshal([]byte(body),&jsonData)
	if err != nil {
		c.String(200,"json decode error")
	}
	delete(jsonData,"_url")
	extMap,mainMap := make(map[string]interface{}),make(map[string]interface{})
    jsonData["trigger_id"] = triggerId
    extMap["ip"] = getIp()
    mainMap["req"] = jsonData
    mainMap["ext"] = extMap
    res,_ := json.Marshal(mainMap)
    fmt.Printf(string(res))
    c.String(200,"[]")
}

func Prizepicked(c *gin.Context) {
	jsonData := make(map[string]interface{})
	body,_ := ioutil.ReadAll(c.Request.Body)
	triggerId := c.Param("trigger_id")
	err := json.Unmarshal([]byte(body),&jsonData)
	if err != nil {
		c.String(200,"json decode error")
	}
	delete(jsonData,"_url")
	extMap,mainMap := make(map[string]interface{}),make(map[string]interface{})
    jsonData["trigger_id"] = triggerId
    extMap["ip"] = getIp()
    mainMap["req"] = jsonData
    mainMap["ext"] = extMap
    res,_ := json.Marshal(mainMap)
    fmt.Printf(string(res))
    c.String(200,"[]")
}

func Prizeraffle(c *gin.Context) {
	pid := c.Param("pid")
	triggerId  := c.Param("trigger_id")
	activityId := c.Param("activity_id")
	prizeId := c.Param("prize_id")
	reqMap,extMap,dataMap := make(map[string]string),make(map[string]string),make(map[string]map[string]string)
	reqMap["pid"] = pid
	reqMap["trigger_id"]  = triggerId
	reqMap["activity_id"] = activityId
	reqMap["prize_id"]    = prizeId
	extMap["ip"] = getIp()
	dataMap["req"] = reqMap
	dataMap["ext"] = extMap
	res,_ := json.Marshal(dataMap)
	fmt.Printf(string(res))
    c.String(200,"[]")
}

func Prizeimp(c *gin.Context) {
	pid := c.Param("pid")
	triggerId  := c.Param("trigger_id")
	activityId := c.Param("activity_id")
	prizeId := c.Param("prize_id")
	reqMap,extMap,dataMap := make(map[string]string),make(map[string]string),make(map[string]map[string]string)
	reqMap["pid"] = pid
	reqMap["trigger_id"]  = triggerId
	reqMap["activity_id"] = activityId
	reqMap["prize_id"]    = prizeId
	extMap["ip"] = getIp()
	dataMap["req"] = reqMap
	dataMap["ext"] = extMap
	res,_ := json.Marshal(dataMap)
	fmt.Printf(string(res))
    c.String(200,"[]")
}

func Prizeclick(c *gin.Context) {
	pid := c.Param("pid")
	triggerId  := c.Param("trigger_id")
	activityId := c.Param("activity_id")
	prizeId := c.Param("prize_id")
	reqMap,extMap,dataMap := make(map[string]string),make(map[string]string),make(map[string]map[string]string)
	reqMap["pid"] = pid
	reqMap["trigger_id"]  = triggerId
	reqMap["activity_id"] = activityId
	reqMap["prize_id"]    = prizeId
	extMap["ip"] = getIp()
	dataMap["req"] = reqMap
	dataMap["ext"] = extMap
	res,_ := json.Marshal(dataMap)
	fmt.Printf(string(res))
    c.String(200,"[]")
}

func WinnerList(c *gin.Context) {
	jsonData := make(map[string]interface{})
	body,_ := ioutil.ReadAll(c.Request.Body)
	triggerId := c.Param("trigger_id")
	err := json.Unmarshal([]byte(body),&jsonData)
	if err != nil {
		c.String(200,"json decode error")
	}
	delete(jsonData,"_url")
	extMap,mainMap := make(map[string]interface{}),make(map[string]interface{})
	jsonData["method"] = "winList"
    jsonData["trigger_id"] = triggerId
    extMap["ip"] = getIp()
    mainMap["req"] = jsonData
    mainMap["ext"] = extMap
    res,_ := json.Marshal(mainMap)
    fmt.Printf(string(res))
    c.String(200,"[]")
}

func Wcmnp(c *gin.Context) {
	jsonData := make(map[string]interface{})
	body,_ := ioutil.ReadAll(c.Request.Body)
	triggerId := c.Param("trigger_id")
	err := json.Unmarshal([]byte(body),&jsonData)
	if err != nil {
		c.String(200,"json decode error")
	}
	delete(jsonData,"_url")
	extMap,mainMap := make(map[string]interface{}),make(map[string]interface{})
    jsonData["trigger_id"] = triggerId
    extMap["ip"] = getIp()
    mainMap["req"] = jsonData
    mainMap["ext"] = extMap
    res,_ := json.Marshal(mainMap)
    fmt.Printf(string(res))
    c.String(200,"[]")
}

func Postback(c *gin.Context) {
	jsonData := make(map[string]interface{})
	body,_ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal([]byte(body),&jsonData)
	if err != nil {
		c.String(200,"json decode error")
	}
	delete(jsonData,"_url")
	extMap,mainMap := make(map[string]interface{}),make(map[string]interface{})
    extMap["ip"] = getIp()
    mainMap["req"] = jsonData
    mainMap["ext"] = extMap
    res,_ := json.Marshal(mainMap)
    fmt.Printf(string(res))
    c.String(200,"[]")
}

func Commodity(c *gin.Context) {
	jsonData := make(map[string]interface{})
	body,_ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal([]byte(body),&jsonData)
	if err != nil {
		c.String(200,"json decode error")
	}
	delete(jsonData,"_url")
	extMap,mainMap := make(map[string]interface{}),make(map[string]interface{})
    extMap["ip"] = getIp()
    mainMap["req"] = jsonData
    mainMap["ext"] = extMap
    res,_ := json.Marshal(mainMap)
    result := "{result:0}"
    if jsonData["method"] == "list" {
    } else if jsonData["method"] == "detail" {
    } else {
    	result = "{result:0}"

    }
    fmt.Printf(string(res))
    c.String(200,string(result))
}

func AdCommit(c *gin.Context) {
	jsonData := make(map[string]interface{})
	body,_ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal([]byte(body),&jsonData)
	if err != nil {
		c.String(200,"json decode error")
	}
	province,ok := jsonData["province"].(string)
	if !ok {
		c.String(200," error")
	}
	city,ok := jsonData["city"].(string)
	if !ok {
		c.String(200," error")
	}
	county,ok := jsonData["county"].(string)
	if !ok {
		c.String(200," error")
	}
	address,ok := jsonData["address"].(string)
	if !ok {
		c.String(200," error")
	}
	username,ok := jsonData["username"].(string)
	if !ok {
		c.String(200," error")
	}
	phone,ok := jsonData["phone"].(string)
	if !ok {
		c.String(200," error")
	}
	var addrs = [] string {province,city,county,address}
	deliveryAddress := make(map[string]string)
	deliveryAddress["address"]    = strings.Join(addrs,"")
	deliveryAddress["mobile"]     = phone
	deliveryAddress["consignee"]  = username
	jsonData["delivery_address"] = deliveryAddress
	delete(jsonData,"_url")
	delete(jsonData,"address")
	extMap,mainMap := make(map[string]interface{}),make(map[string]interface{})
    extMap["ip"] = getIp()
    mainMap["req"] = jsonData
    mainMap["ext"] = extMap
    res,_ := json.Marshal(mainMap)
    c.String(200,string(res))
}

func OrderList(c *gin.Context) {
	jsonData := make(map[string]interface{})
	body,_ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal([]byte(body),&jsonData)
	if err != nil {
		c.String(200,"json decode error")
	}
	delete(jsonData,"_url")
	extMap,mainMap := make(map[string]interface{}),make(map[string]interface{})
    extMap["ip"] = getIp()
    mainMap["req"] = jsonData
    mainMap["ext"] = extMap
    res,_ := json.Marshal(mainMap)
    c.String(200,string(res))
}
func Adetail(c *gin.Context) {
	jsonData := make(map[string]interface{})
	body,_ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal([]byte(body),&jsonData)
	if err != nil {
		c.String(200,"json decode error")
	}
	delete(jsonData,"_url")
	extMap,mainMap := make(map[string]interface{}),make(map[string]interface{})
    extMap["ip"] = getIp()
    mainMap["req"] = jsonData
    mainMap["ext"] = extMap
    res,_ := json.Marshal(mainMap)
    c.String(200,string(res))
}

func Afinish(c *gin.Context) {
	jsonData := make(map[string]interface{})
	body,_ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal([]byte(body),&jsonData)
	if err != nil {
		c.String(200,"json decode error")
	}
	delete(jsonData,"_url")
	extMap,mainMap := make(map[string]interface{}),make(map[string]interface{})
    extMap["ip"] = getIp()
    mainMap["req"] = jsonData
    mainMap["ext"] = extMap
    res,_ := json.Marshal(mainMap)
    c.String(200,string(res))
}

func SspTrace(c *gin.Context) {
	fd    := c.Query("fd")
	price := c.Query("price")
	key   := c.Query("key")
	stype  := c.Param("type")
	pid   := c.Param("pid")
	adunit    := c.Param("adunit")
	cid       := c.Param("cid")
	content   := c.Param("content")
	jsonData := make(map[string]interface{})
	body,_ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal([]byte(body),&jsonData)
	if err != nil {
		c.String(200,"json decode error")
	}
	jsonData["type"]    = stype
	jsonData["pid"]     = pid
	jsonData["adunit"]  = adunit
	jsonData["cid"]     = cid
	jsonData["content"] = content
	extMap := make(map[string]string)
	extMap["ip"]    = getIp()
	extMap["price"] = price
	extMap["key"]   = key
	dataMap := make(map[string]interface{})
	dataMap["req"] = jsonData
	dataMap["ext"] = extMap
	res,_ := json.Marshal(dataMap)
	if fd != "" {
		decodeurl,err := url.QueryUnescape(fd)
		if err !=  nil {
			c.String(200,"[]")
		}
		locationUrl := marProReplace(decodeurl)
		c.Redirect(http.StatusMovedPermanently,locationUrl)
		return
	}
	c.String(200,string(res))
}

func FillClick(c *gin.Context) {
	jsonData := make(map[string]interface{})
	body,_ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal([]byte(body),&jsonData)
	if err != nil {
		c.String(200,"json decode error")
	}
	extMap,mainMap := make(map[string]interface{}),make(map[string]interface{})
    extMap["ip"] = getIp()
    mainMap["req"] = jsonData
    mainMap["ext"] = extMap
    res,_ := json.Marshal(mainMap)
    c.String(200,string(res))
}

func Postback(c *gin.Context) {
	jsonData := make(map[string]interface{})
	body,_ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal([]byte(body),&jsonData)
	if err != nil {
		c.String(200,"json decode error")
	}
	extMap,mainMap := make(map[string]interface{}),make(map[string]interface{})
    extMap["ip"] = getIp()
    mainMap["req"] = jsonData
    mainMap["ext"] = extMap
    res,_ := json.Marshal(mainMap)
    c.String(200,string(res))
}

/*宏替换*/
func marProReplace(s string) string {	
	if strings.Contains(s,"{CLICK_TS}") {
		s = strings.Replace(s,"{CLICK_TS}",strconv.FormatInt(time.Now().UnixNano() / 1e6,10),-1)
	} 
	if strings.Contains(s,"{CURRENT_TS}") {
		s =  strings.Replace(s,"{CURRENT_TS}",strconv.FormatInt(time.Now().UnixNano() / 1e6,10),-1)
	}
	if strings.Contains(s,"{CLICK_SS}") {
		s =  strings.Replace(s,"{CLICK_SS}",strconv.FormatInt(time.Now().Unix(),10),-1)
	} 
	if strings.Contains(s,"{CLINET_IP}") {
		s = strings.Replace(s,"{CLINET_IP}",getIp(),-1)
	}
	return s
}
/**
* 获取客户端ip
*/
func getIp() string {
	var ipString string;
	ipString = ""
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        fmt.Println(err)
    }
    for _, address := range addrs {
        if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
            	ipString = ipnet.IP.String()
            }

        }
    }
    return ipString;
}

/*func t(i interface{}) int{
	var tag int = 0;
	switch i.(type) {
	case string:
		tag = 1
	case int:
		tag = 2 

    default :
    	tag = 0

	}
	return tag
}*/





