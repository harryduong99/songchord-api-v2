package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/harryduong99/songchord-api-v2/graph/generated"
	"github.com/harryduong99/songchord-api-v2/graph/model"
	"github.com/harryduong99/songchord-api-v2/repository/commentRepo"
	"github.com/harryduong99/songchord-api-v2/repository/songRepo"
)

func (r *mutationResolver) CreateSong(ctx context.Context, input model.NewSong) (*model.Song, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateComment(ctx context.Context, comment model.NewComment) (*bool, error) {
	/*
		mutation {
			createComment(comment:{name:"harry", email:"harry.duong99@gmail.com", content: "test comment", songId: "613cfd22fcc103b39ba31ebc"})
		}
	*/
	result, err := commentRepo.CreateComment(context.Background(), comment)

	return &result, err
}

func (r *queryResolver) Songs(ctx context.Context, start int, limit int) ([]*model.Song, error) {
	result, err := songRepo.GetSongList(context.Background(), start, limit)

	return result, err
}

func (r *queryResolver) Song(ctx context.Context, title string) (*model.Song, error) {
	// query getSong($title: String!){
	// 	song(title: $title) {
	// 		content,
	// 		author
	// 	}
	// }

	// {
	// 	"title": "Minhnam"
	// }
	song, err := songRepo.GetSongByName(context.Background(), title)

	return &song, err
}

func (r *queryResolver) SongIds(ctx context.Context) ([]string, error) {
	result, err := songRepo.GetSongIds(context.Background())

	return result, err
}

func (r *queryResolver) SongByID(ctx context.Context, id string) (*model.Song, error) {
	song, err := songRepo.GetSongById(context.Background(), id)
	return &song, err
}

func (r *queryResolver) SongsRecommend(ctx context.Context, id string, number int) ([]*model.Song, error) {
	songs, err := songRepo.GetSongsRecommend(context.Background(), id, number)

	return songs, err
}

func (r *queryResolver) Comment(ctx context.Context, id string) ([]*model.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
