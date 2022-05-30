# Database

## Property

```json
{
    "id": "id: PK,number",
    "name": "物件名: string",
    "created_at": "作成日: date",
    "updated_at": "最終更新日: date",
}
```

## Section

```json
{
    "id": "id: PK,number",
    "property_id": "物件ID: number",
    "type": "区画タイプ(1:部屋,2:駐車場): number",
    "name": "区画名: string",
    "created_at": "作成日: date",
    "updated_at": "最終更新日: date",
}
```