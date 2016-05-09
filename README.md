*Coffer is very much in development - and should not be used seriously.*

![](https://a.pomf.cat/pxqfjn.png)

Coffer is a light-weight password manager built ontop of Go's secure backend, with Electron + React serving a beautiful front-end with an  intuitive UI.

## Data retention
Coffer stores all passwords within a local file. All transations to the file are encrypted with AES (tbc), so all data kept is safe.

# TODO

### Backend
- [X] Rewrite endpoints in gorilla/websocket
- [ ] Encrypt disk transaction
- [ ] Write test + fuzz the endpoint

### Frontend
- [X] Write up html structure
- [X] Write React components
- [X] Experiment with AJAX vs WebSockets
- [ ] Automatically hide passwords, show on click/hover
- [ ] Implement Flux data store
  - [ ] Bind WS/AJAX calls to flux store for seemless data sync

### Meta
- [ ] Create unison build system


## notes - delete me
* componentName = connect() (componentName), Now you use {dispatch} as an argument to the component and then you can use it how ever you like within that component
