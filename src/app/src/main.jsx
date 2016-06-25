import ReactDOM from 'react-dom'
import { Router, hashHistory, Route, IndexRoute } from 'react-router'

import Login from './pages/login'
import Manager from './pages/manager'
import About from './pages/about'
import Container from './container'

ReactDOM.render(
  <Router history={hashHistory}>
    <Route path="/" component={Container}>
      <IndexRoute component={Login} />
      <Route path="about" component={About} />
      <Route path="manager" component={Manager} />
    </Route>
  </Router>,
  document.getElementById('app')
)
