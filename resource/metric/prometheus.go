package metric

import (
	"github.com/chaihaobo/gocommon/metric"

	"github.com/chaihaobo/boice-blog-api/resource/config"
)

type (
	PrometheusMetric metric.PrometheusMetric
)

func NewPrometheusMetric(config *config.Configuration) (PrometheusMetric, error) {
	return metric.NewPrometheusMetric(metric.Config{
		Port:        config.Service.MetricPort,
		ServiceName: config.Service.Name,
	})
}
