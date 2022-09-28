package filemanager

import (
	"bookhub/internal/entity"
	"context"
	"errors"
	"os"
)

type FileManager struct {
	root string
}

func NewFileManager(root string) *FileManager {
	return &FileManager{
		root: root,
	}
}

func (fm *FileManager) CreateFile(ctx context.Context, file entity.File) (path string, err error) {
	if len(file.File) == 0 {
		return "", entity.ErrEmptyFile
	}
	if file.Name == "" {
		return "", errors.New("file name is empty")
	}

	switch file.Type {
	case entity.Image:
		path += fm.root + "image/" + file.Name + ".jpg"
	case entity.PDF:
		path += fm.root + "pdf/" + file.Name + ".pdf"
	}

	if _, err := os.Stat(path); err == nil {
		return path, entity.ErrFileAlreadyExists
	}
	if err := os.WriteFile(path, file.File, 0666); err != nil {
		return "", err
	}
	return path, nil
}
func (fm *FileManager) GetFile(ctx context.Context, path string) (file entity.File, err error) {
	file.File, err = os.ReadFile(path)
	if err != nil {
		return entity.File{}, err
	}
	return file, nil
}
func (fm *FileManager) UpdateFile(ctx context.Context, file entity.File) (err error) {
	if len(file.File) == 0 {
		return entity.ErrEmptyFile
	}
	if err = os.WriteFile(file.Path, file.File, 0666); err != nil {
		return err
	}
	return nil
}
func (fm *FileManager) DeleteFile(ctx context.Context, path string) (err error) {
	if err = os.Remove(path); err != nil {
		return err
	}
	return nil
}
