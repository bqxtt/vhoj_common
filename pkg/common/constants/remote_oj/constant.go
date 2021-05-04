package remote_oj

type RemoteOJ int32

const (
	POJ RemoteOJ = 1
	HDU RemoteOJ = 2
	MXT RemoteOJ = 26
	JSK RemoteOJ = 25
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
		SubmitUrl:      "/submit",
		LoginUrl:       "/login",
		StatusUrl:      "/status?user_id=%v&problem_id=%v",
		ResultUrl:      "/showsource?solution_id=%v",
		CompileInfoUrl: "/showcompileinfo?solution_id=%v",
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
