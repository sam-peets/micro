import Cookies from "js-cookie";
import { ApiCall } from "./api";
import { User } from "../types";

export async function ValidateSession(session: string): Promise<User | null> {
    let sid = session;
    if (sid == null) {
        return null;
    }
    let r = await ApiCall("/api/auth/validate", {"sid": sid});
    if (r.status != 200) {
        return null;
    }
    
    return r.data;
}

export async function GetSession(): Promise<string | null> {
    let sess = Cookies.get("session")
    if (sess != undefined) {
        return sess
    }
    return null;

}