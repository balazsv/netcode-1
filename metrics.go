package netcode

import (
	"fmt"
	"github.com/rcrowley/go-metrics"
)

func getMeter(name string, registry metrics.Registry) metrics.Meter {
	defer func() {
		if r := recover(); r!=nil {
		}
	}()

	meter := registry.Get(name)

	if meter == nil {
		meter = metrics.NewMeter()
		registry.Register(name, meter)
	}

	return meter.(metrics.Meter)
}

func getHistogram(name string, size int, registry metrics.Registry) metrics.UniformSample {
	defer func() {
		if r := recover(); r!=nil {
		}
	}()

	sample := registry.Get(name)

	if sample == nil {
		sample = metrics.NewUniformSample(size)
		registry.Register(name, sample)
	}

	return sample.(metrics.UniformSample)
}

func getPacketMetricsName(packetType PacketType, severity, prefix string, postfix string) string {
	return fmt.Sprintf("udpclient:%s/%s%s%s", severity, prefix, packetTypeMap[packetType], postfix)
}
