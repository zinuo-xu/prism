# 🌈 prism

> Pretty structural diff for JSON/YAML/TOML — diff for humans, not machines

[![Go Version](https://img.shields.io/github/go-mod/go-version/zinuo-xu/prism)](https://github.com/zinuo-xu/prism)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![CI](https://github.com/zinuo-xu/prism/actions/workflows/ci.yml/badge.svg)](https://github.com/zinuo-xu/prism/actions/workflows/ci.yml)

Shows what actually changed — keys added, values modified, arrays reordered — not raw text diffs.

## ✨ Features

- 🔍 **Structural diff** — Understands JSON/YAML/TOML structure
- 🎨 **Color-coded output** — Green=added, Red=removed, Yellow=changed
- 📋 **Multiple formats** — Terminal, JSON, TUI, Markdown
- ⚡ **Fast** — Written in Go

## 🚀 Quickstart

```bash
go install github.com/zinuo-xu/prism/cmd/prism@latest
prism old.json new.json
prism --format json a.yaml b.yaml
curl api.com/data | prism - expected.json
```

## 📄 License

MIT (c) [zinuo-xu](https://github.com/zinuo-xu)
