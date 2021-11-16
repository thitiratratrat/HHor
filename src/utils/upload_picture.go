package utils

import (
	"io"
	"mime/multipart"
	"net/http"
	"net/url"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
	"google.golang.org/appengine"
)

func UploadPicture(file multipart.File, foldername string, filename string, request *http.Request) (string, error) {
	bucket := "hhor_pictures"

	var err error

	ctx := appengine.NewContext(request)

	storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile("keys.json"))

	if err != nil {
		return "", err
	}

	defer file.Close()

	storageWriter := storageClient.Bucket(bucket).Object(foldername + filename).NewWriter(ctx)

	if _, err := io.Copy(storageWriter, file); err != nil {
		return "", err
	}

	if err := storageWriter.Close(); err != nil {
		return "", err
	}

	uploaded, err := url.Parse("/" + bucket + "/" + storageWriter.Attrs().Name)

	if err != nil {
		return "", err
	}

	return "https://storage.googleapis.com" + uploaded.Path, nil
}
