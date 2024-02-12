package helpers

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func S3Download(folderName string) {
    // Create a new AWS session
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String("ap-south-1"), // Replace with your S3 bucket region
    })
    if err != nil {
        fmt.Println("Error creating session:", err)
        return
    }

    // Create a new S3 client
    svc := s3.New(sess)

    // Specify the bucket and folder (prefix) you want to download
    bucket := "go-vercel"
    folder := folderName+"/"

    // List objects in the specified bucket and folder
    resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{
        Bucket: aws.String(bucket),
        Prefix: aws.String(folder),
    })
    if err != nil {
        fmt.Println("Error listing objects:", err)
        return
    }

    // Download each object in the folder
    for _, item := range resp.Contents {
        // Generate a file path to save the object locally
		
        filePath := filepath.Join("../output/", *item.Key)

		err := os.MkdirAll(filepath.Dir(filePath), 0755)
        if err != nil {
            fmt.Println("Error creating directory:", err)
            return
        }

        // Create a file to write the object data
        file, err := os.Create(filePath)
        if err != nil {
            fmt.Println("Error creating file:", err)
            return
        }
        defer file.Close()

        // Download the object
		result, err := svc.GetObject(&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    item.Key,
			})
		if err != nil {
			fmt.Println("Error downloading object:", err)
			return
		}
		
		// Write object data to the file
		_, err = io.Copy(file, result.Body)
		if err != nil {
			fmt.Println("Error writing object data:", err)
			return
		}
		

        fmt.Println("Downloaded:", *item.Key)
    }

    fmt.Println("Download completed.")
}
