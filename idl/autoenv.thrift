namespace go chatgpt.model.autoenv

include "base.thrift"


// 间隔<=255min
struct CronJob {
    1: EtcdEnv env
    2: i8 interval
}

struct CronJobReq {
    1: CronJob cron
}

struct CronJobResp {
    255: base.BaseResp baseResp
}

struct EtcdEnv {
    1: string name
    2: string value
}

struct GetEtcdEnvReq {
    1: string name
}

struct GetEtcdEnvResp {
    1: EtcdEnv env

    255: base.BaseResp baseResp
}

struct SetEtcdEnvReq {
    1: EtcdEnv env
}

struct SetEtcdEnvResp {
    255: base.BaseResp baseResp
}

service EtcdEnvService {
    GetEtcdEnvResp GetEtcdEnv(1: GetEtcdEnvReq req)
    SetEtcdEnvResp SetEtcdEnv(1: SetEtcdEnvReq req)
    CronJobResp CronJob(1: CronJobReq req)
}
