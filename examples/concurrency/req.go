package concurrency

type Request struct {
	Num  int
	Resp chan Response
}

type Response struct {
	Num      int
	WorkerID int
}

func PlusOneService(reqs <-chan Request, workerID int) {
	for req := range reqs {
		go func(req Request) {
			defer close(req.Resp)
			req.Resp <- Response{req.Num + 1, workerID}
		}(req)
	}
}
