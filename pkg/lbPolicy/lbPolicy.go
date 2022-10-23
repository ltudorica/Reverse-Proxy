package lbPolicy

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"log"
	math_rand "math/rand"
	"net/url"
	"sync"
	"sync/atomic"

	"bitbucket.org/ltudorica/App/pkg/parsing"
)

type LB interface {
	Next() *Server
	AddServers(server *url.URL)
	GetServers() *[]Server
	CountServers() int
}

type loadBalancer struct {
	servers []Server
	next    uint32
}

type Server struct {
	URL    *url.URL
	IsDead bool
	mu     sync.RWMutex
}

func New(urls []*url.URL) (LB, error) {

	var lb LB = &loadBalancer{}

	for i := 0; i < len(urls); i++ {
		lb.AddServers(urls[i])
	}

	return lb, nil
}

func (lb *loadBalancer) Next() *Server {
	if parsing.GetLB_Policy() == "ROUND_ROBIN" {
		log.Println("In ROUND_ROBIN policy")
		n := atomic.AddUint32(&lb.next, 1)
		return &lb.servers[(int(n)-1)%len(lb.servers)]
	} else {
		log.Println("In RANDOM policy / default")
		var b [8]byte
		_, err := crypto_rand.Read(b[:])
		if err != nil {
			panic("cannot seed math/rand package with cryptographically secure random number generator")
		}
		math_rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
		max := lb.CountServers()
		n := 0 + math_rand.Intn(max-0)
		return &lb.servers[n]
	}
}

func (lb *loadBalancer) AddServers(url *url.URL) {
	lb.servers = append(lb.servers, Server{URL: url})
}

func (lb *loadBalancer) GetServers() *[]Server {
	return &lb.servers
}

func (lb *loadBalancer) CountServers() int {
	return len(lb.servers)
}

func (s *Server) SetDead(d bool) {
	s.mu.Lock()
	s.IsDead = d
	s.mu.Unlock()
}

func (s *Server) GetIsDead() bool {
	s.mu.RLock()
	isAlive := s.IsDead
	s.mu.RUnlock()
	return isAlive
}
