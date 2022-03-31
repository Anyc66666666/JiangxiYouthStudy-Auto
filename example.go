package main

import (
	"fmt"
	"github.com/robfig/cron"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)
//江西省高校青年大学习 自动学习

func main() {
c:=cron.New()
 c.AddFunc("0 0 13 * * 3", func() {//定时每周三下午1点青年大学习
	 id:=GetLatest()  //这个是自动获取最新一期的青年大学习


	 //****   nid 和 cardNo需要修改成自己对应的nid和cardNo
	 nid:="N0013************"  //这里填自己所在的江西省内的大学 学院 专业 班级 对应的nid
	                          //(N开头，具体可去这里看--> http://osscache.vol.jxmfkj.com/pub/vol/config/organization?pid=N0013
	 cardNo:="姓名+学号"   //cardNo 就是自己的个人信息
	 //***



     PostQNDXX(id,nid,cardNo)



 })

c.Start()

	select {

}


}

func GetLatest()(id string){
	resp,err:=http.Get("http://osscache.vol.jxmfkj.com/html/assets/js/course_data.js")
	if err!=nil{fmt.Println("get failed,err:",err)}
	defer resp.Body.Close()

	b,err:=ioutil.ReadAll(resp.Body)
	if err!=nil{fmt.Println("read body failed,err:",err)}
//	fmt.Println(string(b))
m:=strings.Split(string(b),":")[2]
id=strings.Split(m,"\"")[1]
//id=strings.TrimSpace(id)
fmt.Println(id)
return id

}

func PostQNDXX(id,nid,cardNo string){
	url:="http://osscache.vol.jxmfkj.com/pub/vol/volClass/join?accessToken="
	client:=&http.Client{}
	param:=fmt.Sprintf(`{"course":"%v","subOrg":null,"nid":"%v","cardNo":"%v"}`,id,nid,cardNo)

	req,err:=http.NewRequest(http.MethodPost,url,strings.NewReader(param))
	if err!=nil{fmt.Println("new req failed,err:",err)}
	req.Header= map[string][]string{
		"Accept": {"application/json,text/javascript,*/*","q=0.01"},
	//"Accept-Encoding": {"gzip,deflate"},
	"Accept-Language": {"zh-CN,zh","q=0.9,en","q=0.8,en-GB","q=0.7,en-US","q=0.6"},
	"Content-Length": {"82"},
	"Content-Type": {"application/json,","charset=UTF-8"},
	"Proxy-Connection": {"keep-alive"},
	"User-Agent": {"Mozilla/5.0(Linux;Android6.0;Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML,like Gecko) Chrome/99.0.4844.74 Mobile Safari/537.36 Edg/99.0.1150.46"},
	}
	var n int
	for i:=0;i<8;i++{
		time.Sleep(time.Second*3)
	resp,err:=client.Do(req)
	if err!=nil{fmt.Println("client do failed,err:",err)}
      defer resp.Body.Close()
    b,err:=ioutil.ReadAll(resp.Body)
	if err!=nil{fmt.Println("read all failed,err:",err)}
	fmt.Println(string(b))
	s:=strings.Split(string(b),":")[2]
	if s==`200}`{
		fmt.Println("status:200")
		n=6
		break
	}else {
		fmt.Println("qndxx,failed")
		n=4
	continue
	}
}
	  if n==4{
		  if err!=nil{fmt.Println("send mail failed,err:",err)
		  }
	  }else{
		if err!=nil{fmt.Println("send mail failed,err:",err)}
	}
}




