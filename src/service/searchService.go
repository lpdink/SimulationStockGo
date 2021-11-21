package service

import (
	"bytes"
	"encoding/json"
	"example.com/m/src/domain"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
)
//var token="7050090aab373e93c870c9533cba7bc9d396255960ee5d8b52412eb2"
var token="d2144df3d2c4f703454e331d9c95c0dc8685101ca04795d337cd5de1"
const url = "http://api.tushare.pro"


type DataMsg struct {
	Fields   []string   `json:"fields"`
	Has_more bool       `json:"has_more"`
	Items    [][]interface{} `json:"items"`
}

type TushareMsg struct {
	Code    string             `json:"code"`
	Data   DataMsg   `json:"data"`
}

func callAPI(apiName, fields string,
	params map[string]interface{}) []byte {
	m := map[string]interface{}{
		"api_name": apiName,
		"token":    token,
		"fields":   fields,
	}
	if params == nil {
		m["params"] = ""
	} else {
		m["params"] = params
	}
	b, _ := json.Marshal(m)
	resp, _ := http.Post(url, "application/json", bytes.NewBuffer(b))
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return result
}


func Login(userId string, password string)(int, float64,[] domain.UserHolding){
	var user = domain.UserInformation{}
	var user_holdings []domain.UserHolding
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return -1,-1,user_holdings
	}
	var is_exist = db.Where("user_id=?",userId).First(&user).Error==nil
	if is_exist{
		if user.PassWord==password{
			//查询剩余资产和持仓情况
			var funds=user.FreeMoneyAmount
			db.Where("user_id=?", userId).Find(&user_holdings)
			return 1, funds, user_holdings //登录成功
		}else {
			return 0,-1,user_holdings //密码输入错误
		}
	} else {
		return -1, -1,user_holdings//用户不存在
	}
}

func SearchWithId(stockid string){
	params:=map[string]interface{}{
		"ts_code":stockid,
		"start_date":"20211015",
		"end_date":"20211020",
	}
	var result = callAPI("daily", "", params)

	var config TushareMsg
	_= json.Unmarshal(result,&config)
	fmt.Println(config.Data)
	fmt.Println(config.Data.Fields[0],config.Data.Items[0][0])

}

func DrawK(stockid string, startdate string,enddate string)[]map[string]interface{}{
	params := map[string]interface{}{
		"ts_code":stockid,
		"start_date":startdate,
		"end_date":enddate,
	}
	var result = callAPI("daily","",params)
	var msg TushareMsg
	_ = json.Unmarshal(result, &msg)
	log.Println(msg)
	var i int
	//begin,_:=strconv.Atoi(startdate)
	//end, _:=strconv.Atoi(enddate)

	var rst []map[string]interface{}
	for i=0;i<len(msg.Data.Items);i++{
		m:=make(map[string]interface{})
		m["open"]=msg.Data.Items[i][2]
		m["close"]=msg.Data.Items[i][5]
		m["high"]=msg.Data.Items[i][3]
		m["low"]=msg.Data.Items[i][4]
		m["volume"]=msg.Data.Items[i][10]
		m["trade_data"]=msg.Data.Items[i][1]
		rst = append(rst,m)
		//关于时间戳，建议前端用js去获取
	}
	return rst
}

func SearchHoldings(user_id string)([]domain.UserHolding)  {
	var user_holdings []domain.UserHolding
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return user_holdings
	}
	db.Where("user_id=?",user_id).Find(&user_holdings)
	return user_holdings
}
