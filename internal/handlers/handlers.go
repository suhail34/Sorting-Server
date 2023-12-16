package handlers

import (
	"encoding/json"
	"strings"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type RequestBody struct {
  To_Sort [][]int `json:"to_sort"`
}

type RespBody struct {
  Sorted_Arrays [][]int `json:"sorted_arrays"`
  Time_ns int64 `json:"time_ns"`
}

func ProcessSingle(ctx *fiber.Ctx) error {
  rawBody := ctx.Body()
  var reqbody *RequestBody
  resbody := &RespBody{
    Sorted_Arrays: make([][]int, 0),
    Time_ns: 0,
  }
  decoder := json.NewDecoder(strings.NewReader(string(rawBody)))
  decoder.DisallowUnknownFields()
  err := decoder.Decode(&reqbody)
  if err != nil {
    logrus.Error("Error Unmarshalling json body : ",err)
    return ctx.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{"error":"Invalid JSON format"})
  }
  startTime := time.Now()
  for _, arr := range reqbody.To_Sort {
    temp := sortArr(arr)
    resbody.Sorted_Arrays = append(resbody.Sorted_Arrays, temp)
  }
  elapsedTime := time.Since(startTime);
  resbody.Time_ns = elapsedTime.Nanoseconds()
  return ctx.Status(fiber.StatusOK).JSON(resbody)
}

func ProcessConcurrent(ctx *fiber.Ctx) error {
  rawBody := ctx.Body()
  var reqbody *RequestBody
  resbody := &RespBody{
    Sorted_Arrays: make([][]int, 0),
    Time_ns: 0,
  }
  decoder := json.NewDecoder(strings.NewReader(string(rawBody)))
  decoder.DisallowUnknownFields()
  err := decoder.Decode(&reqbody)
  if err != nil {
    logrus.Error("Error Unmarshalling json body : ", err)
    return ctx.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{"error":"Invalid JSON format"})
  }
  ch := make(chan []int, len(reqbody.To_Sort))
  var wg sync.WaitGroup
  startTime := time.Now()
  for _, arr := range reqbody.To_Sort {
    wg.Add(1)
    go func(arr []int) {
      defer wg.Done()
      ch <- sortArr(arr)
    }(arr)
  }
  wg.Wait()
  close(ch)
  for arr := range ch {
    resbody.Sorted_Arrays = append(resbody.Sorted_Arrays, arr)
  }
  elapsedTime := time.Since(startTime)
  resbody.Time_ns = elapsedTime.Nanoseconds()
  return ctx.Status(fiber.StatusOK).JSON(resbody)
}
