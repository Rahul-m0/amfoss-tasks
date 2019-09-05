package main

import (
    "github.com/dghubble/go-twitter/twitter"
    "github.com/dghubble/oauth1"
    "os"
    "flag"
    "fmt"
	
)

func main() {
    
	var name string
	fmt.Println("What is the twitter username you want to get the followers of?")
    fmt.Scanln(&name)
    k := flag.String("twitterHandle", name , "twitter handle of a user")
    flag.Parse()
    config := oauth1.NewConfig("Tts8r6QfmHku8UFWLm8G1BCfA", "4qz2TS5Y5HmLWKaeGln6rjWoOCBMhc0wx9TdHWjO2PQixKVUb9")
    token := oauth1.NewToken("1168486308882173952-wL37ykXc8O700Uh0Bv3BbxgRmrYfX9", "T4DwOuFtzfQSxGRlKyldthmKtHWgXcKmXvuh8GVtoNwf9")
    httpClient := config.Client(oauth1.NoContext, token)
    client := twitter.NewClient(httpClient)
    f, err := os.Create("Marco.txt")

    params := &twitter.FollowerListParams{
        ScreenName: *k,
    }
    
    followers, resp, err := client.Followers.List(params)
    var count int = 0;
    fmt.Println(resp, err)
    f.WriteString("Followers - " + *k)
    for _, follower := range followers.Users {
    count++
    f.WriteString("\n" + follower.ScreenName)
    }
    f.WriteString("\n")
    f.WriteString(fmt.Sprintf("\n", count))
    f.Close()
}
