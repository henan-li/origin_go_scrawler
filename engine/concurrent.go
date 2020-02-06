package engine

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan Item
}

type Scheduler interface {
	Submit(Request)
	WorkerChan() chan Request
	ReadyNotifier
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func NewConcurrentEngine(scheduler Scheduler, workerCount int, ItemChan chan Item) *ConcurrentEngine {

	e := &ConcurrentEngine{
		Scheduler:   scheduler,
		WorkerCount: workerCount,
		ItemChan:    ItemChan,
	}
	return e
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	//in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {

		if isDuplicate(r.Url) {
			continue
		}
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			go func() { e.ItemChan <- item }()
		}

		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

var urlVisitRecord = make(map[string]bool)

func isDuplicate(url string) bool {
	if urlVisitRecord[url] {
		return true
	}
	urlVisitRecord[url] = true
	return false
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {

	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)

			if err != nil {
				continue
			}

			out <- result
		}
	}()
}
