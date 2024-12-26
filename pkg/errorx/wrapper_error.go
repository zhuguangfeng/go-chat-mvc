package errorx

type BizError interface {
	Error() string
	ErrorCode() ErrorCode
}

type WrapperError struct {
	err     error
	errCode ErrorCode
}

func NewBizError(errCode ErrorCode) *WrapperError {
	return &WrapperError{
		errCode: errCode,
	}
}

func (e *WrapperError) Error() string {
	if e.err != nil {
		return e.err.Error()
	}
	_, msg := e.errCode.GetCodeMsg()
	return msg
}

func (e *WrapperError) ErrorCode() ErrorCode {
	return e.errCode
}

func (e *WrapperError) WithError(err error) *WrapperError {
	e.err = err
	return e
}
