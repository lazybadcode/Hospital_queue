NOTE:
- priority ของ queue per day เรียงจาก วันหยุดเสาร์อาทิตย์ วันหยุดพิเศษ วันธรรมดา
- checking config/config.yaml before use

Feature:
- graphql
  - create queue
  - query queue (all, by queue id, by user id)
  - update queue
  - delete queue
- batch
  - send sms at 22:00 (only log in console)

HOW TO USE
1. start api
```shell
TZ=Asia/Bangkok go run .
```
2. start database
```shell
 docker compose up
```

3. add index to mongo for avoid race condition
```javascript
db.queues.createIndex({ date: 1, no: 1 }, { unique: true })
```

4. Go to http://localhost:8080/ for test

Create
```graphql
mutation {
    createQueue(name: "Queue Name", id_card: "1234567890", mobile_no: "1234567890", note:"อาการ ประวัติการรักษาที่ รพ") {
        _id
        user {
            name
            id_card
        }
    }
}
```

Query All
```graphql
query queue {
    queue (input:{user_id:"", queue_id:""}) {
        _id
        no
        slot
        user_id
        user {
            name
            id_card
        }
    }
}
```
Query 1 User
```graphql
query queue {
    queue (input:{user_id:"657c7810254a5cfcf2cbeaf3", queue_id:""}) {
        _id
        no
        slot
        user_id
        user {
            name
            id_card
        }
    }
}
```
Update Queue (specific date,slot,note) 
```graphql
mutation {
    updateQueue(id: "657d797b3b47387f3397a6b0", date: "20231217", slot: 7, note: "notee") {
        _id
        no
        slot
        note
        date
        user {
            name
            id_card
        }
    }
}
```

Update Queue only note
```graphql
mutation {
    updateQueue(id: "657d797b3b47387f3397a6b0", date: "", slot: 0, note: "notee") {
        _id
        no
        slot
        note
        date
        user {
            name
            id_card
        }
    }
}
```

Delete queue
```graphql
mutation {
    deleteQueue(id:"657d4b82d67529915249c768")
}
```



mock special days
```javascript
db = db.getSiblingDB('test');

var specialDays = [
    { "date": "20230101", "description": "New Year's Day" },
    { "date": "20230208", "description": "Makha Bucha" },
    { "date": "20230406", "description": "Chakri Day" },
    { "date": "20230413", "description": "Songkran Festival" },
    { "date": "20230414", "description": "Songkran Festival" },
    { "date": "20230415", "description": "Songkran Festival" },
    { "date": "20230501", "description": "Labor Day" },
    { "date": "20230504", "description": "Coronation Day" },
    { "date": "20230506", "description": "Vesak Day" },
    { "date": "20230603", "description": "H.M. Queen Suthida's Birthday" },
    { "date": "20230703", "description": "Asalha Bucha Day" },
    { "date": "20230704", "description": "Buddhist Lent Day" },
    { "date": "20230728", "description": "H.M. King's Birthday" },
    { "date": "20230812", "description": "Mother's Day" },
    { "date": "20231023", "description": "Chulalongkorn Day" },
    { "date": "20231205", "description": "Father's Day" },
    { "date": "20231210", "description": "Constitution Day" },
    { "date": "20231231", "description": "New Year's Eve" }
];

db.special_days.insertMany(specialDays);
```

# ขั้นตอนในการเพิ่ม Query และ Mutation
หากคุณต้องการเพิ่ม Query หรือ Mutation ในโครงการของคุณ กรุณาปฏิบัติตามขั้นตอนดังนี้:
1. ไปที่ไฟล์ `graph/schema.graphqls`: เพิ่มฟังก์ชันใน `type Query` หรือ `type Mutation` ตามที่คุณต้องการให้มีอยู่ใน GraphQL schema ของคุณ.
2. รันคำสั่ง `go generate ./...`: การรันคำสั่งนี้จะทำให้เพิ่มฟังก์ชันที่คุณได้เพิ่มเข้ามาในไฟล์ `graph/schema.resolvers.go` อัตโนมัติครับ.
3. ไปที่ไฟล์ `graph/schema.resolvers.go`: ที่นี่คุณจะต้อง implement logic สำหรับแต่ละฟังก์ชันที่คุณเพิ่มเข้ามาใหม่ โดยทั่วไปแล้วเรียก function ที่ implement logic ที่ package ชื่อ usecase ของคุณ.
4. คุณสามารถหาข้อมูลเพิ่มเติมได้ที่:
  - [เริ่มต้นกับ GraphQL ใน Golang](https://www.howtographql.com/graphql-go/1-getting-started/)
  - [การใช้ GraphQL กับ Golang โดย Apollo](https://www.apollographql.com/blog/using-graphql-with-golang)

หลังจากปฏิบัติตามขั้นตอนดังกล่าวแล้ว คุณควรจะสามารถเพิ่ม Query และ Mutation ใหม่ในโปรเจ็คของคุณได้เรียบร้อยแล้วครับ!