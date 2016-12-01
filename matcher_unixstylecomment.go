package golexer

// #开头的行注释
type UnixStyleCommentMatcher struct {
	baseMatcher
}

func (self *UnixStyleCommentMatcher) Match(tz *Tokenizer) (*Token, error) {
	if tz.Current() != '#' {
		return nil, nil
	}

	tz.ConsumeOne()

	begin := tz.Index()

	for {

		tz.ConsumeOne()

		if tz.Current() == '\n' || tz.Current() == 0 {
			break
		}

	}

	return NewToken(self, tz, tz.StringRange(begin, tz.index), ""), nil
}

func NewUnixStyleCommentMatcher(id int) TokenMatcher {
	return &UnixStyleCommentMatcher{
		baseMatcher{id},
	}
}
