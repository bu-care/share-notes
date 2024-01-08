package boto3

import (
	"strings"
	// _ "thinking-test/practice/import_lib/lib_b"
)

var (
	G_SNS_CLIENT    ClientSimulatorI
	G_S3_CLIENT     ClientSimulatorI
	G_SQS_CLIENT    ClientSimulatorI
	G_LAMBDA_CLINET ClientSimulatorI
	G_STS_CLIENT    ClientSimulatorI
)

type S3_interface interface {
	Put_object(bucket, key string, body string) any
}

type Sqs_interface interface {
	Send_message(QueueUrl, MessageBody string)
}

type ClientSimulatorI interface {
	Put_object(bucket, key string, body string) any
	Send_message(QueueUrl, MessageBody string)
}

func Client(key string) ClientSimulatorI {
	switch strings.ToLower(key) {
	case "sns":
		return G_SNS_CLIENT
	case "s3":
		return G_S3_CLIENT
	case "sqs":
		return G_SQS_CLIENT
	case "lambda":
		return G_LAMBDA_CLINET
	}

	return G_STS_CLIENT
}

func Set_client(client_type string, client ClientSimulatorI) {
	switch client_type {
	case "sns":
		G_SNS_CLIENT = client
	case "s3":
		G_S3_CLIENT = client
	case "sqs":
		G_SQS_CLIENT = client
	case "lambda":
		G_LAMBDA_CLINET = client
	}
}

func Set_client01(sns, s3, sqs, lambda, sts ClientSimulatorI) {
	G_SNS_CLIENT = sns
	G_S3_CLIENT = s3
	G_SQS_CLIENT = sqs
	G_LAMBDA_CLINET = lambda

	G_STS_CLIENT = sts
}
