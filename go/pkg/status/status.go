package status

type Status struct {
	code uint8
	info string
}

var (
	OK = &Status{code: 0,info: "ok"}
	ArticleSave = &Status{code: 1, info: "article save status"}

	CommentCheck = &Status{code: 1, info: "wait to check"}
	CommentNotPass = &Status{code: 2, info: "comment not pass"}
	CommentReport = &Status{code: 3, info: "be reported"}
	CommentDelete = &Status{code: 9, info: "be deleted"}
)

func (s *Status) GetCode() uint8 {
	return s.code
}
