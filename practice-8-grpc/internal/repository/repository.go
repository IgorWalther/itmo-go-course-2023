package repository

import (
    "context"
    "fmt"
    "grpc-lection/model"
    "sync"
)

type NoteRepository interface {
    CreateNote(ctx context.Context, note model.Note)
    GetNoteByID(ctx context.Context, noteID int) (model.Note, error)
}

type NoteRepositoryImplementation struct {
    data  map[int]model.Note
    mutex sync.Mutex
}

func NewNoteRepository() *NoteRepositoryImplementation {
    return &NoteRepositoryImplementation{
        data: make(map[int]model.Note),
    }
}

func (n *NoteRepositoryImplementation) CreateNote(ctx context.Context, note model.Note) {
    n.mutex.Lock()
    defer n.mutex.Unlock()

    n.data[note.ID] = note
}

func (n *NoteRepositoryImplementation) GetNoteByID(ctx context.Context, noteID int) (model.Note, error) {
    n.mutex.Lock()
    defer n.mutex.Unlock()

    if data, ok := n.data[noteID]; !ok {
        return model.Note{}, fmt.Errorf("can not find note by id %d", noteID)
    } else {
        return data, nil
    }
}
