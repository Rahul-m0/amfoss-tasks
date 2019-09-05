package main


import (
    
    
    "github.com/dghubble/go-twitter/twitter"
    "github.com/dghubble/oauth1"
    
    
	"log"
	"os"
	"fmt"
)

type Credentials struct {
    ConsumerKey       string
    ConsumerSecret    string
    AccessToken       string
    AccessTokenSecret string
}
func getClient(creds *Credentials) (*twitter.Client, error) {
    
    config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
    
    token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)

    httpClient := config.Client(oauth1.NoContext, token)
    
    client := twitter.NewClient(httpClient)

    
    verifyParams := &twitter.AccountVerifyParams{
        SkipStatus:   twitter.Bool(true),
        IncludeEmail: twitter.Bool(true),
    }
      
    user, _, err := client.Accounts.VerifyCredentials(verifyParams)
    if err != nil {
        return nil, err
    }

    log.Printf("User's ACCOUNT:\n%+v\n", user)
    return client, nil
}
func main() {
    
    creds := Credentials{
        AccessToken:       os.Getenv("1168486308882173952-wL37ykXc8O700Uh0Bv3BbxgRmrYfX9"),
        AccessTokenSecret: os.Getenv("T4DwOuFtzfQSxGRlKyldthmKtHWgXcKmXvuh8GVtoNwf9"),
        ConsumerKey:       os.Getenv("Tts8r6QfmHku8UFWLm8G1BCfA"),
        ConsumerSecret:    os.Getenv("4qz2TS5Y5HmLWKaeGln6rjWoOCBMhc0wx9TdHWjO2PQixKVUb9"),
    }

    client, err := getClient(&creds)
    if err != nil {
        log.Println("Error getting Twitter Client")
        log.Println(err)
    }

    
    fmt.Printf("%+v\n", client)

}
