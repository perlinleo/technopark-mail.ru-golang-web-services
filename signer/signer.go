package main

import (
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func SingleHash(in, out chan interface{}) {
  localWg := &sync.WaitGroup{}
  for val := range in {
    data := strconv.Itoa(val.(int))
    md5Result := DataSignerMd5(data)

    localWg.Add(1)
    go func(data string, md5Result string, out chan interface{}, wg *sync.WaitGroup) {
      defer wg.Done()
    
      var crc32DataResult string
      var crc32Md5Result string
    
      crc32Wg := &sync.WaitGroup{}
      crc32Wg.Add(2)
      go asyncCrc32(data, &crc32DataResult, crc32Wg)
      go asyncCrc32(md5Result, &crc32Md5Result, crc32Wg)
      crc32Wg.Wait()
    
      out <- crc32DataResult + "~" + crc32Md5Result
    }(data, md5Result, out, localWg)
  }

  localWg.Wait()
}



func MultiHash(in, out chan interface{}) {
  localWg := &sync.WaitGroup{}
  for val := range in {
    localWg.Add(1)
    go func(val string, out chan interface{}, wg *sync.WaitGroup) {
      var result string;
    
      var crc32Results [6]string;
      
      crc32Wg := &sync.WaitGroup{}
      crc32Wg.Add(6)
    
      for i := range crc32Results {
          go asyncCrc32(strconv.Itoa(i)+val, &crc32Results[i], crc32Wg)
      }
      crc32Wg.Wait()
    
      
    
      for _, val := range crc32Results {
        result+=val;
      }
    
      defer wg.Done()
    
      out <- result;
    }(val.(string), out, localWg)
  }
  localWg.Wait()
}



func CombineResults(in, out chan interface{}) {
  var result []string
  for val := range in {
    result = append(result, val.(string))
  }
  sort.Strings(result)
  out <- strings.Join(result, "_")
}

func asyncCrc32(data string, result *string, wg *sync.WaitGroup) {
  *result = DataSignerCrc32(data)
  defer wg.Done()
}

func ExecutePipeline(jobs ...job) {
  runtime.GOMAXPROCS(0)
  wg := &sync.WaitGroup{}
  in := make(chan interface{})
  for _, job := range jobs {
    wg.Add(1)
    out := make(chan interface{})
    go func(worker func(in, out chan interface{}), in chan interface{}, out chan interface{}, wg *sync.WaitGroup) {
      worker(in, out)
      defer close(out)
      defer wg.Done()
    }(job, in, out, wg)
    in = out
  }

  wg.Wait()
}