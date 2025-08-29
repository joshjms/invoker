package invoker

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/joshjms/invoker/api/workerpb"
	"google.golang.org/grpc"
)

type Invoker struct {
	Options InvokerOptions
}

func NewInvoker(opts InvokerOptions) *Invoker {
	return &Invoker{
		Options: opts,
	}
}

func (inv *Invoker) Run() ([]Report, error) {
	switch inv.Options.Distribution {
	case DistributionUniform:
		return inv.runUniform()
	case DistributionPoisson:
		// return inv.runPoisson()
	default:
		return nil, fmt.Errorf("unknown distribution: %s", inv.Options.Distribution)
	}

	return nil, nil
}

func (inv *Invoker) runUniform() ([]Report, error) {
	conn, err := grpc.Dial(inv.Options.Endpoint, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := workerpb.NewWorkerServiceClient(conn)

	interval := time.Duration(1e9 / inv.Options.Rps)
	ticker := time.NewTicker(interval)

	wg := sync.WaitGroup{}
	var mu sync.Mutex

	reports := make([]Report, 0)
	reqs := int64(0)
	totalReqs := inv.Options.RunTime * inv.Options.Rps / 1000

	for range ticker.C {
		wg.Add(1)
		go func() {
			defer wg.Done()

			var r Report
			r.RequestAt = time.Now().UnixNano()
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			res, err := client.Work(ctx, &workerpb.WorkRequest{
				DurationMs: inv.Options.DurationMs,
			})
			cancel()
			if err != nil {
				fmt.Println(err)
				return
			}
			r.ResponseAt = time.Now().UnixNano()
			r.WorkerStartAt = res.GetStartAt()
			r.WorkerEndAt = res.GetEndAt()

			mu.Lock()
			reports = append(reports, r)
			mu.Unlock()
		}()
		reqs++

		if reqs >= totalReqs {
			break
		}
	}

	wg.Wait()
	ticker.Stop()

	return reports, nil
}
