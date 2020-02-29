import React, { Component } from "react";
import {HashRouter as Router, Switch, Route, Link} from "react-router-dom";
import Home from "./Home";
import Login from "./Login";
import Signup from "./Signup";
import Posts from "./posts/Posts";
import RouteNotFound from "./RouteNotFound";

class App extends Component {

    render() {
        return (
            <Router>
                <nav>
                  <ul>
                    <li>
                      <Link to="/">Home</Link>
                    </li>
                    <li>
                      <Link to="/posts">Posts</Link>
                    </li>
                    <li>
                      <Link to="/login">Login</Link>
                    </li>
                    <li>
                      <Link to="/signup">Sign up</Link>
                    </li>
                  </ul>
                </nav>
                <Switch>
                  <Route path="/login">
                    <Login />
                  </Route>
                  <Route path="/signup">
                    <Signup />
                  </Route>
                  <Route path="/posts">
                    <Posts />
                  </Route>
                  <Route path="/" exact={true}>
                    <Home />
                  </Route>
                  <Route path="*">
                    <RouteNotFound />
                  </Route>
                </Switch>
            </Router>
        );
    }
}

export default App;
