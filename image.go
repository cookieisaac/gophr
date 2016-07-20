package main

import (
	"time"
	"net/http"
	"io"
	"os"
	"mime"
	"mime/multipart"
	"path/filepath"
)

type Image struct {
	ID			string
	UserID		string
	Name		string
	Size		int64
	Location	string
	Description	string
	CreatedAt	time.Time
}

const imageIDLength = 10

func NewImage(user *User) *Image{
	return &Image{
		ID: GenerateID("img", imageIDLength),
		UserID: user.ID,
		CreatedAt: time.Now(),
	}
}

var mimeExtensions = map[string]string{
	"image/png": ".png",
	"image/jpeg": ".jpg",
	"image/gif": ".gif",
}

func (image *Image) CreateFromURL(imageURL string) error {
	response, err := http.Get(imageURL)
	if err != nil {
		return err
	}
	
	if response.StatusCode != http.StatusOK {
		return errImageURLInvalid
	}
	defer response.Body.Close()
	
	mimeType, _, err := mime.ParseMediaType(response.Header.Get("Content-Type"))
	if err != nil {
		return errInvalidImageType
	}
	
	ext, valid := mimeExtensions[mimeType]
	if !valid {
		return errInvalidImageType
	}
	
	image.Name = filepath.Base(imageURL)
	image.Location = image.ID + ext
	
	savedFile, err := os.Create("./data/images/" + image.Location)
	if err != nil {
		return err
	}
	defer savedFile.Close()
	
	size, err := io.Copy(savedFile, response.Body)
	if err != nil {
		return err
	}
	image.Size = size
	
	return globalImageStore.Save(image)
}

func (image *Image) CreateFromFile(file multipart.File, headers *multipart.FileHeader) error {
	image.Name = headers.Filename
	image.Location = image.ID + filepath.Ext(image.Name)
	
	savedFile, err := os.Create("./data/images/"+ image.Location)
	if err != nil {
		return err
	}
	defer savedFile.Close()
	
	size, err := io.Copy(savedFile, file)
	if err != nil {
		return err
	}
	image.Size = size
	
	return globalImageStore.Save(image)
}

func (image *Image) StaticRoute() string {
	return "/im/" + image.Location
}

func (image *Image) ShowRoute() string {
	return "/image/" + image.ID
}