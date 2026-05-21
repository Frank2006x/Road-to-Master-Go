package process

import (
	"context"
	"fmt"
	"os"
	"time"
)

func ProcessFile(ctx context.Context,path string) error{

	select{
		case <-ctx.Done():
			return ctx.Err()
		default:
	}

	data,err:=os.ReadFile(path)
	if err!=nil{
		return err
	}
	time.Sleep(2*time.Second)
	fmt.Printf("Processed %s | size=%d bytes\n", path, len(data))
	return nil
}