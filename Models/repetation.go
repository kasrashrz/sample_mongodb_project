package Models

type Repetition struct {
	RandomRepetitionUuId string `bson:"RandomRepetitionUuid"`
	StartPreActiveTime   int64  `bson:"StartPreActiveTime" json:"startPreActiveTime"`
	StartTime            int64  `bson:"StartTime" json:"startTime"`
	EndTime              int64  `bson:"EndTime" json:"endTime"`
	Terminate            bool   `bson:"Terminate" json:"Terminate"`
	StartJoinTime        int64  `bson:"StartJoinTime" json:"startJoinTime"`
	EndJoinTime          int64  `bson:"EndJoinTime" json:"endJoinTime"`
}
