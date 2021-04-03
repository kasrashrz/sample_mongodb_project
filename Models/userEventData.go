package Models

type UserEventData struct {
	EventId string
	UserEventStage string
	//UserMetaData types.Array
	Score  int
	JoinTime int64
	EndTime int64
	StartTime int64
	PreActiveTime int64
}
