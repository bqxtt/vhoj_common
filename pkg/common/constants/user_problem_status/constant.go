package user_problem_status

type UserProblemStatus int32

const (
	ACCEPTED   UserProblemStatus = 1
	ATTEMPTED  UserProblemStatus = 2
	NOT_SUBMIT UserProblemStatus = 3
)
