package worker

import (
	"concurrent-file-engine/logger"
	"concurrent-file-engine/process"
	"concurrent-file-engine/retry"
	"context"
	"sync"
)

type Job struct {
	ID       int
	FilePath string
}

func Worker(id int, jobs <-chan Job, ctx context.Context,wg *sync.WaitGroup){

	defer wg.Done()
	for{

		select {
		case <-ctx.Done():
			logger.Log.Infof("Worker %d shutting down",id)
			return

		case job, ok := <-jobs:
			if !ok {
				logger.Log.Infof("Worker %d job closed",id)
				return
			}

			logger.Log.Infof("Worker %d processing file: %s", id, job.FilePath)

			err:=retry.Do(3,func() error{
				return process.ProcessFile(ctx,job.FilePath)
			})

			if err != nil {
				logger.Log.Errorf("Worker %d failed to process file: %s", id, job.FilePath)
			}


		}
	}
}