package pack

import (
	"douyin-micro/cmd/video/dal/db"
	"douyin-micro/kitex_gen/video"
)

// Video pack video info
func Video(m *db.Video) *video.Video {
	if m == nil {
		return nil
	}

	return &video.Video{
		Id:            int64(m.UserId),
		PlayUrl:       m.PlayUrl,
		CoverUrl:      m.CoverUrl,
		FavoriteCount: int64(m.FavoriteCount),
		CommentCount:  int64(m.CommentCount),
		Title:         m.Title,
		CreateTime:    m.CreateTime,
	}
}

// Videos pack list of video info
func Videos(ms []*db.Video) ([]*video.Video, int64) {
	videos := make([]*video.Video, 0)
	for _, m := range ms {
		if v := Video(m); v != nil {
			videos = append(videos, v)
		}
	}
	next_time := videos[len(videos)-1].CreateTime // earliest video timestamp

	// reverse the videos so that the latest comes first
	// for i, j := 0, len(videos)-1; i < j; i, j = i+1, j-1 {
	// 	videos[i], videos[j] = videos[j], videos[i]
	// }

	return videos, next_time
}
