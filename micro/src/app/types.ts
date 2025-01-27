export interface User {
    username: string;
    uid: string;
}

export interface Post {
    user: User;
    content: string;
    timestamp: string;
    id: number;
}

export interface Session {
    sid: string,
    uid: number,
    expires: string
}