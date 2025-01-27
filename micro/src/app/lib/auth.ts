import Cookies from "js-cookie";
import { ApiCall } from "./api";
import { User } from "../types";

export async function ValidateSession(session: string): Promise<User | null> {
    if (session == null) {
        return null;
    }
    const r = await ApiCall("/api/auth/validate", {"sid": session});
    if (r.status != 200) {
        return null;
    }
    
    return r.data;
}

export async function GetSession(): Promise<string | null> {
    const sess = Cookies.get("session")
    if (sess != undefined) {
        return sess
    }
    return null;
}