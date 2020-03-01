import React, { Component } from "react";
import { PostTable } from "./PostTable";
import axios from "axios";

// #TODO: get base url from environment
// think of a better place to put this declaration
// to be available for other components as well
const publicAPI = axios.create({
    baseURL: "http://localhost:3000/api/public/v1",
    timeout: 30000 // 30 secs
});

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
        publicAPI.get("/posts")
            .then((response) => {
                this.setState({
                    posts: response.data,
                    isLoaded: true,
                });
            })
            .catch((err) => {
                this.setState({
                    posts: [],
                    error: err.message,
                    isLoaded: true,
                });
            })
        ;
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

        return table;
    }

}

export default Posts;
