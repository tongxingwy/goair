package goair

type MirrorStreamInfo struct {
    sessionID uint64 `plist:"sessionID"`
    latencyMs uint64 `plist:"latencyMs"`
    fpsInfo  []map[string]string //`plist:"fpsInfo"`
    version  string `plist:"version"`
    channel uint64  `plist:"channel"`
    connectTime float64 `plist:"connectTime"`
    deviceInfoTime float64 `plist:"deviceInfoTime"`
    prepareTime float64 `plist:"prepareTime"`
    configTime float64 `plist:"configTime"`
    param2 []byte `plist:"param2"`
    authTime float64 `plist:"authTime"`
    timestampInfo []map[string]string //`plist:"timestampInfo"`
    param1 []byte `plist:"param1"`
    deviceID uint64 `plist:"deviceID"`
    resolveDNSTime float64 `plist:"resolveDNSTime"`
    querySRVTime float64 `plist:"querySRVTime"`
    macAddress float64 `plist:"macAddress"`
}
