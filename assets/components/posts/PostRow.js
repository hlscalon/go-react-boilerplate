import React, { Component } from "react";

export class PostRow extends Component {

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
                <td></td>
            </tr>
        );
    }

}
