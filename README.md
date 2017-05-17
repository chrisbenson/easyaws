# EasyAWS

## Convenience utilities for Amazon Web Services

```
import "github.com/chrisbenson/easyaws/pkg/easyaws"
```

#### EasyAWS makes the common use cases trivial to implement, without having to remember or lookup documentation details.  It uses the official AWS SDK for Go 'under the hood'.

The library has been implemented, as described below, but the CLI has not been implemented yet.

Before making any AWS service call, you will need to acquire a pointer
to an AWS Session.  With this utility, that requires only one line of
code.

Once an AWS Session has been acquired, you can make an AWS service call,
passing your AWS Session pointer into the service call function.
Ideally, it should be done without an allocation, as in the examples below.

### Sessions

Convenience methods that return a pointer to an AWS session.Session.

If (and only if) you need to allocate a variable for the *session.Session, you must:

```
import "github.com/aws/aws-sdk-go/aws/session"
```

Four approaches to return a pointer to an AWS *session.Session are supported:

#### AWS Session From Environmental Variables
###### func SessionFromEnvVars() *session.Session

```
sessPtr := easyaws.SessionFromEnvVars()
```

Supports these environmental variables:

- AWS_ACCESS_KEY_ID
- AWS_SECRET_ACCESS_KEY
- AWS_SESSION_TOKEN
- AWS_REGION
- AWS_PROFILE

#### AWS Session from Profile (file)
###### func SessionFromProfile(Profile string) *session.Session

```
sessPtr := easyaws.SessionFromProfile("gopher")
```

Looks for profile settings in two locations - both required:
- ~/.aws/credentials
- ~/.aws/config

#### SessionFromStaticCreds()
###### func SessionFromStaticCreds(AwsAccessKeyID string, AwsSecretAccessKey string, AwsRegion string) *session.Session

```
AwsAccessKeyID := "AKIAJDFYZLXQ7NR4FG7A"  // Example only.  Not valid.
AwsSecretAccessKey := "o6Aahq0dEJMvBWd0QmNb4a0zrEoTyBsKMdwxo54F"  // Example only.  Not valid.
AwsRegion := "us-east-1"
sessPtr := easyaws.SessionFromStaticCreds(AwsAccessKeyID, AwsSecretAccessKey, AwsRegion)
```

### SES

#### Send Email
###### func SendMail(to []string, from string, cc []string, bcc []string, subject string, body string, awsSession *session.Session) (string, error)

```
to 	:= []string{"larry@stooges.com"}
from 	:= "mo@stooges.com"
cc 	:= []string{"curly@stooges.com"}
bcc 	:= []string{}
subject := "Joke"
body 	:= "Why did the chicken cross the road?"
profile := "default"
messageID, err := aws.SendMail(to, from, cc, bcc, subject, body, easyaws.SessionFromEnvVars())
if err != nil{
    fmt.Println(err)
}
fmt.Println("Message ID: " + messageID)
```

### S3

#### Download Files from S3
###### func FilesFromS3(localDir string, bucket string, keys []string, awsSession *session.Session) error

```
localDir := "/Users/chrisbenson"
bucket := "chrisbenson"
keys := []string{"myTestFiles/test1.txt", "myTestFiles/test2.txt", "myTestFiles/test3.txt"}
err := FilesFromS3(localDir, bucket, keys, easyaws.SessionFromEnvVars())
if err != nil {
    fmt.Println(err)
}
```

#### Download Bytes from S3
###### func BytesFromS3(bucket string, keys []string, awsSession *session.Session) (map[string][]byte, error)

```
bucket := "chrisbenson"
keys := []string{"myTestFiles/test1.txt", "myTestFiles/test2.txt", "myTestFiles/test3.txt"}
bytesMap, err := BytesFromS3(bucket, keys, easyaws.SessionFromEnvVars())
if err != nil {
    fmt.Println(err)
}
```

#### Upload Bytes to S3
###### func BytesToS3(byteMap map[string][]byte, bucket string, prefix string, awsSession *session.Session) error

```
var byteMap = make(map[string][]byte, 3)
byteMap["luke"] = []byte("skywalker")
byteMap["princess"] = []byte("leia")
byteMap["darth"] = []byte("vader")
bucket := "chrisbenson"
prefix := "myTestFiles/" //If present, full S3 file path prefix without file name.  Include a trailing slash, but not a leading slash.
profile := "default"
err := easyaws.BytesToS3(byteMap, bucket, prefix, easyaws.SessionFromEnvVars()) error
if err != nil {
    fmt.Println(err)
}
```

#### Upload Files to S3
###### func FilesToS3(files []string, bucket string, prefix string, awsSession *session.Session) error

```
files := []string{"/Users/chrisbenson/test1.txt", "/Users/chrisbenson/test2.txt", "/Users/chrisbenson/test3.txt"} // Full local file path
bucket := "chrisbenson"
prefix := "myTestFiles/" //If present, full S3 file path prefix without file name.  Include a trailing slash, but not a leading slash.
profile := "default"
err := aws.FilesToS3(files, bucket, prefix, easyaws.SessionFromEnvVars())
if err != nil {
    fmt.Println(err)
}
```

#### Delete From S3
###### func DeleteFromS3(bucket string, keys []string, awsSession *session.Session) error

```
bucket := "chrisbenson"
keys := []string{"myTestFiles/test1.txt", "myTestFiles/test2.txt", "myTestFiles/test3.txt"}
err := aws.DeleteFromS3(bucket, keys, easyaws.SessionFromEnvVars())
if err != nil {
    fmt.Println(err)
}
```

#### List S3 Keys
###### func ListS3Keys(bucket string, prefix string, awsSession *session.Session) ([]string, error)

```
bucket := "chrisbenson"
prefix := "documents/" // Leave empty to list the entire bucket, or include a path prefix for a subset of keys.  Include a trailing slash, but not a leading slash.
profile := "default"
keys, err := aws.ListS3Keys(bucket, prefix, easyaws.SessionFromEnvVars())
if err != nil {
    fmt.Println(err)
} else {
    for _, key := range keys {
    fmt.Println(key)
}
```

### TODO

This is an early development version.  New features are being added all the time.
