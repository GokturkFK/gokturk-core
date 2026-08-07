# gokturk-core

Göktürk platformunun paylaşılan Go çekirdeği: `trap` sözleşmesi (`Provider`,
`TripEvent`, `Trap`, `Artifacts`, `Decoder`, `ErrNotATrip`) ve `correlate`
korelasyon motoru (sıfır-FP, kaynak başına High→Critical kampanya birleşmesi).

Bu depo tek başına bir ürün değildir. Onu import eden ürünler:

| Ürün | Depo | Durum |
|---|---|---|
| GÖKTÜRK — deception mesh | [gokturk-deception-mesh](https://github.com/GokturkFK/gokturk-deception-mesh) | v0.1 tamamlandı |
| GÖKKALKAN — agentic AI runtime security | [gokkalkan](https://github.com/GokturkFK/gokkalkan) | aktif geliştirme |

Tasarım ilkesi: omurga bir kez burada kurulur, her ürün onun **sınırlı bir
uzantısı** olur — tuzak türü ve sensör ürüne özeldir, sözleşme ve korelasyon
ortaktır.

## Sahiplik

`trap` ve `correlate` **@fetihcakmak**'ındır (güvenlik çekirdeği) — GÖKTÜRK'te
de `internal/trap` ve `internal/correlate` onun sahipliğindeydi, bu depoya
çıkarılırken bölüşüm değişmedi. CI/tooling **@uzunkubra50**'de. Bkz.
[CODEOWNERS](CODEOWNERS).

**Burada aktif bir görev yok** — bu depo v0.1.0'da donmuş durumda. Yapılacak
işler aktif üründe: [gokkalkan/issues](https://github.com/GokturkFK/gokkalkan/issues).

## Sözleşme kuralı

`trap` paketindeki tipler (özellikle `TripEvent`'in JSON alan adları) donmuş
wire contract'tir. Değişiklik ancak yeni bir subject versiyonuyla
(`trip.events.v2` gibi) yapılır — aksi halde bu paketi import eden tüm
ürünler aynı anda kırılır.

## Paketler

- `trap` — tuzak yaşam döngüsü sözleşmesi. Somut tuzak türleri (credential
  canary, agent honeypot, ...) bu pakette DEĞİL, her ürünün kendi reposunda
  yaşar; sadece `Provider`/`Decoder` arayüzlerini uygular.
- `correlate` — `Evaluate(trips, technique)`: kaynak başına tek alarm,
  deterministik, ML/anomali yok. `technique` çağıran tarafından verilir
  (her ürün kendi ATT&CK eşlemesini seçer).

## Kullanım

```sh
go get github.com/GokturkFK/gokturk-core
```

```go
alerts := correlate.Evaluate(trips, "T1078")
```

## Geliştirme

```sh
go build ./...
go vet ./...
go test -race ./...
```
