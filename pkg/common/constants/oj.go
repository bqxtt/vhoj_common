package constants

type RemoteOJ int32

const (
	POJ RemoteOJ = 1
	HDU RemoteOJ = 2
)

type RemoteOJInfo struct {
	RemoteOJ   RemoteOJ
	Host       string
	ProblemUrl string
	SubmitUrl  string
	LoginUrl   string
	StatusUrl  string
}

var (
	POJInfo = &RemoteOJInfo{
		RemoteOJ:   POJ,
		Host:       "http://poj.org",
		ProblemUrl: "/problem?id=",
		SubmitUrl:  "",
		LoginUrl:   "/login",
		StatusUrl:  "",
	}
	HDUInfo = &RemoteOJInfo{
		RemoteOJ:   HDU,
		Host:       "http://acm.hdu.edu.cn",
		ProblemUrl: "/showproblem.php?pid=",
		SubmitUrl:  "/submit.php?action=submit",
		LoginUrl:   "/userloginex.php?action=login&cid=0&notice=0",
		StatusUrl:  "/status.php?user=%v&pid=%v",
	}
)

type SubmissionStatus int32

const (
	PENDING SubmissionStatus = 0

	SUBMIT_FAILED_TEMP SubmissionStatus = 1
	SUBMIT_FAILED_PERM SubmissionStatus = 2

	SUBMITTED SubmissionStatus = 3

	QUEUEING  SubmissionStatus = 4
	COMPILING SubmissionStatus = 5
	JUDGING   SubmissionStatus = 6

	AC           SubmissionStatus = 7
	PE           SubmissionStatus = 8
	WA           SubmissionStatus = 9
	TLE          SubmissionStatus = 10
	MLE          SubmissionStatus = 11
	OLE          SubmissionStatus = 12
	RE           SubmissionStatus = 13
	CE           SubmissionStatus = 14
	FAILED_OTHER SubmissionStatus = 15
)
