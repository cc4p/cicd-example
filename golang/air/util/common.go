package util

import (
	"bytes"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"log"
	"time"
)

func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}



func PrintJson(title string, body []byte) {
	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, body, "", "\t")
	if error != nil {
		log.Println("JSON parse error: ", error)
		return
	}
	log.Println(title, prettyJSON.String())

}

func PutMetric2CW(metric *cloudwatch.PutMetricDataInput) {

	conf := aws.Config{
		Region:      aws.String("us-east-1"),
		DisableSSL: 	aws.Bool(true),
		//Credentials: credentials.NewSharedCredentials("~/.aws/credentials", "default"),
	}
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: conf,
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := cloudwatch.New(sess)
	_, err := svc.PutMetricData(metric)
	if err != nil {
		log.Println("Error adding metrics:", err.Error())
		return
	}

}
//AWS_CA_BUNDLE=$HOME/my_custom_ca_bundle
/*
&cloudwatch.PutMetricDataInput{
		Namespace: aws.String("DynamoDBFoo/Traffic"),
		MetricData: []*cloudwatch.MetricDatum{
			&cloudwatch.MetricDatum{
				MetricName: aws.String("ElapsedWithTrans"),
				Unit:       aws.String("Seconds"),
				Value:      aws.Float64(val),
				Dimensions: []*cloudwatch.Dimension{
					&cloudwatch.Dimension{
						Name:  aws.String("ServiceName"),
						Value: aws.String("AmazonDynamoDB"),
					},
					&cloudwatch.Dimension{
						Name:  aws.String("Method"),
						Value: aws.String("TransactWriteItems"),
					},
				},
			},
		},
	}
*/