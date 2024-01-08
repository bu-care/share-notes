package clients

import (
	"fmt"
	boto3 "thinking-test/practice/import_lib/lib_a"
)

type SQS_CLIENT struct{}

var G_SQS_CLIENT *SQS_CLIENT

func init() {
	G_SQS_CLIENT = &SQS_CLIENT{}
	boto3.Set_client("sqs", G_SQS_CLIENT)
}

func (s *SQS_CLIENT) Put_object(bucket, key string, body string) any {
	fmt.Println("SQS_CLIENT Put_object")
	return nil
}
func (s *SQS_CLIENT) Send_message(QueueUrl, MessageBody string) {
	fmt.Println("SQS_CLIENT Send_message", QueueUrl)
}
