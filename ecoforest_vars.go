package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"reflect"
	"fmt"
)

type dataPoint struct {
	funcType int
	function int
	ini      int
	reg      int
	index    int
	metric   prometheus.Gauge
}

func setCounter(funcType int, function int, ini int, reg int, index int, val float64) {
	var metric prometheus.Gauge
	for _, dataPoint := range dataPoints {
		if dataPoint.funcType == funcType && dataPoint.function == function && dataPoint.ini == ini && dataPoint.reg == reg && dataPoint.index == index {
			metric = dataPoint.metric
			fmt.Println(reflect.TypeOf(&metric))
			metric.Set(val)
		}
	}
}

var dataPoints = []dataPoint{
	dataPoint{
		funcType: 1,
		function: 2002,
		ini:      1,
		reg:      39,
		index:    2,
		metric: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "ecoforest_brine_input_gauge",
			Help: "Temperature in C of the brine input temperature",
		}),
	},
	dataPoint{
		funcType: 1,
		function: 2002,
		ini:      1,
		reg:      39,
		index:    3,
		metric: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "ecoforest_brine_output_gauge",
			Help: "Temperature in C of the brine output temperature",
		}),
	},
	dataPoint{
		funcType: 1,
		function: 2002,
		ini:      1,
		reg:      39,
		index:    5,
		metric: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "ecoforest_production_input_gauge",
			Help: "Temperature in C of the production input temperature",
		}),
	},
	dataPoint{
		funcType: 1,
		function: 2002,
		ini:      1,
		reg:      39,
		index:    4,
		metric: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "ecoforest_production_output_gauge",
			Help: "Temperature in C of the production output temperature",
		}),
	},
}
