import React, { Component } from "react";
import { Link } from "react-router-dom";

export class PostsListRow extends Component {

    constructor(props) {
        super(props);
    }

    render() {
        const post = this.props.post;

        return (
            <tr id={ post.id }>
                <td>{ post.id }</td>
                <td>{ post.author }</td>
                <td>{ post.title }</td>
                <td>
                    <Link to={"/posts/" + post.id} className="button">show</Link>
                </td>
            </tr>
        );
    }

}
