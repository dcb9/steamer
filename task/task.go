package task

import (
	"log"

	"github.com/dcb9/steamer/task/common"
	"github.com/dcb9/steamer/task/youku"
	"github.com/dcb9/steamer/task/youtube"
)

func FetchResourceInfo(url string) common.ResourceInfo {
	var resourceInfo common.ResourceInfo
	if youku.IsYouku(url) {
		resourceInfo = youku.FetchInfo(url)
	} else if youtube.IsYoutube(url) {
		resourceInfo = youtube.FetchInfo(url)
	} else {
		log.Fatal("Known platform url: " + url)
	}

	return resourceInfo
}

func AddDownloadTask(url string, id int64, sIndex string) common.DownloadTask {
	var infoR common.ResourceInfoRepository
	var entity common.ResourceInfo

	if youku.IsYouku(url) {
		infoR.CustomStreamsFunc = youku.GenerateStreams
	} else if youtube.IsYoutube(url) {
		infoR.CustomStreamsFunc = youtube.GenerateStreams
	} else {
		log.Fatal("Known platform url: " + url)
	}

	if id != 0 {
		entity = infoR.Entity(id)
	} else {
		entity = infoR.EntityByUrl(url)
	}
	if entity.Id == 0 {
		return common.DownloadTask{}
	}

	var taskR common.DownloadTaskRepository
	return taskR.Add(entity, sIndex)
}
