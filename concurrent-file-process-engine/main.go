package main

import (
	"concurrent-file-engine/logger"
	"concurrent-file-engine/worker"
	"context"
	"os"
	"sync"
)

func main() {
	logger.InitLogger()

	ctx,cancel:=context.WithCancel(context.Background())

	defer cancel()

	jobs := make(chan worker.Job)

	workerCount := 5

	var wg sync.WaitGroup

	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go worker.Worker(i, jobs,ctx,&wg)
	}

	files,err:=os.ReadDir("./data")

	if err != nil {
		logger.Log.Errorf("Failed to read directory: %v", err)
		return
	}

	go func (){
		for i, file := range files {

			select {
				case <-ctx.Done():
					logger.Log.Infof("Job producer shutting down")
					return

				case jobs <- worker.Job{ID: i + 1, FilePath: "./data/" + file.Name()}:
					logger.Log.Infof("Job producer added file: %s", file.Name())
			}
		}

		close(jobs)
		logger.Log.Infof("Job producer finished adding jobs")
	}()
	
	wg.Wait()
	logger.Log.Infof("All workers completed")




}