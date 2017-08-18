export default class websocket {
  constructor() {
    this.ws = new WebSocket("ws://localhost:5050")
  }

  send(payload) {
    console.debug(payload)
    this.ws.send(JSON.stringify(payload))
  }

  listen() {
    return new Promise((resolve, reject) => {

      this.ws.addEventListener('message', (event) => {
        let response = JSON.parse(event.data)

        if (response.error) {
            reject(response.error)
        }

        console.debug(response)
        resolve(response)
      })

    })
  }
}
