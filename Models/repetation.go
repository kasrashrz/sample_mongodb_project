package Models

type Repetition struct {
	StartPreActiveTime int64 	`json:"startPreActiveTime"`
	StartTime int64				`json:"startTime"`
	EndTime int64				`json:"endTime"`
	Terminate bool				`json:"Terminate"`
	StartJoinTime int64			`json:"startJoinTime"`
	EndJoinTime int64			`json:"endJoinTime"`
}