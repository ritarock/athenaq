package aws

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/athena"
)

func Session(profile string) *session.Session {
	return session.Must(session.NewSessionWithOptions(session.Options{
		Profile:           profile,
		SharedConfigState: session.SharedConfigEnable,
	}))
}

func AthenaQuery(profile, bucket, key, query string) {
	sess := Session(profile)
	athenaClient := athena.New(sess)

	resultConf := &athena.ResultConfiguration{}
	resultConf.SetOutputLocation("s3://" + bucket + key)
	input := &athena.StartQueryExecutionInput{
		QueryString:         &query,
		ResultConfiguration: resultConf,
	}
	output, err := athenaClient.StartQueryExecution(input)
	if err != nil {
		log.Fatal(err)
	}

	queryId := &athena.GetQueryExecutionInput{
		QueryExecutionId: output.QueryExecutionId,
	}

	WaitDone(queryId, athenaClient, output)
	DownloadResult(sess, output, bucket, key)
}

func WaitDone(queryId *athena.GetQueryExecutionInput, athenaClient *athena.Athena, output *athena.StartQueryExecutionOutput) (id *string, err error) {
	for {
		executionOutput, err := athenaClient.GetQueryExecution(queryId)
		if err != nil {
			return nil, err
		}
		switch *executionOutput.QueryExecution.Status.State {
		case athena.QueryExecutionStateQueued, athena.QueryExecutionStateRunning:
			time.Sleep(5 * time.Second)
		case athena.QueryExecutionStateSucceeded:
			return output.QueryExecutionId, nil
		default:
			return nil, errors.New(executionOutput.String())
		}
	}
}

func DownloadResult(sess *session.Session, output *athena.StartQueryExecutionOutput, bucket, key string) {
	downloader := s3manager.NewDownloader(sess)
	f, err := os.Create(*output.QueryExecutionId + ".csv")
	dl, err := downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result file %v [bytes]", dl)
}
