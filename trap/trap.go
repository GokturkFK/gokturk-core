// Package trap, Göktürk platformundaki tüm ürünlerin paylaştığı donmuş
// deception sözleşmesini tanımlar (bkz. GOKTURKFK_ROADMAP.md böl. 0-1).
//
// TripEvent JSON şeması donmuş wire contract'tir. Şema değişikliği ancak
// subject version bump (örn. trip.events.v2) ile yapılır — bu paketi
// import eden hiçbir ürün (GÖKTÜRK, GÖKKALKAN, GÖKZİNCİR, ...) alan
// adlarını/tiplerini tek başına değiştiremez.
package trap

import (
	"context"
	"encoding/json"
	"errors"
	"time"
)

// SubjectTripEvents, sensörlerin/kaynakların TripEvent yayınladığı NATS subject'i.
const SubjectTripEvents = "trip.events.v1"

// ErrNotATrip, gözlemin bir tuzak tetiklemesi olmadığını belirtir.
// Decode bu hatayı döndürdüğünde hiçbir event yayınlanmaz — sıfır
// false-positive tezinin kod düzeyindeki karşılığı.
var ErrNotATrip = errors.New("trap: gozlem bir tuzak tetiklemesi degil")

// TripEvent, bir tuzağın tetiklendiğini bildiren wire event'idir.
// trip.events.v1 subject'ine JSON olarak yayınlanır.
type TripEvent struct {
	EventID    string          `json:"event_id"`
	TrapID     string          `json:"trap_id"`
	Sensor     string          `json:"sensor"`
	Source     string          `json:"source"`
	ObservedAt time.Time       `json:"observed_at"`
	Raw        json.RawMessage `json:"raw,omitempty"`
}

// Trap, provision edilmiş bir tuzağın kaydidir. SecretHash API cevaplarına
// asla çıkmaz; düz metin secret yalnızca provision anında Artifacts içinde
// bir kez döner.
type Trap struct {
	ID         string          `json:"id"`
	Type       string          `json:"type"`
	Username   string          `json:"username"`
	SecretHash string          `json:"-"`
	Metadata   json.RawMessage `json:"metadata,omitempty"`
	CreatedBy  string          `json:"created_by,omitempty"`
	CreatedAt  time.Time       `json:"created_at"`
	RevokedAt  *time.Time      `json:"revoked_at,omitempty"`
}

// Artifacts, hedefe ekilecek malzemedir; provision cevabında bir kez döner.
// Secret kalıcı depoya yazılmaz, yalnızca hash'i saklanır.
type Artifacts struct {
	Username string `json:"username"`
	Secret   string `json:"secret"`
}

// RawObservation, sensörün decode öncesi ham gözlemidir (örneğin bir
// auth.log satırı veya bir MCP tool çağrısı). Wire'a çıkmaz; sensör
// içinde Decode'a girdi olur.
type RawObservation struct {
	Sensor     string
	Source     string
	Line       string
	ObservedAt time.Time
}

// Provider, bir tuzak türünün yaşam döngüsünü yönetir. Her ürün kendi tuzak
// türü için bu arayüzü uygular (GÖKTÜRK: credential canary, GÖKKALKAN: agent
// honeypot, ...); dönen Trap/Artifacts şekli her zaman aynı kalır.
type Provider interface {
	Provision(ctx context.Context, createdBy string) (*Trap, *Artifacts, error)
}

// Decoder, ham gözlemi TripEvent'e çevirir. Gözlem bir tuzak tetiklemesi
// değilse ErrNotATrip döner ve hiçbir şey yayınlanmaz (sıfır-FP).
type Decoder interface {
	Decode(obs RawObservation) (*TripEvent, error)
}
