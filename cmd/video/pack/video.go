package pack

import (
	"douyin-micro/cmd/video/dal/db"
	"douyin-micro/kitex_gen/user"
	"douyin-micro/kitex_gen/video"
)

// Video pack video info
func Video(m *db.Video, u *user.User) *video.Video {
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
		Author:        u,
	}
}

// Videos pack list of video info
func Videos(ms []*db.Video, us map[int64]*user.User) []*video.Video {
	videos := make([]*video.Video, 0)
	for _, m := range ms {
		if v := Video(m, us[m.UserId]); v != nil {
			videos = append(videos, v)
		}
	}
	// reverse the videos so that the latest comes first
	// for i, j := 0, len(videos)-1; i < j; i, j = i+1, j-1 {
	// 	videos[i], videos[j] = videos[j], videos[i]
	// }

	return videos
}

func UserIds(ms []*db.Video) []int64 {
	if len(ms) == 0 {
		return []int64{}
	}
	uIds := make([]int64, len(ms))
	uIdMap := make(map[int64]any)
	for _, m := range ms {
		if m != nil {
			uIdMap[m.UserId] = struct{}{}
		}
	}
	for uId := range uIdMap {
		uIds = append(uIds, uId)
	}
	return uIds
}
