package snowfake
/*
import (
	"fmt"
	"sync"
	"time"
)

var (
	UNUSED_BITS   uint8 = 1
	EPOCH_BITS    uint8 = 41
	NODE_ID_BITS  uint8 = 10
	SEQUENCE_BITS uint8 = 12

	timeShift uint8 = SEQUENCE_BITS + NODE_ID_BITS
	nodeShift uint8 = SEQUENCE_BITS
	seqShift  uint8 = 0

	epochTime   uint64 = 1288834974657
	maxTimeBit  uint64 = (1 << EPOCH_BITS) - 1
	maxNodeId   uint64 = (1 << NODE_ID_BITS) - 1
	maxSequence uint64 = (1 << SEQUENCE_BITS) - 1

	timeMask uint64 = maxTimeBit << uint64(timeShift)
	nodeMask uint64 = maxNodeId << uint64(nodeShift)
	seqMask  uint64 = maxSequence << uint64(seqShift)

	//id int64

)

type Snowflake struct {
	mutex *sync.Mutex
	node  uint64
	time  uint64
	seq   uint64
}

// func CreateNodeId()  {
// }

func CreateNewNode(nodeID uint64) (*Snowflake, error) {
	if nodeID >= maxNodeId {
		return nil, fmt.Errorf("NodeID too big")
	}

	s := &Snowflake{}
	s.mutex = &sync.Mutex{}
	s.node = nodeID
	s.time = 0
	s.seq = 0
	return s, nil
}

func (s *Snowflake) Generate() uint64 {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	curTime := s.currentTime()
	if s.time == curTime {
		s.seq++
		s.seq &= maxSequence
	} else {
		s.seq = 0
	}
	s.time = curTime
	r := (s.time << timeShift) & timeMask
	r |= (s.node << nodeShift) & nodeMask
	r |= (s.seq << seqShift) & seqMask

	return r

}

func (s *Snowflake) currentTime() uint64 {

	t := uint64(time.Now().Unix())
	t -= epochTime

	return (1<<EPOCH_BITS - 1) & t
}
*/
