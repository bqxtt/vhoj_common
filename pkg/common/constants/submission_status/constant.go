package submission_status

import "github.com/ecnuvj/vhoj_common/pkg/common/constants/status_type"

type SubmissionStatus struct {
	StatusType status_type.SubmissionStatusType
	Finished   bool
}

var (
	PENDING = &SubmissionStatus{status_type.PENDING, false}

	SUBMIT_FAILED_TEMP = &SubmissionStatus{status_type.SUBMIT_FAILED_TEMP, false}
	SUBMIT_FAILED_PERM = &SubmissionStatus{status_type.SUBMIT_FAILED_PERM, true}

	SUBMITTED = &SubmissionStatus{status_type.SUBMITTED, false}

	QUEUEING  = &SubmissionStatus{status_type.QUEUEING, false}
	COMPILING = &SubmissionStatus{status_type.COMPILING, false}
	JUDGING   = &SubmissionStatus{status_type.JUDGING, false}

	AC           = &SubmissionStatus{status_type.AC, true}
	PE           = &SubmissionStatus{status_type.PE, true}
	WA           = &SubmissionStatus{status_type.WA, true}
	TLE          = &SubmissionStatus{status_type.TLE, true}
	MLE          = &SubmissionStatus{status_type.MLE, true}
	OLE          = &SubmissionStatus{status_type.OLE, true}
	RE           = &SubmissionStatus{status_type.RE, true}
	CE           = &SubmissionStatus{status_type.CE, true}
	FAILED_OTHER = &SubmissionStatus{status_type.FAILED_OTHER, true}
)
