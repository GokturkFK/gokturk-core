// Package correlate, trip'lerden alarm üreten korelasyon kurallarını içerir.
//
// Alert şekli donmuş wire contract'tir. Bu katmanda bilinçli olarak
// ML/anomali yoktur: kurallar deterministiktir, sıfır-FP tezi ancak böyle
// savunulur — hangi ürün/sensör kullanırsa kullansın aynı ilke geçerlidir.
package correlate

import (
	"sort"
	"time"

	"github.com/GokturkFK/gokturk-core/trap"
)

// Severity ve status değerleri tüm ürünlerde aynı tutulur.
const (
	SeverityHigh     = "High"
	SeverityCritical = "Critical"

	StatusOpen   = "open"
	StatusAck    = "ack"
	StatusClosed = "closed"
)

// SubjectAlerts, opsiyonel alarm yayın subject'idir.
const SubjectAlerts = "alerts.v1"

// Alert, SOC paneline ve SIEM'e akan alarm kaydidir.
type Alert struct {
	ID        string    `json:"id,omitempty"`
	Severity  string    `json:"severity"`
	Technique string    `json:"technique,omitempty"`
	Source    string    `json:"source"`
	Status    string    `json:"status"`
	FirstSeen time.Time `json:"first_seen"`
	LastSeen  time.Time `json:"last_seen"`
	TripCount int       `json:"trip_count"`
}

// Evaluate, verilen zaman penceresindeki trip'leri kaynağa (Source) göre
// gruplayıp kaynak başına tek alarm üretir:
//
//   - 1 trip   -> High
//   - >=2 trip -> tek Critical (kampanya birleşmesi)
//
// technique, üretilen alarmların ATT&CK teknik alanına yazılır (çağıran
// hangi teknikle eşleştiğini bilir — örn. GÖKTÜRK için T1078, GÖKKALKAN
// için farklı bir MCP/agent tekniği). FirstSeen/LastSeen gruptaki en
// erken/en geç ObservedAt'tir. Çıktı kaynak adına göre sıralıdır
// (deterministik). Zaman penceresini seçmek ve alarmları persist etmek
// çağıranın sorumluluğudur.
func Evaluate(trips []trap.TripEvent, technique string) []Alert {
	if len(trips) == 0 {
		return nil
	}

	bySource := make(map[string][]trap.TripEvent)
	for _, tr := range trips {
		bySource[tr.Source] = append(bySource[tr.Source], tr)
	}

	sources := make([]string, 0, len(bySource))
	for src := range bySource {
		sources = append(sources, src)
	}
	sort.Strings(sources)

	alerts := make([]Alert, 0, len(sources))
	for _, src := range sources {
		group := bySource[src]

		first, last := group[0].ObservedAt, group[0].ObservedAt
		for _, tr := range group[1:] {
			if tr.ObservedAt.Before(first) {
				first = tr.ObservedAt
			}
			if tr.ObservedAt.After(last) {
				last = tr.ObservedAt
			}
		}

		severity := SeverityHigh
		if len(group) >= 2 {
			severity = SeverityCritical
		}

		alerts = append(alerts, Alert{
			Severity:  severity,
			Technique: technique,
			Source:    src,
			Status:    StatusOpen,
			FirstSeen: first,
			LastSeen:  last,
			TripCount: len(group),
		})
	}

	return alerts
}
