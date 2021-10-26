package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)


func main() {
	sess := session.Must(session.NewSession())

	svc := sts.New(sess)


	opt := sts.GetSessionTokenInput{
		SerialNumber:    serialNumber(),
		TokenCode:       code(),
		DurationSeconds: ttl(),
	}

	t, err := svc.GetSessionToken(&opt)
	if err != nil {
		fmt.Println("token failed", err)
		os.Exit(1)
		return
	}

	Print(t.Credentials)

}

// Print bash code for export variables
func Print(c * sts.Credentials)  {
	fmt.Printf("export AWS_SESSION_TOKEN=%q\n", *c.SessionToken)
	fmt.Printf("export AWS_SECURITY_TOKEN=%q\n", *c.SessionToken)

	fmt.Printf("export AWS_SECRET_ACCESS_KEY=%q\n", *c.SecretAccessKey)

	fmt.Printf("export AWS_ACCESS_KEY=%q\n", *c.AccessKeyId)
	fmt.Printf("export AWS_ACCESS_KEY_ID=%q\n", *c.AccessKeyId)
}




// serialNumber 
func serialNumber() *string {
	env := os.Getenv("AWS_ARN")

	if len(env) > 0 {
		return &env
	}

	return nil
}

// ttl time to live token
func ttl() *int64 {
	env := os.Getenv("TTL")

	if len(env) > 0 {
		i64 , _ := strconv.ParseInt(env, 10, 64)
		return &i64
	}

	return nil
}

// code is tokenCode auths
func code() *string {
	if len(os.Args) >= 2 {
		return &os.Args[1]
	}

	return nil
}
