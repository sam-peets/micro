import { User } from "../types";

export default function usernameLink({user}: {user: User | null}) {
    if (user) {
        return (
            <a href={`/user?id=${user.uid}`}>{user.username}</a>
        )
    } else {
        return (<></>)
    }
    
}