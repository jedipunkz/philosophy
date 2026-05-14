# CLAUDE.md

このリポジトリは哲学者・著作ごとの個人ナレッジベース。Obsidian Vault として管理し、Go 製スクレイパーが収集した論文情報を `research-inbox/` に、書籍情報を `book-inbox/` に蓄積し、Codex / Claude Code が `研究動向/` および `書籍/` の日本語ノートへ統合する。
プロジェクト全体像・運用方針の詳細は `README.md` を参照すること。このファイルは AI エージェントが作業時に従う最小限のルールに絞る。

## 作業前チェック

- 新規ファイルを作る前に既存ファイル（特に同じ時代区分・同じ著者）を1〜2本読み、文体・分量・章立ての粒度を合わせる
- 著者名表記が既存ファイルと揺れないか確認する（例: ハンナ・アーレント、ユング）
- `research-inbox/` または `book-inbox/` を使う作業では、まず `status: raw` / `status: processed` / `status: ignored` を確認する

## 新規ファイル作成ルール

- 配置: `書籍/<西洋哲学|東洋哲学>/<古代|近代|現代|（東洋はサブフォルダ未設定なら直下）>/`
- ファイル名: `著者名-著作名.md`（区切りはハイフンのみ。スペース・アンダースコア不可。著作名は省略しない）
- フロントマター（YAML）をファイル冒頭に付与する。形式は `README.md` の「ファイルフォーマット」節に従う
  - 必須キー: `著者` / `著作` / `時代` / `地域` / `分野` / `キーワード` / `関連ファイル`
  - `関連ファイル` には同一著者または思想的に近い著作を列挙する
- 既存ファイルにフロントマターが無いものがあるが、**新規追加時は必ず付ける**

## research-inbox 統合ルール（論文）

- `research-inbox/` はスクレイパーが生成した未処理論文素材キューであり、正式ノートではない
- `research-inbox/` の内容をそのまま移動・翻訳するのではなく、既存ノートと照合して日本語で再構成する
- 統合作業では、最初に `README.md` と関連する既存ノートを読む
- 対象は原則として `status: raw` のノート。`processed` は必要な場合だけ再確認する
- 正式ノートに反映した `research-inbox/` ノートは、Front Matter を次の形に更新する

```yaml
status: processed
processed_to: "研究動向/<反映先ファイル>.md"
```

- 関係が薄い、品質が低い、重複している素材は `status: ignored` にし、必要なら `ignore_reason` を付ける
- PDF本文は `## PDF Text` に入るが、抽出ノイズがあるため、必ず Abstract / Citation / 既存ノートと照合して使う
- `source`、`citation`、`BibTeX` がある場合は、正式ノートの末尾または参照節に出典として反映する
- スクレイパーは正式ノートを直接更新しない。`research-inbox/` 由来の研究動向ノート化は Codex / Claude Code の編集作業として `研究動向/` に行う

## book-inbox 統合ルール（書籍）

- `book-inbox/` は書籍スクレイパー（Open Library / Project Gutenberg）が生成した未処理書籍素材キューであり、正式ノートではない
- `capture_tool: scrapem-book` のフロントマターを持つ
- `public_domain: true` のノートは Project Gutenberg 由来で、`## Full Text` セクションに本文が含まれる（`max_book_chars` 上限で truncate される場合は `book_text_truncated: true` が立つ）
- Open Library 由来のノートは `## 概要` に description が入る（取得できた場合のみ）
- 取得元 URL は `source`、本文 URL は `plain_text_url` / `gutenberg_url` フィールドを参照する
- 統合先は `書籍/<西洋哲学|東洋哲学>/.../<著者名-著作名>.md`
- 正式ノートに反映した `book-inbox/` ノートは、Front Matter を次の形に更新する

```yaml
status: processed
processed_to: "書籍/<反映先ファイル>.md"
```

- 著作権が切れていない書籍（`public_domain` フィールドなし）は書誌情報・概要のみ利用し、本文は引用しない
- スクレイパーは `書籍/` を直接更新しない。`book-inbox/` 由来の書籍ノート化は Codex / Claude Code の編集作業として `書籍/` に行う

## フォルダ運用の閾値

- `書籍/東洋哲学` のファイルが 5 本を超えたら `インド・仏教/` `中国/` `日本/` のサブフォルダに分割する
- 一人の著者のファイルが 5 本に達したら著者名サブフォルダ化（例: `書籍/西洋哲学/現代/ユング/`）を提案する
- 現時点でユングが 5 本あるため、次に同著者を追加する際はユーザーに移行を確認する

## 横断的タスク（比較・要約）の指針

- 既存ファイルの内容を引用・参照する場合は、必ずファイルパスを明示する
- `research-inbox/` または `book-inbox/` の素材を根拠にする場合も、参照したファイルのパスを作業結果に明示する
- ファクトチェックを依頼された場合は Web 検索で裏取りし、どこを修正したかを diff で示す

## スクレイパー運用

### 論文スクレイパー（research-inbox 向け）

- 設定ファイル: `research-scrape.yaml`
- 単発実行: `scripts/scrape.sh`
- 定期実行: `docker compose up -d research-scheduler`
- 運用段階では論文ソースの更新頻度を考慮し、週1回程度を基本にする。テスト中のみ高頻度でよい

### 書籍スクレイパー（book-inbox 向け）

- 設定ファイル: `book-scrape.yaml`（Open Library API / Project Gutenberg Gutendex API）
- 単発実行: `scripts/scrape.sh --config book-scrape.yaml`
- 定期実行: `docker compose up -d book-scheduler`
- 書籍情報の更新頻度は低いため、週1回（168h）で十分
- 対象著者・著作は `book-scrape.yaml` の `keywords:` セクションで管理する

### 共通

- `scripts/scrape-and-commit.sh` は `research-inbox/`・`book-inbox/`・`.scrapem/seen-*.json` だけを自動コミットするホスト側スクリプト
- Docker コンテナ内に GitHub 認証情報や SSH 鍵を持たせない

## Git 運用

- コミットメッセージは `Add: 著者-著作名` / `Update: 著者-著作名` / `Fix: ...` / `Reorg: ...` のいずれかの prefix を使う
- `git add` / `git commit` までは AI エージェントが実行してよい
- `git push` は **実行しない**。サンドボックスから SSH push できないため、ユーザーがローカル端末で手動実行する
- スクレイプ結果だけの自動同期は `scripts/scrape-and-commit.sh --push` で行う。正式ノートの変更とはコミットを分ける

## 編集時の注意

- 既存ファイル内に検索ノイズらしき文字列（広告コピー的な英字混入など）を見つけた場合は、勝手に削除せずユーザーに確認を取ってから修正する
- ファイル末尾に余計な空行を増やさない
- Obsidian が生成する `.obsidian/workspace.json` や canvas などのUI状態ファイルは、依頼がない限り編集・整理対象にしない
