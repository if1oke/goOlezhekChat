package storage

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/if1oke/storage/internal/file"
)

type Storage struct {
	// Имена с маленькой буквы не экспортируются
	filesBackup map[uuid.UUID]*file.File
	Files       map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{
		Files: make(map[uuid.UUID]*file.File),
	}
}

func (s *Storage) GetByID(fileID uuid.UUID) (*file.File, error) {
	foundFile, ok := s.Files[fileID]
	if !ok {
		return nil, fmt.Errorf("file %v not found: %v", fileID, ok)
	}
	return foundFile, nil
}

func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	newFile, err := file.NewFile(filename, blob)
	if err != nil {
		return nil, err
	}

	s.Files[newFile.ID] = newFile

	return newFile, err
}
