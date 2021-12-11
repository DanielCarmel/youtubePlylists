package main

import (
	api "github.com/DanielCarmel/youtubePlylists"
	"google.golang.org/api/youtube/v3"
	"log"
)

const MaxResults = 5
const MyRating = "like"

func main() {
	client := api.GetClient()
	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating YouTube client: %v", err)
	}

	var likeVideos []*youtube.Video
	youtubeItemPart := []string{"topicDetails", "id"}
	response, err := service.Videos.List(youtubeItemPart).MaxResults(MaxResults).MyRating(MyRating).Do()
	likeVideos = append(likeVideos, response.Items...)

	for response.NextPageToken != "" {
		response, err := service.Videos.List(youtubeItemPart).MaxResults(MaxResults).MyRating(MyRating).PageToken(response.NextPageToken).Do()
		if err != nil {
			log.Fatalf("Error listing videos: %v", err)
		}
		likeVideos = append(likeVideos, response.Items...)
	}
	// Iterate through each item and add it to the correct list.
	for _, item := range response.Items {
		if item.Snippet.CategoryId == "10" { // Youtube`s id for music videos
			log.Println(item)
		} else {
			log.Println("item has ", item.Snippet.CategoryId)
			continue
		}
	}

}
