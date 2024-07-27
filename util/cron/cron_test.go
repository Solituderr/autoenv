package cron

import (
	"fmt"
	"testing"
	"time"
)

func TestCronJobImpl_Send(t *testing.T) {
	c := NewCronJobService()
	type args struct {
		name string
		spec string
		f    func()
	}
	tests_send := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case_1",
			args: args{
				name: "test",
				spec: "1 * * * *",
				f: func() {
					fmt.Println("dododo")
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests_send {
		t.Run(tt.name, func(t *testing.T) {
			err := c.Send(tt.args.name, tt.args.spec, tt.args.f)
			//time.Sleep(10 * time.Minute)
			if (err != nil) != tt.wantErr {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCronJobImpl_SendByInterval(t *testing.T) {
	c := NewCronJobService()
	type args struct {
		name     string
		interval time.Duration
		f        func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case_1",
			args: args{
				name:     "test",
				interval: 1 * time.Second,
				f: func() {
					fmt.Println("dododo")
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := c.SendByInterval(tt.args.name, tt.args.interval, tt.args.f)
			//time.Sleep(10 * time.Second)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendByInterval() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
