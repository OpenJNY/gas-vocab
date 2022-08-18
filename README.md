# GAS - Vocab

Google SpreadSheet で簡単な英単語帳を構築するための GAS コードとそのクライアントアプリ

## セットアップ

Google SpreadSheet

1. スプレッドシートを作成
2. `Sheet1` シートの A1-D1 に `created_at, word, meaning, example` と記入
3. GAS (Extensions > App Script) にて `gas/` 配下の js を作成
4. Deploy > New deployment にて GAS をデプロイ

クライアント

1. ビルドして生成されたバイナリ `vocab` をインストール
2. GAS のデプロイ時に発行された URL を `GAS_VOCAB_URL` 環境変数として設定

## 使い方

```bash
vocab "hooked on"

# meaning
vocab "dazy spells" -m "めまい"

# example
vocab "salt-of-the-earth" -e "Thanks, you guys are salt of the earth."
```
