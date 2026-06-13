package utils

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func SaveFile(file *multipart.FileHeader, path string) error {
	srcFile, err := file.Open()
	if err != nil {
		return err
	}
	defer srcFile.Close()
	
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}

	fileName := filepath.Base(file.Filename)
	dstPath := filepath.Join(path, fileName)

	dstFile, err := os.Create(dstPath) 
	if err != nil{
		return err
	}
	defer dstFile.Close()

	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return err
	}
	return nil
}