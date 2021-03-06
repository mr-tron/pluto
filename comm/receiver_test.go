package comm

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"sync"
	"testing"
	"time"

	"io/ioutil"
	"path/filepath"

	"github.com/go-kit/kit/log"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

// Variables

var (
	confStatus = uint32(0)

	inc1 = []byte("5;hello")
	inc2 = []byte("53;\nwhat\r\tabout!\"§$%&/()=strange#+?`?`?°°°characters")
	inc3 = []byte("13;∰☕✔😉")
	inc4 = []byte("10;1234567890")
	inc5 = []byte(fmt.Sprintf("23;%g", math.MaxFloat64))

	/*
		payload1 := Msg{
			Replica:   "worker-1",
			Vclock:    map[string]uint32{"worker-1": uint32(1), "storage": uint32(0)},
			Operation: "create",
			Create: &Msg_CREATE{
				User:    "user1",
				Mailbox: "university",
				AddTag:  "aa59585f-5a5f-4ea9-887c-74ab2e3f1f4a",
			},
		}

		data, err := proto.Marshal(&payload1)
		if err != nil {
			return
		}

		data = append([]byte(fmt.Sprintf("%d;", len(data))), data...)

		for _, d := range data {
			fmt.Printf("%#x ", d)
		}
	*/
	writeApply1 = []byte{0x31, 0x30, 0x34, 0x3b, 0xa, 0x8, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2d, 0x31, 0x12, 0xc, 0xa, 0x8, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2d, 0x31, 0x10, 0x1, 0x12, 0xb, 0xa, 0x7, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x10, 0x0, 0x1a, 0x6, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x22, 0x39, 0xa, 0x5, 0x75, 0x73, 0x65, 0x72, 0x31, 0x12, 0xa, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x69, 0x74, 0x79, 0x1a, 0x24, 0x61, 0x61, 0x35, 0x39, 0x35, 0x38, 0x35, 0x66, 0x2d, 0x35, 0x61, 0x35, 0x66, 0x2d, 0x34, 0x65, 0x61, 0x39, 0x2d, 0x38, 0x38, 0x37, 0x63, 0x2d, 0x37, 0x34, 0x61, 0x62, 0x32, 0x65, 0x33, 0x66, 0x31, 0x66, 0x34, 0x61}

	/*
			payload2 := Msg{
			Replica:   "worker-1",
			Vclock:    map[string]uint32{"worker-1": uint32(1), "storage": uint32(0)},
			Operation: "create",
			Create: &Msg_CREATE{
				User:    "user2",
				Mailbox: "LongAndInterestingName",
				AddTag:  "525a3f40-7c2c-4b9a-94c8-a3432f25a28a",
			},
		}

		data, err := proto.Marshal(&payload2)
		if err != nil {
			return
		}

		data = append([]byte(fmt.Sprintf("%d;", len(data))), data...)

		for _, d := range data {
			fmt.Printf("%#x ", d)
		}

		fmt.Printf("\n---\n")

		payload3 := Msg{
			Replica:   "worker-1",
			Vclock:    map[string]uint32{"worker-1": uint32(2), "storage": uint32(0)},
			Operation: "delete",
			Delete: &Msg_DELETE{
				User:     "user2",
				Mailbox:  "LongAndInterestingName",
				RmvTags:  []string{"525a3f40-7c2c-4b9a-94c8-a3432f25a28a"},
				RmvMails: []string{"mail-message-name-generated-by-maildir"},
			},
		}

		data, err = proto.Marshal(&payload3)
		if err != nil {
			return
		}

		data = append([]byte(fmt.Sprintf("%d;", len(data))), data...)

		for _, d := range data {
			fmt.Printf("%#x ", d)
		}
	*/
	writeApply2 = []byte{0x31, 0x31, 0x36, 0x3b, 0xa, 0x8, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2d, 0x31, 0x12, 0xc, 0xa, 0x8, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2d, 0x31, 0x10, 0x1, 0x12, 0xb, 0xa, 0x7, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x10, 0x0, 0x1a, 0x6, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x22, 0x45, 0xa, 0x5, 0x75, 0x73, 0x65, 0x72, 0x32, 0x12, 0x16, 0x4c, 0x6f, 0x6e, 0x67, 0x41, 0x6e, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x24, 0x35, 0x32, 0x35, 0x61, 0x33, 0x66, 0x34, 0x30, 0x2d, 0x37, 0x63, 0x32, 0x63, 0x2d, 0x34, 0x62, 0x39, 0x61, 0x2d, 0x39, 0x34, 0x63, 0x38, 0x2d, 0x61, 0x33, 0x34, 0x33, 0x32, 0x66, 0x32, 0x35, 0x61, 0x32, 0x38, 0x61, 0x31, 0x35, 0x36, 0x3b, 0xa, 0x8, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2d, 0x31, 0x12, 0xc, 0xa, 0x8, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2d, 0x31, 0x10, 0x2, 0x12, 0xb, 0xa, 0x7, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x10, 0x0, 0x1a, 0x6, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2a, 0x6d, 0xa, 0x5, 0x75, 0x73, 0x65, 0x72, 0x32, 0x12, 0x16, 0x4c, 0x6f, 0x6e, 0x67, 0x41, 0x6e, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x24, 0x35, 0x32, 0x35, 0x61, 0x33, 0x66, 0x34, 0x30, 0x2d, 0x37, 0x63, 0x32, 0x63, 0x2d, 0x34, 0x62, 0x39, 0x61, 0x2d, 0x39, 0x34, 0x63, 0x38, 0x2d, 0x61, 0x33, 0x34, 0x33, 0x32, 0x66, 0x32, 0x35, 0x61, 0x32, 0x38, 0x61, 0x22, 0x26, 0x6d, 0x61, 0x69, 0x6c, 0x2d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2d, 0x6e, 0x61, 0x6d, 0x65, 0x2d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x62, 0x79, 0x2d, 0x6d, 0x61, 0x69, 0x6c, 0x64, 0x69, 0x72}
)

// Functions

// TestTriggerMsgApplier executes a white-box unit
// test on implemented TriggerMsgApplier() function.
func TestTriggerMsgApplier(t *testing.T) {

	// Create logger.
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	logger = log.With(logger,
		"caller", log.DefaultCaller,
	)

	// Bundle information in Receiver struct.
	recv := &Receiver{
		logger:        logger,
		name:          "worker-1",
		msgInLog:      make(chan struct{}, 1),
		stopTrigger:   make(chan struct{}),
		updateLogLock: &sync.Mutex{},
	}

	// Run trigger function.
	go func() {
		recv.TriggerMsgApplier(2)
	}()

	// Stop trigger function after specific number of seconds.
	go func() {
		<-time.After(7 * time.Second)
		recv.stopTrigger <- struct{}{}
		close(recv.msgInLog)
	}()

	numSignals := 0

	for range recv.msgInLog {
		numSignals++
	}

	assert.Equalf(t, 3, numSignals, "expected to receive 3 triggers but actually received %d", numSignals)
}

// TestIncoming executes a white-box unit
// test on implemented Incoming() function.
func TestIncoming(t *testing.T) {

	// Create logger.
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	logger = log.With(logger,
		"caller", log.DefaultCaller,
	)

	// Create temporary directory.
	dir, err := ioutil.TempDir("", "TestIncoming-")
	assert.Nilf(t, err, "failed to create temporary directory: %v", err)
	defer os.RemoveAll(dir)

	// Create path to temporary log file.
	tmpLogFile := filepath.Join(dir, "log")
	tmpMetaFile := filepath.Join(dir, "meta")

	// Open log file for writing.
	write, err := os.OpenFile(tmpLogFile, (os.O_CREATE | os.O_WRONLY | os.O_APPEND), 0600)
	assert.Nilf(t, err, "failed to open temporary log file for writing: %v", err)

	// Open log file for tracking meta data about already
	// applied parts of the CRDT update messages log file.
	meta, err := os.OpenFile(tmpMetaFile, (os.O_CREATE | os.O_WRONLY), 0600)
	assert.Nilf(t, err, "failed to open temporary log file for meta data: %v", err)

	// Bundle information in Receiver struct.
	recv := &Receiver{
		logger:        logger,
		name:          "worker-1",
		msgInLog:      make(chan struct{}, 1),
		updateLogPath: tmpLogFile,
		updateLogLock: &sync.Mutex{},
		updateLog:     write,
		metaFilePath:  tmpMetaFile,
		metaLog:       meta,
	}

	// Reset position in meta log file to beginning.
	_, err = recv.metaLog.Seek(0, os.SEEK_SET)
	assert.Nilf(t, err, "expected resetting of position in meta log not to fail but received: %v", err)

	// Value 1.
	// Write first value to log file.
	conf, err := recv.Incoming(context.Background(), &BinMsgs{
		Data: inc1,
	})
	assert.Nilf(t, err, "expected nil error for Incoming() but received: %v", err)

	// Wait for signal that new message was written to log.
	<-recv.msgInLog

	// Validate received confirmation struct.
	assert.Equalf(t, confStatus, conf.Status, "expected conf to carry Status=0 but found: %v", conf.Status)

	// Read content of log file for inspection.
	content, err := ioutil.ReadFile(tmpLogFile)
	assert.Nilf(t, err, "expected nil error for ReadFile() but received: %v", err)
	assert.Equalf(t, inc1, content, "expected '%s' in log file but found: %v", inc1, content)

	// Value 2.
	// Write second value to file.
	conf, err = recv.Incoming(context.Background(), &BinMsgs{
		Data: inc2,
	})
	assert.Nilf(t, err, "expected nil error for Incoming() but received: %v", err)

	<-recv.msgInLog

	assert.Equalf(t, confStatus, conf.Status, "expected conf to carry Status=0 but found: %v", conf.Status)

	content, err = ioutil.ReadFile(tmpLogFile)
	assert.Nilf(t, err, "expected nil error for ReadFile() but received: %v", err)

	content = bytes.TrimPrefix(content, inc1)
	assert.Equalf(t, inc2, content, "expected '%s' in log file but found: %v", inc2, content)

	// Value 3.
	// Write third value to file.
	conf, err = recv.Incoming(context.Background(), &BinMsgs{
		Data: inc3,
	})
	assert.Nilf(t, err, "expected nil error for Incoming() but received: %v", err)

	<-recv.msgInLog

	assert.Equalf(t, confStatus, conf.Status, "expected conf to carry Status=0 but found: %v", conf.Status)

	content, err = ioutil.ReadFile(tmpLogFile)
	assert.Nilf(t, err, "expected nil error for ReadFile() but received: %v", err)

	content = bytes.TrimPrefix(content, inc1)
	content = bytes.TrimPrefix(content, inc2)
	assert.Equalf(t, inc3, content, "expected '%s' in log file but found: %v", inc3, content)

	// Value 4.
	// Write fourth value to file.
	conf, err = recv.Incoming(context.Background(), &BinMsgs{
		Data: inc4,
	})
	assert.Nilf(t, err, "expected nil error for Incoming() but received: %v", err)

	<-recv.msgInLog

	assert.Equalf(t, confStatus, conf.Status, "expected conf to carry Status=0 but found: %v", conf.Status)

	content, err = ioutil.ReadFile(tmpLogFile)
	assert.Nilf(t, err, "expected nil error for ReadFile() but received: %v", err)

	content = bytes.TrimPrefix(content, inc1)
	content = bytes.TrimPrefix(content, inc2)
	content = bytes.TrimPrefix(content, inc3)
	assert.Equalf(t, inc4, content, "expected '%s' in log file but found: %v", inc4, content)

	// Value 5.
	// Write fifth value to file.
	conf, err = recv.Incoming(context.Background(), &BinMsgs{
		Data: inc5,
	})
	assert.Nilf(t, err, "expected nil error for Incoming() but received: %v", err)

	<-recv.msgInLog

	assert.Equalf(t, confStatus, conf.Status, "expected conf to carry Status=0 but found: %v", conf.Status)

	content, err = ioutil.ReadFile(tmpLogFile)
	assert.Nilf(t, err, "expected nil error for ReadFile() but received: %v", err)

	content = bytes.TrimPrefix(content, inc1)
	content = bytes.TrimPrefix(content, inc2)
	content = bytes.TrimPrefix(content, inc3)
	content = bytes.TrimPrefix(content, inc4)
	assert.Equalf(t, inc5, content, "expected '%s' in log file but found: %v", inc5, content)
}

// TestApplyStoredMsgs executes a white-box unit
// test on implemented ApplyStoredMsgs() function.
func TestApplyStoredMsgs(t *testing.T) {

	// Create logger.
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	logger = log.With(logger,
		"caller", log.Caller(4),
	)

	// Create temporary directory.
	dir, err := ioutil.TempDir("", "TestApplyStoredMsgs-")
	assert.Nilf(t, err, "failed to create temporary directory: %v", err)
	defer os.RemoveAll(dir)

	// Create path to temporary log files.
	tmpLogFile := filepath.Join(dir, "log")
	tmpMetaFile := filepath.Join(dir, "meta")
	tmpVClockFile := filepath.Join(dir, "vclock")

	// Write binary encoded test message to log file.
	err = ioutil.WriteFile(tmpLogFile, writeApply1, 0600)
	assert.Nilf(t, err, "expected writing test content 1 to log file not to fail but received: %v", err)

	// Open log file for writing.
	write, err := os.OpenFile(tmpLogFile, (os.O_CREATE | os.O_WRONLY | os.O_APPEND), 0600)
	assert.Nilf(t, err, "failed to open temporary log file for writing: %v", err)

	// Open log file for tracking meta data about already
	// applied parts of the CRDT update messages log file.
	meta, err := os.OpenFile(tmpMetaFile, (os.O_CREATE | os.O_WRONLY), 0600)
	assert.Nilf(t, err, "failed to open temporary log file for meta data: %v", err)

	// Open log file of last known vector clock values.
	vclockLog, err := os.OpenFile(tmpVClockFile, (os.O_CREATE | os.O_RDWR), 0600)
	assert.Nilf(t, err, "failed to open temporary vector clock file: %v", err)

	// Simulate nodes.
	nodes := map[string]string{
		"other-node-1": "10.0.0.1",
		"other-node-2": "10.10.0.23",
		"other-node-3": "10.255.0.91",
	}

	// Bundle information in Receiver struct.
	recv := &Receiver{
		logger:           logger,
		name:             "worker-1",
		msgInLog:         make(chan struct{}, 1),
		updateLogPath:    tmpLogFile,
		updateLogLock:    &sync.Mutex{},
		updateLog:        write,
		metaFilePath:     tmpMetaFile,
		metaLog:          meta,
		vclock:           make(map[string]uint32),
		vclockLock:       &sync.Mutex{},
		vclockLog:        vclockLog,
		stopApply:        make(chan struct{}),
		applyCRDTUpdChan: make(chan Msg),
		doneCRDTUpdChan:  make(chan struct{}),
		nodes:            nodes,
	}

	// Reset position in meta log file to beginning.
	_, err = recv.metaLog.Seek(0, os.SEEK_SET)
	assert.Nilf(t, err, "expected resetting of position in meta log not to fail but received: %v", err)

	// Reset position in vector clock file to beginning.
	_, err = recv.vclockLog.Seek(0, os.SEEK_SET)
	assert.Nilf(t, err, "expected resetting of position in vector clock file not to fail but received: %v", err)

	// Set vector clock entries to 0.
	for node := range nodes {
		recv.vclock[node] = 0
	}

	// Including the entry of this node.
	recv.vclock[recv.name] = 0

	// Run apply function to test.
	go func() {
		recv.ApplyStoredMsgs()
	}()

	// Send msgInLog trigger to start apply function.
	recv.msgInLog <- struct{}{}

	// Receive message to apply in correct channel.
	msg, ok := <-recv.applyCRDTUpdChan
	assert.Equalf(t, true, ok, "expected waiting for message on channel to succeed but received: %v", ok)

	// Signal waiting apply function that message was
	// applied successfully at CRDT level.
	recv.doneCRDTUpdChan <- struct{}{}

	// Stop apply function.
	recv.stopApply <- struct{}{}

	// Check received message for correctness.
	assert.Equalf(t, "worker-1", msg.Replica, "expected 'worker-1' as Replica in msg but received: %v", msg.Replica)
	assert.Equalf(t, map[string]uint32{"worker-1": uint32(1), "storage": uint32(0)}, msg.Vclock, "expected 'worker-1:1 storage:0' as Vclock in msg but received: %v", msg.Vclock)
	assert.Equalf(t, "create", msg.Operation, "expected 'create' as Operation in msg but received: %v", msg.Operation)
	assert.Equalf(t, (*Msg_DELETE)(nil), msg.Delete, "expected no Delete entry in msg but received: %v", msg.Delete)
	assert.Equalf(t, (*Msg_APPEND)(nil), msg.Append, "expected no Append entry in msg but received: %v", msg.Append)
	assert.Equalf(t, (*Msg_EXPUNGE)(nil), msg.Expunge, "expected no Expunge entry in msg but received: %v", msg.Expunge)
	assert.Equalf(t, (*Msg_STORE)(nil), msg.Store, "expected no Store entry in msg but received: %v", msg.Store)
	assert.Equalf(t, "user1", msg.Create.User, "expected 'user1' as msg.Create.User but received: %v", msg.Create.User)
	assert.Equalf(t, "university", msg.Create.Mailbox, "expected 'university' as msg.Create.Mailbox but received: %v", msg.Create.Mailbox)
	assert.Equalf(t, "aa59585f-5a5f-4ea9-887c-74ab2e3f1f4a", msg.Create.AddTag, "expected 'aa59585f-5a5f-4ea9-887c-74ab2e3f1f4a' as msg.Create.AddMailbox.Tag but received: %v", msg.Create.AddTag)

	// Check file system content of log file.
	content, err := ioutil.ReadFile(tmpLogFile)
	assert.Nilf(t, err, "expected nil error for ReadFile() but received: %v", err)
	assert.Equalf(t, writeApply1, content, "expected '%s' in log file but found: %v", writeApply1, content)

	// Check file system content of vector clock file.
	content, err = ioutil.ReadFile(tmpVClockFile)
	assert.Nilf(t, err, "expected nil error for ReadFile() but received: %v", err)
	assert.True(t, bytes.Contains(content, []byte("worker-1:1")), "expected 'worker-1:1' to be present in vector clock file but was not")

	// Write second binary encoded test message to log file.
	err = ioutil.WriteFile(tmpLogFile, writeApply2, 0600)
	assert.Nilf(t, err, "expected writing test content 2 to log file not to fail but received: %v", err)

	// Reset meta data log file.
	err = recv.metaLog.Truncate(0)
	assert.Nilf(t, err, "expected truncating meta data log file not to fail but received: %v", err)

	// Reset vector clock internally.
	recv.vclock["worker-1"] = uint32(0)
	err = recv.SaveVClockEntries()
	assert.Nilf(t, err, "expected reset writing of vector clock file not to fail but received: %v", err)

	// Run apply function again for second test.
	go func() {
		recv.ApplyStoredMsgs()
	}()

	// Send msgInLog trigger to start apply function.
	recv.msgInLog <- struct{}{}

	// Receive message to apply in correct channel.
	msg, ok = <-recv.applyCRDTUpdChan
	assert.Equalf(t, true, ok, "expected waiting for message on channel to succeed but received: %v", ok)

	// Signal waiting apply function that message was
	// applied successfully at CRDT level.
	recv.doneCRDTUpdChan <- struct{}{}

	time.Sleep(1 * time.Second)

	// Check received message for correctness.
	assert.Equalf(t, "worker-1", msg.Replica, "expected 'worker-1' as Replica in msg but received: %v", msg.Replica)
	assert.Equalf(t, map[string]uint32{"worker-1": uint32(1), "storage": uint32(0)}, msg.Vclock, "expected 'worker-1:1 storage:0' as Vclock in msg but received: %v", msg.Vclock)
	assert.Equalf(t, "create", msg.Operation, "expected 'create' as Operation in msg but received: %v", msg.Operation)
	assert.Equalf(t, (*Msg_DELETE)(nil), msg.Delete, "expected no Delete entry in msg but received: %v", msg.Delete)
	assert.Equalf(t, (*Msg_APPEND)(nil), msg.Append, "expected no Append entry in msg but received: %v", msg.Append)
	assert.Equalf(t, (*Msg_EXPUNGE)(nil), msg.Expunge, "expected no Expunge entry in msg but received: %v", msg.Expunge)
	assert.Equalf(t, (*Msg_STORE)(nil), msg.Store, "expected no Store entry in msg but received: %v", msg.Store)
	assert.Equalf(t, "user2", msg.Create.User, "expected 'user2' as msg.Create.User but received: %v", msg.Create.User)
	assert.Equalf(t, "LongAndInterestingName", msg.Create.Mailbox, "expected 'LongAndInterestingName' as msg.Create.Mailbox but received: %v", msg.Create.Mailbox)
	assert.Equalf(t, "525a3f40-7c2c-4b9a-94c8-a3432f25a28a", msg.Create.AddTag, "expected '525a3f40-7c2c-4b9a-94c8-a3432f25a28a' as msg.Create.AddMailbox.Tag but received: %v", msg.Create.AddTag)

	// Check file system content of log file.
	content, err = ioutil.ReadFile(tmpLogFile)
	assert.Nilf(t, err, "expected nil error for ReadFile() but received: %v", err)
	assert.Equalf(t, writeApply2, content, "expected '%s' in log file but found: %v", writeApply2, content)

	// Check file system content of vector clock file.
	content, err = ioutil.ReadFile(tmpVClockFile)
	assert.Nilf(t, err, "expected nil error for ReadFile() but received: %v", err)
	assert.True(t, bytes.Contains(content, []byte("worker-1:1")), "expected 'worker-1:1' to be present in vector clock file but was not")

	// Send msgInLog trigger to start apply function.
	recv.msgInLog <- struct{}{}

	// Receive message to apply in correct channel.
	msg, ok = <-recv.applyCRDTUpdChan
	assert.Equalf(t, true, ok, "expected waiting for message on channel to succeed but received: %v", ok)

	// Signal waiting apply function that message was
	// applied successfully at CRDT level.
	recv.doneCRDTUpdChan <- struct{}{}

	time.Sleep(1 * time.Second)

	// Check received message for correctness.
	assert.Equalf(t, "worker-1", msg.Replica, "expected 'worker-1' as Replica in msg but received: %v", msg.Replica)
	assert.Equalf(t, map[string]uint32{"worker-1": uint32(2), "storage": uint32(0)}, msg.Vclock, "expected 'worker-1:2 storage:0' as Vclock in msg but received: %v", msg.Vclock)
	assert.Equalf(t, "delete", msg.Operation, "expected 'delete' as Operation in msg but received: %v", msg.Operation)
	assert.Equalf(t, (*Msg_CREATE)(nil), msg.Create, "expected no Create entry in msg but received: %v", msg.Create)
	assert.Equalf(t, (*Msg_APPEND)(nil), msg.Append, "expected no Append entry in msg but received: %v", msg.Append)
	assert.Equalf(t, (*Msg_EXPUNGE)(nil), msg.Expunge, "expected no Expunge entry in msg but received: %v", msg.Expunge)
	assert.Equalf(t, (*Msg_STORE)(nil), msg.Store, "expected no Store entry in msg but received: %v", msg.Store)
	assert.Equalf(t, "user2", msg.Delete.User, "expected 'user2' as msg.Delete.User but received: %v", msg.Delete.User)
	assert.Equalf(t, "LongAndInterestingName", msg.Delete.Mailbox, "expected 'LongAndInterestingName' as msg.Delete.Mailbox but received: %v", msg.Delete.Mailbox)
	assert.Equalf(t, "525a3f40-7c2c-4b9a-94c8-a3432f25a28a", msg.Delete.RmvTags[0], "expected '525a3f40-7c2c-4b9a-94c8-a3432f25a28a' as msg.Delete.RmvTags[0] but received: %v", msg.Delete.RmvTags[0])
	assert.Equalf(t, "mail-message-name-generated-by-maildir", msg.Delete.RmvMails[0], "expected 'mail-message-name-generated-by-maildir' as msg.Delete.RmvMails[0] but received: %v", msg.Delete.RmvMails[0])

	// Check file system content of log file.
	content, err = ioutil.ReadFile(tmpLogFile)
	assert.Nilf(t, err, "expected nil error for ReadFile() but received: %v", err)
	assert.Equalf(t, writeApply2, content, "expected '%s' in log file but found: %v", writeApply2, content)

	// Check file system content of vector clock file.
	content, err = ioutil.ReadFile(tmpVClockFile)
	assert.Nilf(t, err, "expected nil error for ReadFile() but received: %v", err)
	assert.True(t, bytes.Contains(content, []byte("worker-1:2")), "expected 'worker-1:2' to be present in vector clock file but was not")
}
