package main

import (
        "strings"
        "context"
        "fmt"
        "github.com/aws/aws-lambda-go/lambda"
        "os"
       )


func handler(ctx context.Context) (string, error) {
    cost, err := getCost()
        if err != nil {
            return "", err
        }
    phoneNumberArray := strings.Split(os.Getenv("PHONENUMBER"), ",")
        for _, i := range phoneNumberArray {
            if i == "" {
                return "", fmt.Errorf("Please specify env variable PHONENUMBER including the country code")
            }
     
        err = sendSms(&sendSmsInput{
             message:     fmt.Sprintf("The aws cost so far for today is $%s", cost),
             phoneNumber:  i,
        })
        if err != nil {
            return "", err
        }
    }
    return "v1", err
}

func main() {
    lambda.Start(handler)
}
