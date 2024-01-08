package import_lib

import (
	"fmt"
	"sync"
	"testing"
	boto3 "thinking-test/practice/import_lib/lib_a"

	"thinking-test/practice/import_lib/config"
)

func TestS3(t *testing.T) {
	config.Config()
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			client := boto3.Client("s3")
			client.Put_object(fmt.Sprintf("buket-%v", i), "key", "body")
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			client := boto3.Client("sqs")
			client.Send_message(fmt.Sprintf("url-%v", i), "body")
		}
	}()

	wg.Wait()
}
