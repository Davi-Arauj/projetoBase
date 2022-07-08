package health

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"go.uber.org/zap"

	"github.com/projetoBase/config"
)

var hb *heartbeat

type heartbeat struct {
	ApplicationName    string        `json:"application_name"`
	ApplicationVersion string        `json:"application_version"`
	BinaryCheckSum     string        `json:"binary_checksum"`
	ServiceTag         string        `json:"service_tag"`
	Check              bool          `json:"check"`
	DependsOn          []string      `json:"depends_on"`
	StartedAt          time.Time     `json:"started_at"`
	Requests           int64         `json:"requests"`
	Latency            time.Duration `json:"latency_mean"`
	WithError          int64         `json:"request_with_error"`
	Date               time.Time     `json:"date"`
}

// HandledRequest trata uma requisição realizada
func HandledRequest(latency time.Duration, error bool) {
	if hb == nil {
		log.Println("O heartbeat não foi inicializado")
		return
	}

	hb.Requests++
	hb.Latency = (latency + hb.Latency) / 2
	if error {
		hb.WithError++
	}
}

// Init inicializa o health
func Init(version string) {
	hb = &heartbeat{
		ApplicationName:    "academia",
		ApplicationVersion: version,
		BinaryCheckSum:     getChecksum(),
		ServiceTag:         getServiceTag(),
		StartedAt:          time.Now(),
		DependsOn: []string{
			"sql",
		},
	}

	go doTick()
}

func doTick() {
	var errChan = make(chan error)
	tck := time.NewTicker(time.Duration(config.ObterConfiguracao().HeartbeatIntervalo) * time.Second)
	for {
		select {
		case <-tck.C:
			go func() {
				zap.L().Info("Enviando Beat")
				beat(errChan)
				zap.L().Info("Beat enviada")
			}()

		case e := <-errChan:
			zap.L().Error("Erro ao enviar beat:", zap.Error(e))
		}
	}
}

func beat(errChan chan error) {
	hb.Check = true

	if err := checkDBs(); err != nil {
		errChan <- err
		hb.Check = false
	}

	hb.Date = time.Now()

	data, err := json.Marshal(hb)
	if err != nil {
		errChan <- err
		return
	}

	res, err := (&http.Client{Timeout: 3 * time.Second}).Post(
		config.ObterConfiguracao().HeartbeatEndpoint,
		"application/json",
		bytes.NewBuffer(data),
	)
	if err != nil {
		errChan <- err
		return
	}
	defer func() { _ = res.Body.Close() }()

	if res.StatusCode >= 400 {
		d, _ := ioutil.ReadAll(res.Body)
		errChan <- errors.New(string(d))
		return
	}

	hb.Latency = 0
	hb.Requests = 0
	hb.WithError = 0
}

func getChecksum() string {
	f, err := os.Open(os.Args[0])
	if err != nil {
		log.Println(err)
	}
	defer func() { _ = f.Close() }()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Println(err)
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}

func getServiceTag() string {
	data, err := ioutil.ReadFile("/tmp/product_serial")
	if err != nil {
		log.Println(err)
	}

	return strings.Replace(string(data), "\n", "", 1)
}

func checkDBs() error {
	return nil
}
