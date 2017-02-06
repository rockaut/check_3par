package hpe3par

/*
SystemInformation returned by api/v1/system
*/
type SystemInformation struct {
    ID                   uint32
    Name                 string
    IPv4Addr             string
    IPv6Addr             string
    Model                string
    SerialNumber         string
    SystemVersion        string
    TotalNodes           uint32
    MasterNode           uint32
    OnlineNodes          []uint32
    ClusterNodes         []uint32
    ChunkletSize         uint32
    TotalCapacityMiB     uint32
    AllocatedCapacityMiB uint32
    FreeCapacityMiB      uint32
    FailedCapacityMiB    uint32
    Location             string
    Owner                string
    Contact              string
    Comment              string
    TimeZone             string
    FlashCachePolicy     FlashCachePolicy
    LicenseInfo          string
    Parameters           SystemParameter
}

/*
FlashCachePolicy enumeration
*/
type FlashCachePolicy int8

/*
   FlashCachePolicy enumeration states
   FLASHCACHEPOLICYVALUE_UNKNOWN  = 0
   FLASHCACHEPOLICYVALUE_ENABLED  = 1
   FLASHCACHEPOLICYVALUE_DISABLED = 2
   FLASHCACHEPOLICYVALUE_CLEARED  = 3
*/
const (
    FLASHCACHEPOLICYVALUE_UNKNOWN  FlashCachePolicy = 0
    FLASHCACHEPOLICYVALUE_ENABLED                   = 1
    FLASHCACHEPOLICYVALUE_DISABLED                  = 2
    FLASHCACHEPOLICYVALUE_CLEARED                   = 3
)

var flashCachePolicyDescriptions = [...]string{
    "Unknown",
    "Enabled",
    "Disabled",
    "Cleared",
}

func (fcp FlashCachePolicy) String() string {
    return flashCachePolicyDescriptions[fcp]
}

/*
LicenseInfo object regarding the system license(s)
*/
type LicenseInfo struct {
    IssueTimeSec int
    DiskCount    int32
    WWNBASE      string
    Licenses     []License
    LicenseState LicenseState
}

/*
License object defining the license itself
*/
type License struct {
    Name           string
    ExpiryTimeSec  int
    ExpiryTime8601 string
}

/*
LicenseState object defining the features of the license
*/
type LicenseState struct {
    RemoteCopy           bool
    ThinProvisioning     bool
    Domains              bool
    DynamicOptimization  bool
    VirtualLock          bool
    ThinPersistence      bool
    ThinConversion       bool
    AdaptiveOptimization bool
    PeerVirtualization   bool
    Qos                  bool
    SystemReporter       bool
    DarEncryption        bool
    FileServices         bool
}

/*
SystemParameter defining the available system parameters
*/
type SystemParameter struct {
    RawSpaceAlertFC        uint32
    RawSpaceAlertNL        uint32
    RawSpaceAlertSSD       uint32
    RemoteSyslog           bool
    RemoteSyslogHost       string
    SparingAlgorithm       string
    EventLogSize           uint32
    VVRetentionTimeMax     uint32
    UpgradeNote            string
    PortFailoverEnabled    bool
    AutoExportAfterReboot  bool
    AllowR5OnNLDrives      bool
    AllowR0                bool
    ThermalShutdown        bool
    FailoverMatchedSet     bool
    SessionTimeout         uint32
    HostDIF                bool
    AllowWrtbackSingleNode bool
    AllowWrtbackUpgrade    bool
    DisableDedup           bool
}
