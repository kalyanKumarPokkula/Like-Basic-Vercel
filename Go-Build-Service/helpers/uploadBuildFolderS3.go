package helpers

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func UploadBuildFolder(folderPath string , userId string) error{
	// Initialize a new AWS session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create an S3 service client
	svc := s3.New(sess)

	// Define the bucket name and folder path
	bucketName := "go-vercel"
	

	// Walk through the folder and its subdirectories
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Open the file
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		// Create the object key by removing the folderPath prefix
		key := filepath.ToSlash(path[len(folderPath)+1:])

		key = "build/"+userId+"/"+key

		// Upload the file to S3
		_, err = svc.PutObject(&s3.PutObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(key),
			Body:   file,
		})
		if err != nil {
			return err
		}

		// fmt.Printf("Uploaded file: %s\n", key)
		return nil
	})
	if err != nil {
		fmt.Println("Error uploading folder to S3:", err)
		return err
	}

	fmt.Println("Folder uploaded successfully!")

	return nil
}