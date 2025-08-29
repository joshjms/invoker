package invoker

type Report struct {
	RequestAt     int64
	WorkerStartAt int64
	WorkerEndAt   int64
	ResponseAt    int64
}
