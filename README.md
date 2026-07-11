# 哲学ノート

哲学者・著作ごとに内容をまとめた個人ナレッジベース。Obsidian Vault として管理し、Go 製スクレイパーが収集した論文・書籍情報を `inbox/<著者>/`（スクレイプのキーワードである哲学者ごと）に蓄積し、Codex / Claude Code が `notes/` の日本語ノートへ統合する。

---

## フォルダ構成

```
philosophy/
├── inbox/             # スクレイパーが保存する未処理素材（論文・書籍を一本化）
│   └── <著者>/        # スクレイプのキーワード（哲学者）ごとにサブフォルダ分け
├── notes/             # Codex / Claude Code がまとめる整理済み日本語ノート
│   ├── 西洋哲学/
│   │   ├── 古代/      # ～紀元後5世紀ごろまで（ソクラテス、プラトン等）
│   │   ├── 近代/      # 17〜19世紀（カント、ヘーゲル、マルクス、ニーチェ等）
│   │   └── 現代/      # 20世紀以降（ユング、アーレント、サルトル等）
│   └── 東洋哲学/
│       ├── インド・仏教/  # ※ファイルが増えた時点でサブフォルダを切る
│       ├── 中国/          # ※同上
│       └── 日本/          # ※同上
├── scrape.yaml        # スクレイピング設定（論文＋書籍を一本化）
├── cmd/               # Go 製 scrapem CLI
├── internal/          # scrapem の内部実装
├── docker-compose.yml
└── .scrapem/
    └── seen-urls.json # 重複排除用の取得済み URL インデックス
```

`notes/` 配下では、書籍ノートも研究動向ノートも著者の地域・時代に応じて同じツリーに同居させる。両者は Front Matter の `種別`（`書籍` / `研究動向`）で区別する。

時代区分は研究者によって異なる場合がある。迷った場合はその哲学者が最も活躍した時期を基準にし、ファイル冒頭に時代的背景を記す。

---

## ファイル命名規則

```
著者名-著作名.md       # 種別: 書籍
著者名-トピック.md     # 種別: 研究動向（例: ユング-現代研究動向.md）
```

**例:**
- `ユング-元型論.md`
- `カント-純粋理性批判.md`
- `中村元-ブッダのことばスッタニパータ.md`
- `ハンナ・アーレント-現代研究動向.md`

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
種別: 書籍
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

研究動向ノートの場合は `種別: 研究動向` とし、`著作` の代わりに `分類`（扱うトピック）を用いてもよい。

`分野` と `キーワード` は複数指定可。`関連ファイル` に同一著者または思想的に近い著作を列挙すると、AI エージェントが横断的なリサーチをする際に使いやすくなる。

---

## 自動収集アーキテクチャ

```text
scrape.yaml
  ↓
Go Scraper（論文: Crossref / arXiv、書籍: Wikipedia ja / SEP / Internet Archive）
  ↓
inbox/<著者>/*.md   （未処理素材。スクレイプのキーワード＝哲学者ごとにサブフォルダ分け。論文は capture_tool: scrapem、書籍は capture_tool: scrapem-book）
  ↓
Codex / Claude Code
  ↓
notes/ 配下の整理済み日本語ノート（書籍ノート・研究動向ノート）
```

情報源の役割分担:

- **Crossref / arXiv**: 論文・草稿のメタデータ・Abstract・PDF 本文。`## PDF Text` に抽出本文を保存
- **Wikipedia ja**: 日本語で哲学者・著作・概念の解説。現代著者もカバー。`## Full Text` に extract を保存
- **Stanford Encyclopedia of Philosophy (SEP)**: 英語の学術的に信頼性の高い長文記事。`## Full Text` に本文を保存
- **Internet Archive**: パブリックドメイン書籍の**フルテキスト**（OCR）。古典のみ対象。`public_domain: true`（`access-restricted-item` が false の場合のみ）

PhilArchive / PhilPapers / gutendex.com（Gutenberg のサードパーティ API）は robots.txt で ClaudeBot を名指しで Disallow しており、Cloudflare 側で GitHub Actions の IP レンジもブロックされていたため使用を停止した。

各層の役割:

- `scrape.yaml`: 収集対象の著者・クエリ・情報源を指定する。`keywords:` セクションで対象を管理し、著者ごとに `sources` で使うソースを選ぶ
- Go Scraper: 各 API / サイトからメタデータ・Abstract・Subjects・本文・公開 URL などを収集する
- `inbox/`: 論文・書籍の未処理素材キュー。スクレイパーは Front Matter の `keyword`（哲学者名）ごとに `inbox/<著者>/` サブフォルダへ振り分けて保存する。`## Full Text` / `## PDF Text` に本文が入り、`max_book_chars` / `max_pdf_chars` で truncate された場合は `book_text_truncated: true` / PDF text truncated のマーカーが付く
- `notes/`: 整理済み日本語ノート（地域・時代で整理し、`種別` で書籍／研究動向を区別）
- Codex / Claude Code: `inbox/` と既存ノートを読み、重複を避けて統合する

### スクレイピングの実行

```bash
# 単発実行
scripts/scrape.sh

# 定期実行（scrape.yaml の interval に従う）
docker compose up -d scheduler
```

スクレイプ後に inbox をコミットする場合:

```bash
scripts/scrape-and-commit.sh
```

push まで行う場合:

```bash
scripts/scrape-and-commit.sh --push
```

自動コミット対象は `inbox/`・`.scrapem/seen-urls.json` のみ。正式ノートの更新は、AI エージェントによる別作業・別 PR として扱う。

**GitHub Actions による定期実行:**
`.github/workflows/scrape.yml` が 6 時間ごと（UTC `0 */6 * * *`）にスクレイパーを走らせ、結果を `main` へ直接 commit/push する。手動実行は GitHub UI から `workflow_dispatch` で起動できる。重複防止のため inbox 内の `source:` URL から `seen-urls.json` を毎回再構築し、書き込みごとに逐次保存する。

**重複ノートのクリーンアップ:**
過去の partial 実行で同じ source URL を持つ `-N.md` が残った場合、以下で整理できる。

```bash
SCRAPEM_VAULT_ROOT="$PWD" go run ./cmd/scrapem dedupe --config scrape.yaml --dry-run
SCRAPEM_VAULT_ROOT="$PWD" go run ./cmd/scrapem dedupe --config scrape.yaml
```

### inbox の扱い

`inbox/` は正式ノートではなく、自動収集素材の置き場である。Codex / Claude Code が参照してまとめる場合は `notes/` 配下に整理する。`public_domain: true` のノートは Internet Archive から本文取得が可能。著作権が切れていないものは書誌情報・概要のみ利用する。

AI エージェントが正式ノートへ反映したら、該当する inbox ノートを次のように更新する。

```yaml
status: processed
processed_to: "notes/西洋哲学/近代/カント-純粋理性批判.md"
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

> 「inbox/ のハンナ・アーレント関連ノートを読んで、既存の notes/西洋哲学/現代/ハンナ・アーレント-*.md と照合し、研究動向ノートに重複しない形で統合して」

> 「inbox/ のカント関連ノートを読んで、notes/西洋哲学/近代/ に新規書籍ノートとして追加して」

> 「inbox/ のうち status: raw のものを確認し、正式ノート化すべきもの、無視してよいものに分類して」

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

`scripts/scrape-and-commit.sh --push` は、ホスト側の Git 認証を使って `inbox/` の自動収集結果だけを同期するための運用スクリプトである。正式ノートの編集結果は通常の `git diff` 確認後、意味のある単位でコミットする。
