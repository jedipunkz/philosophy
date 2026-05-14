# 哲学ノート

哲学者・著作ごとに内容をまとめた個人ナレッジベース。Obsidian Vault として管理し、Go 製スクレイパーで収集した論文情報を `research-inbox/` に、書籍情報を `book-inbox/` に蓄積し、Codex / Claude Code で `研究動向/` および `書籍/` の日本語ノートへ統合する。

---

## フォルダ構成

```
philosophy/
├── research-inbox/   # 論文スクレイパーが保存する未処理素材
├── book-inbox/       # 書籍スクレイパーが保存する未処理素材
├── 書籍/
│   ├── 西洋哲学/
│   │   ├── 古代/      # ～紀元後5世紀ごろまで（ソクラテス、プラトン等）
│   │   ├── 近代/      # 17〜19世紀（カント、ヘーゲル、マルクス、ニーチェ等）
│   │   └── 現代/      # 20世紀以降（ユング、アーレント、サルトル等）
│   └── 東洋哲学/
│       ├── インド・仏教/  # ※ファイルが増えた時点でサブフォルダを切る
│       ├── 中国/          # ※同上
│       └── 日本/          # ※同上
├── 研究動向/         # research-inbox をもとに Codex / Claude Code がまとめる研究動向ノート
├── scrape.yaml       # 論文スクレイピング設定（PhilArchive、PhilPapers、arXiv）
├── book-scrape.yaml  # 書籍スクレイピング設定（Open Library、Project Gutenberg）
├── cmd/              # Go 製 scrapem CLI
├── internal/         # scrapem の内部実装
└── docker-compose.yml
```

時代区分は研究者によって異なる場合がある。迷った場合はその哲学者が最も活躍した時期を基準にし、ファイル冒頭に時代的背景を記す。

---

## ファイル命名規則

```
著者名-著作名.md
```

**例:**
- `ユング-元型論.md`
- `カント-純粋理性批判.md`
- `中村元-ブッダのことばスッタニパータ.md`

**ルール:**
- 区切りはハイフン `-` を使う（スペース・アンダースコア不可）
- 著者名は通称・カナ表記を優先する（例: ハンナ・アーレント）
- 著作名が長い場合でも省略しない（検索性のため）
- 一人の著者を複数ファイルで扱う場合はすべて同じ著者名表記に統一する

---

## ファイルフォーマット

各ファイルの冒頭に以下のフロントマター（YAML）を記載する。Codex / Claude Code が内容を横断的に検索・比較する際の手がかりになる。

```yaml
---
著者: ユング
著作: 元型論
時代: 現代
地域: 西洋
分野: [深層心理学, 精神分析, 神話学]
キーワード: [集合的無意識, 元型, 個性化, アニマ, 影]
関連ファイル:
  - ユング-心理学的類型.md
  - ユング-赤の書.md
---
```

`分野` と `キーワード` は複数指定可。`関連ファイル` に同一著者または思想的に近い著作を列挙すると、AI エージェントが横断的なリサーチをする際に使いやすくなる。

---

## 自動収集アーキテクチャ

### 論文収集パイプライン

```text
scrape.yaml
  ↓
Go Scraper（PhilArchive / PhilPapers / arXiv）
  ↓
research-inbox/*.md
  ↓
Codex / Claude Code
  ↓
研究動向/ の日本語まとめノート
```

### 書籍収集パイプライン

```text
book-scrape.yaml
  ↓
Go Scraper（Open Library API / Project Gutenberg Gutendex API）
  ↓
book-inbox/*.md
  ↓
Codex / Claude Code
  ↓
書籍/ 配下の整理済み日本語ノート
```

各層の役割:

- `scrape.yaml`: 論文収集の対象キーワード・クエリ・情報源を指定する
- `book-scrape.yaml`: 書籍収集の対象著者・著作・情報源を指定する。`keywords:` セクションで対象を管理する
- Go Scraper: 各 API からメタデータ・Abstract・Subjects・公開 URL などを収集する
- `research-inbox/`: 論文の未処理素材キュー。`capture_tool: scrapem`
- `book-inbox/`: 書籍の未処理素材キュー。`capture_tool: scrapem-book`。`public_domain: true` のノートは Gutenberg 由来で本文取得可能
- `書籍/`: 整理済み日本語ノート
- `研究動向/`: 論文素材を整理した研究動向ノート
- Codex / Claude Code: 各 inbox と既存ノートを読み、重複を避けて統合する

### スクレイピングの実行

**論文スクレイパー:**

```bash
# 単発実行
scripts/scrape.sh

# 定期実行
docker compose up -d scheduler
```

**書籍スクレイパー:**

```bash
# 単発実行
scripts/scrape.sh --config book-scrape.yaml

# 定期実行（週1回、168h間隔）
docker compose up -d book-scheduler
```

スクレイプ後に両 inbox をコミットする場合:

```bash
scripts/scrape-and-commit.sh
```

push まで行う場合:

```bash
scripts/scrape-and-commit.sh --push
```

自動コミット対象は `research-inbox/`・`book-inbox/`・`.scrapem/seen-*.json` のみ。正式ノートの更新は、AI エージェントによる別作業・別 PR として扱う。

### inbox の扱い

**research-inbox:**
`research-inbox/` は正式ノートではなく、論文収集素材の置き場である。Codex / Claude Code が参照してまとめる場合は `研究動向/` に整理する。

**book-inbox:**
`book-inbox/` は書籍収集素材の置き場である。Codex / Claude Code が参照して書籍ノートを生成する場合は `書籍/` 配下に整理する。`public_domain: true` のノートは Gutenberg から本文取得が可能。著作権が切れていないものは書誌情報・概要のみ利用する。

AIエージェントが正式ノートへ反映したら、該当するノートを次のように更新する。

```yaml
# 論文 (research-inbox → 研究動向/)
status: processed
processed_to: "研究動向/ハンナ・アーレント-現代研究動向.md"

# 書籍 (book-inbox → 書籍/)
status: processed
processed_to: "書籍/西洋哲学/近代/カント-純粋理性批判.md"
```

不要または関係が薄い素材は `status: ignored` にする。

---

## Codex / Claude Code での管理指針

### セッション開始時に伝えること

新しいセッションで作業を開始するとき、最初のメッセージでこのフォルダを指定すると AI エージェントが文脈を把握しやすい。

> 「philosophy フォルダの哲学ノートを管理しています。README.md を参照して構成を確認してください」

### ファイル追加の依頼例

> 「サルトル『存在と無』について既存ファイルと同じ分量でリサーチして追加して。フロントマターも付けて」

### inbox 整理の依頼例

> 「research-inbox/ のハンナ・アーレント関連ノートを読んで、既存の書籍/西洋哲学/現代/ハンナ・アーレント-*.md と照合し、研究動向/ハンナ・アーレント-現代研究動向.md に重複しない形で統合して」

> 「book-inbox/ のカント関連ノートを読んで、書籍/西洋哲学/近代/ に新規ノートとして追加して」

> 「research-inbox/ または book-inbox/ のうち status: raw のものを確認し、正式ノート化すべきもの、無視してよいものに分類して」

### 横断的な質問例

> 「ユングとフロイトの無意識の捉え方の違いをまとめて」
> 「東洋哲学と西洋現代哲学で『自己』の概念はどう違う？」
> 「ニーチェとアーレントの政治観を比較して」

### 内容レビューの依頼例

> 「追加したファイルをWeb検索でリサーチしながらファクトチェックして」

---

## Git 運用ルール

| 操作 | コマンド |
|------|----------|
| 変更を確認 | `git status` / `git diff` |
| ステージング | `git add <ファイル>` |
| コミット | `git commit -m "Add: 著者-著作名"` |
| プッシュ | `git push origin main` |

**コミットメッセージの規則:**

```
Add: ユング-共時性     # 新規ファイル追加
Update: ユング-元型論  # 内容更新
Fix: 誤字・事実誤り修正
Reorg: フォルダ構成変更
```

通常の正式ノート編集では、変更内容を確認してから手動で `push` する。スクレイプ結果だけを同期する場合は、下記の専用スクリプトを使う。

`scripts/scrape-and-commit.sh --push` は、ホスト側の Git 認証を使って `research-inbox/`・`book-inbox/` の自動収集結果だけを同期するための運用スクリプトである。正式ノートの編集結果は通常の `git diff` 確認後、意味のある単位でコミットする。
