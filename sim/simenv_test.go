package sim

import (
	"testing"
)

func TestEnv_envFile(t *testing.T) {
	env,err:=envFile("./testdata/foo.json")
	if err == nil {
		t.Errorf("expected to see error because file dosen't exist")
	}
	env,err=envFile("./testdata/env.json")
	if env==nil{
		t.Errorf("expected new env")
	}

}
