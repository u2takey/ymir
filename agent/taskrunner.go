package agent

import (
	"sort"
	"time"

	"github.com/Sirupsen/logrus"

	"github.com/arlert/ymir/model"
	"github.com/arlert/ymir/task"
	_ "github.com/arlert/ymir/taskset"
	"github.com/arlert/ymir/utils"
	utilhttp "github.com/arlert/ymir/utils/http"
)

// TaskRunner ...
type TaskRunner struct {
	config            *model.AgentConfig
	tasklistforchoose []taskforchoose
	resultqueue       chan *model.TReqResult
	resultlist        []*model.TReqResult
	resultmap         map[string]*model.TResult
	stopped           bool
}

type taskforchoose struct {
	tf     task.NewFunc
	weight int
}

// New ....
func New(cfg *model.AgentConfig) *TaskRunner {
	return &TaskRunner{
		config:      cfg,
		stopped:     false,
		resultqueue: make(chan *model.TReqResult, 100),
		resultmap:   make(map[string]*model.TResult, 0),
	}
}

// RunTasks ...
func (r *TaskRunner) RunTasks() {
	logrus.Debug("RunTasks")
	weight := 0
	runtime := task.GetRunTime()
	for _, f := range task.Sets() {
		taskset := f()
		if taskset.Weight() == 0 {
			weight++
		} else {
			weight += taskset.Weight()
		}
		r.tasklistforchoose = append(
			r.tasklistforchoose,
			taskforchoose{
				tf:     f,
				weight: weight,
			})
	}

	if len(r.tasklistforchoose) == 0 {
		return
	}

	// for now donot need cancel because RunTasks() exit will cause program exit
	if runtime == 0 {
		runtime = r.config.TaskSetTimeout
	}

	logrus.Debugf("timeout: %d, taskset count:%d", runtime, len(r.tasklistforchoose))

	stop := make(chan struct{}, 1)
	go func() {
		time.Sleep(runtime)
		stop <- struct{}{}
	}()

	r.run(stop)

	r.summarize()

	r.submit()
	// done
	utils.PrintJSON(r.resultmap)

}

func (r *TaskRunner) summarize() {
	for _, val := range r.resultmap {
		pctls := []int{10, 25, 50, 75, 90, 95, 99}
		data := make([]model.Bucket, len(pctls))
		for i := range data {
			data[i].Count = pctls[i]
		}
		j := 0
		sort.Float64s(val.Lats)
		for i := 0; i < len(val.Lats) && j < len(pctls); i++ {
			current := i * 100 / len(val.Lats)
			if current >= pctls[j] {
				data[j].Cost = time.Duration(val.Lats[i])
				j++
			}
		}

		bc := 10
		buckets := make([]model.Bucket, bc+1)

		bs := float64(val.Max-val.Min) / float64(bc)
		for i := 0; i < bc; i++ {
			buckets[i].Cost = val.Min + time.Duration(bs*float64(i))
		}
		buckets[bc].Cost = val.Max
		var bi int
		for i := 0; i < len(val.Lats); {
			if val.Lats[i] <= float64(buckets[bi].Cost) {
				i++
				buckets[bi].Count++
			} else if bi < len(buckets)-1 {
				bi++
			}
		}
		val.End = time.Now()
		val.Buckets = buckets
		val.Avg = time.Duration(int64(val.Sum) / val.Count)
		val.Percentile = data
		val.Lats = val.Lats[:0]
	}
}

func (r *TaskRunner) submit() {
	work := &model.TWorkInstance{
		InstanceID: r.config.InstanceID,
		JobName:    r.config.JobName,
		WorkID:     r.config.WorkID,
		NodeName:   r.config.NodeName,
	}
	for _, val := range r.resultmap {
		work.Results = append(work.Results, *val)
	}

	err := utilhttp.DefaultClient.CallWithJson(nil, nil, "POST", r.config.MasterAddr+"/api/v1/tresult", work)
	if err != nil {
		logrus.Error("submit error ", err)
	} else {
		logrus.Debug("submit success ", work)
	}
}

// RunSingleTask ...
func (r *TaskRunner) run(stop chan struct{}) {
	count := task.GetRoutineCount()
	for i := 0; i < count; i++ {
		f := r.chooseTask(i, count)
		// logrus.Debug("start worker ", i, f().Name())
		go r.worker(f)
	}

	for {
		select {
		case _ = <-stop:
			r.stopped = true
			return
		case ret := <-r.resultqueue:
			// logrus.Debugln("get ret", ret)

			r.addret(ret)
			r.resultlist = append(r.resultlist, ret)
		}
	}
}

func (r *TaskRunner) addret(ret *model.TReqResult) {
	if _, ok := r.resultmap[ret.Name]; !ok {
		logrus.Debugln("add resultmap for", ret.Name)
		r.resultmap[ret.Name] = &model.TResult{
			Name:    ret.Name,
			Start:   ret.Time,
			CodeMap: make(map[int]int64, 0),
			Min:     time.Hour,
		}
	}

	tresult := r.resultmap[ret.Name]
	tresult.Lats = append(tresult.Lats, float64(ret.Cost))
	tresult.CodeMap[ret.Code]++
	tresult.Count++
	tresult.Sum += ret.Cost
	if ret.Cost > tresult.Max {
		tresult.Max = ret.Cost
	} else if ret.Cost < tresult.Min {
		tresult.Min = ret.Cost
	}
}

func (r *TaskRunner) chooseTask(index, max int) task.NewFunc {
	maxweight := r.tasklistforchoose[len(r.tasklistforchoose)-1].weight
	intn := float64(maxweight) / float64(max) * float64(index)
	for _, task := range r.tasklistforchoose {
		if intn < float64(task.weight) {
			return task.tf
		}
	}
	return r.tasklistforchoose[0].tf
}

func (r *TaskRunner) worker(f task.NewFunc) {
	for {
		if r.stopped {
			return
		}
		workset := f()
		for _, work := range workset.Tasks() {
			start := time.Now()
			code := work.Run()
			end := time.Now()
			r.resultqueue <- &model.TReqResult{
				Name: workset.Name() + "." + work.Name(),
				Time: start,
				Cost: end.Sub(start),
				Code: code,
			}
		}
	}
}
