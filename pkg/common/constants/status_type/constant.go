package status_type

type SubmissionStatusType int32

var (
	PENDING SubmissionStatusType = 0

	SUBMIT_FAILED_TEMP SubmissionStatusType = 1
	SUBMIT_FAILED_PERM SubmissionStatusType = 2

	SUBMITTED SubmissionStatusType = 3

	QUEUEING  SubmissionStatusType = 4
	COMPILING SubmissionStatusType = 5
	JUDGING   SubmissionStatusType = 6

	AC           SubmissionStatusType = 7
	PE           SubmissionStatusType = 8
	WA           SubmissionStatusType = 9
	TLE          SubmissionStatusType = 10
	MLE          SubmissionStatusType = 11
	OLE          SubmissionStatusType = 12
	RE           SubmissionStatusType = 13
	CE           SubmissionStatusType = 14
	FAILED_OTHER SubmissionStatusType = 15
)
