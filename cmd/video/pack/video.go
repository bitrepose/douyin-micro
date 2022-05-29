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
		IsFavorite:    false,
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
	return videos
}

func UserIds(ms []*db.Video) []int64 {
	if len(ms) == 0 {
		return []int64{}
	}
	uIdMap := make(map[int64]any)
	for _, m := range ms {
		if m != nil {
			uIdMap[m.UserId] = struct{}{}
		}
	}
	uIds := make([]int64, len(uIdMap))
	i := 0
	for uId := range uIdMap {
		uIds[i] = uId
		i++
	}
	return uIds
}
