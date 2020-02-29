import React, { Component } from "react";
import { PostTable } from "./PostTable";

import "../../styles/css/index.css";

class Posts extends Component {

    constructor(props) {
        super(props);

        this.state = {
            posts: [],
            isLoaded: false,
            error: null,
        };

        this.listPosts = this.listPosts.bind(this);
    }

    listPosts() {
        let posts = [
            {"id": 1, "author": "hlscalon", "title": "this is the first post ever", "description": "this post is awewsome"},
            {"id": 2, "author": "batman", "title": "gotham needs you", "description": "I am out of business"},
            {"id": 3, "author": "robin", "title": "nooooo", "description": "Please Batman, don't go!!!"},
        ];
        this.setState({
            posts: posts,
            isLoaded: true,
        });
    }

    componentDidMount() {
        this.listPosts();
    }

    render() {
        let table;

        if (!this.state.isLoaded) {
            table = <div>Loading...</div>;
        } else if (this.state.error) {
            table = <div>Problem loading information:<br/>{this.state.error}</div>;
        } else {
            table = <PostTable posts={ this.state.posts } />;
        }

        return (
            <div>
                <section className="section">
                    <div className="container">
                        { table }
                    </div>
                </section>
            </div>
        );
    }

}

export default Posts;
