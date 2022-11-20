package system

import (
    "github.com/docker/docker/pkg/parsers/kernel"
    "github.com/shirou/gopsutil/v3/cpu"
    "runtime"
)

type Information struct {
    Version       string `json:"version"`
    KernelVersion string `json:"kernel_version"`
    Architecture  string `json:"architecture"`
    OS            string `json:"os"`
    CpuModel      string `json:"cpu_model"`
    CpuCount      int    `json:"cpu_count"`
}

func GetSystemInformation() (*Information, error) {
    	k, err := kernel.GetKernelVersion()
    	if err != nil {
       	    return nil, err
    	}

   	c, err := cpu.Info()
    	if err != nil {
            return nil, err
    	}

    	s := &Information{
            Version:       Version,
            KernelVersion: k.String(),
            Architecture:  runtime.GOARCH,
            OS:            runtime.GOOS,
            CpuModel:      c[0].ModelName,
            CpuCount:      runtime.NumCPU(),
    	}

    return s, nil
}
