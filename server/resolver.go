package main

import (
	"context"

	prisma "github.com/prismaTutorial/generated/prisma-client"
)

type Resolver struct {
	Prisma *prisma.Client
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Post() PostResolver {
	return &postResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) User() UserResolver {
	return &userResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, name string) (*prisma.User, error) {
	return r.Prisma.CreateUser(prisma.UserCreateInput{
		Name: name,
	}).Exec(ctx)
}
func (r *mutationResolver) CreateDraft(ctx context.Context, title string, userId string) (*prisma.Post, error) {
	return r.Prisma.CreatePost(prisma.PostCreateInput{
		Title: title,
		Author: &prisma.UserCreateOneWithoutPostsInput{
			Connect: &prisma.UserWhereUniqueInput{ID: &userId},
		},
	}).Exec(ctx)
}
func (r *mutationResolver) Publish(ctx context.Context, postId string) (*prisma.Post, error) {
	published := true
	return r.Prisma.UpdatePost(prisma.PostUpdateParams{
		Where: prisma.PostWhereUniqueInput{ID: &postId},
		Data:  prisma.PostUpdateInput{Published: &published},
	}).Exec(ctx)
}

type postResolver struct{ *Resolver }

func (r *postResolver) Author(ctx context.Context, obj *prisma.Post) (*prisma.User, error) {
	return r.Prisma.Post(prisma.PostWhereUniqueInput{ID: &obj.ID}).Author().Exec(ctx)
}

type queryResolver struct{ *Resolver }

//func (r *queryResolver) PublishedPosts(ctx context.Context) ([]prisma.Post, error) {
//	published := true
//	return r.Prisma.Posts(&prisma.PostsParams{
//		Where: &prisma.PostWhereInput{Published: &published},
//	}).Exec(ctx)
//}

// convert to an array of pointers
func (r *queryResolver) PublishedPosts(ctx context.Context) ([]*prisma.Post, error) {
	published := true
	posts, err := r.Prisma.Posts(&prisma.PostsParams{
		Where: &prisma.PostWhereInput{Published: &published},
	}).Exec(ctx)

	if err != nil {
		panic(err)

	}

	postsPointers := make([]*prisma.Post, 0)
	for _, p := range posts {
		postsPointers = append(postsPointers, &p)

	}

	return postsPointers, err
}

func (r *queryResolver) Post(ctx context.Context, postId string) (*prisma.Post, error) {
	return r.Prisma.Post(prisma.PostWhereUniqueInput{ID: &postId}).Exec(ctx)
}
func (r *queryResolver) PostsByUser(ctx context.Context, userId string) ([]prisma.Post, error) {
	return r.Prisma.Posts(&prisma.PostsParams{
		Where: &prisma.PostWhereInput{
			Author: &prisma.UserWhereInput{
				ID: &userId,
			}},
	}).Exec(ctx)
}

type userResolver struct{ *Resolver }

func (r *userResolver) Posts(ctx context.Context, obj *prisma.User) ([]prisma.Post, error) {
	return r.Prisma.User(prisma.UserWhereUniqueInput{ID: &obj.ID}).Posts(nil).Exec(ctx)
}
