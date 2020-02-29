import React, { Component } from "react";
import { PostRow } from "./PostRow";

export class PostTable extends Component {

    constructor(props) {
        super(props);
    }

    render() {
        let rows = [];

        this.props.posts.forEach(function(post) {
            rows.push(<PostRow key={ post.id } post={ post } />);
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
