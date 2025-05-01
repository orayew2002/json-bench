# 📊 Go JSON Benchmarking

This repository benchmarks various Go JSON libraries to evaluate their `Marshal` and `Unmarshal` performance on Apple
M2 (ARM64) architecture.

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

Lower ns/op = Faster Performance

These benchmarks measure the time it takes to marshal (encode) and unmarshal (decode) JSON using different Go libraries

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