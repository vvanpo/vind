package vind

import (
	"io"

	"github.com/google/uuid"
	"github.com/vvanpo/vind/internal/state"
)

// Filesystem ...
type Filesystem struct {
	storage Storage
}

func Load(s Storage) (Filesystem, error) {
	if db, err := s.DB(); err != nil {
		return Filesystem{}, err
	} else {
		state.Init(db)
		db.Close()
	}

	return Filesystem{s}, nil
}

func (fs Filesystem) Add(content io.Reader) error {
	id := uuid.New()

	if err := fs.storage.Add(id.String(), content); err != nil {
		return err
	}

	db, err := fs.storage.DB()

	if err != nil {
		fs.storage.Delete(id.String())

		return err
	}

	return state.AddFile(db, id)
}

func (fs Filesystem) Select(filter Filter, sort Sort) (<-chan File, error) {
	out := make(chan File)

	close(out)

	return out, nil
}

type Filter struct{}

type Sort struct{}

type File struct{}

// Property ...
func (f File) Property(group, name string, params ...any) (any, error) {
	return nil, nil
}
