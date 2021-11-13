package main

import (
	"bufio"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type ResponseContent struct {
	Addr string `json:"address"`
}

type addContent struct {
	AreaID    string `json:"area_id"`
	AddCarNum int    `json:"add_car_num""`
}

type timeline struct {
	AddContent []addContent `json:"add_content"`
}

// フェーズで行っていく
// 1. mainの中で定義した車両からリクエストを行う...DONE
// 2. jsonから設定を読み込んで、それを元に車両を増やしてリクエストを行う
// 3. 終了時間を設定して台数を調整(増減を再現できるようにする)...ここは必要かどうかは後ほど考える
// 		→ ランダム値で削除されるたみんぐとその数を定義してできるようにする or 指定数に達した時に、これ以上増えないように....いるのか?

var (
	fazeF  int
	logger *log.Logger
)

func init() {
	flag.IntVar(&fazeF, "faze", 1, "decide faze of implementation")
}

func main() {
	// parse flag
	flag.Parse()

	// create logger
	f, _ := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR, 0755)
	logger = log.New(f, "", log.Ldate)

	// launch server
	go launchMockServer()

	// exec Parallel
	ctx, _ := context.WithTimeout(context.Background(), 7*time.Second)
	fmt.Printf("faze: %d\n", fazeF)
	switch fazeF {
	case 1:
		fazeOne(ctx)
	case 2:
		fazeTwo(ctx)
	case 3:
		fazeThree(ctx)
	default:
		log.Fatal("no exist faze")
	}

	// wait 5 second
	<-ctx.Done()
	f.Close()

	switch fazeF {
	case 1:
		// TODO
	case 2:
		resultMap := make(map[string]struct{})
		f, _ := os.Open("log.txt")
		defer f.Close()
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			texts := scanner.Text()
			carID := texts[11:]
			if _, ok := resultMap[carID]; !ok {
				resultMap[carID] = struct{}{}
			}
		}
		timelines := readSetting()
		for i, timeline := range timelines {
			for _, c := range timeline.AddContent {
				for j := 0; j < c.AddCarNum; j++ {
					carID := fmt.Sprintf("timeline_%d:area_%s:car_num_%d", i, c.AreaID, j)
					if _, ok := resultMap[carID]; !ok {
						log.Fatalf("carID[%s] didn't request\ncontents: %v", carID, resultMap)
					}
				}
			}
		}
	case 3:
		// TODO
	default:
		log.Fatal("no exist faze")
	}
	fmt.Println("all car did request")
	fmt.Println("time out")
}

func fazeOne(ctx context.Context) {
	carIDs := []string{"c1", "c2", "c3", "c4", "c5"}
	for _, carID := range carIDs {
		carID := carID
		go func() {
			for {
				time.Sleep(100 * time.Millisecond)
				execRequest(ctx, carID)
			}
		}()
	}
}

func fazeTwo(ctx context.Context) {
	timelines := readSetting()
	for i, timeline := range timelines {
		fmt.Println("*****************add thread*****************")
		i := i
		timeline := timeline
		for _, c := range timeline.AddContent {
			c := c
			for j := 0; j < c.AddCarNum; j++ {
				j := j
				go func() {
					for {
						time.Sleep(100 * time.Millisecond)
						execRequest(ctx, fmt.Sprintf("timeline_%d:area_%s:car_num_%d", i, c.AreaID, j))
					}
				}()
			}
		}
		time.Sleep(1 * time.Second)
	}
}

func readSetting() []timeline {
	buf, _ := ioutil.ReadFile("setting.json")
	var timelines []timeline
	json.Unmarshal(buf, &timelines)
	return timelines
}

func fazeThree(ctx context.Context) {
	// TODO
}

func launchMockServer() {
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		longExec()
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"address": "sample"}`))
	}))
}

func longExec() {
	time.Sleep(100 * time.Millisecond)
}

func execRequest(ctx context.Context, carID string) {
	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, "http://localhost:8080", nil)
	c := &http.Client{}
	start := time.Now()
	res, err := c.Do(req)
	if err != nil {
		fmt.Printf("error request exec: %v\n", err)
		return
	}
	duration := time.Since(start).Milliseconds() //ms
	fmt.Printf("carID[%s]...time duration: %dms\n", carID, duration)
	defer res.Body.Close()
	var data ResponseContent
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		fmt.Printf("error decode: %v\n", err)
		return
	}
	logger.Println(carID)
	// fmt.Printf("carID[%s]...parsed content: %+v\n", carID, data)
}
