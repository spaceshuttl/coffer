# Coffer
*Coffer is very much in development - and should not be used seriously.*
Coffer is a light-weight password manager built ontop of Go's secure backend, with Electron + React serving a beautiful front-end for intuitive UI.

## Data retention
Coffer stores all passwords within a local file. All transations to the file are encrypted with AES (tbc), so all data kept is safe.

# TODO

### Backend
- [ ] Encrypt disk transactions
- [ ] Write test + fuzz the endpoints

### Frontend
- [ ] Write up html structure
- [ ] Write React components
- [ ] Experiment with AJAX vs WebSockets

### Meta
- [ ] Create unison build system
