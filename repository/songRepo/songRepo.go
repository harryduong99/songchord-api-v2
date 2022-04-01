package songRepo

import (
	"context"
	"log"

	"github.com/harryduong99/songchord-api-v2/graph/model"

	"github.com/harryduong99/songchord-api-v2/driver"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetSongByName(ctx context.Context, title string) (model.Song, error) {
	var song model.Song
	data := driver.Mongo.ConnectCollection("song_chords", "songs").FindOne(ctx, bson.M{"title": title})
	error := data.Decode(&song)
	return song, error
}

func GetSongById(ctx context.Context, id string) (model.Song, error) {
	var song model.Song
	data := driver.Mongo.ConnectCollection("song_chords", "songs").FindOne(ctx, bson.M{"id": id})
	error := data.Decode(&song)
	return song, error
}

func GetSongList(ctx context.Context, limit int) (result interface{}) {
	var song model.Song
	var songs []model.Song

	option := options.Find().SetLimit(int64(limit))
	cur, err := driver.Mongo.ConnectCollection("song_chords", "songs").Find(ctx, bson.M{}, option)
	defer cur.Close(ctx)
	if err != nil {
		log.Println(err)
		return nil
	}
	for cur.Next(ctx) {
		cur.Decode(&song)
		songs = append(songs, song)
	}
	return songs
}

func GetSongIds(ctx context.Context) ([]string, error) {
	var song model.Song
	var songIds []string

	option := options.Find()
	cur, err := driver.Mongo.ConnectCollection("song_chords", "songs").Find(ctx, bson.M{}, option)

	defer cur.Close(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for cur.Next(ctx) {
		cur.Decode(&song)
		songIds = append(songIds, song.ID)
	}

	return songIds, nil
}

func InsertSong(ctx context.Context, song model.Song) error {
	_, err := driver.Mongo.ConnectCollection("song_chords", "songs").InsertOne(ctx, song)
	return err
}

func UpdateSong(ctx context.Context, song model.Song) error {
	var songModel model.Song
	col := driver.Mongo.ConnectCollection("song_chords", "songs")
	filter := bson.M{"title": song.Title}

	// get existing comment
	record := col.FindOne(ctx, bson.M{"title": song.Title})
	record.Decode(&songModel)
	songModel.Comment = append(songModel.Comment, song.Comment[0])
	song.Comment = songModel.Comment

	update := bson.M{"$set": song}
	upsertBool := true
	updateOption := options.UpdateOptions{
		Upsert: &upsertBool,
	}
	_, err := driver.Mongo.ConnectCollection("song_chords", "songs").UpdateOne(ctx, filter, update, &updateOption)
	return err
}
func DeleteSong(ctx context.Context, title string) error {
	_, err := driver.Mongo.ConnectCollection("song_chords", "songs").DeleteOne(ctx, bson.M{"title": title})
	return err
}
