package main

import (
	"context"
	"google.golang.org/grpc"
	ac "google.golang.org/grpc/connectivity"
	"testing"
	"time"
)

func customReadyCheck(ctx context.Context, conn *grpc.ClientConn) ac.State {

}

func TestNewWithOption(t *testing.T) {
	type want struct {
		timeout           time.Duration
		checkReadyTimout  time.Duration
		heartbeatInterval time.Duration
		readyCheckFunc    ReadyCheckFunc
	}

	td := 123 * time.Second

	tables := []struct {
		name string
		args []TrackerOption
		want want
	}{

		{
			name: "default",
			args: []TrackerOption,
			want: want{
				timeout:           defaultTimeout,
				checkReadyTimout:  checkReadyTimeout,
				heartbeatInterval: heartbeatInterval,
				readyCheckFunc:    defaultReadyCheck,
			},
		},

		{
			name: "with timeout",
			args: []TrackerOption{WithTimeout(td)},
			want: want{
				timeout:           defaultTimeout,
				checkReadyTimout:  td,
				heartbeatInterval: heartbeatInterval,
				readyCheckFunc:    defaultReadyCheck,
			},
		},

		{
			name: "with heartbeat interval",
			args: []TrackerOption{WithHeartbeatInterval(td)},
			want: want{
				timeout:           defaultTimeout,
				checkReadyTimout:  checkReadyTimeout,
				heartbeatInterval: td,
				readyCheckFunc:    defaultReadyCheck,
			},
		},

		{
			name: "with custom ready check",
			args: []TrackerOption{WithCustomReadyCheckFunc(customReadyCheck)},
			want: want{
				timeout:           defaultTimeout,
				checkReadyTimout:  checkReadyTimeout,
				heartbeatInterval: heartbeatInterval,
				readyCheckFunc:    customReadyCheck,
			},
		},
	}

	for _, table := range tables {

		t.Run(table.name, func(t *testing.T) {

		})
	}

}
