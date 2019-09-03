package main

import (
	"log"
	"sync/atomic"
)

// SequenceCache contains a cache for sequences
type SequenceCache struct {
	cur   uint64
	upper uint64
}

// the uuid of the sequence
const seqId string = "663f3432-02c5-4371-a3b5-c3d7eda604ca"

// every time cache 1000 steps
const stepSize uint64 = 1000


func createSequenceCache(cassandraConfig CassandraConfig) (*SequenceCache, error) {
	var err error
	session, err = createCqlSession(cassandraConfig)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var seq uint64
	iter := session.Query("select seq from sequences").Iter()
	for iter.Scan(&seq) {
	}
	if err := iter.Close(); err != nil {
		log.Fatal("failed to get id from sequences", err)
		return nil, err
	}
	var next = seq + stepSize + 1
	if err := session.Query("update sequences set seq = ? where id = ?", next, seqId).Exec(); err != nil {
		log.Fatal("failed to update sequences", err)
		return nil, err
	}

	return &SequenceCache{seq, seq + stepSize}, nil
}

func (sequenceCache *SequenceCache) getSeq() uint64 {
	seq := atomic.AddUint64(&sequenceCache.cur, 1)
	// TODO  can we avoid blocking here?
	if seq == sequenceCache.upper {
		sequenceCache.refresh()
	}
	return seq
}

func (sequenceCache *SequenceCache) refresh() {

}
