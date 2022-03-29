package commentRepo

import (
	"context"

	"github.com/harryduong99/songchord-api-v2/graph/model"

	"github.com/harryduong99/songchord-api-v2/config"
	"github.com/harryduong99/songchord-api-v2/driver"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateComment(ctx context.Context, comment model.NewComment) (bool, error) {
	var song model.Song

	objID, err := primitive.ObjectIDFromHex(comment.SongID)
	if err != nil {
		panic(err)
	}
	collection := driver.Mongo.ConnectCollection(config.DB_NAME, "songs")

	songRecord := driver.Mongo.ConnectCollection("song_chords", "songs").FindOne(ctx, bson.M{"_id": objID})
	songRecord.Decode(&song)

	newComment := &model.Comment{
		Name:    comment.Name,
		Email:   comment.Email,
		Content: comment.Content,
	}
	song.Comment = append(song.Comment, newComment)

	filter := bson.M{"_id": bson.M{"$eq": objID}}
	update := bson.M{"$set": bson.M{"comments": song.Comment}}

	_, error := collection.UpdateOne(
		ctx,
		filter,
		update,
	)

	return error == nil, error
}
