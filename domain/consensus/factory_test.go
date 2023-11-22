package consensus

import (
	"io/ioutil"
	"testing"

	"github.com/Pyrinpyi/pyipad/domain/prefixmanager/prefix"

	"github.com/Pyrinpyi/pyipad/domain/dagconfig"
	"github.com/Pyrinpyi/pyipad/infrastructure/db/database/ldb"
)

func TestNewConsensus(t *testing.T) {
	f := NewFactory()

	config := &Config{Params: dagconfig.DevnetParams}

	tmpDir, err := ioutil.TempDir("", "TestNewConsensus")
	if err != nil {
		return
	}

	db, err := ldb.NewLevelDB(tmpDir, 8)
	if err != nil {
		t.Fatalf("error in NewLevelDB: %s", err)
	}

	_, shouldMigrate, err := f.NewConsensus(config, db, &prefix.Prefix{}, nil)
	if err != nil {
		t.Fatalf("error in NewConsensus: %+v", err)
	}

	if shouldMigrate {
		t.Fatalf("A fresh consensus should never return shouldMigrate=true")
	}
}
