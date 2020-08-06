package status

type Status struct {
	code uint8
	info string
}

var (
	OK = &Status{code: 0,info: "ok"}
	ArticleSave = &Status{code: 1, info: "article save status"}

	CommentCheck = &Status{code: 1, info: "sensitive word is included"}
)

func (s *Status) GetCode() uint8 {
	return s.code
}
