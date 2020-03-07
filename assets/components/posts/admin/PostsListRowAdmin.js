import React, { Component } from "react";
import { Link } from "react-router-dom";
import axios from "axios";

// #TODO: get base url from environment
// think of a better place to put this declaration
// to be available for other components as well
const adminAPI = axios.create({
    baseURL: "http://localhost:3000/api/admin/v1",
    timeout: 30000 // 30 secs
});

export class PostsListRowAdmin extends Component {

    constructor(props) {
        super(props);

        this.deletePost = this.deletePost.bind(this);
    }

    deletePost() {
        if (confirm("Are you sure to delete this post?")) {
            adminAPI.delete("/posts/" + this.props.post.id)
                .then((res) => {
                    this.props.listPosts();
                })
                .catch((err) => {
                    alert(err.message);
                })
            ;
        }
    }

    render() {
        const post = this.props.post;

        return (
            <tr id={ post.id }>
                <td>{ post.id }</td>
                <td>{ post.author }</td>
                <td>{ post.title }</td>
                <td>
                    <Link to={"/admin/posts/" + post.id} className="button">edit</Link>
                    <button className="button" onClick={ this.deletePost }>delete</button>
                </td>
            </tr>
        );
    }

}
