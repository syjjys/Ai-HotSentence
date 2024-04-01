import React from 'react';
import { Router, Route, Switch,Redirect } from 'dva/router';
import SomePage from './routes/SomePage';

function RouterConfig({ history }) {
  return (
    <Router history={history}>
      <Switch>
        {/* 重定向根路径到 /main */}
        <Redirect from="/" exact to="/main" />
        <Route path="/main" exact component={SomePage} />
        {/* 可以在这里添加更多的路由 */}
      </Switch>
    </Router>
  );
}

export default RouterConfig;