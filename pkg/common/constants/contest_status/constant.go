package contest_status

type ContestStatus int32

const (
	ALL       ContestStatus = 0
	SCHEDULED ContestStatus = 1
	RUNNING   ContestStatus = 2
	ENDED     ContestStatus = 3
)
