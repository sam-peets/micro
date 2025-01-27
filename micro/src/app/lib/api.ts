import axios, { AxiosResponse } from "axios";
import { Post, Session, User } from "../types";

export async function ApiCall(endpoint: string, data: any): Promise<AxiosResponse<any, any>> {
    let res = await axios.post(endpoint, data);

    return res;
}

export async function GetUser(uid: number): Promise<User> {
    let x = await ApiCall("/api/users", {
        "sid": "57559913b3ace9a5a5298ed5a8542c69bc07bbeb3679babc3e840ed647c3400d",
        "uid": uid,
    })

    return x.data;
}

interface PostResponse {
    uid: number,
    postid: number,
    content: string,
    timestamp: string,
}

async function postResponseToPost(res: PostResponse): Promise<Post> {
    let u = await GetUser(res.uid);
    
    return {
        "user": u,
        "content": res.content,
        "timestamp": res.timestamp,
        "id": res.postid,
    }
}

export async function GetRecent(limit: number, skip: number): Promise<Post[]> {
    let x: PostResponse[] = (await ApiCall("/api/posts/recent", {
        "sid": "57559913b3ace9a5a5298ed5a8542c69bc07bbeb3679babc3e840ed647c3400d",
        "limit": limit,
        "skip": skip
    })).data

    let posts: Post[] = [];
    for (let pr of x) {
        posts.push(await postResponseToPost(pr))
    }

    return posts
}



export async function SignIn(username: string, password: string): Promise<Session | null> {
    let x = (await ApiCall("/api/auth", {
        "username": username,
        "password": password
    }))

    if (x.status != 200) {
        return null;
    }

    return x.data;
}