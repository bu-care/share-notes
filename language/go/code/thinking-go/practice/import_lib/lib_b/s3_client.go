package clients

import (
	"fmt"
	boto3 "thinking-test/practice/import_lib/lib_a"
)

type S3_CLIENT struct{}

var G_S3_CLIENT *S3_CLIENT

func init() {
	G_S3_CLIENT = &S3_CLIENT{}
	boto3.Set_client("s3", G_S3_CLIENT)
}

func (s *S3_CLIENT) Put_object(bucket, key string, body string) any {
	fmt.Println("S3 CLIENT Put_object", bucket)
	return nil
}

func (s *S3_CLIENT) Send_message(QueueUrl, MessageBody string) {}
