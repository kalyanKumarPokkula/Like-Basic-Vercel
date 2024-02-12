package helpers

// import (
// 	"context"
// 	"fmt"

// 	"os"
// 	"path/filepath"

// 	"github.com/aws/aws-sdk-go-v2/aws"
// 	"github.com/aws/aws-sdk-go-v2/service/s3"
// )

// type AWS_Service struct{
// 	S3_Client *s3.Client
// }

// func uploadFileToS3(s3Client AWS_Service, bucketName, filePath string , fileName string) error {
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	_, err = s3Client.S3_Client.PutObject(context.TODO(), &s3.PutObjectInput{
// 		Bucket: aws.String(bucketName),
// 		Key:    aws.String(fileName), // Use the file name as the S3 object key
// 		Body:   file,
// 	})
// 	return err
// }

// func UploadFolderToS3(s3Client AWS_Service, bucketName string, folderPath string, folderName string) error {
// 	files, err := os.ReadDir(folderPath)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println(files)

// 	for _, file := range files {
// 		filePath := filepath.Join(folderPath, file.Name())
// 		fmt.Println(filePath)
// 		fileName := folderName + "/" +file.Name()

// 		if !file.IsDir() {
// 			err := uploadFileToS3(s3Client, bucketName, filePath,fileName )
// 			if err != nil {
// 				return err
// 			}
// 			fmt.Printf("Uploaded file: %s\n", filePath)
// 		}
// 	}

// 	return nil
// }