package main

import (
	"eventing/pkg/kncloudevents"
	"github.com/prometheus/common/log"
	"context"
	"fmt"
	"github.com/cloudevents/sdk-go-master/pkg/cloudevents"
	"net/http"
	"io/ioutil"
)
func main(){
	c, err := kncloudevents.NewDefaultClient()
	if err !=nil{
		log.Fatal("Failed to create client, ", err)
	}
	log.Fatal(c.StartReceiver(context.Background(), display))
}
func display(event cloudevents.Event) {

	fmt.Printf("☁️  cloudevents.Event\n%s", event.String())
	requestUrl := ""
	client := &http.Client{Transport:&http.Transport{}}
	req, err := http.NewRequest("GET",requestUrl,nil)
	if err != nil {
		log.Errorf("req:%v", err)
	}
	query:= req.URL.Query()
	req.URL.RawQuery = query.Encode()
	response, err := client.Do(req)
	if err != nil {
		log.Error("Response error ", err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error("Get response body fail ", err)

	}
	log.Debug("string result : %s", string(body))
	log.Debug("end!")
}

