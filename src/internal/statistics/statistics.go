package statistics

import (
	"log/slog"
	"path/filepath"

	"github.com/sunbk201/ua3f/internal/paths"
)

type Recorder struct {
	RewriteRecordList     *RewriteRecordList
	PassThroughRecordList *PassThroughRecordList
	ConnectionRecordList  *ConnectionRecordList
}

func New() *Recorder {
	logDir, err := paths.EnsureLogDir()
	if err != nil {
		slog.Warn("statistics.New ensure log dir", slog.Any("error", err))
		logDir = "."
	}

	return &Recorder{
		RewriteRecordList:     NewRewriteRecordList(filepath.Join(logDir, "rewrite_stats")),
		PassThroughRecordList: NewPassThroughRecordList(filepath.Join(logDir, "pass_stats")),
		ConnectionRecordList:  NewConnectionRecordList(filepath.Join(logDir, "conn_stats")),
	}
}

func (r *Recorder) Start() {
	r.RewriteRecordList.Run()
	r.PassThroughRecordList.Run()
	r.ConnectionRecordList.Run()
}

func (r *Recorder) AddRecord(record any) {
	switch rec := record.(type) {
	case *RewriteRecord:
		select {
		case r.RewriteRecordList.recordAddChan <- rec:
		default:
		}
	case *PassThroughRecord:
		select {
		case r.PassThroughRecordList.recordAddChan <- rec:
		default:
		}
	case *ConnectionRecord:
		select {
		case r.ConnectionRecordList.recordAddChan <- rec:
		default:
		}
	}
}

func (r *Recorder) RemoveRecord(record any) {
	switch rec := record.(type) {
	case *ConnectionRecord:
		select {
		case r.ConnectionRecordList.recordRemoveChan <- rec:
		default:
		}
	}
}
