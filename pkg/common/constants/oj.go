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
}

var (
	POJInfo = &RemoteOJInfo{
		RemoteOJ:   POJ,
		Host:       "http://poj.org",
		ProblemUrl: "/problem?id=",
		SubmitUrl:  "",
		LoginUrl:   "/login",
	}
	HDUInfo = &RemoteOJInfo{
		RemoteOJ:   HDU,
		Host:       "http://acm.hdu.edu.cn",
		ProblemUrl: "/showproblem.php?pid=",
		SubmitUrl:  "",
		LoginUrl:   "/userloginex.php?action=login&cid=0&notice=0",
	}
)
