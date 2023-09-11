# Channel

![channel](./public/channel.png)

> 채널

- 정보 전달
- GoRoutine 사이에는 서로 독립적으로 실행
- GoRoutine간에 데이터가 공유가 될려면 Channel을 사용해야 한다.
- 즉, 비동기 통신간의 데이터 공유를 하려면 Channel이 필요 함

> Select

- GoRoutine이 여러 통신작업을 기다리게한다.
