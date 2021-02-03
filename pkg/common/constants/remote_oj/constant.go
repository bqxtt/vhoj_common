package remote_oj

type RemoteOJ int32

const (
	POJ RemoteOJ = 1
	HDU RemoteOJ = 2
)

type RemoteOJInfo struct {
	RemoteOJ       RemoteOJ
	Host           string
	ProblemUrl     string
	SubmitUrl      string
	LoginUrl       string
	StatusUrl      string
	ResultUrl      string
	CompileInfoUrl string
}

var (
	POJInfo = &RemoteOJInfo{
		RemoteOJ:       POJ,
		Host:           "http://poj.org",
		ProblemUrl:     "/problem?id=",
		SubmitUrl:      "",
		LoginUrl:       "/login",
		StatusUrl:      "",
		ResultUrl:      "",
		CompileInfoUrl: "",
	}
	HDUInfo = &RemoteOJInfo{
		RemoteOJ:       HDU,
		Host:           "http://acm.hdu.edu.cn",
		ProblemUrl:     "/showproblem.php?pid=",
		SubmitUrl:      "/submit.php?action=submit",
		LoginUrl:       "/userloginex.php?action=login&cid=0&notice=0",
		StatusUrl:      "/status.php?user=%v&pid=%v",
		ResultUrl:      "/status.php?first=%v",
		CompileInfoUrl: "/viewerror.php?rid=%v",
	}
)
