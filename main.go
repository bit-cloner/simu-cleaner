package main

import (
	"fmt"
	"os"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	regions := []string{
		"ap-south-1",
		"eu-west-3",
		"eu-north-1",
		"eu-west-2",
		"eu-west-1",
		"ap-northeast-3",
		"ap-northeast-2",
		"ap-northeast-1",
		"sa-east-1",
		"ca-central-1",
		"ap-southeast-1",
		"ap-southeast-2",
		"eu-central-1",
		"us-east-1",
		"us-east-2",
		"us-west-1",
		"us-west-2",
		"cn-north-1",
		"cn-northwest-1",
	}
	region := ""
	prompt := &survey.Select{
		Message: "Select one of the follwing AWS Regions",
		Options: regions,
	}
	survey.AskOne(prompt, &region, nil)
	//fmt.Println(region)
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		fmt.Println("Error creating session", err)
		os.Exit(1)
	}
	// Create an S3 service client
	svc := s3.New(sess)

	// List all S3 buckets
	var allbuckets = []string{}
	var imusize int64
	var totalcost float64
	result, err := svc.ListBuckets(nil)
	if err != nil {
		fmt.Println("Error listing buckets", err)
		os.Exit(1)
	}
	//fmt.Println("S3 Buckets:")
	for _, bucket := range result.Buckets {
		//append bucket name to allbuckets
		allbuckets = append(allbuckets, *bucket.Name)

	}
	//print allbuckets
	//fmt.Println(allbuckets)
	//for each in allbuckets
	for _, bucket := range allbuckets {
		//list all multipart uploads
		result, err := svc.ListMultipartUploads(&s3.ListMultipartUploadsInput{
			Bucket: aws.String(bucket),
		})
		if err != nil {
			if aerr, ok := err.(s3.RequestFailure); ok && aerr.StatusCode() == 301 && aerr.Code() == "BucketRegionError" {
				//if bucket is in another region, skip it
				fmt.Println("Bucket", bucket, "is in another region, skipping")
				continue
			} else {
				fmt.Println("Error listing multipart uploads", err)
				continue
			}
		}
		//for each multipart upload
		for _, upload := range result.Uploads {
			// list all parts
			result, err := svc.ListParts(&s3.ListPartsInput{
				Bucket:   aws.String(bucket),
				Key:      upload.Key,
				UploadId: upload.UploadId,
			})
			if err != nil {
				fmt.Println("Error listing parts", err)
				continue
			}
			//for each part
			for _, part := range result.Parts {
				//add size to total
				imusize += *part.Size
			}
		}
		if imusize/1024/1024 > 0 { // Only take action if there are incomplete multipart uploads
			fmt.Println("\nTotal size of all incomplete multipart uploads in bucket", bucket, "is", imusize/1024/1024, "MB")
			//convert int64 to float64
			var floatimusize = float64(imusize) * 1.0
			// truncateflot to 2 decimal places
			var cost = fmt.Sprintf("%.2f", floatimusize/1024/1024/1024*0.023)
			fmt.Println("Approximately this is incuring ", cost, "USD per month")
			var action = "skip"
			prompt := &survey.Select{
				Message: "\nSelect one of the follwing actions",
				Options: []string{"skip", "clean"},
			}
			survey.AskOne(prompt, &action, nil)
			if action == "clean" {
				//for each multipart upload
				for _, upload := range result.Uploads {
					// list all parts
					result, err := svc.ListParts(&s3.ListPartsInput{
						Bucket:   aws.String(bucket),
						Key:      upload.Key,
						UploadId: upload.UploadId,
					})
					if err != nil {
						fmt.Println("Error listing parts", err)
						continue
					}
					//for each part
					for _, part := range result.Parts {
						//abort uploads that are older than 24 hours
						oneDayAgo := time.Now().Add(-24 * time.Hour)
						if part.LastModified.Before(oneDayAgo) {
							_, err := svc.AbortMultipartUpload(&s3.AbortMultipartUploadInput{
								Bucket:   aws.String(bucket),
								Key:      upload.Key,
								UploadId: upload.UploadId,
							})
							if err != nil {
								fmt.Println("Error aborting multipart upload", err)
								continue
							}
							fmt.Println("deleted multipart upload of ", *part.Size/1024/1024, "MB from bucket", bucket)
							var savedcost = float64(*part.Size) / 1024 / 1024 / 1024 * 0.023
							totalcost += savedcost
						}
					}
				}
			}
		} else {
			continue
		}
	}
	fmt.Println("\nTotal cost saved is approximately", fmt.Sprintf("%.2f", totalcost), "USD per month")
}
