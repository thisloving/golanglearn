package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Payload struct {
	name string
}

func (p *Payload) Play() {
	fmt.Printf("%s 打LOL游戏...当前任务完成\n", p.name)
}

type Job struct {
	PayloadData Payload
}

var JobQueue chan Job

type Worker struct {
	name       string
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

func (w *Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobChannel
			fmt.Printf("[%s] 把自己注册到 对象池中 当前任务的长度是[%d] \n", w.name, len(w.WorkerPool))
			select {
			case job := <-w.JobChannel:
				fmt.Printf("[%s] 工人接收到了任务 当前任务的长度是[%d]\n", w.name, len(w.WorkerPool))
				job.PayloadData.Play()
				time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
			case <-w.quit:
			}
		}
	}()
}

func (w *Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

type Dispatcher struct {
	name       string
	maxWorkers int
	WorkerPool chan chan Job
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{
		WorkerPool: pool,
		maxWorkers: maxWorkers,
	}
}

func NewWorker(workerPool chan chan Job, name string) Worker {
	fmt.Printf("创建了一个工人,它的名字是:%s \n", name)
	return Worker{
		name:       name,
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool),
	}
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.maxWorkers; i++ {
		name := fmt.Sprintf("worker-%s", strconv.Itoa(i))
		worker := NewWorker(d.WorkerPool, name)
		worker.name = name
		worker.Start()
	}
	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			fmt.Println("调度者,接收到一个工作任务")
			time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
			go func(job Job) {
				jobChannel := <-d.WorkerPool
				jobChannel <- job
			}(job)
		default:
			//fmt.Printf("OK!\n")
		}
	}
}

func intialize() {
	maxWorkers := 2
	maxQueue := 4
	dispatch := NewDispatcher(maxWorkers)
	JobQueue = make(chan Job, maxQueue)
	dispatch.Run()
}

func main() {
	intialize()

	for i := 0; i < 10; i++ {
		p := Payload{
			fmt.Sprintf("[Player-%s]", strconv.Itoa(i)),
		}

		JobQueue <- Job{
			PayloadData: p,
		}

		time.Sleep(time.Second)
	}

	close(JobQueue)
}
