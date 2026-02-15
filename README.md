# 🗄️ Build Your Own RDB in Go

**B+Treeから始めるRDBMS自作プロジェクト**

> 参考書籍: [Build Your Own Database From Scratch in Go](https://build-your-own.org/database/) by James Smith

---

## 📖 プロジェクト概要

このプロジェクトは、Go言語でリレーショナルデータベースをゼロから自作することを目標としています。
シンプルなKVストアから始め、最終的にSQL風のクエリ言語を持つミニRDBの完成を目指します。

---

## 🎯 学習目標

- データベースの内部構造を**ボトムアップ**で理解する
- **B+Tree** の実装を通じてインデックス構造を習得する
- **クラッシュリカバリ** や **トランザクション** など、実務で必要な概念を実装する
- **コピーオンライト** を用いた並行制御の仕組みを学ぶ
- SQL風クエリパーサーの実装を経験する

---

## 🗺️ ロードマップ

### Part I: シンプルなKVストア

KVストアをゼロから作り上げる最初のフェーズです。
（各章は単独で学習可能・無料公開）

| Chapter | テーマ | 概要 |
|---------|--------|------|
| 0 | イントロダクション | プロジェクト全体像の把握 |
| 1 | ファイルからDBへ | なぜファイル操作だけではDBにならないのか |
| 2 | インデックスデータ構造 | 検索を高速化するデータ構造の基礎 |
| 3 | B-Tree & クラッシュリカバリ | B-Treeの概念とデータ耐久性の入門 |
| 4 | B+Treeノードと挿入処理 | B+Treeの実装（挿入編） |
| 5 | B+Treeの削除とテスト | B+Treeの実装（削除・テスト編） |
| 6 | 追記型KVストア | ディスク上でのKVストア実装 |
| 7 | フリーリスト：ページの再利用 | 不要ページを管理し効率よく再利用する仕組み |

### Part II: ミニ・リレーショナルDB

KVストアを土台に、本格的なRDBを構築するフェーズです。

| Chapter | テーマ | 概要 |
|---------|--------|------|
| 8 | KV上のテーブル | KVストアの上にテーブル構造を実装する |
| 9 | 範囲クエリ | 効率的な範囲スキャンの実装 |
| 10 | セカンダリインデックス | 複数インデックスによる高速検索 |
| 11 | アトミックトランザクション | ACID特性の「A」を実現する |
| 12 | 並行制御 | 複数クライアントからの安全なアクセス制御 |
| 13 | SQLパーサー | SQL風クエリ言語の構文解析器 |
| 14 | クエリ言語 | パーサーを使ったクエリ実行エンジン |

---

## 🏗️ アーキテクチャ概要

```
┌─────────────────────────────────────────┐
│           Query Language (SQL風)         │  ← Chapter 13-14
├─────────────────────────────────────────┤
│         Concurrent Transactions          │  ← Chapter 11-12
├─────────────────────────────────────────┤
│          Relational DB (Tables)          │  ← Chapter 8-10
├─────────────────────────────────────────┤
│      Copy-on-Write KV Store (B+Tree)     │  ← Chapter 3-7
├─────────────────────────────────────────┤
│         File I/O & Durability            │  ← Chapter 1-2
└─────────────────────────────────────────┘
```

---

## 🛠️ 技術スタック

| 項目 | 詳細 |
|------|------|
| 言語 | Go |
| データ構造 | B+Tree（コピーオンライト） |
| ストレージ | ページベースのファイルI/O |
| 並行制御 | MVCC (Multi-Version Concurrency Control) |
| クエリ | SQL風独自クエリ言語 |

---

## 📁 ディレクトリ構成（予定）

```
.
├── README.md
├── go.mod
├── go.sum
│
├── kv/                  # Part I: KVストア
│   ├── btree/           # B+Tree実装
│   ├── freelist/        # フリーリスト管理
│   └── store/           # ディスクKVストア
│
├── db/                  # Part II: リレーショナルDB
│   ├── table/           # テーブル管理
│   ├── index/           # セカンダリインデックス
│   ├── txn/             # トランザクション制御
│   └── concurrent/      # 並行制御
│
├── sql/                 # クエリ言語
│   ├── parser/          # SQLパーサー
│   └── executor/        # クエリ実行エンジン
│
└── tests/               # 統合テスト
```

---

## 🚀 Getting Started

### 必要環境

- Go 1.21以上

### セットアップ

```bash
git clone https://github.com/<your-username>/build-your-own-rdb.git
cd build-your-own-rdb
go mod tidy
```

### テスト実行

```bash
go test ./...
```

---

## 📚 参考資料

| 書籍/リソース | 内容 |
|--------------|------|
| [Build Your Own Database From Scratch in Go](https://build-your-own.org/database/) | 本プロジェクトのメイン参考書 |
| [Database Internals / Alex Petrov](https://www.databass.dev/) | ストレージエンジンの詳細な解説 |
| [Designing Data-Intensive Applications / Martin Kleppmann](https://dataintensive.net/) | 分散システムを含む幅広いDB設計の知識 |

---

## 📝 進捗記録

- [ ] Chapter 0: イントロダクション
- [ ] Chapter 1: ファイルからDBへ
- [ ] Chapter 2: インデックスデータ構造
- [ ] Chapter 3: B-Tree & クラッシュリカバリ
- [ ] Chapter 4: B+Treeノードと挿入処理
- [ ] Chapter 5: B+Treeの削除とテスト
- [ ] Chapter 6: 追記型KVストア
- [ ] Chapter 7: フリーリスト：ページの再利用
- [ ] Chapter 8: KV上のテーブル
- [ ] Chapter 9: 範囲クエリ
- [ ] Chapter 10: セカンダリインデックス
- [ ] Chapter 11: アトミックトランザクション
- [ ] Chapter 12: 並行制御
- [ ] Chapter 13: SQLパーサー
- [ ] Chapter 14: クエリ言語

---

## 📄 ライセンス

MIT
