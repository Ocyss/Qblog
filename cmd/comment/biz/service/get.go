package service

import (
	"github.com/Ocyss/Qblog/cmd/comment/biz/dal/db"
	"github.com/Ocyss/Qblog/kitex_gen/comment"
	"github.com/Ocyss/Qblog/kitex_gen/user"
	"github.com/Ocyss/Qblog/pkg/conv"
)

func (s *CommentService) Get(request *comment.CommentListReq) (resp *comment.CommentListResp, err error) {
	resp = new(comment.CommentListResp)
	commentList, err := db.GetComment(s.ctx, uint(request.Aid))
	if err != nil {
		return
	}
	resp.Data = make([]*comment.Comment, len(commentList))
	for i, c := range commentList {
		u := &user.User{
			Name:  c.Name,
			Email: c.Email,
			Qq:    c.Qq,
		}
		if c.Uid != nil {
			u.Uid = conv.ToInt64(c.Uid)
		}
		resp.Data[i] = &comment.Comment{
			User:    u,
			Content: c.Content,
			Reply:   conv.ToInt64(c.ReplyID),
			Father:  conv.ToInt64(c.FatherID),
		}
	}
	return
}
