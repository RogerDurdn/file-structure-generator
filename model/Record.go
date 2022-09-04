package model

type Record struct {
	Topic string `csv:"Topic"`
	Name  string `csv:"Name"`
	Unit  string `csv:"Unit"`
}

type TreeTopic struct {
	Topics map[string]Topic
}

func NewTreeTopic() *TreeTopic {
	return &TreeTopic{make(map[string]Topic)}
}

type Topic struct {
	Units map[string][]*Record
}

func NewTopic() Topic {
	return Topic{make(map[string][]*Record)}
}
