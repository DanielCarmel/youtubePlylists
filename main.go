package main

import (
	api "github.com/DanielCarmel/youtubePlylists"
	"google.golang.org/api/youtube/v3"
	"log"
)

const MaxResults = 5
const MyRating = "like"
const VideoCategoryId = "10" // Youtube api id for music videos

func main() {
	// Create client
	client := api.GetClient()
	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating YouTube client: %v", err)
	}

	likedVideos, err := GetLikedVideos(service)
	if err != nil {
		log.Fatalf("Error listing videos: %v", err)
	}

	// Iterate through each item and add it to the correct list.
	for _, item := range likedVideos {
		println(item)
	}
}

func GetLikedVideos(service *youtube.Service) ([]*youtube.Video, error) {
	//var isItStillGoing = true
	var videosList []*youtube.Video
	youtubeItemPart := []string{"id", "snippet", "topicDetails"}

	response, err := service.Videos.List(youtubeItemPart).MaxResults(MaxResults).VideoCategoryId(VideoCategoryId).MyRating(MyRating).Do()
	if err != nil {
		return nil, err
	}

	//NextPageToken := response.NextPageToken
	videosList = append(videosList, response.Items...)

	/*
		for isItStillGoing {
			response, err := service.Videos.List(youtubeItemPart).MaxResults(MaxResults).VideoCategoryId(VideoCategoryId).MyRating(MyRating).PageToken(NextPageToken).Do()
			if err != nil {
				return nil, err
			}
			videosList = append(videosList, response.Items...)

			if response.NextPageToken == "" {
				isItStillGoing = false
			} else {
				NextPageToken = response.NextPageToken
			}
		}
	*/
	return videosList, nil
}
