package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/authzed/authzed-go/proto/authzed/api/v1"
	authzed "github.com/authzed/authzed-go/v1"
	"github.com/authzed/grpcutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	client, err := authzed.NewClient(
		"localhost:50051",
		grpcutil.WithInsecureBearerToken("somerandomkeyhere"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("unable to initialize client: %s", err)
	}

	request := &pb.WriteRelationshipsRequest{Updates: []*pb.RelationshipUpdate{
		{ // Emilia is a Writer on Post 1
			Operation: pb.RelationshipUpdate_OPERATION_CREATE,
			Relationship: &pb.Relationship{
				Resource: &pb.ObjectReference{
					ObjectType: "blog/post",
					ObjectId:   "1",
				},
				Relation: "writer",
				Subject: &pb.SubjectReference{
					Object: &pb.ObjectReference{
						ObjectType: "blog/user",
						ObjectId:   "emilia",
					},
				},
			},
		},
		{ // Beatrice is a Reader on Post 1
			Operation: pb.RelationshipUpdate_OPERATION_CREATE,
			Relationship: &pb.Relationship{
				Resource: &pb.ObjectReference{
					ObjectType: "blog/post",
					ObjectId:   "1",
				},
				Relation: "reader",
				Subject: &pb.SubjectReference{
					Object: &pb.ObjectReference{
						ObjectType: "blog/user",
						ObjectId:   "beatrice",
					},
				},
			},
		},
	}}

	_, err = client.WriteRelationships(context.Background(), request)
	if err != nil {
		log.Fatalf("failed to write relations: %s", err)
	}

	fmt.Println("Added emilia as a writer")
	fmt.Println("Added beatrice as a reader")
}
