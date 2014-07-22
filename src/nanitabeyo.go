package hello

import (
	"net/http"
	"net/url"
	"appengine"
	"appengine/urlfetch"
	"math/rand"
	"time"
)

var api_tokens = []string{
	"773bb293-8ff1-931b-625d-3941d40e4b44",
	"29cf40d0-4a78-918a-800d-4881639522ac",
	"07a04821-7da1-551d-6e04-e33fb28b7307",
	"208c4c53-f986-cb0d-d07f-71bc765e0b54",
	"57a78e37-3c47-af8e-f79c-470e1358d028",
	"23ebeee8-e7eb-4fa7-bff1-a786d94a09bf",
	"51df1ced-c7a1-5311-681a-bf035229b5cb",
	"212d42a0-a7e0-ce73-4d61-7251dd44d2a3",
	"eb77681e-0705-28ed-3ec2-a0eb378a7e28",
	"0af398b4-cb41-3534-cd60-803580582d17"}

const YO_USER string = "http://api.justyo.co/yo/"

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query()["username"]
	rand.Seed(time.Now().UnixNano())
	randomtoken := api_tokens[rand.Intn(len(api_tokens))]
	
	c := appengine.NewContext(r)
	client := http.Client{Transport: &urlfetch.Transport{Context:c}}
	//&& apikey != ""
	if username != nil {
		res, err := client.PostForm(YO_USER, url.Values{"api_token":{randomtoken},"username":username})
		if err != nil {
			c.Errorf("Error %s",err)
		}
		if res.StatusCode != 200{
			c.Warningf("yo failed with: %s", res)
		}else
		{
			
		}
	}else
	{
		c.Errorf("username and key required")
	}
}


