package game

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/arielril/go-gl-collision-detection/collision"

	"github.com/go-gl/glfw/v3.3/glfw"
)

const benchMaxLines = 10000

type benchmarkResult []*result

type result struct {
	cType          collision.Type
	t              time.Time
	nLines         int
	fps            float64
	hSplit, vSplit uint8
}

func _newResult(cType collision.Type, t time.Time, nL int, fps float64, hSplit, vSplit uint8) *result {
	return &result{cType, t, nL, fps, hSplit, vSplit}
}

func _addRes(l *benchmarkResult, r *result) {
	*l = append(*l, r)
}

func _clearEnv() {
	maxLines = 250
	Init()
}

func _printComputation(r *result) {
	padding := 2
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.TabIndent)
	fmt.Fprintf(
		w,
		"Time: %v\tLines: %5d\tFPS: %v\n",
		r.t.Format(time.StampMilli),
		r.nLines,
		r.fps,
	)
	w.Flush()
}

func _computeFps() float64 {
	baseTime := glfw.GetTime()
	for i := 0; i < qtdFrames; i++ {
		displayScenario()
	}
	endTime := glfw.GetTime()

	fpsVal := qtdFrames / (endTime - baseTime)

	return fpsVal
}

func _increaseAndCreateObjects() {
	maxLines += 500
}

func _getConfig(step int) *gameConfig {
	cfg := newGameConfig()
	cfg.collisionType = collision.Provider.Me
	cfg.hSplit = 10
	cfg.vSplit = 10
	cfg.nLines = 500 * step
	cfg.showCar = true
	cfg.showLines = true
	cfg.showMenu = false

	return cfg
}

func _getConfigW(t collision.Type, s uint8) func(int) *gameConfig {
	return func(step int) *gameConfig {
		cfg := _getConfig(step)

		cfg.collisionType = t
		cfg.hSplit = 10 * s
		cfg.vSplit = 10 * s
		return cfg
	}
}

func _saveBenchmarkResult(br benchmarkResult, t collision.Type) {
	now := time.Now()
	path := fmt.Sprintf(
		"./results/bmk-2-2-%s-%d%v%d%d%d%d.log",
		t,
		now.Year(),
		now.Month(),
		now.Day(),
		now.Hour(),
		now.Minute(),
		now.Second(),
	)

	f, err := os.Create(path)
	if err != nil {
		fmt.Printf("failed to save results to file %s. Err: %v\n", path, err)
		return
	}
	defer f.Close()

	for _, r := range br {
		str := fmt.Sprintf(
			"{\"time\": \"%v\", \"lines\": %d, \"collisionType\": \"%s\", \"fps\": %.5f}\n",
			r.t.Format(time.StampMilli),
			r.nLines,
			r.cType,
			r.fps,
		)
		_, err = f.WriteString(str)
		if err != nil {
			fmt.Printf("failed to save (%#v) result\n", r)
			return
		}
	}
}

func _saveBenchmarkResultCSV(bbr [][]*result) {
	now := time.Now()
	path := fmt.Sprintf(
		"./results/%d%v%d%d%d%d.csv",
		now.Year(),
		now.Month(),
		now.Day(),
		now.Hour(),
		now.Minute(),
		now.Second(),
	)

	f, err := os.Create(path)
	if err != nil {
		fmt.Printf("failed to save results to file %s. Err: %v\n", path, err)
		return
	}
	defer f.Close()

	header := "'#lines','H-V',dummy,subspace\n"
	_, err = f.WriteString(header)
	if err != nil {
		fmt.Printf("failed to write header to file. Err: %v\n", err)
		return
	}

	for _, br := range bbr {
		for _, r := range br {
			nLines := r.nLines
			hSplit := r.hSplit
			vSplit := r.vSplit

			var dummyRes float64
			var subspaceRes float64

			if r.cType == collision.Provider.Me {
				subspaceRes = r.fps
			}
			if r.cType == collision.Provider.Professor {
				dummyRes = r.fps
			}

			s := fmt.Sprintf("%d,'%d-%d',%.4f,%.4f\n", nLines, hSplit, vSplit, dummyRes, subspaceRes)
			_, err = f.WriteString(s)
			if err != nil {
				fmt.Printf("failed to write res (%v) to file. Err: %v\n", s, err)
				return
			}
		}
	}
}

func _run(getConfigFunc func(int) *gameConfig) benchmarkResult {
	benchRes := make(benchmarkResult, 0)

	step := 1
	lastComputedFps := time.Now()
	for maxLines < benchMaxLines {
		now := time.Now()

		if now.Sub(lastComputedFps).Seconds() >= 1 {
			cfg := getConfigFunc(step)
			InitCustom(cfg)

			fps := _computeFps()

			r := _newResult(
				cfg.collisionType,
				now,
				cfg.nLines,
				fps,
				cfg.hSplit,
				cfg.vSplit,
			)
			_addRes(&benchRes, r)
			_printComputation(r)

			lastComputedFps = time.Now()
			step++
		}
	}

	return benchRes
}

// RunBenchmark do what it says
func RunBenchmark() {
	bbr := make([][]*result, 0)

	startTime := time.Now()
	fmt.Printf("Benchmark started at:\t%v\n", startTime.Format(time.StampMilli))

	fmt.Printf("\nRunning with the professor collision...\n")
	getConfigFunc := _getConfigW(collision.Provider.Professor, 1)
	res := _run(getConfigFunc)
	bbr = append(bbr, res)
	// _saveBenchmarkResult(res, collision.Provider.Professor)
	_clearEnv()

	for i := uint8(1); i <= 5; i++ {
		fmt.Printf("\nRunning with the student collision...\n")
		getConfigFunc = _getConfigW(collision.Provider.Me, i)
		res = _run(getConfigFunc)
		bbr = append(bbr, res)
		// _saveBenchmarkResult(res, collision.Provider.Me)
		_clearEnv()
	}
	fmt.Printf("BBR %v\n", bbr)

	endTime := time.Now()
	fmt.Printf("Benchmark finished at:\t%v\n", endTime.Format(time.StampMilli))

	_saveBenchmarkResultCSV(bbr)
}
