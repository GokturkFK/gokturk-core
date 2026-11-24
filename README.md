<div align="center">

# gokturk-core

**Göktürk Güvenlik Platformu için Paylaşılan Go Çekirdeği**

[![Go](https://img.shields.io/badge/Go-1.25-00ADD8?style=flat-square&logo=go&logoColor=white)](https://go.dev/)
[![Version](https://img.shields.io/badge/S%C3%B6zle%C5%9Fme-v0.1.0%20Dondurulmu%C5%9F-22c55e?style=flat-square)](https://github.com/GokturkFK/gokturk-core)
[![Zero-FP](https://img.shields.io/badge/Korelasyon-S%C4%B1f%C4%B1r--FP%20Belirlenimci-7c3aed?style=flat-square)](https://github.com/GokturkFK/gokturk-core)
[![Org](https://img.shields.io/badge/Org-GokturkFK-0d1117?style=flat-square)](https://github.com/GokturkFK)

*Göktürk platformundaki tüm ürünleri besleyen dondurulmuş tuzak sözleşmesi (`trap`) ve belirlenimci, sıfır yanlış pozitifli kampanya korelasyon motoru (`correlate`).*

**🇹🇷 Türkçe** · [🇬🇧 English](#-english)

</div>

---

## Genel Bakış

**`gokturk-core`**, Göktürk siber güvenlik ekosistemindeki tüm Go ürünleri tarafından import edilen ortak çekirdek kütüphanedir. Tek başına bir ürün değildir; her ürünün üzerine kurulduğu minimal ve sürüm kilitli omurgayı tanımlar:

- **`trap`:** Tuzak yaşam döngüsü sözleşmesi (`Provider`, `TripEvent`, `Trap`, `Artifacts`, `Decoder`, `ErrNotATrip`). Somut tuzak türleri (kimlik kanaryası, ajan honeypot, sahte NHI) bu pakette değil, her ürünün kendi deposunda yaşar ve bu arayüzleri uygular.
- **`correlate`:** Belirlenimci korelasyon motoru (`Evaluate(trips, technique)`). ML/anomali gürültüsü olmadan kaynak başına tek alarm üretir: tek doğrulanmış tetiklenme **High**, aynı kaynaktan ≥2 tetiklenme tek bir **Critical** kampanya alarmında birleşir. MITRE ATT&CK / ATLAS `technique` kimliği çağıran ürün tarafından verilir.

---

## Platform Ekosistemi

| Katman | Depo | Alan | Durum |
| :---: | :--- | :--- | :--- |
| **P1** | [**`gokturk-deception-mesh`**](https://github.com/GokturkFK/gokturk-deception-mesh) | Kimlik kanaryası aldatma & SSH sensörü | ✅ v0.1 tamamlandı |
| **P2** | [**`gokkalkan`**](https://github.com/GokturkFK/gokkalkan) | Ajan tabanlı yapay zekâ güvenliği (MCP egress proxy + honeypot) | ✅ Uçtan uca tamamlandı |
| **P3** | [**`gokzincir`**](https://github.com/GokturkFK/gokzincir) | İnsan-olmayan kimlik (NHI) yönetişimi & blast-radius grafiği | ✅ Uçtan uca tamamlandı |

---

## Sözleşme Kararlılığı

`trap` paketindeki tipler — özellikle `TripEvent`'in JSON alan adları — **dondurulmuş wire contract** (`v0.1.0`) niteliğindedir. Kırıcı şema değişiklikleri yerinde yapılmaz; yalnızca yeni sürümlenmiş bir subject (örn. `trip.events.v2`) ile ilerlenir.

---

## Kullanım

```sh
go get github.com/GokturkFK/gokturk-core
```

```go
package main

import (
    "github.com/GokturkFK/gokturk-core/correlate"
    "github.com/GokturkFK/gokturk-core/trap"
)

func evaluate(trips []trap.TripEvent) []correlate.Alert {
    return correlate.Evaluate(trips, "T1078")
}
```

---

## Geliştirme ve Doğrulama

```sh
go build ./...
go vet ./...
go test -race ./...
```

---

## Sahiplik

- `trap` ve `correlate` güvenlik çekirdeği: [@fetihcakmak](https://github.com/fetihcakmak)
- CI ve depo araçları: [@uzunkubra50](https://github.com/uzunkubra50)
- Ayrıntılar: [`CODEOWNERS`](CODEOWNERS)

---

<a id="-english"></a>

<details>
<summary><b>🇬🇧 English</b> — click to expand</summary>

<br/>

## Overview

**`gokturk-core`** is the shared Go kernel imported across the Göktürk security ecosystem. It defines the version-locked backbone that every platform pillar extends:

- **`trap`:** Lifecycle wire contract (`Provider`, `TripEvent`, `Trap`, `Artifacts`, `Decoder`, `ErrNotATrip`). Concrete trap implementations live in each product repository.
- **`correlate`:** Deterministic correlation engine (`Evaluate(trips, technique)`). Zero ML/anomaly noise: a single verified trip raises a **High** alert, while ≥2 trips from the same source merge into a single **Critical** campaign alert.

## Ecosystem

- **P1:** [`gokturk-deception-mesh`](https://github.com/GokturkFK/gokturk-deception-mesh) — Credential canary deception & SSH detection
- **P2:** [`gokkalkan`](https://github.com/GokturkFK/gokkalkan) — Agentic AI runtime security (MCP egress proxy + honeypot)
- **P3:** [`gokzincir`](https://github.com/GokturkFK/gokzincir) — Non-Human Identity (NHI) governance & blast-radius graph

## Wire Contract Stability

Types in `trap` (especially `TripEvent` JSON fields) are frozen at `v0.1.0`. Breaking changes require a new versioned subject (e.g., `trip.events.v2`).

</details>

---

<div align="center">
  <sub><a href="https://github.com/GokturkFK">GokturkFK</a> Güvenlik Platformu parçasıdır</sub>
</div>
