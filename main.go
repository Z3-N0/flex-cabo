package main

import (
	"time"
)

const (
	defaultSessionTimeout = 15 * time.Minute
	defaultPort           = "8080"
)

type Params struct {
	AppConfigPath   string
	InfraConfigPath string
	SessionTimeout  time.
	Mode            string
	StdErrFlag      bool
	TraceFlag       bool
}

func main(){
	Params := ParseFlags()
	if err:= appconfig.


}


func ParseFlags() *Params{
	return &Params{}
}
