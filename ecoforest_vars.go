package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	// "reflect"
	// "fmt"
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
			// fmt.Println(reflect.TypeOf(&metric))
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
	dataPoint{
		funcType: 1,
		function: 2002,
		ini:      1,
		reg:      39,
		index:    18,
		metric: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "ecoforest_dhw_output_gauge",
			Help: "Temperature in C of the domestic hot water",
		}),
	},
	dataPoint{
		funcType: 1,
		function: 2002,
		ini:      1,
		reg:      39,
		index:    12,
		metric: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "ecoforest_outdoor_temperature_gauge",
			Help: "Temperature in C of the external sensor",
		}),
	},
	dataPoint{
		funcType: 1,
		function: 2002,
		ini:      1,
		reg:      39,
		index:    14,
		metric: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "ecoforest_brine_pressure_gauge",
			Help: "Pressure in bar of the brine loop",
		}),
	},
	dataPoint{
		funcType: 1,
		function: 2002,
		ini:      1,
		reg:      39,
		index:    15,
		metric: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "ecoforest_production_pressure_gauge",
			Help: "Pressure in bar of the production loop",
		}),
	},
	dataPoint{
		funcType: 1,
		function: 2002,
		ini:      176,
		reg:      29,
		index:    23,
		metric: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "ecoforest_production_group_1_gauge",
			Help: "Production Group 1 Current",
		}),
	},
	dataPoint{
		funcType: 1,
		function: 2002,
		ini:      176,
		reg:      29,
		index:    22,
		metric: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "ecoforest_production_group_2_gauge",
			Help: "Production Group 2 Current",
		}),
	},
	dataPoint{
		funcType: 1,
		function: 2002,
		ini:      176,
		reg:      29,
		index:    21,
		metric: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "ecoforest_production_group_3_gauge",
			Help: "Production Group 3 Current",
		}),
	},
	dataPoint{
		funcType: 1,
		function: 2002,
		ini:      176,
		reg:      29,
		index:    20,
		metric: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "ecoforest_production_group_4_gauge",
			Help: "Production Group 4 Current",
		}),
	},

	dataPoint{
		funcType: 1,
		function: 2002,
		ini:      176,
		reg:      29,
		index:    23,
		metric: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "ecoforest_cooling_group_1_gauge",
			Help: "Cooling Group 1 Current",
		}),
	},
	dataPoint{
		funcType: 1,
		function: 2002,
		ini:      176,
		reg:      29,
		index:    22,
		metric: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "ecoforest_cooling_group_2_gauge",
			Help: "Cooling Group 2 Current",
		}),
	},
	dataPoint{
		funcType: 1,
		function: 2002,
		ini:      176,
		reg:      29,
		index:    21,
		metric: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "ecoforest_cooling_group_3_gauge",
			Help: "Cooling Group 3 Current",
		}),
	},
	dataPoint{
		funcType: 1,
		function: 2002,
		ini:      176,
		reg:      29,
		index:    20,
		metric: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "ecoforest_cooling_group_4_gauge",
			Help: "Cooling Group 4 Current",
		}),
	},
}
