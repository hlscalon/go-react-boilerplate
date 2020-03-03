import React, { Component } from "react";
import { PostsListRow } from "./PostsListRow";

export class PostsListTable extends Component {

    constructor(props) {
        super(props);
    }

    render() {
        let rows = [];

        this.props.posts.forEach(function(post) {
            rows.push(<PostsListRow key={ post.id } post={ post } />);
        });

        return (
            <table className="table is-hoverable is-fullwidth">
                <thead>
                    <tr>
                        <th>ID</th>
                        <th>Author</th>
                        <th>Title</th>
                        <th></th>
                    </tr>
                </thead>
                <tbody>{ rows }</tbody>
            </table>
        );
    }

}
