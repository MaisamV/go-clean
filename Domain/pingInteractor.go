package Domain

type PingInteractor interface {
	Ping() (BaseResponse, error)
}
