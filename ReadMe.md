## 使い方

docker・docker-composeが必須です

1. src/api/.env.localのファイル名を.envに変更
1. `docker-compose up -d`を実行
2. [ここ](http://localhost/sections)にアクセス
   - apiのコンパイルに時間がかかるので少し待つかもしれません
3. 以上…で取りあえず動くはずです

### おまけ

[ここ](http://localhost:3000/graphql)にアクセスするとgraphqlにquery投げたりできます

[ここ](http://localhost:8001)にアクセスするとローカルのdynamodbのテーブルが見れます

何故かplantumlのコンテナまで入ってます

documentsとかありますが。あまり意味はありません

## 注意点

お試しで作ったものなので、ボタンとかが反応なかったりします

### apiについて

wireとか入れてますが使ってません

error処理が結構適当です

### bffについて

これはいずれファットフロントになるので参考にしないでください

### frontについて

センスがなくてすいません

## エラー対応

### `docker-compose up -d`の際のエラー

`docker-compose up -d`の際に以下エラーが出た場合

```
ERROR: An HTTP request took too long to complete. Retry with --verbose to obtain debug information.
If you encounter this issue regularly because of slow network conditions, consider setting COMPOSE_HTTP_TIMEOUT to a higher value (current value: 60).
```

以下コマンドを試す

`COMPOSE_HTTP_TIMEOUT=600 docker-compose up -d`