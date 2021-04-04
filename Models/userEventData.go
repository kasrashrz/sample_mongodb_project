package Models

type UserEventData struct {
	EventId string				`bson:"EventId" json:"eventId"`
	UserEventStage string		`bson:"UserEventStage" json:"userEventStage"`
	//UserMetaData types.Array	`json:"marketName"`
	Score  int					`bson:"Score" json:"score"`
	JoinTime int64				`bson:"JoinTime" json:"joinTime"`
	EndTime int64				`bson:"EndTime" json:"endTime"`
	StartTime int64				`bson:"StartTime" json:"startTime"`
	PreActiveTime int64			`bson:"PreActiveTime" json:"preActiveTime"`
}
