import { Post } from "../types";
import UsernameLink from "./usernameLink";

interface PostProps {
    post: Post
}

export default function PostElem({post}: PostProps) {
    return (
        <div className="border">
            <UsernameLink user={post.user}/>
            <p>{post.timestamp}</p>
            <p>{post.content}</p>
        </div>
    )
}