import ReactDOM from 'react-dom'
import { Router, hashHistory, Route, IndexRoute } from 'react-router'

import Container from './container'

import Login from './pages/login'
import Manager from './pages/manager'
import About from './pages/about'

ReactDOM.render(
  <Router history={hashHistory}>
    <Route path="/" component={Container}>
      <IndexRoute   component={Login} />
      <Route        component={Manager}  path="manager"/>
      <Route        component={About}    path="about"/>
    </Route>
  </Router>,
  document.getElementById('frontend')
)
