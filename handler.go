package main

import (
	"context"
	autoenv "github.com/Solituderr/autoenv/kitex_gen/chatgpt/model/autoenv"
)

// EtcdEnvServiceImpl implements the last service interface defined in the IDL.
type EtcdEnvServiceImpl struct{}

// GetEtcdEnv implements the EtcdEnvServiceImpl interface.
func (s *EtcdEnvServiceImpl) GetEtcdEnv(ctx context.Context, req *autoenv.GetEtcdEnvReq) (resp *autoenv.GetEtcdEnvResp, err error) {
	// TODO: Your code here...
	return
}

// SetEtcdEnv implements the EtcdEnvServiceImpl interface.
func (s *EtcdEnvServiceImpl) SetEtcdEnv(ctx context.Context, req *autoenv.SetEtcdEnvReq) (resp *autoenv.SetEtcdEnvResp, err error) {
	// TODO: Your code here...
	return
}

// CronJob implements the EtcdEnvServiceImpl interface.
func (s *EtcdEnvServiceImpl) CronJob(ctx context.Context, req *autoenv.CronJobReq) (resp *autoenv.CronJobResp, err error) {
	// TODO: Your code here...
	return
}
