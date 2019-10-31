package main

import (
	"context"
	"fmt"

	prisma "github.com/prismaTutorial/generated/prisma-client"
)

func main() {
	client := prisma.New(nil)
	ctx := context.TODO()

	// Create a new user with two posts
	//name := "Bob"
	//email := "bob@prisma.io"
	//title1 := "Join us for GraphQL Conf in 2019"
	//title2 := "Subscribe to GraphQL Weekly for GraphQL news"
	//newUser, err := client.CreateUser(prisma.UserCreateInput{
	//	Name:  name,
	//	Email: &email,
	//	Posts: &prisma.PostCreateManyWithoutAuthorInput{
	//		Create: []prisma.PostCreateWithoutAuthorInput{
	//			prisma.PostCreateWithoutAuthorInput{
	//				Title: title1,
	//			},
	//			prisma.PostCreateWithoutAuthorInput{
	//				Title: title2,
	//			},
	//		},
	//	},
	//}).Exec(ctx)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Created new user: %+v\n", newUser)

	//allUsers, err := client.Users(nil).Exec(ctx)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%+v\n", allUsers)

	//allPosts, err := client.Posts(nil).Exec(ctx)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%+v\n", allPosts)

	// Query data
	email := "bob@prisma.io"
	postsByUser, err := client.User(prisma.UserWhereUniqueInput{
		Email: &email,
	}).Posts(nil).Exec(ctx)

	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", postsByUser)
}
