package model

type Topic struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var topics = []Topic{
	{Id: 1, Name: "Software Engineering"},
	{Id: 2, Name: "Cloud Computing"},
	{Id: 3, Name: "DevOps"},
}

func GetAllTopics() []Topic {
	return topics
}

func GetTopicById(id int) (Topic, bool) {
	for _, topic := range topics {
		if topic.Id == id {
			return topic, true
		}
	}
	return Topic{}, false
}
