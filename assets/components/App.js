import React, { Component } from "react";
import {HashRouter as Router, Switch, Route, Link} from "react-router-dom";
import Home from "./Home";
import Login from "./Login";
import Signup from "./Signup";
import PostsList from "./posts/PostsList";
import PostsDetail from "./posts/PostsDetail";
import PostsDetailAdmin from "./posts/PostsDetailAdmin";
import RouteNotFound from "./RouteNotFound";

import "../styles/css/index.css";

class App extends Component {

    render() {
        return (
            <Router>
                <div>
                    <section className="hero is-light">
                        <nav className="navbar" role="navigation" aria-label="main navigation">
                            <div className="navbar-menu is-active">
                                <div className="navbar-start">
                                    <Link to="/" className="navbar-item">Home</Link>
                                    <Link to="/posts" className="navbar-item">Posts</Link>
                                </div>
                                <div className="navbar-end">
                                    <div className="navbar-item">
                                        <div className="buttons">
                                            <Link to="/signup" className="button is-primary"><strong>Sign up</strong></Link>
                                            <Link to="/login" className="button is-light">Log in</Link>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </nav>
                    </section>
                </div>

                <section className="section">
                    <div className="container">
                        <Switch>
                            <Route path="/login">
                                <Login />
                            </Route>
                            <Route path="/signup">
                                <Signup />
                            </Route>
                            <Route path="/admin/posts/:postID">
                                <PostsDetailAdmin />
                            </Route>
                            <Route path="/posts/:postID">
                                <PostsDetail />
                            </Route>
                            <Route path="/posts">
                                <PostsList />
                            </Route>
                            <Route path="/" exact={true}>
                                <Home />
                            </Route>
                            <Route path="*">
                                <RouteNotFound />
                            </Route>
                        </Switch>
                    </div>
                </section>
            </Router>
        );
    }
}

export default App;
