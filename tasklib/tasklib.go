package tasklib

import (
	"github.com/nutanix-core/acs-prom-metrics/parser"
	"github.com/prometheus/client_golang/prometheus"
)

type taskManagerMetrics struct {
	NumberOfTasks *prometheus.CounterVec `metric:"name:number_of_tasks;help:Total number of tasks created;labels:partition"`
}

var input *taskManagerMetrics = &taskManagerMetrics{}

func Dosomething() {
	err := parser.InitializeMetrics(input, "some_namespace")
	if err != nil {
		panic(err)
	}
	go Random()
}

func Random() {
	go func() {
		for {
			input.NumberOfTasks.WithLabelValues("test1").Inc()
		}
	}()
}