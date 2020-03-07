import React, { Component } from "react";
import { withRouter, Link } from "react-router-dom";
import axios from "axios";

// #TODO: get base url from environment
// think of a better place to put this declaration
// to be available for other components as well
const adminAPI = axios.create({
    baseURL: "http://localhost:3000/api/admin/v1",
    timeout: 30000 // 30 secs
});

class PostsDetailAdmin extends Component {

    constructor(props) {
        super(props);

        this.state = {
            post: {},
            id: props.match.params.postID || 0,
            isLoaded: false,
            error: null,
        };

        this.history = props.history;
        this.location = props.location;
        this.getPost = this.getPost.bind(this);
        this.savePost = this.savePost.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
        this.handleInputChange = this.handleInputChange.bind(this);
    }

    handleInputChange(event) {
        const target = event.target;
        const value = target.type === "checkbox" ? target.checked : target.value;
        const name = target.name;

        this.setState({
            post: {
                ...this.state.post,
                [name]: value
            }
        });
    }

    handleSubmit(event) {
        this.savePost();
        event.preventDefault();
    }

    savePost() {
        var promise = null;
        if (this.state.id > 0) {
            promise = adminAPI.put("/posts/" + this.state.id, this.state.post);
        } else {
            promise = adminAPI.post("/posts", this.state.post);
        }

        promise
            .then((response) => {
                this.history.push("/admin/posts");
            })
            .catch((err) => {
                alert(err.message);
            })
        ;
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

        adminAPI.get("/posts/" + this.state.id)
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
        if (this.location.pathname != "/admin/posts/create") {
            this.getPost();
        } else {
            this.setState({
                isLoaded: true,
            });
        }
    }

    render() {
        let html;

        if (!this.state.isLoaded) {
            html = <div>Loading...</div>;
        } else if (this.state.error) {
            html = <div>Problem loading information:<br/>{ this.state.error }</div>;
        } else {
            html =
                <form onSubmit={ this.handleSubmit }>
                    <div className="field">
                        <label className="label">Author</label>
                        <div className="control">
                            <input className="input" type="text" name="author" value={ this.state.post.author } onChange={ this.handleInputChange } />
                        </div>
                    </div>

                    <div className="field">
                        <label className="label">Title</label>
                        <div className="control">
                            <input className="input" type="text" name="title" value={ this.state.post.title } onChange={ this.handleInputChange } />
                        </div>
                    </div>

                    <div className="field">
                        <label className="label">Description</label>
                        <div className="control">
                            <textarea className="textarea" name="description" value={ this.state.post.description } onChange={ this.handleInputChange } ></textarea>
                        </div>
                    </div>

                    <div className="field is-grouped">
                        <p className="control">
                            <button className="button is-primary" type="submit">Submit</button>
                        </p>
                        <p className="control">
                            <Link to={"/admin/posts/"} className="button is-light">Cancel</Link>
                        </p>
                    </div>
                </form>
            ;
        }

        return html;
    }

}

export default withRouter(PostsDetailAdmin);
