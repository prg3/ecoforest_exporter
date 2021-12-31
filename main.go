package main

import (
	"net/http"
	"time"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"fmt"
    "log"
	"crypto/tls"
	"strings"
	"io"
	"strconv"
)

func getData() {
	interval := 10
	dataTags := [][]int{
		// Analog (1)
		{1, 2002, 1, 39},
		{1, 2002, 40, 19},
		{1, 2002, 97, 30},
		{1, 2002, 176, 29},
		{1, 2002, 214, 7},
		// Integers (2)
		// {2, 2001, 5033, 3},
		// {2, 2001, 5066, 18},
		// {2, 2001, 5113, 31},
		// {2, 2001, 5185, 27},
		// {2, 2001, 5241, 34},
		// {2, 2001, 5285, 14}, 
		//  Binary (3)
		// {3, 2001, 61, 25},
		// {3, 2001, 101, 97},
		// {3, 2001, 206, 45},
	}
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	user := os.Getenv("ECOFOREST_USER")
	pass := os.Getenv("ECOFOREST_PASSS")
	ip := os.Getenv("ECOFOREST_IP")
	ecoforestUrl := fmt.Sprintf("https://%s:8000/recepcion_datos_4.cgi", ip)

	go func() {
		for {

			for _, tag := range dataTags {
				body := fmt.Sprintf("idOperacion=%d&dir=%d&num=%d", tag[1], tag[2], tag[3])
				client := &http.Client{}
				req, err := http.NewRequest("POST", ecoforestUrl, strings.NewReader(body) )
				if err != nil {
					log.Fatal(err)
				}
				req.SetBasicAuth(user, pass)
				resp, err := client.Do(req)
				if err != nil {
					log.Fatal(err)
				}
				defer resp.Body.Close()

				if resp.StatusCode == http.StatusOK {
					bodyBytes, err := io.ReadAll(resp.Body)
					if err != nil {
						log.Fatal(err)
					}
					bodyString := string(bodyBytes)
					ecoData := strings.Split(bodyString,"\n")
					for i, v := range strings.Split(ecoData[1],"&") {
						// fmt.Println(v)

						if i >= 2 {
							val, err := strconv.ParseUint(v, 16, 16)
							if err != nil {
								log.Fatal(err)
							}
							f := float64(int16(val)) / 10
							fmt.Println(fmt.Sprintf("%d:%d:%d - %f", tag[1],tag[2], i,  f))
							setCounter(tag[0], tag[1], tag[2], tag[3], i, f)
						}
					}
					// fmt.Println(strings.Split(ecoData[1],"&"))

				}
			}
			time.Sleep(time.Duration(interval) * time.Second)
		}
	}()
}



func recordMetrics() {
	go func() {
			for {
					opsProcessed.Inc()
					time.Sleep(2 * time.Second)
			}
	}()
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
			Name: "myapp_processed_ops_total",
			Help: "The total number of processed events",
	})
)

func main() {
	recordMetrics()
	getData()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
