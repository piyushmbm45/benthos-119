package pure_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/nehal119/benthos-119/pkg/component/input"
	"github.com/nehal119/benthos-119/pkg/manager"
	bmock "github.com/nehal119/benthos-119/pkg/manager/mock"
	"github.com/nehal119/benthos-119/pkg/message"
)

func TestInprocDryRun(t *testing.T) {
	ctx, done := context.WithTimeout(context.Background(), time.Second*30)
	defer done()

	t.Parallel()

	mgr, err := manager.New(manager.NewResourceConfig())
	if err != nil {
		t.Fatal(err)
	}

	mgr.SetPipe("foo", make(chan message.Transaction))

	conf := input.NewConfig()
	conf.Type = "inproc"
	conf.Inproc = "foo"

	ip, err := mgr.NewInput(conf)
	require.NoError(t, err)

	<-time.After(time.Millisecond * 100)

	ip.TriggerStopConsuming()
	if err = ip.WaitForClose(ctx); err != nil {
		t.Error(err)
	}
}

func TestInprocDryRunNoConn(t *testing.T) {
	ctx, done := context.WithTimeout(context.Background(), time.Second*30)
	defer done()

	t.Parallel()

	conf := input.NewConfig()
	conf.Type = "inproc"
	conf.Inproc = "foo"

	ip, err := bmock.NewManager().NewInput(conf)
	require.NoError(t, err)

	<-time.After(time.Millisecond * 100)

	ip.TriggerStopConsuming()
	if err = ip.WaitForClose(ctx); err != nil {
		t.Error(err)
	}
}