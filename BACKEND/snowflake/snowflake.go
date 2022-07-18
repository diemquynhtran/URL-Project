package snowflake

import (
	"fmt"
	"sync"
	"time"
)

// Snowfake is an object to generate the ID
type Snowfake struct {
	mu *sync.Mutex

	node uint64
	time uint64
	seq  uint64
}

// New creates new snowflake instance based on config.
// It returns an error when nodeID is greater than or equal to 2^nodeBits
func CreateNode(nodeID uint64) (*Snowfake, error) {

	if nodeID >= maxNode {
		return nil, fmt.Errorf("nodeID should less than %d", maxNode)
	}

	s := &Snowfake{}

	s.mu = &sync.Mutex{}
	s.node = nodeID
	s.time = 0
	s.seq = 0

	return s, nil
}

func (s *Snowfake) GenerateID() uint64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	t := s.now()
	if s.time == t {
		s.seq++
		s.seq &= 1<<seqBits - 1
	} else {
		s.seq = 0
	}

	s.time = t

	r := (s.time << timeShift) & timeMask
	r |= (s.node << nodeShift) & nodeMask
	r |= (s.seq << seqShift) & seqMask

	return r
}

func (s *Snowfake) now() uint64 {

	t := uint64(time.Now().Unix())
	t -= epoch

	return (1<<timeBits - 1) & t
}
