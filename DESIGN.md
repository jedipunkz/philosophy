# DESIGN.md

この文書は、哲学ノート Vault を中心にした自動収集・蓄積・AI参照システムの構成案をまとめる。

目的は、ユーザーが `scrape.yaml` に関心キーワードを記述するだけで、Go 製スクレイパーが関連情報を収集し、Obsidian Vault に Markdown として蓄積し、Codex や Claude Code がその Vault を読んで回答・比較・要約できる状態を作ること。

## 全体像

```text
scrape.yaml
  ↓
Go Scraper
  ↓
Markdown Normalizer
  ↓
Obsidian Vault
  ↓
Git Sync
  ↓
Codex / Claude Code
```

Docker Compose では、Go スクレイパー、スケジューラ、必要に応じた補助サービスをコンテナ単位で管理する。Obsidian Vault はホスト側のディレクトリを bind mount し、コンテナから Markdown ファイルを書き込めるようにする。

## レイヤー構成

### 1. Capture Layer

外部情報を取得する層。Go で実装する。

主な責務:

- `scrape.yaml` を読み込む
- キーワードごとに検索・スクレイピングを実行する
- 取得元 URL、タイトル、本文、取得日時、キーワードを保持する
- 同一 URL や同一タイトルの重複取得を避ける
- 取得失敗時にログを残し、Vault 書き込みを壊さない

初期実装では、対象ソースを広げすぎない。哲学論文・研究資料が集まるサイトを優先し、静的 HTML または公開 API/RSS で取得できる範囲から始める。JavaScript 実行が必要なサイト、ログインが必要なサイト、購読契約が必要な本文は対象外とする。

#### 情報源の優先順位

哲学 Vault の情報源は、以下の順で扱う。

1. `PhilArchive`
   - 哲学分野の Open Access 論文・草稿の主要アーカイブ
   - 本文にアクセスできる可能性が高く、この Vault の主情報源に向いている
   - 初期実装の第一候補

2. `PhilPapers`
   - 哲学文献の包括的な索引・ bibliography
   - 論文本文そのものではなく、メタデータや外部リンク中心の結果も多い
   - 収集対象としては、タイトル、著者、要旨、リンク、カテゴリを取り込む

3. `Stanford Encyclopedia of Philosophy`
   - 論文アーカイブではないが、哲学トピックの高品質な概説として有用
   - 既存ノートの背景理解や関連概念の補助資料として使う

4. `arXiv`
   - `axiv.org` ではなく `arxiv.org`
   - 哲学専用のアーカイブではない
   - 科学哲学、数理哲学、物理学の哲学、AI倫理、認知科学、論理学など、隣接領域の論文収集に向いている
   - 公式 API があるため、スクレイピングより API 連携を優先する

このため、初期方針は `PhilArchive` を主情報源、`PhilPapers` を索引、`arXiv` を隣接領域の補助ソースとして設計する。

### 2. Configuration Layer

人間が関心領域を指定する層。設定ファイルは `scrape.yaml` とする。

想定例:

```yaml
vault:
  root: /vault
  inbox: inbox
  notes: notes

scrape:
  interval: 24h
  user_agent: philosophy-scraper/0.1

sources:
  - name: philarchive
    enabled: true
    type: html
    base_url: https://philarchive.org

  - name: philpapers
    enabled: true
    type: html
    base_url: https://philpapers.org

  - name: arxiv
    enabled: false
    type: api
    base_url: https://export.arxiv.org/api/query

keywords:
  - name: ユング
    queries:
      - ユング 元型論
      - ユング 集合的無意識
    sources:
      - philarchive
      - philpapers
    tags:
      - 深層心理学
      - 現代思想

  - name: ハンナ・アーレント
    queries:
      - アーレント 全体主義
      - アーレント 人間の条件
    sources:
      - philarchive
      - philpapers
    tags:
      - 政治哲学
      - 現代思想

  - name: AI倫理
    queries:
      - artificial intelligence ethics philosophy
      - machine ethics philosophy
    sources:
      - philarchive
      - philpapers
      - arxiv
    tags:
      - 応用倫理学
      - 技術哲学
```

`keywords[].name` は保存先ファイル名や Front Matter の基準に使う。`queries` は実際の検索語、`sources` は検索対象、`tags` は Obsidian 上の検索性を上げるための補助メタデータとして扱う。

`arXiv` は既定では `enabled: false` とし、キーワード単位で必要な場合だけ有効にする。古典哲学や哲学史の収集には向かないが、論理学、科学哲学、AI倫理、計算論的認知科学のような領域では有効な補助ソースになる。

### 3. Pipeline Layer

取得した情報を Vault に保存可能な Markdown に正規化する層。

主な責務:

- HTML から本文候補を抽出する
- Markdown に変換する
- Obsidian 用 Front Matter を付与する
- Obsidian graph 用の `[[...]]` リンクとタグを本文へ付与する
- 保存先パスを決定する
- ファイル名を安全な文字列に正規化する
- 既存ファイルとの衝突時に上書きせず、スキップまたは別名保存する

生成される Markdown の例:

```markdown
---
source: https://example.com/article
title: ユングの元型論について
captured_at: 2026-05-09T06:00:00+09:00
capture_tool: scrapem
keyword: ユング
queries:
  - ユング 元型論
tags:
  - 深層心理学
  - 現代思想
status: raw
---

# ユングの元型論について

## Obsidian Links

- 研究動向: [[研究動向/ユング-現代研究動向|ユング-現代研究動向]]
- キーワード: [[ユング]]
- 関連分野: [[深層心理学]]
- 関連分野: [[現代思想]]
- 関連タグ: #深層心理学 #現代思想

取得本文...
```

初期保存先は `inbox/` とする。人間または AI が内容確認した後、著作別ノートは `書籍/西洋哲学/` や `書籍/東洋哲学/` 配下へ、Codex / Claude Code がまとめる研究動向は `研究動向/` 配下へ再編集する。

### 4. Storage Layer

Obsidian Vault を永続データ層として使う。

このリポジトリ自体を Vault と見なし、既存の哲学ノートと同じ Markdown ベースで管理する。

想定構成:

```text
philosophy/
├── CLAUDE.md
├── README.md
├── DESIGN.md
├── scrape.yaml
├── inbox/
│   └── 2026-05-09-ユングの元型論について.md
├── 書籍/
│   ├── 西洋哲学/
│   └── 東洋哲学/
├── 研究動向/
└── .scrapem/
    ├── seen-urls.json
    └── scraper.log
```

`inbox/` は自動収集された未処理メモの置き場とする。`.scrapem/` はスクレイパーの状態管理用で、URL の重複排除やログに使う。Git 管理対象にするかどうかは、実装時に運用方針を決める。

### 5. Intelligence Layer

Codex と Claude Code が Vault を読む層。

主な使い方:

- `inbox/` の未処理スクレイピング結果を `研究動向/` に要約する
- 既存ノートと関連するファイルを探す
- 新規ノート候補を既存フォーマットに整形する
- 複数の哲学者・著作を横断比較する
- `CLAUDE.md` と `README.md` に従って Vault の構成を保つ

AI はスクレイパーの代替ではなく、Vault に蓄積された Markdown を読む分析・編集レイヤーとして扱う。

### 6. Git Sync Layer

スクレイピング結果を GitHub に同期する層。スクレイパーコンテナの中ではなく、ホスト側スクリプトで実行する。

理由:

- Docker コンテナに SSH 鍵や GitHub 認証情報を渡さずに済む
- 既存のローカル Git 設定をそのまま使える
- `inbox/` と `.scrapem/seen-urls.json` だけを最小単位でコミットできる
- Claude Mobile App や別環境の Claude Code / Codex が GitHub 上の最新 `inbox/` を参照できる

想定フロー:

```text
scripts/scrape-and-commit.sh --push
  ↓
scripts/scrape.sh
  ↓
inbox/ と .scrapem/seen-urls.json に変更があれば git commit
  ↓
origin の現在ブランチへ push
  ↓
Claude Mobile / Claude Code が inbox を参照
  ↓
整理済みノート更新用の PR を作成
```

自動コミット対象は以下に限定する。

- `inbox/`
- `.scrapem/seen-urls.json`

スクレイパー本体、設定ファイル、正式ノートの変更は自動コミットしない。正式ノート化は Codex / Claude Code による別作業として扱い、PR でレビューできる形にする。

## Docker Compose 構成案

初期構成では、サービスを小さく分ける。

```yaml
services:
  scraper:
    build:
      context: .
      dockerfile: docker/scraper/Dockerfile
    volumes:
      - .:/vault
    command: ["scrapem", "run", "--config", "/vault/scrape.yaml"]

  scheduler:
    build:
      context: .
      dockerfile: docker/scraper/Dockerfile
    volumes:
      - .:/vault
    command: ["scrapem", "watch", "--config", "/vault/scrape.yaml"]
```

`scraper` は単発実行用、`scheduler` は定期実行用とする。開発初期は cron コンテナや外部ジョブ管理を増やさず、Go アプリ側に `watch` サブコマンドを持たせる構成が単純。

将来的に必要であれば、以下を追加する。

- `browser`: Playwright などを使った動的ページ取得用
- `rss`: RSS 収集専用ワーカー
- `api`: スクレイピング状態やキーワード一覧を確認する軽量 API

## Go アプリ構成案

```text
cmd/
└── scrapem/
    └── main.go
internal/
├── config/
│   └── config.go
├── scraper/
│   ├── client.go
│   ├── search.go
│   └── extract.go
├── markdown/
│   └── render.go
├── vault/
│   ├── writer.go
│   └── dedupe.go
└── scheduler/
    └── scheduler.go
```

責務:

- `config`: `scrape.yaml` の読み込みと検証
- `scraper`: HTTP 取得、検索、本文抽出
- `markdown`: Front Matter と本文の Markdown 生成
- `vault`: 保存先パス、重複排除、ファイル書き込み
- `scheduler`: 定期実行

## データフロー

```text
1. ユーザーが scrape.yaml にキーワードを書く
2. Docker Compose で scraper または scheduler を起動する
3. Go アプリが queries をもとに対象ページを探索する
4. ページ本文とメタデータを取得する
5. Markdown に変換する
6. Obsidian Vault の inbox/ に保存する
7. Codex / Claude Code が inbox/ と既存ノートを読んで要約・比較・編集する
8. 必要に応じて人間が著作別ノートとして 書籍/西洋哲学/ または 書籍/東洋哲学/ に整理し、AI は研究動向として 研究動向/ に整理する
```

## 運用ルール

- `scrape.yaml` は人間が編集する唯一の収集設定ファイルとする
- 自動収集結果はまず `inbox/` に入れる
- 既存の正式ノートをスクレイパーが直接上書きしない
- スクレイピング結果の自動同期は `inbox/` と `.scrapem/seen-urls.json` のみに限定する
- 運用段階ではスクレイピング間隔を週1回程度にし、テスト中のみ高頻度で回す
- 取得元 URL と取得日時は必ず Front Matter に残す
- robots.txt、利用規約、アクセス頻度に配慮する
- ログインが必要なサイト、個人情報、非公開情報は対象外とする
- AI が正式ノート化する場合も、既存 `README.md` と `CLAUDE.md` の形式に合わせる

## 初期実装スコープ

最初の実装では以下に絞る。

- `scrape.yaml` の読み込み
- キーワードと query の定義
- HTTP GET による静的ページ取得
- タイトルと本文の簡易抽出
- Obsidian Front Matter 付き Markdown 生成
- `inbox/` へのファイル保存
- URL 重複排除
- Docker Compose による単発実行

以下は後続対応とする。

- 動的サイトのブラウザレンダリング
- RSS 連携
- 検索 API 連携
- ベクトル検索
- Web UI
- 自動分類による正式ノート化
- Codex / Claude Code の定期実行連携

## 期待する完成形

この構成により、Vault は手動で書いた哲学ノートだけでなく、関心キーワードに沿って自動収集された素材も蓄積する知識基盤になる。

Go スクレイパーは情報を集める。Obsidian Vault は永続保存する。Codex と Claude Code は保存された Markdown を読み、既存ノートとの接続、要約、比較、正式ノート化を支援する。
