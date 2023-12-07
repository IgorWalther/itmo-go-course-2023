package api

import (
    "context"
    "fmt"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "grpc-lection/internal/repository"
    "grpc-lection/model"
    "grpc-lection/pkg/generated/proto/note"
    "math/rand"
)

type NoteApi struct {
    note.UnimplementedNoteManagerServer
    noteRepository repository.NoteRepository
}

func NewNoteApi(noteRepository repository.NoteRepository) *NoteApi {
    return &NoteApi{
        noteRepository: noteRepository,
    }
}

func (n *NoteApi) CreateNote(ctx context.Context, request *note.CreateNoteRequest) (*note.CreateNoteResponse, error) {
    noteID := rand.Int()

    n.noteRepository.CreateNote(ctx, model.Note{
        ID:     noteID,
        UserID: int(request.UserId),
        Info:   request.Note.Data,
    })

    return &note.CreateNoteResponse{
        NoteId: int64(noteID),
    }, nil
}

func (n *NoteApi) GetAllNotes(ctx context.Context, request *note.GetAllNotesRequest) (*note.GetAllNotesResponse, error) {
    panic("not implemented")
}

func (n *NoteApi) GetNoteByID(ctx context.Context, request *note.GetNoteByIDRequest) (*note.GetNoteByIDResponse, error) {
    resultNote, err := n.noteRepository.GetNoteByID(ctx, int(request.NoteId))

    if err != nil {
        return nil, status.Error(codes.NotFound, fmt.Sprintf("repository error %v", err))
    }

    return &note.GetNoteByIDResponse{
        UserId: int64(resultNote.UserID),
        Note: &note.Note{
            Data: resultNote.Info,
        },
    }, nil
}

func (n *NoteApi) DeleteNoteByID(ctx context.Context, request *note.DeleteNoteByIDRequest) (*note.DeleteNoteByIDResponse, error) {
    panic("not implemented")
}
