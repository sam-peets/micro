"use client";
import Cookies from "js-cookie"
import { useEffect, useState } from "react";
import { GetSession, ValidateSession } from "../lib/auth";
import { User } from "../types";

function login() {
    window.location.href = "/signin"
}

function logout() {
    Cookies.remove("session")
    window.location.reload()
}

export default function UserStatus() {
    const [data, setData] = useState<User | null>(null);

    useEffect(() => {
        GetSession().then(session => {
            if (session) {
                ValidateSession(session).then(user => {
                    setData(user)
                })
            }
        })
    }, [])
    let labelText, buttonText, buttonAction;
    if (data) {
        labelText = data.username;
        buttonText = "Sign Out"
        buttonAction = logout;
    } else {
        labelText = "not signed in";
        buttonText = "Sign In"
        buttonAction = login;
    }

    return (
        <div style={{float: "right"}}>
            <label style={{padding: "2px"}}>{labelText}</label>
            <button onClick={login}>{buttonText}</button>
        </div>
    )
}