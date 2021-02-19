package library
const (
    ErrnoSuccess = 0
    ErrnoError   = 1
    ErrnoUnknown = 2
)
var ErrNoToMsgMap = map[int32]string{
    ErrnoSuccess: "success",
    ErrnoError:   "failed",
    ErrnoUnknown: "unknown",
}
func GetErrMsg(errNo int32) string {
    if errMsg, ok := ErrNoToMsgMap[errNo]; ok {
        return errMsg
    }
    return "unknown error"
}
