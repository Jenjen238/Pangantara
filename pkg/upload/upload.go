package upload

import (
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	MaxFileSize    = 5 * 1024 * 1024 // 5 MB
	UploadDir      = "./uploads"
	DocumentSubDir = "documents"
	ImageSubDir    = "images"
)

var allowedDocumentTypes = map[string]bool{
	".jpg": true, ".jpeg": true, ".png": true,
}

var allowedImageTypes = map[string]bool{
	".jpg": true, ".jpeg": true, ".png": true, ".webp": true,
}

func SaveDocument(file *multipart.FileHeader, folder string) (string, error) {
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedDocumentTypes[ext] {
		return "", errors.New("tipe file tidak diizinkan, gunakan PDF, JPG, atau PNG")
	}
	if file.Size > MaxFileSize {
		return "", fmt.Errorf("ukuran file melebihi batas %d MB", MaxFileSize/(1024*1024))
	}
	return saveFile(file, DocumentSubDir, folder)
}

func SaveImage(file *multipart.FileHeader, folder string) (string, error) {
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedImageTypes[ext] {
		return "", errors.New("tipe file tidak diizinkan, gunakan JPG, PNG, atau WEBP")
	}
	if file.Size > MaxFileSize {
		return "", fmt.Errorf("ukuran file melebihi batas %d MB", MaxFileSize/(1024*1024))
	}
	return saveFile(file, ImageSubDir, folder)
}

func saveFile(file *multipart.FileHeader, subDir, folder string) (string, error) {
	ext := strings.ToLower(filepath.Ext(file.Filename))
	dir := filepath.Join(UploadDir, subDir, folder)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return "", fmt.Errorf("gagal membuat direktori: %v", err)
	}
	filename := fmt.Sprintf("%s_%d%s", uuid.New().String(), time.Now().UnixMilli(), ext)
	fullPath := filepath.Join(dir, filename)

	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("gagal membuka file: %v", err)
	}
	defer src.Close()

	dst, err := os.Create(fullPath)
	if err != nil {
		return "", fmt.Errorf("gagal menyimpan file: %v", err)
	}
	defer dst.Close()

	buf := make([]byte, 32*1024)
	for {
		n, readErr := src.Read(buf)
		if n > 0 {
			if _, writeErr := dst.Write(buf[:n]); writeErr != nil {
				return "", fmt.Errorf("gagal menulis file: %v", writeErr)
			}
		}
		if readErr != nil {
			break
		}
	}

	relativePath := filepath.Join("/uploads", subDir, folder, filename)
	return filepath.ToSlash(relativePath), nil
}

func DeleteFile(relativePath string) error {
	if relativePath == "" {
		return nil
	}
	fullPath := filepath.Join(UploadDir, strings.TrimPrefix(relativePath, "/uploads/"))
	if err := os.Remove(fullPath); err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
}