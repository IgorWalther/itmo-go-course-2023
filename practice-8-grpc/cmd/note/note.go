package main

import (
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "grpc-lection/internal/api"
    "grpc-lection/internal/repository"
    "grpc-lection/pkg/generated/proto/note"
    "log"
    "net"
)

func main() {
    noteRepository := repository.NewNoteRepository()
    noteApi := api.NewNoteApi(noteRepository)

    lsn, err := net.Listen("tcp", ":8082")

    if err != nil {
        panic(err)
    }

    var noteServer = grpc.NewServer()

    reflection.Register(noteServer)
    note.RegisterNoteManagerServer(noteServer, noteApi)

    if err := noteServer.Serve(lsn); err != nil {
        log.Fatalf("serving error %v", err)
    }
}
