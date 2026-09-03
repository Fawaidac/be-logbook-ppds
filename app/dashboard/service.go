package dashboard

import (
	"context"
	"math"
	"strings"
	"time"

	"be-logbook-ppds/app/pendidikan"
)

type Service interface {
	GetDashboardSummary(ctx context.Context, year int, username, name string) (*DashboardSummaryResponse, error)
	GetLaporanSummary(ctx context.Context, periode, stase string, username, name string) (*LaporanSummaryResponse, error)
}

type service struct {
	repo           Repository
	kompetensiRepo pendidikan.KompetensiRepository
}

func NewService(repo Repository, kompetensiRepo pendidikan.KompetensiRepository) Service {
	return &service{
		repo:           repo,
		kompetensiRepo: kompetensiRepo,
	}
}

func (s *service) GetDashboardSummary(ctx context.Context, year int, username, name string) (*DashboardSummaryResponse, error) {
	if year == 0 {
		year = time.Now().Year()
	}

	totalTindakan, _ := s.repo.GetTotalTindakan(ctx, username, name)
	menungguValidasi, _ := s.repo.GetMenungguValidasi(ctx, username, name)
	perluRevisi, _ := s.repo.GetPerluRevisi(ctx, username, name)
	upcomingJadwals, _ := s.repo.GetUpcomingJadwals(ctx, username, name)
	recentEntries, _ := s.repo.GetRecentEntries(ctx, username, name)

	months := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	mandiriData := make([]int, 12)
	dibimbingData := make([]int, 12)
	observasiData := make([]int, 12)
	seminarData := make([]int, 12)

	var mandiriCum, dibimbingCum, observasiCum, seminarCum int

	for m := 1; m <= 12; m++ {
		mCount, _ := s.repo.CountTindakanByMonth(ctx, year, m, "mandiri", username, name)
		dCount, _ := s.repo.CountTindakanByMonth(ctx, year, m, "dibimbing", username, name)
		oCount, _ := s.repo.CountTindakanByMonth(ctx, year, m, "observasi", username, name)
		sCount, _ := s.repo.CountKegiatanByMonth(ctx, year, m, username, name)

		mandiriCum += mCount
		dibimbingCum += dCount
		observasiCum += oCount
		seminarCum += sCount

		mandiriData[m-1] = mandiriCum
		dibimbingData[m-1] = dibimbingCum
		observasiData[m-1] = observasiCum
		seminarData[m-1] = seminarCum
	}

	return &DashboardSummaryResponse{
		TotalTindakan:    totalTindakan,
		MenungguValidasi: menungguValidasi,
		PerluRevisi:      perluRevisi,
		UpcomingJadwals:  upcomingJadwals,
		RecentEntries:    recentEntries,
		ChartData: ChartDataResponse{
			Months:    months,
			Mandiri:   mandiriData,
			Dibimbing: dibimbingData,
			Observasi: observasiData,
			Seminar:   seminarData,
		},
	}, nil
}

func (s *service) GetLaporanSummary(ctx context.Context, periode, stase string, username, name string) (*LaporanSummaryResponse, error) {
	rekapTindakan, _ := s.repo.GetRekapTindakan(ctx, periode, stase, username, name)

	var kompetensiList interface{}
	if s.kompetensiRepo != nil {
		if username != "" || name != "" {
			list, err := s.kompetensiRepo.FindByUsername(ctx, username)
			if (err != nil || len(list) == 0) && name != "" {
				list, err = s.kompetensiRepo.FindByUsername(ctx, name)
			}
			if err == nil {
				kompetensiList = list
			} else {
				kompetensiList = []string{}
			}
		} else {
			list, err := s.kompetensiRepo.FindAll(ctx)
			if err == nil {
				kompetensiList = list
			} else {
				kompetensiList = []string{}
			}
		}
	} else {
		kompetensiList = []string{}
	}

	simposium, _ := s.repo.GetKegiatanCountByKategori(ctx, "simposium", username, name)
	workshop, _ := s.repo.GetKegiatanCountByKategori(ctx, "workshop", username, name)
	multidisiplin, _ := s.repo.GetKegiatanCountByKategori(ctx, "multidisiplin", username, name)
	ilmiahLain, _ := s.repo.GetKegiatanCountByKategori(ctx, "ilmiah_lain", username, name)

	totalBimbingan, _ := s.repo.GetTotalBimbingan(ctx, username, name)
	menungguCount, _ := s.repo.GetMenungguBimbinganCount(ctx, username, name)
	disetujuiCount, _ := s.repo.GetDisetujuiBimbinganCount(ctx, username, name)

	responRate := 100
	if totalBimbingan > 0 {
		responRate = int(math.Round(float64(disetujuiCount) / float64(totalBimbingan) * 100))
	}

	rawDpjpStats, _ := s.repo.GetDPJPStats(ctx, username, name)
	for i := range rawDpjpStats {
		n := rawDpjpStats[i].Name
		cleaned := strings.TrimSpace(n)
		for _, prefix := range []string{"dr.", "Sp.OT", "Sp.B", "Sp.An-KIC", ","} {
			cleaned = strings.ReplaceAll(cleaned, prefix, "")
		}
		words := strings.Fields(cleaned)
		initials := ""
		if len(words) > 0 {
			initials += strings.ToUpper(words[0][:1])
		}
		if len(words) > 1 {
			initials += strings.ToUpper(words[1][:1])
		}
		if initials == "" {
			initials = "DP"
		}
		rawDpjpStats[i].Initials = initials

		pct := 0
		if rawDpjpStats[i].TotalLogs > 0 {
			pct = int(math.Round(float64(rawDpjpStats[i].ApprovedLogs) / float64(rawDpjpStats[i].TotalLogs) * 100))
		}
		rawDpjpStats[i].Pct = pct
	}

	return &LaporanSummaryResponse{
		RekapTindakan:      rekapTindakan,
		KompetensiList:     kompetensiList,
		SimposiumCount:     simposium,
		WorkshopCount:      workshop,
		MultidisiplinCount: multidisiplin,
		IlmiahLainCount:    ilmiahLain,
		TotalBimbingan:     totalBimbingan,
		MenungguCount:      menungguCount,
		DisetujuiCount:     disetujuiCount,
		ResponRate:         responRate,
		DPJPStats:          rawDpjpStats,
	}, nil
}