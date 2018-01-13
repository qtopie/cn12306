package main

import(
	"net/http"
	"log"
	"io/ioutil"
	"crypto/tls"
)


func q(rawurl string){
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	req, err:= http.NewRequest("GET", rawurl, nil)
	if err!=nil{
		log.Fatal(err)
	}

	q := req.URL.Query()
	q.Set("leftTicketDTO.train_date","2018-02-11")
	q.Set("leftTicketDTO.from_station","IZQ")
	q.Set("leftTicketDTO.to_station","WHN")
	q.Set("purpose_codes","ADULT")
	req.URL.RawQuery = unsortedEncode(q)

	log.Println(req.URL.String())

	resp, err:= client.Do(req)
	if err != nil{
		log.Fatal(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err!=nil{
		log.Fatal(err)
	}
	log.Println("Data", string(data))
}

func main(){
	// q("https://kyfw.12306.cn/otn/leftTicket/log")
	q("https://kyfw.12306.cn/otn/leftTicket/queryZ")
}