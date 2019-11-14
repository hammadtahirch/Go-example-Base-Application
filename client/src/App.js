import React, { useRef, useLayoutEffect } from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import Login from './Components/Pages/Account/Login';
import Dashboard from './Components/Pages/Dashboard/Dashboard';

const App = () => {
  const router = useRef(null);

  // Listening to router for route changes and pushing changes to Shopify
  useLayoutEffect(() => {
    router.current.history.listen((location) => {
      ShopifyApp.pushState(location.pathname);
    });
  }, []);

  return (
    
      <Router ref={router}>
        <Switch>
          <Route exact path="/" component={Login} />
          <Route exact path="/login" component={Login} />
          <Route exact path="/dashboard" component={Dashboard} />
        </Switch>
      </Router>
  );
};

export default App