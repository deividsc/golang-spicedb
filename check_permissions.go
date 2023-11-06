package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/authzed/authzed-go/v1"
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

	ctx := context.Background()

	emilia := &pb.SubjectReference{Object: &pb.ObjectReference{
		ObjectType: "blog/user",
		ObjectId:   "emilia",
	}}

	beatrice := &pb.SubjectReference{Object: &pb.ObjectReference{
		ObjectType: "blog/user",
		ObjectId:   "beatrice",
	}}

	firstPost := &pb.ObjectReference{
		ObjectType: "blog/post",
		ObjectId:   "1",
	}

	resp, err := client.CheckPermission(ctx, &pb.CheckPermissionRequest{
		Resource:   firstPost,
		Permission: "read",
		Subject:    emilia,
	})
	if err != nil {
		log.Fatalf("failed to check permission: %s", err)
	}
	fmt.Printf("emilia has permission to read the first post: %t\n", resp.Permissionship == pb.CheckPermissionResponse_PERMISSIONSHIP_HAS_PERMISSION)

	resp, err = client.CheckPermission(ctx, &pb.CheckPermissionRequest{
		Resource:   firstPost,
		Permission: "write",
		Subject:    emilia,
	})
	if err != nil {
		log.Fatalf("failed to check permission: %s", err)
	}
	fmt.Printf("emilia has permission to write on first post: %t\n", resp.Permissionship == pb.CheckPermissionResponse_PERMISSIONSHIP_HAS_PERMISSION)

	resp, err = client.CheckPermission(ctx, &pb.CheckPermissionRequest{
		Resource:   firstPost,
		Permission: "read",
		Subject:    beatrice,
	})
	if err != nil {
		log.Fatalf("failed to check permission: %s", err)
	}
	fmt.Printf("beatrice has permission to read the first post: %t\n", resp.Permissionship == pb.CheckPermissionResponse_PERMISSIONSHIP_HAS_PERMISSION)

	resp, err = client.CheckPermission(ctx, &pb.CheckPermissionRequest{
		Resource:   firstPost,
		Permission: "write",
		Subject:    beatrice,
	})
	if err != nil {
		log.Fatalf("failed to check permission: %s", err)
	}
	fmt.Printf("beatrice has permission to write on first post: %t\n", resp.Permissionship == pb.CheckPermissionResponse_PERMISSIONSHIP_HAS_PERMISSION)

}
