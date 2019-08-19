package main

import (
	"github.com/dghubble/oauth1"
	"github.com/dghubble/oauth1/twitter"
	"bufio"
    "fmt"
    "net/http"
	"github.com/coreos/pkg/flagutil"
	
	
		
)

func main() {
	config := oauth1.NewConfig("consumerKey", "consumerSecret")
	token := oauth1.NewToken("accessToken", "accessSecret")
	
	httpClient := config.Client(oauth1.NoContext, token)

	
	client := twitter.NewClient(httpClient)	
}
