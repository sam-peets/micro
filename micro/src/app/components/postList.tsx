import { Post } from "../types";
import PostElem from "./post";

export default function PostList({posts}: {posts: Post[]}) {
    const p = posts.map(post => (
        <PostElem post={post} key={post.id}/>
    ));
    return p;
}