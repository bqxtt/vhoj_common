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
	ResultUrl  string
}

var (
	POJInfo = &RemoteOJInfo{
		RemoteOJ:   POJ,
		Host:       "http://poj.org",
		ProblemUrl: "/problem?id=",
		SubmitUrl:  "",
		LoginUrl:   "/login",
		StatusUrl:  "",
		ResultUrl:  "",
	}
	HDUInfo = &RemoteOJInfo{
		RemoteOJ:   HDU,
		Host:       "http://acm.hdu.edu.cn",
		ProblemUrl: "/showproblem.php?pid=",
		SubmitUrl:  "/submit.php?action=submit",
		LoginUrl:   "/userloginex.php?action=login&cid=0&notice=0",
		StatusUrl:  "/status.php?user=%v&pid=%v",
		ResultUrl:  "/status.php?first=%v",
	}
)

type SubmissionStatus struct {
	StatusNum int32
	Finished  bool
}

var (
	PENDING = &SubmissionStatus{0, false}

	SUBMIT_FAILED_TEMP = &SubmissionStatus{1, false}
	SUBMIT_FAILED_PERM = &SubmissionStatus{2, true}

	SUBMITTED = &SubmissionStatus{3, false}

	QUEUEING  = &SubmissionStatus{4, false}
	COMPILING = &SubmissionStatus{5, false}
	JUDGING   = &SubmissionStatus{6, false}

	AC           = &SubmissionStatus{7, true}
	PE           = &SubmissionStatus{8, true}
	WA           = &SubmissionStatus{9, true}
	TLE          = &SubmissionStatus{10, true}
	MLE          = &SubmissionStatus{11, true}
	OLE          = &SubmissionStatus{12, true}
	RE           = &SubmissionStatus{13, true}
	CE           = &SubmissionStatus{14, true}
	FAILED_OTHER = &SubmissionStatus{15, true}
)

type Language int32

const (
	CPP  Language = 1
	JAVA Language = 2
)
