package Models

type UserEventData struct {
	EventId string				`json:"eventId"`
	UserEventStage string		`json:"userEventStage"`
	//UserMetaData types.Array	`json:"marketName"`
	Score  int					`json:"score"`
	JoinTime int64				`json:"joinTime"`
	EndTime int64				`json:"endTime"`
	StartTime int64				`json:"startTime"`
	PreActiveTime int64			`json:"preActiveTime"`
}
