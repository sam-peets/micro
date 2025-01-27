"use client";

import Cookies from "js-cookie";
import { SignIn } from "@/app/lib/api";
import { FormEvent, useState } from "react";

export default function SignInForm() {
    const [error, setError] = useState<string>("");
    async function onSubmit(event: FormEvent<HTMLFormElement>) {
        event.preventDefault()

        const formData = new FormData(event.currentTarget);
        let username = formData.get("username")?.toString()
        let password = formData.get("password")?.toString()
        if (username == undefined || password == undefined) {
            setError("username or password missing")
            return;
        }

        let sess = await SignIn(username, password)
        if (!sess) {
            setError("incorrect username or password");
            return;
        }

        Cookies.set("session", sess.sid)
        window.location.href = "/"
    }

    return (
        <form onSubmit={onSubmit}>
            <p>{error}</p>
            <input type="text" name="username" placeholder="username"></input>
            <input type="password" name="password" placeholder="password"></input>
            <button type="submit">Sign In</button>
        </form>
    )
}