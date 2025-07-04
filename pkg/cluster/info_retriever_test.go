package cluster

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"ricebean/pkg/config"
)

func TestInfoRetrieverRegion(t *testing.T) {
	t.Parallel()

	c := viper.New()
	c.Set("pitaya.cluster.info.region", "us")
	conf := config.NewConfig(c)

	infoRetriever := NewInfoRetriever(*&config.NewPitayaConfig(conf).Cluster.Info)

	assert.Equal(t, "us", infoRetriever.Region())
}
