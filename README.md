# 📊 Go JSON Benchmarking

This repository benchmarks various Go JSON libraries to evaluate their `Marshal` and `Unmarshal` performance

---

## 📌 Benchmarked Libraries

| Library         | Description                                              |
|-----------------|----------------------------------------------------------|
| `encoding/json` | Go standard library (baseline)                           |
| `jsoniter`      | High-performance drop-in replacement for `encoding/json` |
| `sonic`         | Fastest JSON library from Bytedance (TikTok)             |
| `easyjson`      | Code generation-based fast JSON (zero reflection)        |
| `ffjson`        | Older codegen-based lib (now deprecated but tested)      |

---

## 🚀 Benchmark Setup

- **Platform:** macOS (Darwin), Apple M2
- **Architecture:** ARM64
- **Go Version:** 1.21+
- **Benchmark Command:**

```bash
go test -bench=. -benchtime=1000000x 
```

## 📈 Benchmark Results

#### <div style="color:red">When using simple Struct<div>

```
type UserProfile struct {
    ID   string `json:"id"`   // Unique identifier for the user
    Name string `json:"name"` // Name of the user
}
```

#### 📌 Note: ffjson is no longer actively maintained, and easyjson requires code generation (go generate). Use them with caution in production.

Lower ns/op = Faster Performance

These benchmarks measure the time it takes to marshal (encode) and unmarshal (decode) JSON using different Go libraries
on an Apple M2 (goarch: arm64).

### 🔁 Combined Marshal + Unmarshal (1000000 iteration)

| 🏆 Rank | Library       | Benchmark Name         | Time (ns/op) |
|---------|---------------|------------------------|--------------|
| 🥇      | FFJSON        | `BenchmarkFFJSON-8`    | **308.1**    |
| 🥈      | EasyJSON      | `BenchmarkEasyjson-8`  | 320.4        |
| 🥉      | Jsoniter      | `BenchmarkJsoniter-8`  | 629.2        |
|         | encoding/json | `BenchmarkJSON-8`      | 979.9        |
|         | Sonic         | `BenchmarkSonicJSON-8` | 995.5        |

### 🔄 Marshal Only (1000000 iteration)

| 🏆 Rank | Library       | Benchmark Name                | Time (ns/op) |
|---------|---------------|-------------------------------|--------------|
| 🥇      | FFJSON        | `BenchmarkMarshalFFJSON-8`    | **118.4**    |
| 🥈      | EasyJSON      | `BenchmarkMarshalEasyjson-8`  | 152.9        |
| 🥉      | Jsoniter      | `BenchmarkMarshalJsoniter-8`  | 200.9        |
|         | Sonic         | `BenchmarkMarshalSonicJSON-8` | 404.0        |
|         | encoding/json | `BenchmarkMarshalJSON-8`      | 425.8        |

### 🔁 Unmarshal Only (1000000 iteration)

| 🏆 Rank | Library       | Benchmark Name                  | Time (ns/op) |
|---------|---------------|---------------------------------|--------------|
| 🥇      | EasyJSON      | `BenchmarkUnmarshalEasyjson-8`  | **159.8**    |
| 🥈      | FFJSON        | `BenchmarkUnmarshalFFJSON-8`    | 166.9        |
| 🥉      | Jsoniter      | `BenchmarkUnmarshalJsoniter-8`  | 380.9        |
|         | encoding/json | `BenchmarkUnmarshalJSON-8`      | 502.7        |
|         | Sonic         | `BenchmarkUnmarshalSonicJSON-8` | 549.0        |

## 📈 Benchmark Results

#### <div style="color:red"> When using difficult Struct <div>

```
type UserProfile struct {
	ID            string            `json:"id"`
	Name          string            `json:"name"`
	Age           int               `json:"age"`
	Email         *string           `json:"email,omitempty"`
	CreatedAt     time.Time         `json:"created_at"`
	LastLogin     *time.Time        `json:"last_login,omitempty"`
	Preferences   map[string]string `json:"preferences"`
	Friends       []FriendProfile   `json:"friends"`
	Address       Address           `json:"address"`
	PaymentMethod *PaymentMethod    `json:"payment_method,omitempty"`
	Metadata      map[string]any    `json:"metadata"`
	Tags          []string          `json:"tags"`
}

type FriendProfile struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Connected time.Time `json:"connected"`
}

type Address struct {
	Country   string `json:"country"`
	City      string `json:"city"`
	Street    string `json:"street"`
	ZipCode   string `json:"zip_code"`
	IsPrimary bool   `json:"is_primary"`
}

type PaymentMethod struct {
	Type       string `json:"type"`
	CardNumber string `json:"card_number"`
	Expiry     string `json:"expiry"`
}

```

#### 📌 Note: ffjson is no longer actively maintained, and easyjson requires code generation (go generate). Use them with caution in production.

Lower ns/op = Faster Performance

These benchmarks measure the time it takes to marshal (encode) and unmarshal (decode) JSON using different Go libraries
on an Apple M2 (goarch: arm64).

### 🔁 Combined Marshal + Unmarshal (1000000 iteration)

| 🏆 Rank | Library       | Benchmark Name         | Time (ns/op) |
|---------|---------------|------------------------|--------------|
| 🥇      | EasyJSON      | `BenchmarkEasyjson-8`  | **5139**     |
| 🥈      | FFJSON        | `BenchmarkFFJSON-8`    | 5191         |
| 🥉      | Jsoniter      | `BenchmarkJsoniter-8`  | 7417         |
|         | Sonic         | `BenchmarkSonicJSON-8` | 8191         |
|         | encoding/json | `BenchmarkJSON-8`      | 15499        |

### 🔄 Marshal Only (1000000 iteration)

| 🏆 Rank | Library       | Benchmark Name                | Time (ns/op) |
|---------|---------------|-------------------------------|--------------|
| 🥇      | FFJSON        | `BenchmarkMarshalFFJSON-8`    | **1806**     |
| 🥈      | EasyJSON      | `BenchmarkMarshalEasyjson-8`  | 1855         |
| 🥉      | Jsoniter      | `BenchmarkMarshalJsoniter-8`  | 2046         |
|         | Sonic         | `BenchmarkMarshalSonicJSON-8` | 3209         |
|         | encoding/json | `BenchmarkMarshalJSON-8`      | 6553         |

### 🔁 Unmarshal Only (1000000 iteration)

| 🏆 Rank | Library       | Benchmark Name                  | Time (ns/op) |
|---------|---------------|---------------------------------|--------------|
| 🥈      | FFJSON        | `BenchmarkUnmarshalFFJSON-8`    | **2841**     |
| 🥇      | EasyJSON      | `BenchmarkUnmarshalEasyjson-8`  | 2846         |
| 🥉      | Sonic         | `BenchmarkUnmarshalSonicJSON-8` | 4479         |
|         | Jsoniter      | `BenchmarkUnmarshalJsoniter-8`  | 5034         |
|         | encoding/json | `BenchmarkUnmarshalJSON-8`      | 8708         |

## 📌 JSON Serialization Library Performance Comparison

#### In the world of Go, JSON serialization is a common task, and the performance of various libraries can significantly impact your application’s efficiency. Below is a detailed comparison of popular JSON libraries based on their marshal and unmarshal performance when handling complex data structures.

#### Key Takeaways

<i style="color:green">FFJSON</i>:
🏆 Top Performer
FFJSON stands out with the best performance for both marshalling and unmarshalling operations. It efficiently handles
complex structures, making it an excellent choice for performance-critical applications.

<i style="color:green">EasyJSON</i>:
💪 Highly Competitive
EasyJSON offers solid performance in both marshaling and unmarshaling. While it doesn't always outpace FFJSON, it's
still a top contender with a great balance between speed and ease of use.

<i style="color:green">Jsoniter</i>:
⚡ Good, but Not Best
Jsoniter delivers good performance, but in most scenarios, it doesn't outperform FFJSON or EasyJSON. It’s a reliable
option for general use but falls behind the leading libraries in specific performance benchmarks.

<i style="color:green">encoding/json</i>:
🐢 Slower
Go’s standard library (encoding/json) is slower than all the optimized libraries. While it's convenient and included in
the Go standard library, it lags in performance when compared to specialized JSON libraries.

<i style="color:green">Sonic</i>:
🐌 Slightly Slower, But Useful in Some Cases
Sonic performs slower than other libraries but might still be useful, particularly for small, simpler structures where
speed isn't the highest priority. It can be a good option when simplicity and ease of integration are more important
than raw performance.