package easyaws

import (
	"bytes"
	"os"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/pkg/errors"
)

func FilesFromS3(localDir string, bucket string, keys []string, awsSession *session.Session) error {
	if _, err := os.Stat(localDir); os.IsNotExist(err) {
		err = os.MkdirAll(localDir, os.ModePerm)
		if err != nil {
			return errors.Wrap(err, "Error in os.MkdirAll()")
		}
	}
	s3Downloader := s3manager.NewDownloader(awsSession)
	for _, key := range keys {
		f, err := os.Create(filepath.Join(localDir, filepath.Base(key)))
		if err != nil {
			return errors.Wrap(err, "Error in os.Create()")
		}
		input := &s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
		}
		_, err = s3Downloader.Download(f, input)
		if err != nil {
			return errors.Wrap(err, "Error in s3Downloader.Download()")
		}
	}
	return nil
}

func BytesFromS3(bucket string, keys []string, awsSession *session.Session) (map[string][]byte, error) {
	var byteMap = make(map[string][]byte, len(keys))
	s3Downloader := s3manager.NewDownloader(awsSession)
	for _, key := range keys {
		b := &aws.WriteAtBuffer{}
		input := &s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
		}
		_, err := s3Downloader.Download(b, input)
		if err != nil {
			return nil, errors.Wrap(err, "Error in s3Downloader.Download()")
		}
		byteMap[key] = b.Bytes()
	}
	return byteMap, nil
}

func FilesToS3(files []string, bucket string, prefix string, awsSession *session.Session) error {
	s3Uploader := s3manager.NewUploader(awsSession)
	for _, file := range files {
		reader, err := os.Open(file)
		if err != nil {
			return errors.Wrap(err, "Error in os.Open()")
		}
		fi, err := reader.Stat()
		if err != nil {
			return errors.Wrap(err, "Error in reader.Stat()")
		}
		key := prefix + fi.Name()
		defer reader.Close()
		input := &s3manager.UploadInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
			Body:   reader,
		}
		_, err = s3Uploader.Upload(input)
		if err != nil {
			return errors.Wrap(err, "Error in s3Uploader.Upload()")
		}
	}
	return nil
}

func BytesToS3(byteMap map[string][]byte, bucket string, prefix string, awsSession *session.Session) error {
	s3Uploader := s3manager.NewUploader(awsSession)
	for k, v := range byteMap {
		key := filepath.Join(prefix, k)
		input := &s3manager.UploadInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
			Body:   bytes.NewReader(v),
		}
		_, err := s3Uploader.Upload(input)
		if err != nil {
			return errors.Wrap(err, "Error in s3Uploader.Upload()")
		}
	}
	return nil
}

func DeleteFromS3(bucket string, keys []string, awsSession *session.Session) error {
	svc := s3.New(awsSession)
	// svc := s3.New(awsSession, awsConfig)
	objects := []*s3.ObjectIdentifier{}
	for _, key := range keys {
		o := s3.ObjectIdentifier{
			Key:       aws.String(key), // Required
			// VersionId: easyaws.String("ObjectVersionId"),
		}
		objects = append(objects, &o)
	}
	params := &s3.DeleteObjectsInput{
		Bucket: aws.String(bucket),
		Delete: &s3.Delete{
			Objects: objects,
			Quiet: aws.Bool(true),
		},
	}
	_, err := svc.DeleteObjects(params)
	if err != nil {
		return errors.Wrap(err, "Error in S3 DeleteObjects().")
	}
	return nil
}


func ListS3Keys(bucket string, prefix string, awsSession *session.Session) ([]string, error) {
	svc := s3.New(awsSession)
	params := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket), // Required
		Prefix: aws.String(prefix),
	}
	output, err := svc.ListObjectsV2(params)
	if err != nil {
		return nil, errors.Wrap(err, "Error in s3.ListObjectsV2()")
	}
	count := 0
	for _, o := range output.Contents {
		if *o.Key != prefix {
			count++
		}
	}
	keys := make([]string, count)
	i := 0
	for _, o := range output.Contents {
		if *o.Key != prefix {
			keys[i] = *o.Key
			i++
		}
	}
	return keys, nil
}

func PresignedTempS3Url(bucket string, key string, secondsTillExpiration int, awsSession *session.Session) (string, error) {

	//put := S3PutRequest{}
	svc := s3.New(awsSession)
	por, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	url, err := por.Presign(time.Duration(secondsTillExpiration) * time.Second)
	if err != nil {
		return "", errors.Wrap(err, "Error in PutObjectRequest.Presign()")
	}
	return url, nil
}
