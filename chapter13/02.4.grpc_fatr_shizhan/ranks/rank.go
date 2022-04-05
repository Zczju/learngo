package ranks

import (
	gobmi "github.com/armstrongli/go-bmi"
	"learn.go/chapter13/02.4.grpc_fatr_shizhan/apis"
	"learn.go/chapter13/02.4.grpc_fatr_shizhan/frinterface"
	"log"
	"math"
	"sort"
	"sync"
)

var _ frinterface.ServeInterface = &FatRateRank{}

type RankItem struct {
	Name    string
	Sex     string
	FatRate float64
}

type FatRateRank struct {
	items     []RankItem
	itemsLock sync.Mutex
}

func (r *FatRateRank) RegisterPersonalInformation(pi *apis.PersonalInformation) error {
	bmi, err := gobmi.BMI(float64(pi.Weight), float64(pi.Tall))
	if err != nil {
		log.Println("计算BMI失败111：", err)
		return err
	}
	fr := gobmi.CalcFatRate(bmi*100, int(pi.Age), pi.Sex)
	r.inputRecord(pi.Name, pi.Sex, fr)
	return nil
}

func (r *FatRateRank) UpdatePersonalInformation(pi *apis.PersonalInformation) (*apis.PersonalInformationFatRate, error) {
	bmi, err := gobmi.BMI(float64(pi.Weight), float64(pi.Tall))
	if err != nil {
		log.Println("计算BMI失败222：", err)
		return nil, err
	}
	fr := gobmi.CalcFatRate(bmi*100, int(pi.Age), pi.Sex)
	r.inputRecord(pi.Name, pi.Sex, fr)
	return &apis.PersonalInformationFatRate{
		Name:    pi.Name,
		FatRate: float32(fr),
	}, nil
}

func (r *FatRateRank) GetFatRate(name string) (*apis.PersonalRank, error) {
	rankId, sex, fr := r.getRank(name)
	return &apis.PersonalRank{
		Name:       name,
		Sex:        sex,
		RankNumber: int64(rankId),
		FatRate:    float32(fr),
	}, nil
}

func (r *FatRateRank) GetTop() ([]*apis.PersonalRank, error) {
	return r.getRankTop(), nil
}

func NewFatRateRank() *FatRateRank {
	return &FatRateRank{
		items: make([]RankItem, 0, 100),
	}
}

func (r *FatRateRank) inputRecord(name, sex string, fatRate ...float64) {
	r.itemsLock.Lock()
	defer r.itemsLock.Unlock()
	minFatRate := math.MaxFloat64
	for _, item := range fatRate {
		if minFatRate > item {
			minFatRate = item
		}
	}

	found := false
	for i, item := range r.items {
		if item.Name == name {
			if item.FatRate >= minFatRate {
				item.FatRate = minFatRate
			}
			r.items[i] = item
			found = true
			break
		}
	}
	if !found {
		r.items = append(r.items, RankItem{
			Name:    name,
			Sex:     sex,
			FatRate: minFatRate,
		})
	}
}

func (r *FatRateRank) getRank(name string) (rank int, sex string, fatRate float64) {
	r.itemsLock.Lock()
	defer r.itemsLock.Unlock()
	sort.Slice(r.items, func(i, j int) bool {
		return r.items[i].FatRate < r.items[j].FatRate
	})
	frs := map[float64]struct{}{}
	for _, item := range r.items {
		frs[item.FatRate] = struct{}{}
		if item.Name == name {
			fatRate = item.FatRate
		}
	}
	rankArr := make([]float64, 0, len(frs))
	for k := range frs {
		rankArr = append(rankArr, k)
	}
	sort.Float64s(rankArr)
	for i, frItem := range rankArr {
		if frItem == fatRate {
			rank = i + 1
			break
		}
	}
	return
}

func (r *FatRateRank) getRankTop() []*apis.PersonalRank {
	r.itemsLock.Lock()
	defer r.itemsLock.Unlock()
	sort.Slice(r.items, func(i, j int) bool {
		return r.items[i].FatRate < r.items[j].FatRate
	})
	out := make([]*apis.PersonalRank, 0, 10)
	for i := 0; i < 10 && i < len(r.items); i++ {
		item := r.items[i]
		out = append(out, &apis.PersonalRank{
			Name:       item.Name,
			Sex:        item.Sex,
			RankNumber: int64(i + 1),
			FatRate:    float32(item.FatRate),
		})
	}
	return out
}
