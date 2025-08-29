package invoker

type InvokerOptions struct {
	Rps          int64
	RunTime      int64 // total duration of the run in ms
	Distribution Distribution
	DurationMs   int64 // duration of the cpu-spin in ms

	Endpoint string
}

type Distribution string

var (
	DistributionUniform Distribution = "uniform"
	DistributionPoisson Distribution = "poisson"
)
