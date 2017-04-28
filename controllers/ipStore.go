package controllers

import (
	"log"
	"math/rand"
	"net/url"
	"sync"
)

var gIpStore *ipStore

type ipStore struct {
	// ip->urlType->count
	ipCounts map[string]map[string]int
	sync.RWMutex
}

func init() {
	gIpStore = &ipStore{
		ipCounts: make(map[string]map[string]int, 0),
	}
}

// 获取点击量最少的链接
func GetCkLeastUrl(ip string, ckUrls []string) string {
	return gIpStore.getCkLeastUrl(ip, ckUrls)
}

func (s *ipStore) getCount(ip, host string) int {
	s.RLock()
	defer s.RUnlock()
	if counts, ok := s.ipCounts[ip]; !ok {
		return 0
	} else {
		if num, ok := counts[host]; !ok {
			return 0
		} else {
			return num
		}
	}
}

func (s *ipStore) counts(ip, ckUrl string) {
	ckParsedUrl, err := url.Parse(ckUrl)
	if err != nil {
		return
	}

	host := ckParsedUrl.Host
	s.Lock()
	if counts, ok := s.ipCounts[ip]; !ok {
		s.ipCounts[ip] = make(map[string]int, 3)
		s.ipCounts[ip][host] = 1
	} else {
		if curNum, ok := counts[host]; !ok {
			counts[host] = 1
		} else {
			counts[host] = curNum + 1
		}
	}
	s.Unlock()
}

func (s *ipStore) getCkLeastUrl(ip string, ckUrls []string) string {
	if len(ckUrls) == 0 {
		return ""
	}

	var minCountsUrls []string
	var minCounts int
	for _, ckUrl := range ckUrls {
		ckParsedUrl, err := url.Parse(ckUrl)
		if err != nil {
			continue
		}

		host := ckParsedUrl.Host
		num := s.getCount(ip, host)
		if len(minCountsUrls) == 0 {
			minCounts = num
			minCountsUrls = []string{ckUrl}
		} else {
			if num < minCounts {
				minCounts = num
				minCountsUrls = []string{ckUrl}
			} else if num == minCounts {
				minCountsUrls = append(minCountsUrls, ckUrl)
			}
		}
	}

	// 多个点击量一样的就随机找一个
	minCountsUrl := minCountsUrls[rand.Intn(len(minCountsUrls))]
	log.Printf("select link:%s\n", minCountsUrl)
	s.counts(ip, minCountsUrl)
	return minCountsUrl
}
