package common

const (
	BuildStatusPending     = "pending"
	BuildStatusPreparation = "preparation"
	BuildStatusRunning     = "running"

	// 最终状态只能3种
	BuildStatusCancel = "cancel"
	BuildStatusError  = "error"
	BuildStatusOk     = "ok"
)

const (
	BuildEventCheckParam = "err_check_param"
	BuildEventPutJob     = "err_put_job"
)

func BuildStatusEnded(stat string) bool {
	if stat == BuildStatusCancel || stat == BuildStatusError || stat == BuildStatusOk {
		return true
	}
	return false
}
