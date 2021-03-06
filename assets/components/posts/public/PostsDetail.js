import React, { Component } from "react";
import { withRouter } from "react-router-dom";
import axios from "axios";

// #TODO: get base url from environment
// think of a better place to put this declaration
// to be available for other components as well
const publicAPI = axios.create({
    baseURL: "http://localhost:3000/api/public/v1",
    timeout: 30000 // 30 secs
});

class PostsDetail extends Component {

    constructor(props) {
        super(props);

        this.state = {
            post: {},
            id: props.match.params.postID,
            isLoaded: false,
            error: null,
        };

        this.getPost = this.getPost.bind(this);
    }

    getPost() {
        if (!this.state.id || this.state.id <= 0) {
            this.setState({
                post: {},
                isLoaded: true,
                error: "Invalid identifier",
            });

            return;
        }

        publicAPI.get("/posts/" + this.state.id)
            .then((response) => {
                this.setState({
                    post: response.data,
                    isLoaded: true,
                });
            })
            .catch((err) => {
                this.setState({
                    post: {},
                    error: err.message,
                    isLoaded: true,
                });
            })
        ;
    }

    componentDidMount() {
        this.getPost();
    }

    render() {
        let html;

        if (!this.state.isLoaded) {
            html = <div>Loading...</div>;
        } else if (this.state.error) {
            html = <div>Problem loading information:<br/>{ this.state.error }</div>;
        } else {
            html =
                <div className="card">
                    <header className="card-header">
                        <p className="card-header-title">
                            { this.state.post.title }
                        </p>
                    </header>
                    <div className="card-content">
                        <div className="content">
                            { this.state.post.description }
                        </div>
                        <p className="subtitle is-6">
                            @{ this.state.post.author }
                        </p>
                    </div>
                </div>
            ;
        }

        return html;
    }

}

export default withRouter(PostsDetail);
