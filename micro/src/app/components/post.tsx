import { User, Post } from "../types";

interface PostProps {
    post: Post
}

export default function PostElem({post}: PostProps) {
    return (
        <div className="border">
            <p>{post.user.username} - {post.timestamp}</p>
            <p>{post.content}</p>
        </div>
    )
}